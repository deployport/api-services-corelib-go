package signingv1

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	sdk "go.deployport.com/specular-runtime/client"
	"go.deployport.com/specular-runtime/server"
)

type signatureVerifier struct {
	req         *http.Request
	fetcher     CredentialFetcher
	logger      *slog.Logger
	region      string
	op          *sdk.Operation
	serviceName string
	vendorCode  string
	serverTime  func() time.Time
	Creds       *Credentials
}

// NewSignatureVerifier creates a new signature verifier
func NewSignatureVerifier(
	logger *slog.Logger,
	req *http.Request,
	fetcher CredentialFetcher,
	region string,
	op *sdk.Operation,
	serviceName string,
	vendorCode string,
	serverTime func() time.Time,
) *signatureVerifier {
	return &signatureVerifier{
		req:         req,
		fetcher:     fetcher,
		region:      region,
		logger:      logger,
		op:          op,
		serviceName: serviceName,
		vendorCode:  vendorCode,
		serverTime:  serverTime,
	}
}

func (v *signatureVerifier) createAuthError(msg string) *server.AuthenticationError {
	return &server.AuthenticationError{
		Message: "access denied, " + msg,
	}
}

func (v *signatureVerifier) Verify() error {
	serverDateUTC := v.serverTime().UTC()
	headers := v.req.Header
	v.logger.Debug("verifying signature", "headers", headers)
	authHeaderValues := strings.Split(headers.Get(AuthorizationHeader), ", ")
	lenAuthHeaderValues := len(authHeaderValues)
	v.logger.Debug("verifying signature", "lenAuthHeaderValues", lenAuthHeaderValues, "authHeaderValues", authHeaderValues)
	if lenAuthHeaderValues != 3 {
		return v.createAuthError("signature required")
	}
	credentialPart := authHeaderValues[0]
	expectedAuthHeaderPrefix := vendorAuthHeaderPrefix(v.vendorCode) + " "
	v.logger.Debug("verifying signature", "expectedAuthHeader", expectedAuthHeaderPrefix)
	if !strings.HasPrefix(credentialPart, expectedAuthHeaderPrefix) {
		return v.createAuthError("signature mismatch")
	}
	headerWithoutPrefix := credentialPart[len(expectedAuthHeaderPrefix):]
	v.logger.Debug("verifying signature", "headerWithoutPrefix", headerWithoutPrefix)
	if !strings.HasPrefix(headerWithoutPrefix, credentialPartPrefix) {
		return v.createAuthError("signature mismatch")
	}

	headerWithoutPrefix = strings.TrimPrefix(headerWithoutPrefix, credentialPartPrefix)

	credentialParts := strings.Split(headerWithoutPrefix, "/")
	if len(credentialParts) != 5 {
		return v.createAuthError("signature mismatch")
	}
	keyID := credentialParts[credentialPartKeyIDIndex]
	v.logger.Debug("verifying signature", "keyID", keyID)
	creds, err := v.fetcher(keyID)
	if err != nil {
		return err
	}
	if creds == nil {
		return v.createAuthError("unknown or expired access key id")
	}
	v.Creds = creds
	keyDateValue := credentialParts[credentialPartDateIndex]
	v.logger.Debug("verifying signature", "keyDateValue", keyDateValue)
	if keyDateValue != formatShortTime(serverDateUTC) {
		return v.createAuthError("date mismatch")
	}
	region := credentialParts[credentialPartRegionIndex]
	v.logger.Debug("verifying signature", "region", region)
	if region != v.region {
		return v.createAuthError("invalid region")
	}
	service := credentialParts[credentialPartServiceIndex]
	v.logger.Debug("verifying signature", "service", service)
	if service != v.serviceName {
		return v.createAuthError("service mismatch")
	}
	requestCode := credentialParts[credentialPartRequestCodeIndex]
	v.logger.Debug("verifying signature", "requestCode", requestCode)
	if requestCode != vendorRequestCode(v.vendorCode) {
		return v.createAuthError("invalid request code")
	}

	signedHeadersPart := authHeaderValues[1]
	v.logger.Debug("verifying signature", "signedHeadersPart", signedHeadersPart)
	if !strings.HasPrefix(signedHeadersPart, signedHeadersPartPrefix) {
		return v.createAuthError("invalid signed headers")
	}
	// split signedHeadersPart by ;
	signedHeaders := strings.Split(signedHeadersPart[len(signedHeadersPartPrefix):], ";")
	v.logger.Debug("verifying signature", "signedHeaders", signedHeaders)

	headerValues := http.Header{}
	for _, k := range signedHeaders {
		headerValues[http.CanonicalHeaderKey(k)] = headers.Values(k)
	}
	v.logger.Debug("verifying signature", "headerValues", headerValues)
	dateHeaderString := headerValues.Get(vendorDateHeaderKey(v.vendorCode))
	dateHeaderTime, err := time.Parse(timeFormat, dateHeaderString)
	if err != nil {
		return v.createAuthError("invalid date header")
	}

	// check request is same day as server
	if !dateOnlyEqual(dateHeaderTime, serverDateUTC) {
		return v.createAuthError("date mismatch")
	}
	bodyBytes, err := io.ReadAll(v.req.Body)
	if err != nil {
		return err
	}
	v.req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	signaturePart := authHeaderValues[2]
	receivedSignatureDigest := signaturePart[len(signaturePartPrefix):]
	v.logger.Debug("verifying signature from line", "digest", receivedSignatureDigest)

	computedSignature, err := SignHeaders(
		v.logger,
		v.vendorCode,
		v.serviceName,
		*creds,
		v.region,
		v.req,
		bodyBytes,
		v.op,
		dateHeaderTime)
	if err != nil {
		return err
	}
	v.logger.Debug("computed signature", "signature", v.req.Header.Values(AuthorizationHeader))
	// compare signature value at last part of the
	if receivedSignatureDigest != computedSignature.Digest {
		return v.createAuthError("invalid signature")
	}

	return nil
}

// CredentialFetcher is a function that returns a credential for a given keyID
type CredentialFetcher func(keyID string) (*Credentials, error)
