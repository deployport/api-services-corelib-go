package signingv1

import (
	"crypto/hmac"
	"crypto/sha256"
	"log/slog"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type headerMatcher interface {
	MatchHeader(name string) bool
}

type headerMatcherFunc func(name string) bool

func (h headerMatcherFunc) MatchHeader(name string) bool {
	return h(name)
}

func matchHeader(name string) headerMatcherFunc {
	return func(n string) bool {
		return n == name
	}
}

func negateMatcher(m headerMatcher) headerMatcherFunc {
	return func(name string) bool {
		return !m.MatchHeader(name)
	}
}

type headerMatchers []headerMatcher

func (h headerMatchers) MatchHeader(name string) bool {
	for _, m := range h {
		if m.MatchHeader(name) {
			return true
		}
	}
	return false
}

// AuthorizationHeader is the name of the header that contains the signature
const AuthorizationHeader = "Authorization"

type headerSignatureBuilder struct {
	logger *slog.Logger
	// Signed       http.Header
	Request      *http.Request
	Body         []byte
	orderedNames []string

	// separated by ;
	signedHeaderNames string
	canonicalHeaders  string
	signedHeaderVals  http.Header
	query             url.Values
	canonicalString   string
	BodyDigest        string
	stringToSign      string
	VendorCode        string
}

func (ctx *headerSignatureBuilder) collectOrderedNames() {
	requiredSignedHeaders := headerMatchers{
		matchHeader(vendorDateHeaderKey(ctx.VendorCode)),
		matchHeader(vendorHeaderCanonicalNameKey(ctx.VendorCode, "Service")),
		matchHeader(vendorHeaderCanonicalNameKey(ctx.VendorCode, "Resource")),
		matchHeader(vendorHeaderCanonicalNameKey(ctx.VendorCode, "Operation")),
		matchHeader(vendorHeaderCanonicalNameKey(ctx.VendorCode, "Content-Sha256")),
		matchHeader(vendorHeaderCanonicalNameKey(ctx.VendorCode, "Region")),
		matchHeader("Host"),
	}
	ctx.signedHeaderVals = http.Header{}
	for k, v := range ctx.Request.Header {
		if !requiredSignedHeaders.MatchHeader(k) {
			continue
		}
		ctx.orderedNames = append(ctx.orderedNames, k)
		ctx.signedHeaderVals[k] = v
	}
	sort.Strings(ctx.orderedNames)
	signed := make([]string, len(ctx.orderedNames))
	for i, k := range ctx.orderedNames {
		signed[i] = strings.ToLower(k)
	}
	ctx.signedHeaderNames = strings.Join(signed, ";")
}

func (ctx *headerSignatureBuilder) buildCanonicalHeaders() {

	headerItems := make([]string, len(ctx.orderedNames))

	// log.Printf("ctx.orderedNames: %v", ctx.orderedNames)
	for i, k := range ctx.orderedNames {
		if k == "host" {
			if host := ctx.Request.Host; host != "" {
				headerItems[i] = "host:" + host
			} else {
				headerItems[i] = "host:" + ctx.Request.Host
			}
		} else {
			headerValues := make([]string, len(ctx.signedHeaderVals[k]))
			for i, v := range ctx.signedHeaderVals[k] {
				headerValues[i] = strings.TrimSpace(v)
			}
			headerItems[i] = k + ":" +
				strings.Join(headerValues, ",")
		}
	}

	ctx.logger.Debug("headerItems", slog.Any("headerItems", headerItems))
	// stripExcessSpaces(headerItems)
	ctx.canonicalHeaders = strings.Join(headerItems, "\n")
}

func (ctx *headerSignatureBuilder) build() {
	ctx.query = ctx.Request.URL.Query()
	ctx.collectOrderedNames()
	ctx.buildCanonicalHeaders()
	ctx.buildCanonicalString()
}

func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func hmacSHA256(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func (ctx *headerSignatureBuilder) buildCanonicalString() {
	ctx.Request.URL.RawQuery = strings.Replace(ctx.query.Encode(), "+", "%20", -1)

	uri := getURIPath(ctx.Request.URL)

	ctx.canonicalString = strings.Join([]string{
		ctx.Request.Method,
		uri,
		ctx.Request.URL.RawQuery,
		ctx.canonicalHeaders + "\n",
		ctx.signedHeaderNames,
		ctx.BodyDigest,
	}, "\n")
}

// func (ctx *headerSignatureBuilder) buildCanonicalHeaders(r headerMatcher, header http.Header) {
// 	var headers []string
// 	headers = append(headers, "host")
// 	for k, v := range header {
// 		if !r.MatchHeader(k) {
// 			continue // ignored header
// 		}
// 		if ctx.SignedHeaderVals == nil {
// 			ctx.SignedHeaderVals = make(http.Header)
// 		}

// 		lowerCaseKey := strings.ToLower(k)
// 		if _, ok := ctx.SignedHeaderVals[lowerCaseKey]; ok {
// 			// include additional values
// 			ctx.SignedHeaderVals[lowerCaseKey] = append(ctx.SignedHeaderVals[lowerCaseKey], v...)
// 			continue
// 		}

// 		headers = append(headers, lowerCaseKey)
// 		ctx.SignedHeaderVals[lowerCaseKey] = v
// 	}
// 	sort.Strings(headers)

// 	ctx.signedHeaders = strings.Join(headers, ";")

// 	headerItems := make([]string, len(headers))
// 	for i, k := range headers {
// 		if k == "host" {
// 			if ctx.Host != "" {
// 				headerItems[i] = "host:" + ctx.Host
// 			} else {
// 				headerItems[i] = "host:" + ctx.URL.Host
// 			}
// 		} else {
// 			headerValues := make([]string, len(ctx.SignedHeaderVals[k]))
// 			for i, v := range ctx.SignedHeaderVals[k] {
// 				headerValues[i] = strings.TrimSpace(v)
// 			}
// 			headerItems[i] = k + ":" +
// 				strings.Join(headerValues, ",")
// 		}
// 	}
// 	// stripExcessSpaces(headerItems)
// 	ctx.canonicalHeaders = strings.Join(headerItems, "\n")
// }

// const doubleSpace = "  "

// // stripExcessSpaces will rewrite the passed in slice's string values to not
// // contain multiple side-by-side spaces.
// func stripExcessSpaces(vals []string) {
// 	var j, k, l, m, spaces int
// 	for i, str := range vals {
// 		// Trim trailing spaces
// 		for j = len(str) - 1; j >= 0 && str[j] == ' '; j-- {
// 		}

// 		// Trim leading spaces
// 		for k = 0; k < j && str[k] == ' '; k++ {
// 		}
// 		str = str[k : j+1]

// 		// Strip multiple spaces.
// 		j = strings.Index(str, doubleSpace)
// 		if j < 0 {
// 			vals[i] = str
// 			continue
// 		}

// 		buf := []byte(str)
// 		for k, m, l = j, j, len(buf); k < l; k++ {
// 			if buf[k] == ' ' {
// 				if spaces == 0 {
// 					// First space.
// 					buf[m] = buf[k]
// 					m++
// 				}
// 				spaces++
// 			} else {
// 				// End of multiple spaces.
// 				spaces = 0
// 				buf[m] = buf[k]
// 				m++
// 			}
// 		}

// 		vals[i] = string(buf[:m])
// 	}
// }

func getURIPath(u *url.URL) string {
	var uri string

	if len(u.Opaque) > 0 {
		uri = "/" + strings.Join(strings.Split(u.Opaque, "/")[3:], "/")
	} else {
		uri = u.EscapedPath()
	}

	if len(uri) == 0 {
		uri = "/"
	}

	return uri
}

func vendorDateHeaderKey(vendorCode string) string {
	return vendorHeaderCanonicalNameKey(vendorCode, "Date")
}

func vendorRegionHeaderKey(vendorCode string) string {
	return vendorHeaderCanonicalNameKey(vendorCode, "Region")
}

func vendorHeaderCanonicalNameKey(vendorCode, header string) string {
	return http.CanonicalHeaderKey("x-" + vendorCode + "-" + header)
}
