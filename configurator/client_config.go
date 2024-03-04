package configurator

import (
	"context"
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
