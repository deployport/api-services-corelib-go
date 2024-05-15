package corelib

import (
	"errors"
	"sync"

	clientruntime "go.deployport.com/specular-runtime/client"
)

// NewAccessDeniedProblem creates a new AccessDeniedProblem
func NewAccessDeniedProblem() *AccessDeniedProblem {
	s := &AccessDeniedProblem{}
	return s
}

// AccessDeniedProblem entity
type AccessDeniedProblem struct {
	Message string
}

// Error implements the error interface
func (e *AccessDeniedProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccessDeniedProblem]
func (e *AccessDeniedProblem) Is(err error) bool {
	_, ok := err.(*AccessDeniedProblem)
	return ok
}

// IsAccessDeniedProblem indicates whether the given error chain contains an error of type [AccessDeniedProblem]
func IsAccessDeniedProblem(err error) bool {
	return errors.Is(err, &AccessDeniedProblem{})
}

// GetMessage returns the value for the field message
func (e *AccessDeniedProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccessDeniedProblem) SetMessage(message string) {
	e.Message = message
}

func (e *AccessDeniedProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

func (e *AccessDeniedProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetStruct(e.TypeFQTN().String())
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// TypeFQTN returns the Allow Typq Fully Qualified Type Name
func (e *AccessDeniedProblem) TypeFQTN() clientruntime.TypeFQTN {
	return clientruntime.NewTypeFQTN("Deployport/CoreLib", "AccessDeniedProblem")
}

// NewForbiddenProblem creates a new ForbiddenProblem
func NewForbiddenProblem() *ForbiddenProblem {
	s := &ForbiddenProblem{}
	return s
}

// ForbiddenProblem entity
type ForbiddenProblem struct {
	Message string
}

// Error implements the error interface
func (e *ForbiddenProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [ForbiddenProblem]
func (e *ForbiddenProblem) Is(err error) bool {
	_, ok := err.(*ForbiddenProblem)
	return ok
}

// IsForbiddenProblem indicates whether the given error chain contains an error of type [ForbiddenProblem]
func IsForbiddenProblem(err error) bool {
	return errors.Is(err, &ForbiddenProblem{})
}

// GetMessage returns the value for the field message
func (e *ForbiddenProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *ForbiddenProblem) SetMessage(message string) {
	e.Message = message
}

func (e *ForbiddenProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

func (e *ForbiddenProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetStruct(e.TypeFQTN().String())
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// TypeFQTN returns the Allow Typq Fully Qualified Type Name
func (e *ForbiddenProblem) TypeFQTN() clientruntime.TypeFQTN {
	return clientruntime.NewTypeFQTN("Deployport/CoreLib", "ForbiddenProblem")
}

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
	if _, err := pk.NewType(
		"AccessDeniedProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccessDeniedProblem()
		}),
	); err != nil {
		return nil, err
	}
	if _, err := pk.NewType(
		"ForbiddenProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewForbiddenProblem()
		}),
	); err != nil {
		return nil, err
	}

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
func SpecularPackage() *clientruntime.Package {
	return specularPackageOnce()
}
