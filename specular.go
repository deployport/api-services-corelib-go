package corelib

import (
	"errors"

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

// Hydrate implements struct hydrate
func (e *AccessDeniedProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccessDeniedProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *AccessDeniedProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccessDeniedProblem.Path()
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

// Hydrate implements struct hydrate
func (e *ForbiddenProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *ForbiddenProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *ForbiddenProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathForbiddenProblem.Path()
}

// SignedOperationV1 entity
// Marks an operation with digital signature authentication for id4ntity verification
type SignedOperationV1 struct {
}

// ServiceSignatureV1 entity
type ServiceSignatureV1 struct {
	ServiceName string
}

var packagePath = clientruntime.ModulePathFromTrustedValues(
	"deployport",
	"corelib",
)

func newSpecularPackage() (pk *clientruntime.Package, err error) {
	pk = clientruntime.NewPackage(packagePath)
	localSpecularMeta.structPathAccessDeniedProblem, err = pk.NewType(
		"AccessDeniedProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccessDeniedProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathForbiddenProblem, err = pk.NewType(
		"ForbiddenProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewForbiddenProblem()
		}),
	)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func init() {
	initSpecularMeta()
}

func initSpecularMeta() {
	pk, err := newSpecularPackage()
	if err != nil {
		panic(errors.New("failed to initialize shared allow package deployport/corelib"))
	}
	localSpecularMeta.mod = pk
}

// SpecularMetaInfo defines metadata of the specular module
type SpecularMetaInfo struct {
	mod                           *clientruntime.Package
	structPathAccessDeniedProblem *clientruntime.StructDefinition
	structPathForbiddenProblem    *clientruntime.StructDefinition
}

// Module returns the module definition
func (m *SpecularMetaInfo) Module() *clientruntime.Package {
	return m.mod
}

// AccessDeniedProblem allows easy access to structure
func (m *SpecularMetaInfo) AccessDeniedProblem() *clientruntime.StructDefinition {
	return m.structPathAccessDeniedProblem
}

// ForbiddenProblem allows easy access to structure
func (m *SpecularMetaInfo) ForbiddenProblem() *clientruntime.StructDefinition {
	return m.structPathForbiddenProblem
}

var localSpecularMeta *SpecularMetaInfo = &SpecularMetaInfo{}

// SpecularMeta returns metadata of the specular module
func SpecularMeta() *SpecularMetaInfo {
	return localSpecularMeta
}
