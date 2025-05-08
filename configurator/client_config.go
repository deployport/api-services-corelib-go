package configurator

import (
	"context"
	"net/http"
	"strings"
	"time"

	"go.deployport.com/api-services-corelib/configurator/signingv1"
	sdk "go.deployport.com/specular-runtime/client"
)

// NewOnSubmissionForCredentials returns a OnSubmissionHandler that signs the request with the credentials
func NewOnSubmissionForCredentials(
	creds signingv1.Credentials,
	region string,
) sdk.OnSubmissionHandler {
	return func(ctx context.Context, sub *sdk.Submission) error {
		ReplaceRegionInRequest(region, sub.HTTPRequest)
		sa := FindSignedOperationV1Annotation(sub.OperationRequest.Operation.Annotations())
		if sa == nil {
			return nil
		}
		sig := FindSignedServiceSignatureAnnotation(sub.OperationRequest.Operation.Resource().Package().Annotations())
		if sig == nil {
			return nil
		}
		logger := sub.Logger
		logger.Debug("on submission signing credentials", "operation", sub.OperationRequest.Operation.Name())
		_, err := signingv1.SignHeaders(
			logger,
			VendorCode,
			sig.ServiceName,
			creds,
			region,
			sub.HTTPRequest,
			sub.HTTPBody,
			sub.OperationRequest.Operation,
			time.Now().UTC(),
		)
		if err != nil {
			return err
		}
		return nil
	}
}

// NewOnSubmissionForRegionOnly returns a OnSubmissionHandler that sets the region in the request
func NewOnSubmissionForRegionOnly(
	region string,
) sdk.OnSubmissionHandler {
	return func(ctx context.Context, sub *sdk.Submission) error {
		// replace host <region> placeholder with the region
		// in the request URL
		ReplaceRegionInRequest(region, sub.HTTPRequest)
		return nil
	}
}

// ReplaceRegionInRequest replaces the <region> placeholder in the host with the region
func ReplaceRegionInRequest(region string, req *http.Request) {
	// replace host <region> placeholder with the region
	// in the request URL
	req.Host = strings.ReplaceAll(req.Host, "<region>", region)
	hostHeader := req.Header.Get("Host")
	if hostHeader != "" {
		req.Header.Set("Host", strings.ReplaceAll(hostHeader, "<region>", region))
	}
	req.URL.Host = strings.ReplaceAll(req.URL.Host, "<region>", region)
}
