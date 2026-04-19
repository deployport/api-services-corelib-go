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

// NewProblemAction creates a new ProblemAction
func NewProblemAction() *ProblemAction {
	s := &ProblemAction{}
	s.InitializeDefaults()
	return s
}

// ProblemAction - A localized call-to-action surfaced alongside a problem. Title and URL
// are already resolved to a specific locale; clients render them as-is.
type ProblemAction struct {
	// Optional stable identifier. Clients can key on it to apply custom
	// presentation (icon, ordering) per well-known action.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
	// BCP 47 tag of the locale the title and url were resolved to. May
	// differ from the caller's requested locale when a fallback was used.
	Locale *string `json:"locale,omitempty" yaml:"locale,omitempty"`
	Title  string  `json:"title,omitempty" yaml:"title,omitempty"`
	Url    *string `json:"url,omitempty" yaml:"url,omitempty"`
}

// GetId returns the value for the field id
func (e *ProblemAction) GetId() string {
	return e.Id
}

// SetId sets the value for the field id
func (e *ProblemAction) SetId(id string) {
	e.Id = id
}

// GetLocale returns the value for the field locale
func (e *ProblemAction) GetLocale() *string {
	return e.Locale
}

// SetLocale sets the value for the field locale
func (e *ProblemAction) SetLocale(locale *string) {
	e.Locale = locale
}

// GetTitle returns the value for the field title
func (e *ProblemAction) GetTitle() string {
	return e.Title
}

// SetTitle sets the value for the field title
func (e *ProblemAction) SetTitle(title string) {
	e.Title = title
}

// GetUrl returns the value for the field url
func (e *ProblemAction) GetUrl() *string {
	return e.Url
}

// SetUrl sets the value for the field url
func (e *ProblemAction) SetUrl(url *string) {
	e.Url = url
}

// StructPath returns StructPath
func (e *ProblemAction) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathProblemAction.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ProblemAction) InitializeDefaults() {
}

// problemActionAlias is defined to help pre and post JSON marshaling without recursive loops
type problemActionAlias ProblemAction

// UnmarshalJSON implements json.Unmarshaler
func (e *ProblemAction) UnmarshalJSON(data []byte) error {
	var alias problemActionAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ProblemAction)(&alias)).InitializeDefaults()
	*e = ProblemAction(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ProblemAction) MarshalJSON() ([]byte, error) {
	alias := problemActionAlias(e)
	return json.Marshal(alias)
}

// NewThrottledProblem creates a new ThrottledProblem
func NewThrottledProblem() *ThrottledProblem {
	s := &ThrottledProblem{}
	s.InitializeDefaults()
	return s
}

// ThrottledProblem - Raised when the caller is being throttled and should retry later.
type ThrottledProblem struct {
	// Optional calls-to-action to help the caller (e.g. support, upgrade,
	// docs links), ordered most-specific first. Empty when none apply.
	Actions []*ProblemAction `json:"actions,omitempty" yaml:"actions,omitempty"`
	Message string           `json:"message,omitempty" yaml:"message,omitempty"`
	// Milliseconds the caller should wait before retrying.
	RetryAfterMs *int64 `json:"retryAfterMs,omitempty" yaml:"retryAfterMs,omitempty"`
}

// Error implements the error interface
func (e *ThrottledProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [ThrottledProblem]
func (e *ThrottledProblem) Is(err error) bool {
	_, ok := err.(*ThrottledProblem)
	return ok
}

// IsThrottledProblem indicates whether the given error chain contains an error of type [ThrottledProblem]
func IsThrottledProblem(err error) bool {
	return errors.Is(err, &ThrottledProblem{})
}

// GetActions returns the value for the field actions
func (e *ThrottledProblem) GetActions() []*ProblemAction {
	return e.Actions
}

// SetActions sets the value for the field actions
func (e *ThrottledProblem) SetActions(actions []*ProblemAction) {
	e.Actions = actions
}

// GetMessage returns the value for the field message
func (e *ThrottledProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *ThrottledProblem) SetMessage(message string) {
	e.Message = message
}

// GetRetryAfterMs returns the value for the field retryAfterMs
func (e *ThrottledProblem) GetRetryAfterMs() *int64 {
	return e.RetryAfterMs
}

// SetRetryAfterMs sets the value for the field retryAfterMs
func (e *ThrottledProblem) SetRetryAfterMs(retryAfterMs *int64) {
	e.RetryAfterMs = retryAfterMs
}

// StructPath returns StructPath
func (e *ThrottledProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathThrottledProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ThrottledProblem) InitializeDefaults() {
}

// throttledProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type throttledProblemAlias ThrottledProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *ThrottledProblem) UnmarshalJSON(data []byte) error {
	var alias throttledProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ThrottledProblem)(&alias)).InitializeDefaults()
	*e = ThrottledProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ThrottledProblem) MarshalJSON() ([]byte, error) {
	alias := throttledProblemAlias(e)
	return json.Marshal(alias)
}

