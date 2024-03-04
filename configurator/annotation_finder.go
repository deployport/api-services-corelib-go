package configurator

import (
	corelib "go.deployport.com/api-services-corelib"
	sdk "go.deployport.com/specular-runtime/client"
)

// FindSignedOperationV1Annotation returns the first SignedOperationV1 annotation in the list
func FindSignedOperationV1Annotation(list []sdk.Annotation) *corelib.SignedOperationV1 {
	for _, a := range list {
		if sa, ok := a.(*corelib.SignedOperationV1); ok {
			return sa
		}
	}
	return nil
}

// FindSignedServiceSignatureAnnotation returns the first ServiceSignatureV1 annotation in the list
func FindSignedServiceSignatureAnnotation(list []sdk.Annotation) *corelib.ServiceSignatureV1 {
	for _, a := range list {
		if sa, ok := a.(*corelib.ServiceSignatureV1); ok {
			return sa
		}
	}
	return nil
}
