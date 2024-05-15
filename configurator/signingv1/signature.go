package signingv1

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	sdk "go.deployport.com/specular-runtime/client"
)

// Credentials is the credentials for signing
type Credentials struct {
	KeyID  string
	Secret string
}

// Signature is the signature for the request
type Signature struct {
	Digest string
}

// SignHeaders generates the Authorization header value for the request
func SignHeaders(
	logger *slog.Logger,
	vendorCode string,
	serviceName string,
	creds Credentials,
	region string,
	req *http.Request,
	body []byte,
	op *sdk.Operation,
	time time.Time,
) (*Signature, error) {
	sr := signerRequest{
		logger:     logger,
		creds:      creds,
		req:        req,
		region:     region,
		time:       time,
		op:         op,
		vendorCode: vendorCode,
		service:    serviceName,
		body:       body,
	}
	err := sr.generate()
	if err != nil {
		return nil, err
	}
	return &Signature{
		Digest: sr.signature,
	}, nil
}

type signerRequest struct {
	logger       *slog.Logger
	body         []byte
	creds        Credentials
	req          *http.Request
	stringToSign string
	region       string
	time         time.Time
	op           *sdk.Operation
	vendorCode   string
	service      string
	// host             string
	url              *url.URL
	authHeaderPrefix string
	credentialString string
	signature        string
	headerBuilder    headerSignatureBuilder
}

func (r *signerRequest) generate() error {
	logger := r.logger
	logger.Debug("signing request", "vendorCode", r.vendorCode, "service", r.service, "region", r.region, "time", r.time)
	r.credentialString = buildSigningScope(r.vendorCode, r.region, r.service, r.time)
	logger.Debug("signing request", "signing scope", r.credentialString)
	credentialPart := buildCredentialPart(
		r.creds.KeyID,
		r.credentialString,
	)
	logger.Debug("signing request", "credential part", credentialPart)
	r.authHeaderPrefix = vendorAuthHeaderPrefix(r.vendorCode) + " " + credentialPart

	bodyDigest := hex.EncodeToString(hashSHA256(r.body))
	logger.Debug("body digest", "digest", bodyDigest)

	r.req.Body = io.NopCloser(bytes.NewReader(r.body))

	r.req.Header.Set(vendorDateHeaderKey(r.vendorCode), formatTime(r.time))
	r.req.Header.Set(vendorHeaderCanonicalNameKey(r.vendorCode, "Service"), r.service)
	r.req.Header.Set(vendorHeaderCanonicalNameKey(r.vendorCode, "Resource"), r.op.Resource().PackageUniqueName())
	r.req.Header.Set(vendorHeaderCanonicalNameKey(r.vendorCode, "Operation"), r.op.Name())
	r.req.Header.Set(vendorRegionHeaderKey(r.vendorCode), r.region)
	r.req.Header.Set(vendorHeaderCanonicalNameKey(r.vendorCode, "Content-Sha256"), bodyDigest)

	r.headerBuilder = headerSignatureBuilder{
		logger:     logger,
		Request:    r.req,
		Body:       r.body,
		BodyDigest: bodyDigest,
		VendorCode: r.vendorCode,
	}
	r.headerBuilder.build()
	r.buildStringToSign()
	logger.Debug("signing request", "stringToSign", r.stringToSign)
	r.buildSignature()
	logger.Debug("signing request", "ordered header names", r.headerBuilder.signedHeaderNames)
	parts := []string{
		r.authHeaderPrefix,
		signedHeadersPartPrefix + r.headerBuilder.signedHeaderNames,
		signaturePartPrefix + r.signature,
	}
	r.req.Header.Set(AuthorizationHeader, strings.Join(parts, ", "))
	return nil
}

func (r *signerRequest) buildStringToSign() {
	r.logger.Debug("building string to sign", "canonicalString", r.headerBuilder.canonicalString)
	r.stringToSign = strings.Join([]string{
		vendorAuthHeaderPrefix(r.vendorCode),
		formatTime(r.time),
		r.credentialString,
		hex.EncodeToString(hashSHA256([]byte(r.headerBuilder.canonicalString))),
	}, "\n")
}

const signedHeadersPartPrefix = "SignedHeaders="

const signaturePartPrefix = "Signature="

func (r *signerRequest) buildSignature() {
	creds := r.deriveSigningKey(r.region, r.service, r.creds.Secret, r.time)
	signature := hmacSHA256(creds, []byte(r.stringToSign))
	r.signature = hex.EncodeToString(signature)
}

func (r *signerRequest) deriveSigningKey(region, service, secretKey string, dt time.Time) []byte {
	kDate := hmacSHA256([]byte(r.vendorCode+"1"+secretKey), []byte(formatShortTime(dt)))
	kRegion := hmacSHA256(kDate, []byte(region))
	kService := hmacSHA256(kRegion, []byte(service))
	signingKey := hmacSHA256(kService, []byte(vendorRequestCode(r.vendorCode)))
	return signingKey
}

const credentialPartPrefix = "Credential="

func buildCredentialPart(keyID, signingScope string) string {
	return fmt.Sprintf("%s%s/%s", credentialPartPrefix, keyID, signingScope)
}

func formatShortTime(dt time.Time) string {
	return dt.UTC().Format(shortTimeFormat)
}
func formatTime(dt time.Time) string {
	return dt.UTC().Format(timeFormat)
}

const (
	timeFormat                     = "20060102T150405Z"
	shortTimeFormat                = "20060102"
	credentialPartKeyIDIndex       = 0
	credentialPartDateIndex        = 1
	credentialPartRegionIndex      = 2
	credentialPartServiceIndex     = 3
	credentialPartRequestCodeIndex = 4
)

func buildSigningScope(vendorCode string, region, service string, dt time.Time) string {
	return strings.Join([]string{
		formatShortTime(dt),
		region,
		service,
		vendorRequestCode(vendorCode),
	}, "/")
}

func vendorRequestCode(vendorCode string) string {
	return vendorCode + "1_request"
}

func vendorAuthHeaderPrefix(vendorCode string) string {
	return strings.ToUpper(vendorCode) + "1-HMAC-SHA256"
}