// NewQuotaExceededProblem creates a new QuotaExceededProblem
func NewQuotaExceededProblem() *QuotaExceededProblem {
	s := &QuotaExceededProblem{}
	s.InitializeDefaults()
	return s
}

// QuotaExceededProblem - Raised when a resource quota has been exceeded.
type QuotaExceededProblem struct {
	// Optional calls-to-action to help the caller (e.g. request an increase,
	// upgrade, docs links), ordered most-specific first. Empty when none apply.
	Actions []*ProblemAction `json:"actions,omitempty" yaml:"actions,omitempty"`
	Message string           `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *QuotaExceededProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [QuotaExceededProblem]
func (e *QuotaExceededProblem) Is(err error) bool {
	_, ok := err.(*QuotaExceededProblem)
	return ok
}

// IsQuotaExceededProblem indicates whether the given error chain contains an error of type [QuotaExceededProblem]
func IsQuotaExceededProblem(err error) bool {
	return errors.Is(err, &QuotaExceededProblem{})
}

// GetActions returns the value for the field actions
func (e *QuotaExceededProblem) GetActions() []*ProblemAction {
	return e.Actions
}

// SetActions sets the value for the field actions
func (e *QuotaExceededProblem) SetActions(actions []*ProblemAction) {
	e.Actions = actions
}

// GetMessage returns the value for the field message
func (e *QuotaExceededProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *QuotaExceededProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *QuotaExceededProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathQuotaExceededProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *QuotaExceededProblem) InitializeDefaults() {
}

// quotaExceededProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type quotaExceededProblemAlias QuotaExceededProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *QuotaExceededProblem) UnmarshalJSON(data []byte) error {
	var alias quotaExceededProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*QuotaExceededProblem)(&alias)).InitializeDefaults()
	*e = QuotaExceededProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e QuotaExceededProblem) MarshalJSON() ([]byte, error) {
	alias := quotaExceededProblemAlias(e)
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
	localSpecularMeta.structPathProblemAction, err = pk.NewType(
		"ProblemAction",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewProblemAction()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathThrottledProblem, err = pk.NewType(
		"ThrottledProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewThrottledProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathQuotaExceededProblem, err = pk.NewType(
		"QuotaExceededProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewQuotaExceededProblem()
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
	mod                            *clientruntime.Package
	structPathAccessDeniedProblem  *clientruntime.StructDefinition
	structPathForbiddenProblem     *clientruntime.StructDefinition
	structPathProblemAction        *clientruntime.StructDefinition
	structPathThrottledProblem     *clientruntime.StructDefinition
	structPathQuotaExceededProblem *clientruntime.StructDefinition
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

// ProblemActionStruct allows easy access to structure
func (m *SpecularMetaInfo) ProblemActionStruct() *clientruntime.StructDefinition {
	return m.structPathProblemAction
}

// ThrottledProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) ThrottledProblemStruct() *clientruntime.StructDefinition {
	return m.structPathThrottledProblem
}

// QuotaExceededProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) QuotaExceededProblemStruct() *clientruntime.StructDefinition {
	return m.structPathQuotaExceededProblem
}

var localSpecularMeta *SpecularMetaInfo = &SpecularMetaInfo{}

// SpecularMeta returns metadata of the specular module
func SpecularMeta() *SpecularMetaInfo {
	return localSpecularMeta
}
