package corelib

import (
	"encoding/json"
	"errors"

	clientruntime "go.deployport.com/specular-runtime/client"
)

// NewAccessDeniedProblem creates a new AccessDeniedProblem
func NewAccessDeniedProblem() *AccessDeniedProblem {
	s := &AccessDeniedProblem{}
	s.InitializeDefaults()
	return s
}

// AccessDeniedProblem struct
type AccessDeniedProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
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

// StructPath returns StructPath
func (e *AccessDeniedProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccessDeniedProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccessDeniedProblem) InitializeDefaults() {
}

// accessDeniedProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type accessDeniedProblemAlias AccessDeniedProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *AccessDeniedProblem) UnmarshalJSON(data []byte) error {
	var alias accessDeniedProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccessDeniedProblem)(&alias)).InitializeDefaults()
	*e = AccessDeniedProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccessDeniedProblem) MarshalJSON() ([]byte, error) {
	alias := accessDeniedProblemAlias(e)
	return json.Marshal(alias)
}

// NewForbiddenProblem creates a new ForbiddenProblem
func NewForbiddenProblem() *ForbiddenProblem {
	s := &ForbiddenProblem{}
	s.InitializeDefaults()
	return s
}

// ForbiddenProblem struct
type ForbiddenProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
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

// StructPath returns StructPath
func (e *ForbiddenProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathForbiddenProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ForbiddenProblem) InitializeDefaults() {
}

// forbiddenProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type forbiddenProblemAlias ForbiddenProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *ForbiddenProblem) UnmarshalJSON(data []byte) error {
	var alias forbiddenProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ForbiddenProblem)(&alias)).InitializeDefaults()
	*e = ForbiddenProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ForbiddenProblem) MarshalJSON() ([]byte, error) {
	alias := forbiddenProblemAlias(e)
	return json.Marshal(alias)
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

// AccessDeniedProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) AccessDeniedProblemStruct() *clientruntime.StructDefinition {
	return m.structPathAccessDeniedProblem
}

// ForbiddenProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) ForbiddenProblemStruct() *clientruntime.StructDefinition {
	return m.structPathForbiddenProblem
}

var localSpecularMeta *SpecularMetaInfo = &SpecularMetaInfo{}

// SpecularMeta returns metadata of the specular module
func SpecularMeta() *SpecularMetaInfo {
	return localSpecularMeta
}
