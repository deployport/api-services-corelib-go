package corelib
import clientruntime "go.deployport.com/specular-runtime/client"
import "sync"
import "errors"

// SignedOperationV1 entity
// Marks an operation with digital signature authentication for id4ntity verification
type SignedOperationV1 struct {
}

// TypeFQTN returns the Allow Typq Fully Qualified Type Name
func (tp *SignedOperationV1) TypeFQTN() clientruntime.TypeFQTN {
    return clientruntime.NewTypeFQTN("Deployport/CoreLib", "SignedOperationV1")
}

// ServiceSignatureV1 entity
type ServiceSignatureV1 struct {
    ServiceName string
}

// TypeFQTN returns the Allow Typq Fully Qualified Type Name
func (tp *ServiceSignatureV1) TypeFQTN() clientruntime.TypeFQTN {
    return clientruntime.NewTypeFQTN("Deployport/CoreLib", "ServiceSignatureV1")
}

// NewSpecularPackage returns a new package instance for Deployport/CoreLib
func NewSpecularPackage() (*clientruntime.Package, error) {
    pk := clientruntime.NewPackage(
        "Deployport/CoreLib",
    )

    return pk, nil
}

var specularPackageOnce = sync.OnceValue(func() *clientruntime.Package {
    pk, err := NewSpecularPackage()
    if err != nil {
        panic(errors.New("failed to initialize shared allow package Deployport/CoreLib"))
    }
    return pk
})

// SpecularPackage returns a shared package instance for Deployport/CoreLib, panics when errors are found
func SpecularPackage() (*clientruntime.Package) {
    return specularPackageOnce()
}
