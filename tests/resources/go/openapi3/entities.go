package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

// OpenAPI structure is generated from "#".
//
// Validation schema for OpenAPI Specification 3.0.X.
type OpenAPI struct {
	// Value must match pattern: `^3\.0\.\d(-.+)?$`.
	// Required.
	Openapi       string                 `json:"openapi"`
	Info          Info                   `json:"info"`                   // Required.
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	Servers       []Server               `json:"servers,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	Tags          []Tag                  `json:"tags,omitempty"`
	Paths         Paths                  `json:"paths"`                  // Required.
	Components    *Components            `json:"components,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
}

// WithOpenapi sets Openapi value.
func (v *OpenAPI) WithOpenapi(val string) *OpenAPI {
	v.Openapi = val
	return v
}

// WithInfo sets Info value.
func (v *OpenAPI) WithInfo(val Info) *OpenAPI {
	v.Info = val
	return v
}

// WithExternalDocs sets ExternalDocs value.
func (v *OpenAPI) WithExternalDocs(val ExternalDocumentation) *OpenAPI {
	v.ExternalDocs = &val
	return v
}

// WithServers sets Servers value.
func (v *OpenAPI) WithServers(val ...Server) *OpenAPI {
	v.Servers = val
	return v
}

// WithSecurity sets Security value.
func (v *OpenAPI) WithSecurity(val ...map[string][]string) *OpenAPI {
	v.Security = val
	return v
}

// WithTags sets Tags value.
func (v *OpenAPI) WithTags(val ...Tag) *OpenAPI {
	v.Tags = val
	return v
}

// WithPaths sets Paths value.
func (v *OpenAPI) WithPaths(val Paths) *OpenAPI {
	v.Paths = val
	return v
}

// WithComponents sets Components value.
func (v *OpenAPI) WithComponents(val Components) *OpenAPI {
	v.Components = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *OpenAPI) WithMapOfAnything(val map[string]interface{}) *OpenAPI {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *OpenAPI) WithMapOfAnythingItem(key string, val interface{}) *OpenAPI {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalOpenAPI OpenAPI

var ignoreKeysOpenAPI = []string{
	"openapi",
	"info",
	"externalDocs",
	"servers",
	"security",
	"tags",
	"paths",
	"components",
}

// UnmarshalJSON decodes JSON.
func (v *OpenAPI) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOpenAPI(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOpenAPI {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = OpenAPI(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v OpenAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOpenAPI(v), v.MapOfAnything)
}

// Info structure is generated from "#/definitions/Info".
type Info struct {
	Title          string                 `json:"title"`                    // Required.
	Description    *string                `json:"description,omitempty"`
	TermsOfService *string                `json:"termsOfService,omitempty"` // Format: uri-reference.
	Contact        *Contact               `json:"contact,omitempty"`
	License        *License               `json:"license,omitempty"`
	Version        string                 `json:"version"`                  // Required.
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: `^x-`.
}

// WithTitle sets Title value.
func (v *Info) WithTitle(val string) *Info {
	v.Title = val
	return v
}

// WithDescription sets Description value.
func (v *Info) WithDescription(val string) *Info {
	v.Description = &val
	return v
}

// WithTermsOfService sets TermsOfService value.
func (v *Info) WithTermsOfService(val string) *Info {
	v.TermsOfService = &val
	return v
}

// WithContact sets Contact value.
func (v *Info) WithContact(val Contact) *Info {
	v.Contact = &val
	return v
}

// WithLicense sets License value.
func (v *Info) WithLicense(val License) *Info {
	v.License = &val
	return v
}

// WithVersion sets Version value.
func (v *Info) WithVersion(val string) *Info {
	v.Version = val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Info) WithMapOfAnything(val map[string]interface{}) *Info {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Info) WithMapOfAnythingItem(key string, val interface{}) *Info {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalInfo Info

var ignoreKeysInfo = []string{
	"title",
	"description",
	"termsOfService",
	"contact",
	"license",
	"version",
}

// UnmarshalJSON decodes JSON.
func (v *Info) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalInfo(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysInfo {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Info(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(v), v.MapOfAnything)
}

// Contact structure is generated from "#/definitions/Contact".
type Contact struct {
	Name          *string                `json:"name,omitempty"`
	URL           *string                `json:"url,omitempty"`   // Format: uri-reference.
	Email         *string                `json:"email,omitempty"` // Format: email.
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (v *Contact) WithName(val string) *Contact {
	v.Name = &val
	return v
}

// WithURL sets URL value.
func (v *Contact) WithURL(val string) *Contact {
	v.URL = &val
	return v
}

// WithEmail sets Email value.
func (v *Contact) WithEmail(val string) *Contact {
	v.Email = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Contact) WithMapOfAnything(val map[string]interface{}) *Contact {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Contact) WithMapOfAnythingItem(key string, val interface{}) *Contact {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalContact Contact

var ignoreKeysContact = []string{
	"name",
	"url",
	"email",
}

// UnmarshalJSON decodes JSON.
func (v *Contact) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalContact(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysContact {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Contact(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(v), v.MapOfAnything)
}

// License structure is generated from "#/definitions/License".
type License struct {
	Name          string                 `json:"name"`          // Required.
	URL           *string                `json:"url,omitempty"` // Format: uri-reference.
	MapOfAnything map[string]interface{} `json:"-"`             // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (v *License) WithName(val string) *License {
	v.Name = val
	return v
}

// WithURL sets URL value.
func (v *License) WithURL(val string) *License {
	v.URL = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *License) WithMapOfAnything(val map[string]interface{}) *License {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *License) WithMapOfAnythingItem(key string, val interface{}) *License {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalLicense License

var ignoreKeysLicense = []string{
	"name",
	"url",
}

// UnmarshalJSON decodes JSON.
func (v *License) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalLicense(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysLicense {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = License(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(v), v.MapOfAnything)
}

// ExternalDocumentation structure is generated from "#/definitions/ExternalDocumentation".
type ExternalDocumentation struct {
	Description   *string                `json:"description,omitempty"`
	// Format: uri-reference.
	// Required.
	URL           string                 `json:"url"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (v *ExternalDocumentation) WithDescription(val string) *ExternalDocumentation {
	v.Description = &val
	return v
}

// WithURL sets URL value.
func (v *ExternalDocumentation) WithURL(val string) *ExternalDocumentation {
	v.URL = val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *ExternalDocumentation) WithMapOfAnything(val map[string]interface{}) *ExternalDocumentation {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *ExternalDocumentation) WithMapOfAnythingItem(key string, val interface{}) *ExternalDocumentation {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalExternalDocumentation ExternalDocumentation

var ignoreKeysExternalDocumentation = []string{
	"description",
	"url",
}

// UnmarshalJSON decodes JSON.
func (v *ExternalDocumentation) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalExternalDocumentation(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysExternalDocumentation {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ExternalDocumentation(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ExternalDocumentation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocumentation(v), v.MapOfAnything)
}

// Server structure is generated from "#/definitions/Server".
type Server struct {
	URL           string                    `json:"url"`                   // Required.
	Description   *string                   `json:"description,omitempty"`
	Variables     map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnything map[string]interface{}    `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithURL sets URL value.
func (v *Server) WithURL(val string) *Server {
	v.URL = val
	return v
}

// WithDescription sets Description value.
func (v *Server) WithDescription(val string) *Server {
	v.Description = &val
	return v
}

// WithVariables sets Variables value.
func (v *Server) WithVariables(val map[string]ServerVariable) *Server {
	v.Variables = val
	return v
}

// WithVariablesItem sets Variables item value.
func (v *Server) WithVariablesItem(key string, val ServerVariable) *Server {
	if v.Variables == nil {
		v.Variables = make(map[string]ServerVariable, 1)
	}

	v.Variables[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Server) WithMapOfAnything(val map[string]interface{}) *Server {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Server) WithMapOfAnythingItem(key string, val interface{}) *Server {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalServer Server

var ignoreKeysServer = []string{
	"url",
	"description",
	"variables",
}

// UnmarshalJSON decodes JSON.
func (v *Server) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalServer(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysServer {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Server(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(v), v.MapOfAnything)
}

// ServerVariable structure is generated from "#/definitions/ServerVariable".
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default"`               // Required.
	Description   *string                `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithEnum sets Enum value.
func (v *ServerVariable) WithEnum(val ...string) *ServerVariable {
	v.Enum = val
	return v
}

// WithDefault sets Default value.
func (v *ServerVariable) WithDefault(val string) *ServerVariable {
	v.Default = val
	return v
}

// WithDescription sets Description value.
func (v *ServerVariable) WithDescription(val string) *ServerVariable {
	v.Description = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *ServerVariable) WithMapOfAnything(val map[string]interface{}) *ServerVariable {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *ServerVariable) WithMapOfAnythingItem(key string, val interface{}) *ServerVariable {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalServerVariable ServerVariable

var ignoreKeysServerVariable = []string{
	"enum",
	"default",
	"description",
}

// UnmarshalJSON decodes JSON.
func (v *ServerVariable) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalServerVariable(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysServerVariable {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ServerVariable(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(v), v.MapOfAnything)
}

// Tag structure is generated from "#/definitions/Tag".
type Tag struct {
	Name          string                 `json:"name"`                   // Required.
	Description   *string                `json:"description,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (v *Tag) WithName(val string) *Tag {
	v.Name = val
	return v
}

// WithDescription sets Description value.
func (v *Tag) WithDescription(val string) *Tag {
	v.Description = &val
	return v
}

// WithExternalDocs sets ExternalDocs value.
func (v *Tag) WithExternalDocs(val ExternalDocumentation) *Tag {
	v.ExternalDocs = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Tag) WithMapOfAnything(val map[string]interface{}) *Tag {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Tag) WithMapOfAnythingItem(key string, val interface{}) *Tag {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalTag Tag

var ignoreKeysTag = []string{
	"name",
	"description",
	"externalDocs",
}

// UnmarshalJSON decodes JSON.
func (v *Tag) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalTag(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysTag {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Tag(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(v), v.MapOfAnything)
}

// PathItem structure is generated from "#/definitions/PathItem".
type PathItem struct {
	Ref                  *string                `json:"$ref,omitempty"`
	Summary              *string                `json:"summary,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Servers              []Server               `json:"servers,omitempty"`
	Parameters           []ParameterOrRef       `json:"parameters,omitempty"`
	MapOfOperationValues map[string]Operation   `json:"-"`                     // Key must match pattern: `^(get|put|post|delete|options|head|patch|trace)$`.
	MapOfAnything        map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithRef sets Ref value.
func (v *PathItem) WithRef(val string) *PathItem {
	v.Ref = &val
	return v
}

// WithSummary sets Summary value.
func (v *PathItem) WithSummary(val string) *PathItem {
	v.Summary = &val
	return v
}

// WithDescription sets Description value.
func (v *PathItem) WithDescription(val string) *PathItem {
	v.Description = &val
	return v
}

// WithServers sets Servers value.
func (v *PathItem) WithServers(val ...Server) *PathItem {
	v.Servers = val
	return v
}

// WithParameters sets Parameters value.
func (v *PathItem) WithParameters(val ...ParameterOrRef) *PathItem {
	v.Parameters = val
	return v
}

// WithMapOfOperationValues sets MapOfOperationValues value.
func (v *PathItem) WithMapOfOperationValues(val map[string]Operation) *PathItem {
	v.MapOfOperationValues = val
	return v
}

// WithMapOfOperationValuesItem sets MapOfOperationValues item value.
func (v *PathItem) WithMapOfOperationValuesItem(key string, val Operation) *PathItem {
	if v.MapOfOperationValues == nil {
		v.MapOfOperationValues = make(map[string]Operation, 1)
	}

	v.MapOfOperationValues[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *PathItem) WithMapOfAnything(val map[string]interface{}) *PathItem {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *PathItem) WithMapOfAnythingItem(key string, val interface{}) *PathItem {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalPathItem PathItem

var ignoreKeysPathItem = []string{
	"$ref",
	"summary",
	"description",
	"servers",
	"parameters",
}

// UnmarshalJSON decodes JSON.
func (v *PathItem) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalPathItem(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysPathItem {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexGetPutPostDeleteOptionsHeadPatchTrace.MatchString(key) {
			matched = true

			if vv.MapOfOperationValues == nil {
				vv.MapOfOperationValues = make(map[string]Operation, 1)
			}

			var val Operation

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfOperationValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = PathItem(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v PathItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPathItem(v), v.MapOfOperationValues, v.MapOfAnything)
}

// ParameterReference structure is generated from "#/definitions/ParameterReference".
type ParameterReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *ParameterReference) WithRef(val string) *ParameterReference {
	v.Ref = val
	return v
}

// Parameter structure is generated from "#/definitions/Parameter".
type Parameter struct {
	Name             string                  `json:"name"`                      // Required.
	In               ParameterIn             `json:"in"`                        // Required.
	Description      *string                 `json:"description,omitempty"`
	Required         *bool                   `json:"required,omitempty"`
	Deprecated       *bool                   `json:"deprecated,omitempty"`
	AllowEmptyValue  *bool                   `json:"allowEmptyValue,omitempty"`
	Style            *string                 `json:"style,omitempty"`
	Explode          *bool                   `json:"explode,omitempty"`
	AllowReserved    *bool                   `json:"allowReserved,omitempty"`
	Schema           *SchemaOrRef            `json:"schema,omitempty"`
	Content          map[string]MediaType    `json:"content,omitempty"`
	Example          *interface{}            `json:"example,omitempty"`
	Examples         map[string]ExampleOrRef `json:"examples,omitempty"`
	SchemaXORContent *SchemaXORContentOneOf1 `json:"-"`
	Location         *ParameterLocation      `json:"-"`
	MapOfAnything    map[string]interface{}  `json:"-"`                         // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (v *Parameter) WithName(val string) *Parameter {
	v.Name = val
	return v
}

// WithIn sets In value.
func (v *Parameter) WithIn(val ParameterIn) *Parameter {
	v.In = val
	return v
}

// WithDescription sets Description value.
func (v *Parameter) WithDescription(val string) *Parameter {
	v.Description = &val
	return v
}

// WithRequired sets Required value.
func (v *Parameter) WithRequired(val bool) *Parameter {
	v.Required = &val
	return v
}

// WithDeprecated sets Deprecated value.
func (v *Parameter) WithDeprecated(val bool) *Parameter {
	v.Deprecated = &val
	return v
}

// WithAllowEmptyValue sets AllowEmptyValue value.
func (v *Parameter) WithAllowEmptyValue(val bool) *Parameter {
	v.AllowEmptyValue = &val
	return v
}

// WithStyle sets Style value.
func (v *Parameter) WithStyle(val string) *Parameter {
	v.Style = &val
	return v
}

// WithExplode sets Explode value.
func (v *Parameter) WithExplode(val bool) *Parameter {
	v.Explode = &val
	return v
}

// WithAllowReserved sets AllowReserved value.
func (v *Parameter) WithAllowReserved(val bool) *Parameter {
	v.AllowReserved = &val
	return v
}

// WithSchema sets Schema value.
func (v *Parameter) WithSchema(val SchemaOrRef) *Parameter {
	v.Schema = &val
	return v
}

// WithContent sets Content value.
func (v *Parameter) WithContent(val map[string]MediaType) *Parameter {
	v.Content = val
	return v
}

// WithContentItem sets Content item value.
func (v *Parameter) WithContentItem(key string, val MediaType) *Parameter {
	if v.Content == nil {
		v.Content = make(map[string]MediaType, 1)
	}

	v.Content[key] = val

	return v
}

// WithExample sets Example value.
func (v *Parameter) WithExample(val interface{}) *Parameter {
	v.Example = &val
	return v
}

// WithExamples sets Examples value.
func (v *Parameter) WithExamples(val map[string]ExampleOrRef) *Parameter {
	v.Examples = val
	return v
}

// WithExamplesItem sets Examples item value.
func (v *Parameter) WithExamplesItem(key string, val ExampleOrRef) *Parameter {
	if v.Examples == nil {
		v.Examples = make(map[string]ExampleOrRef, 1)
	}

	v.Examples[key] = val

	return v
}

// WithSchemaXORContent sets SchemaXORContent value.
func (v *Parameter) WithSchemaXORContent(val SchemaXORContentOneOf1) *Parameter {
	v.SchemaXORContent = &val
	return v
}

// WithLocation sets Location value.
func (v *Parameter) WithLocation(val ParameterLocation) *Parameter {
	v.Location = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Parameter) WithMapOfAnything(val map[string]interface{}) *Parameter {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Parameter) WithMapOfAnythingItem(key string, val interface{}) *Parameter {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalParameter Parameter

var ignoreKeysParameter = []string{
	"name",
	"in",
	"description",
	"required",
	"deprecated",
	"allowEmptyValue",
	"style",
	"explode",
	"allowReserved",
	"schema",
	"content",
	"example",
	"examples",
}

// UnmarshalJSON decodes JSON.
func (v *Parameter) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalParameter(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &vv.SchemaXORContent)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &vv.Location)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			vv.Example = &v
		}
	}

	for _, key := range ignoreKeysParameter {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Parameter(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(v), v.MapOfAnything, v.SchemaXORContent, v.Location)
}

// Schema structure is generated from "#/definitions/Schema".
type Schema struct {
	Title                *string                     `json:"title,omitempty"`
	MultipleOf           *float64                    `json:"multipleOf,omitempty"`
	Maximum              *float64                    `json:"maximum,omitempty"`
	ExclusiveMaximum     *bool                       `json:"exclusiveMaximum,omitempty"`
	Minimum              *float64                    `json:"minimum,omitempty"`
	ExclusiveMinimum     *bool                       `json:"exclusiveMinimum,omitempty"`
	MaxLength            *int64                      `json:"maxLength,omitempty"`
	MinLength            *int64                      `json:"minLength,omitempty"`
	Pattern              *string                     `json:"pattern,omitempty"`              // Format: regex.
	MaxItems             *int64                      `json:"maxItems,omitempty"`
	MinItems             *int64                      `json:"minItems,omitempty"`
	UniqueItems          *bool                       `json:"uniqueItems,omitempty"`
	MaxProperties        *int64                      `json:"maxProperties,omitempty"`
	MinProperties        *int64                      `json:"minProperties,omitempty"`
	Required             []string                    `json:"required,omitempty"`
	Enum                 []interface{}               `json:"enum,omitempty"`
	Type                 *SchemaType                 `json:"type,omitempty"`
	Not                  *SchemaOrRef                `json:"not,omitempty"`
	AllOf                []SchemaOrRef               `json:"allOf,omitempty"`
	OneOf                []SchemaOrRef               `json:"oneOf,omitempty"`
	AnyOf                []SchemaOrRef               `json:"anyOf,omitempty"`
	Items                *SchemaOrRef                `json:"items,omitempty"`
	Properties           map[string]SchemaOrRef      `json:"properties,omitempty"`
	AdditionalProperties *SchemaAdditionalProperties `json:"additionalProperties,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	Format               *string                     `json:"format,omitempty"`
	Default              *interface{}                `json:"default,omitempty"`
	Nullable             *bool                       `json:"nullable,omitempty"`
	Discriminator        *Discriminator              `json:"discriminator,omitempty"`
	ReadOnly             *bool                       `json:"readOnly,omitempty"`
	WriteOnly            *bool                       `json:"writeOnly,omitempty"`
	Example              *interface{}                `json:"example,omitempty"`
	ExternalDocs         *ExternalDocumentation      `json:"externalDocs,omitempty"`
	Deprecated           *bool                       `json:"deprecated,omitempty"`
	XML                  *XML                        `json:"xml,omitempty"`
	MapOfAnything        map[string]interface{}      `json:"-"`                              // Key must match pattern: `^x-`.
}

// WithTitle sets Title value.
func (v *Schema) WithTitle(val string) *Schema {
	v.Title = &val
	return v
}

// WithMultipleOf sets MultipleOf value.
func (v *Schema) WithMultipleOf(val float64) *Schema {
	v.MultipleOf = &val
	return v
}

// WithMaximum sets Maximum value.
func (v *Schema) WithMaximum(val float64) *Schema {
	v.Maximum = &val
	return v
}

// WithExclusiveMaximum sets ExclusiveMaximum value.
func (v *Schema) WithExclusiveMaximum(val bool) *Schema {
	v.ExclusiveMaximum = &val
	return v
}

// WithMinimum sets Minimum value.
func (v *Schema) WithMinimum(val float64) *Schema {
	v.Minimum = &val
	return v
}

// WithExclusiveMinimum sets ExclusiveMinimum value.
func (v *Schema) WithExclusiveMinimum(val bool) *Schema {
	v.ExclusiveMinimum = &val
	return v
}

// WithMaxLength sets MaxLength value.
func (v *Schema) WithMaxLength(val int64) *Schema {
	v.MaxLength = &val
	return v
}

// WithMinLength sets MinLength value.
func (v *Schema) WithMinLength(val int64) *Schema {
	v.MinLength = &val
	return v
}

// WithPattern sets Pattern value.
func (v *Schema) WithPattern(val string) *Schema {
	v.Pattern = &val
	return v
}

// WithMaxItems sets MaxItems value.
func (v *Schema) WithMaxItems(val int64) *Schema {
	v.MaxItems = &val
	return v
}

// WithMinItems sets MinItems value.
func (v *Schema) WithMinItems(val int64) *Schema {
	v.MinItems = &val
	return v
}

// WithUniqueItems sets UniqueItems value.
func (v *Schema) WithUniqueItems(val bool) *Schema {
	v.UniqueItems = &val
	return v
}

// WithMaxProperties sets MaxProperties value.
func (v *Schema) WithMaxProperties(val int64) *Schema {
	v.MaxProperties = &val
	return v
}

// WithMinProperties sets MinProperties value.
func (v *Schema) WithMinProperties(val int64) *Schema {
	v.MinProperties = &val
	return v
}

// WithRequired sets Required value.
func (v *Schema) WithRequired(val ...string) *Schema {
	v.Required = val
	return v
}

// WithEnum sets Enum value.
func (v *Schema) WithEnum(val ...interface{}) *Schema {
	v.Enum = val
	return v
}

// WithType sets Type value.
func (v *Schema) WithType(val SchemaType) *Schema {
	v.Type = &val
	return v
}

// WithNot sets Not value.
func (v *Schema) WithNot(val SchemaOrRef) *Schema {
	v.Not = &val
	return v
}

// WithAllOf sets AllOf value.
func (v *Schema) WithAllOf(val ...SchemaOrRef) *Schema {
	v.AllOf = val
	return v
}

// WithOneOf sets OneOf value.
func (v *Schema) WithOneOf(val ...SchemaOrRef) *Schema {
	v.OneOf = val
	return v
}

// WithAnyOf sets AnyOf value.
func (v *Schema) WithAnyOf(val ...SchemaOrRef) *Schema {
	v.AnyOf = val
	return v
}

// WithItems sets Items value.
func (v *Schema) WithItems(val SchemaOrRef) *Schema {
	v.Items = &val
	return v
}

// WithProperties sets Properties value.
func (v *Schema) WithProperties(val map[string]SchemaOrRef) *Schema {
	v.Properties = val
	return v
}

// WithPropertiesItem sets Properties item value.
func (v *Schema) WithPropertiesItem(key string, val SchemaOrRef) *Schema {
	if v.Properties == nil {
		v.Properties = make(map[string]SchemaOrRef, 1)
	}

	v.Properties[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *Schema) WithAdditionalProperties(val SchemaAdditionalProperties) *Schema {
	v.AdditionalProperties = &val
	return v
}

// WithDescription sets Description value.
func (v *Schema) WithDescription(val string) *Schema {
	v.Description = &val
	return v
}

// WithFormat sets Format value.
func (v *Schema) WithFormat(val string) *Schema {
	v.Format = &val
	return v
}

// WithDefault sets Default value.
func (v *Schema) WithDefault(val interface{}) *Schema {
	v.Default = &val
	return v
}

// WithNullable sets Nullable value.
func (v *Schema) WithNullable(val bool) *Schema {
	v.Nullable = &val
	return v
}

// WithDiscriminator sets Discriminator value.
func (v *Schema) WithDiscriminator(val Discriminator) *Schema {
	v.Discriminator = &val
	return v
}

// WithReadOnly sets ReadOnly value.
func (v *Schema) WithReadOnly(val bool) *Schema {
	v.ReadOnly = &val
	return v
}

// WithWriteOnly sets WriteOnly value.
func (v *Schema) WithWriteOnly(val bool) *Schema {
	v.WriteOnly = &val
	return v
}

// WithExample sets Example value.
func (v *Schema) WithExample(val interface{}) *Schema {
	v.Example = &val
	return v
}

// WithExternalDocs sets ExternalDocs value.
func (v *Schema) WithExternalDocs(val ExternalDocumentation) *Schema {
	v.ExternalDocs = &val
	return v
}

// WithDeprecated sets Deprecated value.
func (v *Schema) WithDeprecated(val bool) *Schema {
	v.Deprecated = &val
	return v
}

// WithXML sets XML value.
func (v *Schema) WithXML(val XML) *Schema {
	v.XML = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Schema) WithMapOfAnything(val map[string]interface{}) *Schema {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Schema) WithMapOfAnythingItem(key string, val interface{}) *Schema {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalSchema Schema

var ignoreKeysSchema = []string{
	"title",
	"multipleOf",
	"maximum",
	"exclusiveMaximum",
	"minimum",
	"exclusiveMinimum",
	"maxLength",
	"minLength",
	"pattern",
	"maxItems",
	"minItems",
	"uniqueItems",
	"maxProperties",
	"minProperties",
	"required",
	"enum",
	"type",
	"not",
	"allOf",
	"oneOf",
	"anyOf",
	"items",
	"properties",
	"additionalProperties",
	"description",
	"format",
	"default",
	"nullable",
	"discriminator",
	"readOnly",
	"writeOnly",
	"example",
	"externalDocs",
	"deprecated",
	"xml",
}

// UnmarshalJSON decodes JSON.
func (v *Schema) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalSchema(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.Default == nil {
		if _, ok := m["default"]; ok {
			var v interface{}
			vv.Default = &v
		}
	}

	if vv.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			vv.Example = &v
		}
	}

	for _, key := range ignoreKeysSchema {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Schema(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchema(v), v.MapOfAnything)
}

// SchemaReference structure is generated from "#/definitions/SchemaReference".
type SchemaReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *SchemaReference) WithRef(val string) *SchemaReference {
	v.Ref = val
	return v
}

// SchemaOrRef structure is generated from "#/definitions/SchemaOrRef".
type SchemaOrRef struct {
	Schema          *Schema          `json:"-"`
	SchemaReference *SchemaReference `json:"-"`
}

// WithSchema sets Schema value.
func (v *SchemaOrRef) WithSchema(val Schema) *SchemaOrRef {
	v.Schema = &val
	return v
}

// WithSchemaReference sets SchemaReference value.
func (v *SchemaOrRef) WithSchemaReference(val SchemaReference) *SchemaOrRef {
	v.SchemaReference = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *SchemaOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.Schema)
	if err != nil {
		v.Schema = nil
	}

	err = json.Unmarshal(data, &v.SchemaReference)
	if err != nil {
		v.SchemaReference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v SchemaOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.Schema, v.SchemaReference)
}

// SchemaAdditionalProperties structure is generated from "#/definitions/Schema->additionalProperties".
type SchemaAdditionalProperties struct {
	SchemaOrRef *SchemaOrRef `json:"-"`
	Bool        *bool        `json:"-"`
}

// WithSchemaOrRef sets SchemaOrRef value.
func (v *SchemaAdditionalProperties) WithSchemaOrRef(val SchemaOrRef) *SchemaAdditionalProperties {
	v.SchemaOrRef = &val
	return v
}

// WithBool sets Bool value.
func (v *SchemaAdditionalProperties) WithBool(val bool) *SchemaAdditionalProperties {
	v.Bool = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *SchemaAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.SchemaOrRef)
	if err != nil {
		v.SchemaOrRef = nil
	}

	err = json.Unmarshal(data, &v.Bool)
	if err != nil {
		v.Bool = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v SchemaAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.SchemaOrRef, v.Bool)
}

// Discriminator structure is generated from "#/definitions/Discriminator".
type Discriminator struct {
	PropertyName         string                 `json:"propertyName"`      // Required.
	Mapping              map[string]string      `json:"mapping,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                 // All unmatched properties.
}

// WithPropertyName sets PropertyName value.
func (v *Discriminator) WithPropertyName(val string) *Discriminator {
	v.PropertyName = val
	return v
}

// WithMapping sets Mapping value.
func (v *Discriminator) WithMapping(val map[string]string) *Discriminator {
	v.Mapping = val
	return v
}

// WithMappingItem sets Mapping item value.
func (v *Discriminator) WithMappingItem(key string, val string) *Discriminator {
	if v.Mapping == nil {
		v.Mapping = make(map[string]string, 1)
	}

	v.Mapping[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *Discriminator) WithAdditionalProperties(val map[string]interface{}) *Discriminator {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *Discriminator) WithAdditionalPropertiesItem(key string, val interface{}) *Discriminator {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

type marshalDiscriminator Discriminator

var ignoreKeysDiscriminator = []string{
	"propertyName",
	"mapping",
}

// UnmarshalJSON decodes JSON.
func (v *Discriminator) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalDiscriminator(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysDiscriminator {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = Discriminator(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Discriminator) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalDiscriminator(v))
	}

	return marshalUnion(marshalDiscriminator(v), v.AdditionalProperties)
}

// XML structure is generated from "#/definitions/XML".
type XML struct {
	Name          *string                `json:"name,omitempty"`
	Namespace     *string                `json:"namespace,omitempty"` // Format: uri.
	Prefix        *string                `json:"prefix,omitempty"`
	Attribute     *bool                  `json:"attribute,omitempty"`
	Wrapped       *bool                  `json:"wrapped,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                   // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (v *XML) WithName(val string) *XML {
	v.Name = &val
	return v
}

// WithNamespace sets Namespace value.
func (v *XML) WithNamespace(val string) *XML {
	v.Namespace = &val
	return v
}

// WithPrefix sets Prefix value.
func (v *XML) WithPrefix(val string) *XML {
	v.Prefix = &val
	return v
}

// WithAttribute sets Attribute value.
func (v *XML) WithAttribute(val bool) *XML {
	v.Attribute = &val
	return v
}

// WithWrapped sets Wrapped value.
func (v *XML) WithWrapped(val bool) *XML {
	v.Wrapped = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *XML) WithMapOfAnything(val map[string]interface{}) *XML {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *XML) WithMapOfAnythingItem(key string, val interface{}) *XML {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalXML XML

var ignoreKeysXML = []string{
	"name",
	"namespace",
	"prefix",
	"attribute",
	"wrapped",
}

// UnmarshalJSON decodes JSON.
func (v *XML) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalXML(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysXML {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = XML(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v XML) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalXML(v), v.MapOfAnything)
}

// MediaType structure is generated from "#/definitions/MediaType".
type MediaType struct {
	Schema        *SchemaOrRef            `json:"schema,omitempty"`
	Example       *interface{}            `json:"example,omitempty"`
	Examples      map[string]ExampleOrRef `json:"examples,omitempty"`
	Encoding      map[string]Encoding     `json:"encoding,omitempty"`
	MapOfAnything map[string]interface{}  `json:"-"`                  // Key must match pattern: `^x-`.
}

// WithSchema sets Schema value.
func (v *MediaType) WithSchema(val SchemaOrRef) *MediaType {
	v.Schema = &val
	return v
}

// WithExample sets Example value.
func (v *MediaType) WithExample(val interface{}) *MediaType {
	v.Example = &val
	return v
}

// WithExamples sets Examples value.
func (v *MediaType) WithExamples(val map[string]ExampleOrRef) *MediaType {
	v.Examples = val
	return v
}

// WithExamplesItem sets Examples item value.
func (v *MediaType) WithExamplesItem(key string, val ExampleOrRef) *MediaType {
	if v.Examples == nil {
		v.Examples = make(map[string]ExampleOrRef, 1)
	}

	v.Examples[key] = val

	return v
}

// WithEncoding sets Encoding value.
func (v *MediaType) WithEncoding(val map[string]Encoding) *MediaType {
	v.Encoding = val
	return v
}

// WithEncodingItem sets Encoding item value.
func (v *MediaType) WithEncodingItem(key string, val Encoding) *MediaType {
	if v.Encoding == nil {
		v.Encoding = make(map[string]Encoding, 1)
	}

	v.Encoding[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *MediaType) WithMapOfAnything(val map[string]interface{}) *MediaType {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *MediaType) WithMapOfAnythingItem(key string, val interface{}) *MediaType {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalMediaType MediaType

var ignoreKeysMediaType = []string{
	"schema",
	"example",
	"examples",
	"encoding",
}

// UnmarshalJSON decodes JSON.
func (v *MediaType) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalMediaType(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			vv.Example = &v
		}
	}

	for _, key := range ignoreKeysMediaType {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = MediaType(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v MediaType) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMediaType(v), v.MapOfAnything)
}

// ExampleReference structure is generated from "#/definitions/ExampleReference".
type ExampleReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *ExampleReference) WithRef(val string) *ExampleReference {
	v.Ref = val
	return v
}

// Example structure is generated from "#/definitions/Example".
type Example struct {
	Summary       *string                `json:"summary,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Value         *interface{}           `json:"value,omitempty"`
	ExternalValue *string                `json:"externalValue,omitempty"` // Format: uri-reference.
	MapOfAnything map[string]interface{} `json:"-"`                       // Key must match pattern: `^x-`.
}

// WithSummary sets Summary value.
func (v *Example) WithSummary(val string) *Example {
	v.Summary = &val
	return v
}

// WithDescription sets Description value.
func (v *Example) WithDescription(val string) *Example {
	v.Description = &val
	return v
}

// WithValue sets Value value.
func (v *Example) WithValue(val interface{}) *Example {
	v.Value = &val
	return v
}

// WithExternalValue sets ExternalValue value.
func (v *Example) WithExternalValue(val string) *Example {
	v.ExternalValue = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Example) WithMapOfAnything(val map[string]interface{}) *Example {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Example) WithMapOfAnythingItem(key string, val interface{}) *Example {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalExample Example

var ignoreKeysExample = []string{
	"summary",
	"description",
	"value",
	"externalValue",
}

// UnmarshalJSON decodes JSON.
func (v *Example) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalExample(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.Value == nil {
		if _, ok := m["value"]; ok {
			var v interface{}
			vv.Value = &v
		}
	}

	for _, key := range ignoreKeysExample {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Example(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Example) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExample(v), v.MapOfAnything)
}

// ExampleOrRef structure is generated from "#/definitions/ExampleOrRef".
type ExampleOrRef struct {
	ExampleReference *ExampleReference `json:"-"`
	Example          *Example          `json:"-"`
}

// WithExampleReference sets ExampleReference value.
func (v *ExampleOrRef) WithExampleReference(val ExampleReference) *ExampleOrRef {
	v.ExampleReference = &val
	return v
}

// WithExample sets Example value.
func (v *ExampleOrRef) WithExample(val Example) *ExampleOrRef {
	v.Example = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *ExampleOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.ExampleReference)
	if err != nil {
		v.ExampleReference = nil
	}

	err = json.Unmarshal(data, &v.Example)
	if err != nil {
		v.Example = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ExampleOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.ExampleReference, v.Example)
}

// Encoding structure is generated from "#/definitions/Encoding".
type Encoding struct {
	ContentType   *string           `json:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty"`
	Style         *EncodingStyle    `json:"style,omitempty"`
	Explode       *bool             `json:"explode,omitempty"`
	AllowReserved *bool             `json:"allowReserved,omitempty"`
}

// WithContentType sets ContentType value.
func (v *Encoding) WithContentType(val string) *Encoding {
	v.ContentType = &val
	return v
}

// WithHeaders sets Headers value.
func (v *Encoding) WithHeaders(val map[string]Header) *Encoding {
	v.Headers = val
	return v
}

// WithHeadersItem sets Headers item value.
func (v *Encoding) WithHeadersItem(key string, val Header) *Encoding {
	if v.Headers == nil {
		v.Headers = make(map[string]Header, 1)
	}

	v.Headers[key] = val

	return v
}

// WithStyle sets Style value.
func (v *Encoding) WithStyle(val EncodingStyle) *Encoding {
	v.Style = &val
	return v
}

// WithExplode sets Explode value.
func (v *Encoding) WithExplode(val bool) *Encoding {
	v.Explode = &val
	return v
}

// WithAllowReserved sets AllowReserved value.
func (v *Encoding) WithAllowReserved(val bool) *Encoding {
	v.AllowReserved = &val
	return v
}

// Header structure is generated from "#/definitions/Header".
type Header struct {
	Description     *string                 `json:"description,omitempty"`
	Required        *bool                   `json:"required,omitempty"`
	Deprecated      *bool                   `json:"deprecated,omitempty"`
	AllowEmptyValue *bool                   `json:"allowEmptyValue,omitempty"`
	Explode         *bool                   `json:"explode,omitempty"`
	AllowReserved   *bool                   `json:"allowReserved,omitempty"`
	Schema          *SchemaOrRef            `json:"schema,omitempty"`
	Content         map[string]MediaType    `json:"content,omitempty"`
	Example         *interface{}            `json:"example,omitempty"`
	Examples        map[string]ExampleOrRef `json:"examples,omitempty"`
	MapOfAnything   map[string]interface{}  `json:"-"`                         // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (v *Header) WithDescription(val string) *Header {
	v.Description = &val
	return v
}

// WithRequired sets Required value.
func (v *Header) WithRequired(val bool) *Header {
	v.Required = &val
	return v
}

// WithDeprecated sets Deprecated value.
func (v *Header) WithDeprecated(val bool) *Header {
	v.Deprecated = &val
	return v
}

// WithAllowEmptyValue sets AllowEmptyValue value.
func (v *Header) WithAllowEmptyValue(val bool) *Header {
	v.AllowEmptyValue = &val
	return v
}

// WithExplode sets Explode value.
func (v *Header) WithExplode(val bool) *Header {
	v.Explode = &val
	return v
}

// WithAllowReserved sets AllowReserved value.
func (v *Header) WithAllowReserved(val bool) *Header {
	v.AllowReserved = &val
	return v
}

// WithSchema sets Schema value.
func (v *Header) WithSchema(val SchemaOrRef) *Header {
	v.Schema = &val
	return v
}

// WithContent sets Content value.
func (v *Header) WithContent(val map[string]MediaType) *Header {
	v.Content = val
	return v
}

// WithContentItem sets Content item value.
func (v *Header) WithContentItem(key string, val MediaType) *Header {
	if v.Content == nil {
		v.Content = make(map[string]MediaType, 1)
	}

	v.Content[key] = val

	return v
}

// WithExample sets Example value.
func (v *Header) WithExample(val interface{}) *Header {
	v.Example = &val
	return v
}

// WithExamples sets Examples value.
func (v *Header) WithExamples(val map[string]ExampleOrRef) *Header {
	v.Examples = val
	return v
}

// WithExamplesItem sets Examples item value.
func (v *Header) WithExamplesItem(key string, val ExampleOrRef) *Header {
	if v.Examples == nil {
		v.Examples = make(map[string]ExampleOrRef, 1)
	}

	v.Examples[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Header) WithMapOfAnything(val map[string]interface{}) *Header {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Header) WithMapOfAnythingItem(key string, val interface{}) *Header {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalHeader Header

var ignoreKeysHeader = []string{
	"description",
	"required",
	"deprecated",
	"allowEmptyValue",
	"explode",
	"allowReserved",
	"schema",
	"content",
	"example",
	"examples",
	"style",
}

// UnmarshalJSON decodes JSON.
func (v *Header) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalHeader(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["style"]; !ok || string(v) != `"simple"` {
		return fmt.Errorf(`bad or missing const value for "style" ("simple" expected, %s received)`, v)
	}

	delete(m, "style")

	if vv.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			vv.Example = &v
		}
	}

	for _, key := range ignoreKeysHeader {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Header(vv)

	return nil
}

var (
	// constHeader is unconditionally added to JSON.
	constHeader = json.RawMessage(`{"style":"simple"}`)
)

// MarshalJSON encodes JSON.
func (v Header) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHeader, marshalHeader(v), v.MapOfAnything)
}

// SchemaXORContentOneOf1 structure is generated from "#/definitions/SchemaXORContent/oneOf/1".
//
// Has Content.
//
// Some properties are not allowed if content is present.
type SchemaXORContentOneOf1 struct {

}

// ParameterLocationOneOf0 structure is generated from "#/definitions/ParameterLocation/oneOf/0".
//
// Path Parameter.
//
// Parameter in path.
type ParameterLocationOneOf0 struct {
	Style                *ParameterLocationOneOf0Style `json:"style,omitempty"`
	AdditionalProperties map[string]interface{}        `json:"-"`               // All unmatched properties.
}

// WithStyle sets Style value.
func (v *ParameterLocationOneOf0) WithStyle(val ParameterLocationOneOf0Style) *ParameterLocationOneOf0 {
	v.Style = &val
	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ParameterLocationOneOf0) WithAdditionalProperties(val map[string]interface{}) *ParameterLocationOneOf0 {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ParameterLocationOneOf0) WithAdditionalPropertiesItem(key string, val interface{}) *ParameterLocationOneOf0 {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

type marshalParameterLocationOneOf0 ParameterLocationOneOf0

var ignoreKeysParameterLocationOneOf0 = []string{
	"style",
	"in",
	"required",
}

// UnmarshalJSON decodes JSON.
func (v *ParameterLocationOneOf0) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalParameterLocationOneOf0(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["in"]; !ok || string(v) != `"path"` {
		return fmt.Errorf(`bad or missing const value for "in" ("path" expected, %s received)`, v)
	}

	delete(m, "in")

	if v, ok := m["required"]; !ok || string(v) != "true" {
		return fmt.Errorf(`bad or missing const value for "required" (true expected, %s received)`, v)
	}

	delete(m, "required")

	for _, key := range ignoreKeysParameterLocationOneOf0 {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = ParameterLocationOneOf0(vv)

	return nil
}

var (
	// constParameterLocationOneOf0 is unconditionally added to JSON.
	constParameterLocationOneOf0 = json.RawMessage(`{"in":"path","required":true}`)
)

// MarshalJSON encodes JSON.
func (v ParameterLocationOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf0, marshalParameterLocationOneOf0(v), v.AdditionalProperties)
}

// ParameterLocationOneOf1 structure is generated from "#/definitions/ParameterLocation/oneOf/1".
//
// Query Parameter.
//
// Parameter in query.
type ParameterLocationOneOf1 struct {
	Style                *ParameterLocationOneOf1Style `json:"style,omitempty"`
	AdditionalProperties map[string]interface{}        `json:"-"`               // All unmatched properties.
}

// WithStyle sets Style value.
func (v *ParameterLocationOneOf1) WithStyle(val ParameterLocationOneOf1Style) *ParameterLocationOneOf1 {
	v.Style = &val
	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ParameterLocationOneOf1) WithAdditionalProperties(val map[string]interface{}) *ParameterLocationOneOf1 {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ParameterLocationOneOf1) WithAdditionalPropertiesItem(key string, val interface{}) *ParameterLocationOneOf1 {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

type marshalParameterLocationOneOf1 ParameterLocationOneOf1

var ignoreKeysParameterLocationOneOf1 = []string{
	"style",
	"in",
}

// UnmarshalJSON decodes JSON.
func (v *ParameterLocationOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalParameterLocationOneOf1(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["in"]; !ok || string(v) != `"query"` {
		return fmt.Errorf(`bad or missing const value for "in" ("query" expected, %s received)`, v)
	}

	delete(m, "in")

	for _, key := range ignoreKeysParameterLocationOneOf1 {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = ParameterLocationOneOf1(vv)

	return nil
}

var (
	// constParameterLocationOneOf1 is unconditionally added to JSON.
	constParameterLocationOneOf1 = json.RawMessage(`{"in":"query"}`)
)

// MarshalJSON encodes JSON.
func (v ParameterLocationOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf1, marshalParameterLocationOneOf1(v), v.AdditionalProperties)
}

// ParameterLocationOneOf2 structure is generated from "#/definitions/ParameterLocation/oneOf/2".
//
// Header Parameter.
//
// Parameter in header.
type ParameterLocationOneOf2 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ParameterLocationOneOf2) WithAdditionalProperties(val map[string]interface{}) *ParameterLocationOneOf2 {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ParameterLocationOneOf2) WithAdditionalPropertiesItem(key string, val interface{}) *ParameterLocationOneOf2 {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ParameterLocationOneOf2) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["in"]; !ok || string(v) != `"header"` {
		return fmt.Errorf(`bad or missing const value for "in" ("header" expected, %s received)`, v)
	}

	delete(m, "in")

	if v, ok := m["style"]; !ok || string(v) != `"simple"` {
		return fmt.Errorf(`bad or missing const value for "style" ("simple" expected, %s received)`, v)
	}

	delete(m, "style")

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constParameterLocationOneOf2 is unconditionally added to JSON.
	constParameterLocationOneOf2 = json.RawMessage(`{"in":"header","style":"simple"}`)
)

// MarshalJSON encodes JSON.
func (v ParameterLocationOneOf2) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf2, v.AdditionalProperties)
}

// ParameterLocationOneOf3 structure is generated from "#/definitions/ParameterLocation/oneOf/3".
//
// Cookie Parameter.
//
// Parameter in cookie.
type ParameterLocationOneOf3 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ParameterLocationOneOf3) WithAdditionalProperties(val map[string]interface{}) *ParameterLocationOneOf3 {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ParameterLocationOneOf3) WithAdditionalPropertiesItem(key string, val interface{}) *ParameterLocationOneOf3 {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ParameterLocationOneOf3) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["in"]; !ok || string(v) != `"cookie"` {
		return fmt.Errorf(`bad or missing const value for "in" ("cookie" expected, %s received)`, v)
	}

	delete(m, "in")

	if v, ok := m["style"]; !ok || string(v) != `"form"` {
		return fmt.Errorf(`bad or missing const value for "style" ("form" expected, %s received)`, v)
	}

	delete(m, "style")

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constParameterLocationOneOf3 is unconditionally added to JSON.
	constParameterLocationOneOf3 = json.RawMessage(`{"in":"cookie","style":"form"}`)
)

// MarshalJSON encodes JSON.
func (v ParameterLocationOneOf3) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf3, v.AdditionalProperties)
}

// ParameterLocation structure is generated from "#/definitions/ParameterLocation".
//
// Parameter location.
type ParameterLocation struct {
	OneOf0 *ParameterLocationOneOf0 `json:"-"`
	OneOf1 *ParameterLocationOneOf1 `json:"-"`
	OneOf2 *ParameterLocationOneOf2 `json:"-"`
	OneOf3 *ParameterLocationOneOf3 `json:"-"`
}

// WithOneOf0 sets OneOf0 value.
func (v *ParameterLocation) WithOneOf0(val ParameterLocationOneOf0) *ParameterLocation {
	v.OneOf0 = &val
	return v
}

// WithOneOf1 sets OneOf1 value.
func (v *ParameterLocation) WithOneOf1(val ParameterLocationOneOf1) *ParameterLocation {
	v.OneOf1 = &val
	return v
}

// WithOneOf2 sets OneOf2 value.
func (v *ParameterLocation) WithOneOf2(val ParameterLocationOneOf2) *ParameterLocation {
	v.OneOf2 = &val
	return v
}

// WithOneOf3 sets OneOf3 value.
func (v *ParameterLocation) WithOneOf3(val ParameterLocationOneOf3) *ParameterLocation {
	v.OneOf3 = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *ParameterLocation) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.OneOf0)
	if err != nil {
		v.OneOf0 = nil
	}

	err = json.Unmarshal(data, &v.OneOf1)
	if err != nil {
		v.OneOf1 = nil
	}

	err = json.Unmarshal(data, &v.OneOf2)
	if err != nil {
		v.OneOf2 = nil
	}

	err = json.Unmarshal(data, &v.OneOf3)
	if err != nil {
		v.OneOf3 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ParameterLocation) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.OneOf0, v.OneOf1, v.OneOf2, v.OneOf3)
}

// ParameterOrRef structure is generated from "#/definitions/ParameterOrRef".
type ParameterOrRef struct {
	ParameterReference *ParameterReference `json:"-"`
	Parameter          *Parameter          `json:"-"`
}

// WithParameterReference sets ParameterReference value.
func (v *ParameterOrRef) WithParameterReference(val ParameterReference) *ParameterOrRef {
	v.ParameterReference = &val
	return v
}

// WithParameter sets Parameter value.
func (v *ParameterOrRef) WithParameter(val Parameter) *ParameterOrRef {
	v.Parameter = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *ParameterOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.ParameterReference)
	if err != nil {
		v.ParameterReference = nil
	}

	err = json.Unmarshal(data, &v.Parameter)
	if err != nil {
		v.Parameter = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ParameterOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.ParameterReference, v.Parameter)
}

// Operation structure is generated from "#/definitions/Operation".
type Operation struct {
	Tags          []string                 `json:"tags,omitempty"`
	Summary       *string                  `json:"summary,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	ExternalDocs  *ExternalDocumentation   `json:"externalDocs,omitempty"`
	ID            *string                  `json:"operationId,omitempty"`
	Parameters    []ParameterOrRef         `json:"parameters,omitempty"`
	RequestBody   *RequestBodyOrRef        `json:"requestBody,omitempty"`
	Responses     Responses                `json:"responses"`              // Required.
	Callbacks     map[string]CallbackOrRef `json:"callbacks,omitempty"`
	Deprecated    *bool                    `json:"deprecated,omitempty"`
	Security      []map[string][]string    `json:"security,omitempty"`
	Servers       []Server                 `json:"servers,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"`                      // Key must match pattern: `^x-`.
}

// WithTags sets Tags value.
func (v *Operation) WithTags(val ...string) *Operation {
	v.Tags = val
	return v
}

// WithSummary sets Summary value.
func (v *Operation) WithSummary(val string) *Operation {
	v.Summary = &val
	return v
}

// WithDescription sets Description value.
func (v *Operation) WithDescription(val string) *Operation {
	v.Description = &val
	return v
}

// WithExternalDocs sets ExternalDocs value.
func (v *Operation) WithExternalDocs(val ExternalDocumentation) *Operation {
	v.ExternalDocs = &val
	return v
}

// WithID sets ID value.
func (v *Operation) WithID(val string) *Operation {
	v.ID = &val
	return v
}

// WithParameters sets Parameters value.
func (v *Operation) WithParameters(val ...ParameterOrRef) *Operation {
	v.Parameters = val
	return v
}

// WithRequestBody sets RequestBody value.
func (v *Operation) WithRequestBody(val RequestBodyOrRef) *Operation {
	v.RequestBody = &val
	return v
}

// WithResponses sets Responses value.
func (v *Operation) WithResponses(val Responses) *Operation {
	v.Responses = val
	return v
}

// WithCallbacks sets Callbacks value.
func (v *Operation) WithCallbacks(val map[string]CallbackOrRef) *Operation {
	v.Callbacks = val
	return v
}

// WithCallbacksItem sets Callbacks item value.
func (v *Operation) WithCallbacksItem(key string, val CallbackOrRef) *Operation {
	if v.Callbacks == nil {
		v.Callbacks = make(map[string]CallbackOrRef, 1)
	}

	v.Callbacks[key] = val

	return v
}

// WithDeprecated sets Deprecated value.
func (v *Operation) WithDeprecated(val bool) *Operation {
	v.Deprecated = &val
	return v
}

// WithSecurity sets Security value.
func (v *Operation) WithSecurity(val ...map[string][]string) *Operation {
	v.Security = val
	return v
}

// WithServers sets Servers value.
func (v *Operation) WithServers(val ...Server) *Operation {
	v.Servers = val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Operation) WithMapOfAnything(val map[string]interface{}) *Operation {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Operation) WithMapOfAnythingItem(key string, val interface{}) *Operation {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalOperation Operation

var ignoreKeysOperation = []string{
	"tags",
	"summary",
	"description",
	"externalDocs",
	"operationId",
	"parameters",
	"requestBody",
	"responses",
	"callbacks",
	"deprecated",
	"security",
	"servers",
}

// UnmarshalJSON decodes JSON.
func (v *Operation) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOperation(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOperation {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Operation(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(v), v.MapOfAnything)
}

// RequestBodyReference structure is generated from "#/definitions/RequestBodyReference".
type RequestBodyReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *RequestBodyReference) WithRef(val string) *RequestBodyReference {
	v.Ref = val
	return v
}

// RequestBody structure is generated from "#/definitions/RequestBody".
type RequestBody struct {
	Description   *string                `json:"description,omitempty"`
	Content       map[string]MediaType   `json:"content"`               // Required.
	Required      *bool                  `json:"required,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (v *RequestBody) WithDescription(val string) *RequestBody {
	v.Description = &val
	return v
}

// WithContent sets Content value.
func (v *RequestBody) WithContent(val map[string]MediaType) *RequestBody {
	v.Content = val
	return v
}

// WithContentItem sets Content item value.
func (v *RequestBody) WithContentItem(key string, val MediaType) *RequestBody {
	if v.Content == nil {
		v.Content = make(map[string]MediaType, 1)
	}

	v.Content[key] = val

	return v
}

// WithRequired sets Required value.
func (v *RequestBody) WithRequired(val bool) *RequestBody {
	v.Required = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *RequestBody) WithMapOfAnything(val map[string]interface{}) *RequestBody {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *RequestBody) WithMapOfAnythingItem(key string, val interface{}) *RequestBody {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalRequestBody RequestBody

var ignoreKeysRequestBody = []string{
	"description",
	"content",
	"required",
}

// UnmarshalJSON decodes JSON.
func (v *RequestBody) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalRequestBody(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysRequestBody {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = RequestBody(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v RequestBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalRequestBody(v), v.MapOfAnything)
}

// RequestBodyOrRef structure is generated from "#/definitions/RequestBodyOrRef".
type RequestBodyOrRef struct {
	RequestBodyReference *RequestBodyReference `json:"-"`
	RequestBody          *RequestBody          `json:"-"`
}

// WithRequestBodyReference sets RequestBodyReference value.
func (v *RequestBodyOrRef) WithRequestBodyReference(val RequestBodyReference) *RequestBodyOrRef {
	v.RequestBodyReference = &val
	return v
}

// WithRequestBody sets RequestBody value.
func (v *RequestBodyOrRef) WithRequestBody(val RequestBody) *RequestBodyOrRef {
	v.RequestBody = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *RequestBodyOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.RequestBodyReference)
	if err != nil {
		v.RequestBodyReference = nil
	}

	err = json.Unmarshal(data, &v.RequestBody)
	if err != nil {
		v.RequestBody = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v RequestBodyOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.RequestBodyReference, v.RequestBody)
}

// Responses structure is generated from "#/definitions/Responses".
type Responses struct {
	Default                  *ResponseOrRef           `json:"default,omitempty"`
	MapOfResponseOrRefValues map[string]ResponseOrRef `json:"-"`                 // Key must match pattern: `^[1-5](?:\d{2}|XX)$`.
	MapOfAnything            map[string]interface{}   `json:"-"`                 // Key must match pattern: `^x-`.
}

// WithDefault sets Default value.
func (v *Responses) WithDefault(val ResponseOrRef) *Responses {
	v.Default = &val
	return v
}

// WithMapOfResponseOrRefValues sets MapOfResponseOrRefValues value.
func (v *Responses) WithMapOfResponseOrRefValues(val map[string]ResponseOrRef) *Responses {
	v.MapOfResponseOrRefValues = val
	return v
}

// WithMapOfResponseOrRefValuesItem sets MapOfResponseOrRefValues item value.
func (v *Responses) WithMapOfResponseOrRefValuesItem(key string, val ResponseOrRef) *Responses {
	if v.MapOfResponseOrRefValues == nil {
		v.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
	}

	v.MapOfResponseOrRefValues[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Responses) WithMapOfAnything(val map[string]interface{}) *Responses {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Responses) WithMapOfAnythingItem(key string, val interface{}) *Responses {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalResponses Responses

var ignoreKeysResponses = []string{
	"default",
}

// UnmarshalJSON decodes JSON.
func (v *Responses) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalResponses(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysResponses {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regex15D2XX.MatchString(key) {
			matched = true

			if vv.MapOfResponseOrRefValues == nil {
				vv.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
			}

			var val ResponseOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfResponseOrRefValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Responses(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Responses) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponses(v), v.MapOfResponseOrRefValues, v.MapOfAnything)
}

// ResponseReference structure is generated from "#/definitions/ResponseReference".
type ResponseReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *ResponseReference) WithRef(val string) *ResponseReference {
	v.Ref = val
	return v
}

// Response structure is generated from "#/definitions/Response".
type Response struct {
	Description   string                 `json:"description"`       // Required.
	Headers       map[string]HeaderOrRef `json:"headers,omitempty"`
	Content       map[string]MediaType   `json:"content,omitempty"`
	Links         map[string]LinkOrRef   `json:"links,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                 // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (v *Response) WithDescription(val string) *Response {
	v.Description = val
	return v
}

// WithHeaders sets Headers value.
func (v *Response) WithHeaders(val map[string]HeaderOrRef) *Response {
	v.Headers = val
	return v
}

// WithHeadersItem sets Headers item value.
func (v *Response) WithHeadersItem(key string, val HeaderOrRef) *Response {
	if v.Headers == nil {
		v.Headers = make(map[string]HeaderOrRef, 1)
	}

	v.Headers[key] = val

	return v
}

// WithContent sets Content value.
func (v *Response) WithContent(val map[string]MediaType) *Response {
	v.Content = val
	return v
}

// WithContentItem sets Content item value.
func (v *Response) WithContentItem(key string, val MediaType) *Response {
	if v.Content == nil {
		v.Content = make(map[string]MediaType, 1)
	}

	v.Content[key] = val

	return v
}

// WithLinks sets Links value.
func (v *Response) WithLinks(val map[string]LinkOrRef) *Response {
	v.Links = val
	return v
}

// WithLinksItem sets Links item value.
func (v *Response) WithLinksItem(key string, val LinkOrRef) *Response {
	if v.Links == nil {
		v.Links = make(map[string]LinkOrRef, 1)
	}

	v.Links[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Response) WithMapOfAnything(val map[string]interface{}) *Response {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Response) WithMapOfAnythingItem(key string, val interface{}) *Response {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalResponse Response

var ignoreKeysResponse = []string{
	"description",
	"headers",
	"content",
	"links",
}

// UnmarshalJSON decodes JSON.
func (v *Response) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalResponse(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysResponse {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Response(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Response) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponse(v), v.MapOfAnything)
}

// HeaderReference structure is generated from "#/definitions/HeaderReference".
type HeaderReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *HeaderReference) WithRef(val string) *HeaderReference {
	v.Ref = val
	return v
}

// HeaderOrRef structure is generated from "#/definitions/HeaderOrRef".
type HeaderOrRef struct {
	HeaderReference *HeaderReference `json:"-"`
	Header          *Header          `json:"-"`
}

// WithHeaderReference sets HeaderReference value.
func (v *HeaderOrRef) WithHeaderReference(val HeaderReference) *HeaderOrRef {
	v.HeaderReference = &val
	return v
}

// WithHeader sets Header value.
func (v *HeaderOrRef) WithHeader(val Header) *HeaderOrRef {
	v.Header = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *HeaderOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.HeaderReference)
	if err != nil {
		v.HeaderReference = nil
	}

	err = json.Unmarshal(data, &v.Header)
	if err != nil {
		v.Header = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v HeaderOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.HeaderReference, v.Header)
}

// LinkReference structure is generated from "#/definitions/LinkReference".
type LinkReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *LinkReference) WithRef(val string) *LinkReference {
	v.Ref = val
	return v
}

// Link structure is generated from "#/definitions/Link".
type Link struct {
	OperationID   *string                `json:"operationId,omitempty"`
	OperationRef  *string                `json:"operationRef,omitempty"` // Format: uri-reference.
	Parameters    map[string]interface{} `json:"parameters,omitempty"`
	RequestBody   *interface{}           `json:"requestBody,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Server        *Server                `json:"server,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: `^x-`.
}

// WithOperationID sets OperationID value.
func (v *Link) WithOperationID(val string) *Link {
	v.OperationID = &val
	return v
}

// WithOperationRef sets OperationRef value.
func (v *Link) WithOperationRef(val string) *Link {
	v.OperationRef = &val
	return v
}

// WithParameters sets Parameters value.
func (v *Link) WithParameters(val map[string]interface{}) *Link {
	v.Parameters = val
	return v
}

// WithParametersItem sets Parameters item value.
func (v *Link) WithParametersItem(key string, val interface{}) *Link {
	if v.Parameters == nil {
		v.Parameters = make(map[string]interface{}, 1)
	}

	v.Parameters[key] = val

	return v
}

// WithRequestBody sets RequestBody value.
func (v *Link) WithRequestBody(val interface{}) *Link {
	v.RequestBody = &val
	return v
}

// WithDescription sets Description value.
func (v *Link) WithDescription(val string) *Link {
	v.Description = &val
	return v
}

// WithServer sets Server value.
func (v *Link) WithServer(val Server) *Link {
	v.Server = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Link) WithMapOfAnything(val map[string]interface{}) *Link {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Link) WithMapOfAnythingItem(key string, val interface{}) *Link {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalLink Link

var ignoreKeysLink = []string{
	"operationId",
	"operationRef",
	"parameters",
	"requestBody",
	"description",
	"server",
}

// UnmarshalJSON decodes JSON.
func (v *Link) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalLink(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.RequestBody == nil {
		if _, ok := m["requestBody"]; ok {
			var v interface{}
			vv.RequestBody = &v
		}
	}

	for _, key := range ignoreKeysLink {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Link(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Link) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLink(v), v.MapOfAnything)
}

// LinkOrRef structure is generated from "#/definitions/LinkOrRef".
type LinkOrRef struct {
	LinkReference *LinkReference `json:"-"`
	Link          *Link          `json:"-"`
}

// WithLinkReference sets LinkReference value.
func (v *LinkOrRef) WithLinkReference(val LinkReference) *LinkOrRef {
	v.LinkReference = &val
	return v
}

// WithLink sets Link value.
func (v *LinkOrRef) WithLink(val Link) *LinkOrRef {
	v.Link = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *LinkOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.LinkReference)
	if err != nil {
		v.LinkReference = nil
	}

	err = json.Unmarshal(data, &v.Link)
	if err != nil {
		v.Link = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v LinkOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.LinkReference, v.Link)
}

// ResponseOrRef structure is generated from "#/definitions/ResponseOrRef".
type ResponseOrRef struct {
	ResponseReference *ResponseReference `json:"-"`
	Response          *Response          `json:"-"`
}

// WithResponseReference sets ResponseReference value.
func (v *ResponseOrRef) WithResponseReference(val ResponseReference) *ResponseOrRef {
	v.ResponseReference = &val
	return v
}

// WithResponse sets Response value.
func (v *ResponseOrRef) WithResponse(val Response) *ResponseOrRef {
	v.Response = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *ResponseOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.ResponseReference)
	if err != nil {
		v.ResponseReference = nil
	}

	err = json.Unmarshal(data, &v.Response)
	if err != nil {
		v.Response = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ResponseOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.ResponseReference, v.Response)
}

// CallbackReference structure is generated from "#/definitions/CallbackReference".
type CallbackReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *CallbackReference) WithRef(val string) *CallbackReference {
	v.Ref = val
	return v
}

// Callback structure is generated from "#/definitions/Callback".
type Callback struct {
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
	AdditionalProperties map[string]PathItem    `json:"-"` // All unmatched properties.
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Callback) WithMapOfAnything(val map[string]interface{}) *Callback {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Callback) WithMapOfAnythingItem(key string, val interface{}) *Callback {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *Callback) WithAdditionalProperties(val map[string]PathItem) *Callback {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *Callback) WithAdditionalPropertiesItem(key string, val PathItem) *Callback {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]PathItem, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *Callback) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if v.MapOfAnything == nil {
				v.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]PathItem, 1)
		}

		var val PathItem

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v Callback) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfAnything, v.AdditionalProperties)
}

// CallbackOrRef structure is generated from "#/definitions/CallbackOrRef".
type CallbackOrRef struct {
	CallbackReference *CallbackReference `json:"-"`
	Callback          *Callback          `json:"-"`
}

// WithCallbackReference sets CallbackReference value.
func (v *CallbackOrRef) WithCallbackReference(val CallbackReference) *CallbackOrRef {
	v.CallbackReference = &val
	return v
}

// WithCallback sets Callback value.
func (v *CallbackOrRef) WithCallback(val Callback) *CallbackOrRef {
	v.Callback = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *CallbackOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.CallbackReference)
	if err != nil {
		v.CallbackReference = nil
	}

	err = json.Unmarshal(data, &v.Callback)
	if err != nil {
		v.Callback = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v CallbackOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.CallbackReference, v.Callback)
}

// Paths structure is generated from "#/definitions/Paths".
type Paths struct {
	MapOfPathItemValues map[string]PathItem    `json:"-"` // Key must match pattern: `^\/`.
	MapOfAnything       map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithMapOfPathItemValues sets MapOfPathItemValues value.
func (v *Paths) WithMapOfPathItemValues(val map[string]PathItem) *Paths {
	v.MapOfPathItemValues = val
	return v
}

// WithMapOfPathItemValuesItem sets MapOfPathItemValues item value.
func (v *Paths) WithMapOfPathItemValuesItem(key string, val PathItem) *Paths {
	if v.MapOfPathItemValues == nil {
		v.MapOfPathItemValues = make(map[string]PathItem, 1)
	}

	v.MapOfPathItemValues[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Paths) WithMapOfAnything(val map[string]interface{}) *Paths {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Paths) WithMapOfAnythingItem(key string, val interface{}) *Paths {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *Paths) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regex.MatchString(key) {
			matched = true

			if v.MapOfPathItemValues == nil {
				v.MapOfPathItemValues = make(map[string]PathItem, 1)
			}

			var val PathItem

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfPathItemValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if v.MapOfAnything == nil {
				v.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v Paths) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfPathItemValues, v.MapOfAnything)
}

// Components structure is generated from "#/definitions/Components".
type Components struct {
	Schemas         *ComponentsSchemas         `json:"schemas,omitempty"`
	Responses       *ComponentsResponses       `json:"responses,omitempty"`
	Parameters      *ComponentsParameters      `json:"parameters,omitempty"`
	Examples        *ComponentsExamples        `json:"examples,omitempty"`
	RequestBodies   *ComponentsRequestBodies   `json:"requestBodies,omitempty"`
	Headers         *ComponentsHeaders         `json:"headers,omitempty"`
	SecuritySchemes *ComponentsSecuritySchemes `json:"securitySchemes,omitempty"`
	Links           *ComponentsLinks           `json:"links,omitempty"`
	Callbacks       *ComponentsCallbacks       `json:"callbacks,omitempty"`
	MapOfAnything   map[string]interface{}     `json:"-"`                         // Key must match pattern: `^x-`.
}

// WithSchemas sets Schemas value.
func (v *Components) WithSchemas(val ComponentsSchemas) *Components {
	v.Schemas = &val
	return v
}

// WithResponses sets Responses value.
func (v *Components) WithResponses(val ComponentsResponses) *Components {
	v.Responses = &val
	return v
}

// WithParameters sets Parameters value.
func (v *Components) WithParameters(val ComponentsParameters) *Components {
	v.Parameters = &val
	return v
}

// WithExamples sets Examples value.
func (v *Components) WithExamples(val ComponentsExamples) *Components {
	v.Examples = &val
	return v
}

// WithRequestBodies sets RequestBodies value.
func (v *Components) WithRequestBodies(val ComponentsRequestBodies) *Components {
	v.RequestBodies = &val
	return v
}

// WithHeaders sets Headers value.
func (v *Components) WithHeaders(val ComponentsHeaders) *Components {
	v.Headers = &val
	return v
}

// WithSecuritySchemes sets SecuritySchemes value.
func (v *Components) WithSecuritySchemes(val ComponentsSecuritySchemes) *Components {
	v.SecuritySchemes = &val
	return v
}

// WithLinks sets Links value.
func (v *Components) WithLinks(val ComponentsLinks) *Components {
	v.Links = &val
	return v
}

// WithCallbacks sets Callbacks value.
func (v *Components) WithCallbacks(val ComponentsCallbacks) *Components {
	v.Callbacks = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *Components) WithMapOfAnything(val map[string]interface{}) *Components {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *Components) WithMapOfAnythingItem(key string, val interface{}) *Components {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalComponents Components

var ignoreKeysComponents = []string{
	"schemas",
	"responses",
	"parameters",
	"examples",
	"requestBodies",
	"headers",
	"securitySchemes",
	"links",
	"callbacks",
}

// UnmarshalJSON decodes JSON.
func (v *Components) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalComponents(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysComponents {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = Components(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v Components) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponents(v), v.MapOfAnything)
}

// ComponentsSchemas structure is generated from "#/definitions/Components->schemas".
type ComponentsSchemas struct {
	MapOfSchemaOrRefValues map[string]SchemaOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties   map[string]interface{} `json:"-"` // All unmatched properties.
}

// WithMapOfSchemaOrRefValues sets MapOfSchemaOrRefValues value.
func (v *ComponentsSchemas) WithMapOfSchemaOrRefValues(val map[string]SchemaOrRef) *ComponentsSchemas {
	v.MapOfSchemaOrRefValues = val
	return v
}

// WithMapOfSchemaOrRefValuesItem sets MapOfSchemaOrRefValues item value.
func (v *ComponentsSchemas) WithMapOfSchemaOrRefValuesItem(key string, val SchemaOrRef) *ComponentsSchemas {
	if v.MapOfSchemaOrRefValues == nil {
		v.MapOfSchemaOrRefValues = make(map[string]SchemaOrRef, 1)
	}

	v.MapOfSchemaOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsSchemas) WithAdditionalProperties(val map[string]interface{}) *ComponentsSchemas {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsSchemas) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsSchemas {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsSchemas) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfSchemaOrRefValues == nil {
				v.MapOfSchemaOrRefValues = make(map[string]SchemaOrRef, 1)
			}

			var val SchemaOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfSchemaOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsSchemas) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfSchemaOrRefValues, v.AdditionalProperties)
}

// ComponentsResponses structure is generated from "#/definitions/Components->responses".
type ComponentsResponses struct {
	MapOfResponseOrRefValues map[string]ResponseOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties     map[string]interface{}   `json:"-"` // All unmatched properties.
}

// WithMapOfResponseOrRefValues sets MapOfResponseOrRefValues value.
func (v *ComponentsResponses) WithMapOfResponseOrRefValues(val map[string]ResponseOrRef) *ComponentsResponses {
	v.MapOfResponseOrRefValues = val
	return v
}

// WithMapOfResponseOrRefValuesItem sets MapOfResponseOrRefValues item value.
func (v *ComponentsResponses) WithMapOfResponseOrRefValuesItem(key string, val ResponseOrRef) *ComponentsResponses {
	if v.MapOfResponseOrRefValues == nil {
		v.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
	}

	v.MapOfResponseOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsResponses) WithAdditionalProperties(val map[string]interface{}) *ComponentsResponses {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsResponses) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsResponses {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsResponses) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfResponseOrRefValues == nil {
				v.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
			}

			var val ResponseOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfResponseOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsResponses) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfResponseOrRefValues, v.AdditionalProperties)
}

// ComponentsParameters structure is generated from "#/definitions/Components->parameters".
type ComponentsParameters struct {
	MapOfParameterOrRefValues map[string]ParameterOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties      map[string]interface{}    `json:"-"` // All unmatched properties.
}

// WithMapOfParameterOrRefValues sets MapOfParameterOrRefValues value.
func (v *ComponentsParameters) WithMapOfParameterOrRefValues(val map[string]ParameterOrRef) *ComponentsParameters {
	v.MapOfParameterOrRefValues = val
	return v
}

// WithMapOfParameterOrRefValuesItem sets MapOfParameterOrRefValues item value.
func (v *ComponentsParameters) WithMapOfParameterOrRefValuesItem(key string, val ParameterOrRef) *ComponentsParameters {
	if v.MapOfParameterOrRefValues == nil {
		v.MapOfParameterOrRefValues = make(map[string]ParameterOrRef, 1)
	}

	v.MapOfParameterOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsParameters) WithAdditionalProperties(val map[string]interface{}) *ComponentsParameters {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsParameters) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsParameters {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsParameters) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfParameterOrRefValues == nil {
				v.MapOfParameterOrRefValues = make(map[string]ParameterOrRef, 1)
			}

			var val ParameterOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfParameterOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsParameters) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfParameterOrRefValues, v.AdditionalProperties)
}

// ComponentsExamples structure is generated from "#/definitions/Components->examples".
type ComponentsExamples struct {
	MapOfExampleOrRefValues map[string]ExampleOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties    map[string]interface{}  `json:"-"` // All unmatched properties.
}

// WithMapOfExampleOrRefValues sets MapOfExampleOrRefValues value.
func (v *ComponentsExamples) WithMapOfExampleOrRefValues(val map[string]ExampleOrRef) *ComponentsExamples {
	v.MapOfExampleOrRefValues = val
	return v
}

// WithMapOfExampleOrRefValuesItem sets MapOfExampleOrRefValues item value.
func (v *ComponentsExamples) WithMapOfExampleOrRefValuesItem(key string, val ExampleOrRef) *ComponentsExamples {
	if v.MapOfExampleOrRefValues == nil {
		v.MapOfExampleOrRefValues = make(map[string]ExampleOrRef, 1)
	}

	v.MapOfExampleOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsExamples) WithAdditionalProperties(val map[string]interface{}) *ComponentsExamples {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsExamples) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsExamples {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsExamples) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfExampleOrRefValues == nil {
				v.MapOfExampleOrRefValues = make(map[string]ExampleOrRef, 1)
			}

			var val ExampleOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfExampleOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsExamples) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfExampleOrRefValues, v.AdditionalProperties)
}

// ComponentsRequestBodies structure is generated from "#/definitions/Components->requestBodies".
type ComponentsRequestBodies struct {
	MapOfRequestBodyOrRefValues map[string]RequestBodyOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties        map[string]interface{}      `json:"-"` // All unmatched properties.
}

// WithMapOfRequestBodyOrRefValues sets MapOfRequestBodyOrRefValues value.
func (v *ComponentsRequestBodies) WithMapOfRequestBodyOrRefValues(val map[string]RequestBodyOrRef) *ComponentsRequestBodies {
	v.MapOfRequestBodyOrRefValues = val
	return v
}

// WithMapOfRequestBodyOrRefValuesItem sets MapOfRequestBodyOrRefValues item value.
func (v *ComponentsRequestBodies) WithMapOfRequestBodyOrRefValuesItem(key string, val RequestBodyOrRef) *ComponentsRequestBodies {
	if v.MapOfRequestBodyOrRefValues == nil {
		v.MapOfRequestBodyOrRefValues = make(map[string]RequestBodyOrRef, 1)
	}

	v.MapOfRequestBodyOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsRequestBodies) WithAdditionalProperties(val map[string]interface{}) *ComponentsRequestBodies {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsRequestBodies) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsRequestBodies {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsRequestBodies) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfRequestBodyOrRefValues == nil {
				v.MapOfRequestBodyOrRefValues = make(map[string]RequestBodyOrRef, 1)
			}

			var val RequestBodyOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfRequestBodyOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsRequestBodies) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfRequestBodyOrRefValues, v.AdditionalProperties)
}

// ComponentsHeaders structure is generated from "#/definitions/Components->headers".
type ComponentsHeaders struct {
	MapOfHeaderOrRefValues map[string]HeaderOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties   map[string]interface{} `json:"-"` // All unmatched properties.
}

// WithMapOfHeaderOrRefValues sets MapOfHeaderOrRefValues value.
func (v *ComponentsHeaders) WithMapOfHeaderOrRefValues(val map[string]HeaderOrRef) *ComponentsHeaders {
	v.MapOfHeaderOrRefValues = val
	return v
}

// WithMapOfHeaderOrRefValuesItem sets MapOfHeaderOrRefValues item value.
func (v *ComponentsHeaders) WithMapOfHeaderOrRefValuesItem(key string, val HeaderOrRef) *ComponentsHeaders {
	if v.MapOfHeaderOrRefValues == nil {
		v.MapOfHeaderOrRefValues = make(map[string]HeaderOrRef, 1)
	}

	v.MapOfHeaderOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsHeaders) WithAdditionalProperties(val map[string]interface{}) *ComponentsHeaders {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsHeaders) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsHeaders {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsHeaders) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfHeaderOrRefValues == nil {
				v.MapOfHeaderOrRefValues = make(map[string]HeaderOrRef, 1)
			}

			var val HeaderOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfHeaderOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfHeaderOrRefValues, v.AdditionalProperties)
}

// SecuritySchemeReference structure is generated from "#/definitions/SecuritySchemeReference".
type SecuritySchemeReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (v *SecuritySchemeReference) WithRef(val string) *SecuritySchemeReference {
	v.Ref = val
	return v
}

// APIKeySecurityScheme structure is generated from "#/definitions/APIKeySecurityScheme".
type APIKeySecurityScheme struct {
	Name          string                 `json:"name"`                  // Required.
	In            APIKeySecuritySchemeIn `json:"in"`                    // Required.
	Description   *string                `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (v *APIKeySecurityScheme) WithName(val string) *APIKeySecurityScheme {
	v.Name = val
	return v
}

// WithIn sets In value.
func (v *APIKeySecurityScheme) WithIn(val APIKeySecuritySchemeIn) *APIKeySecurityScheme {
	v.In = val
	return v
}

// WithDescription sets Description value.
func (v *APIKeySecurityScheme) WithDescription(val string) *APIKeySecurityScheme {
	v.Description = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *APIKeySecurityScheme) WithMapOfAnything(val map[string]interface{}) *APIKeySecurityScheme {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *APIKeySecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *APIKeySecurityScheme {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalAPIKeySecurityScheme APIKeySecurityScheme

var ignoreKeysAPIKeySecurityScheme = []string{
	"name",
	"in",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *APIKeySecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalAPIKeySecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"apiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("apiKey" expected, %s received)`, v)
	}

	delete(m, "type")

	for _, key := range ignoreKeysAPIKeySecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = APIKeySecurityScheme(vv)

	return nil
}

var (
	// constAPIKeySecurityScheme is unconditionally added to JSON.
	constAPIKeySecurityScheme = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON.
func (v APIKeySecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKeySecurityScheme, marshalAPIKeySecurityScheme(v), v.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	Scheme        string                    `json:"scheme"`                 // Required.
	BearerFormat  *string                   `json:"bearerFormat,omitempty"`
	Description   *string                   `json:"description,omitempty"`
	OneOf0        *HTTPSecuritySchemeOneOf0 `json:"-"`
	OneOf1        *HTTPSecuritySchemeOneOf1 `json:"-"`
	MapOfAnything map[string]interface{}    `json:"-"`                      // Key must match pattern: `^x-`.
}

// WithScheme sets Scheme value.
func (v *HTTPSecurityScheme) WithScheme(val string) *HTTPSecurityScheme {
	v.Scheme = val
	return v
}

// WithBearerFormat sets BearerFormat value.
func (v *HTTPSecurityScheme) WithBearerFormat(val string) *HTTPSecurityScheme {
	v.BearerFormat = &val
	return v
}

// WithDescription sets Description value.
func (v *HTTPSecurityScheme) WithDescription(val string) *HTTPSecurityScheme {
	v.Description = &val
	return v
}

// WithOneOf0 sets OneOf0 value.
func (v *HTTPSecurityScheme) WithOneOf0(val HTTPSecuritySchemeOneOf0) *HTTPSecurityScheme {
	v.OneOf0 = &val
	return v
}

// WithOneOf1 sets OneOf1 value.
func (v *HTTPSecurityScheme) WithOneOf1(val HTTPSecuritySchemeOneOf1) *HTTPSecurityScheme {
	v.OneOf1 = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *HTTPSecurityScheme) WithMapOfAnything(val map[string]interface{}) *HTTPSecurityScheme {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *HTTPSecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *HTTPSecurityScheme {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalHTTPSecurityScheme HTTPSecurityScheme

var ignoreKeysHTTPSecurityScheme = []string{
	"scheme",
	"bearerFormat",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalHTTPSecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &vv.OneOf0)
	if err != nil {
		vv.OneOf0 = nil
	}

	err = json.Unmarshal(data, &vv.OneOf1)
	if err != nil {
		vv.OneOf1 = nil
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %s received)`, v)
	}

	delete(m, "type")

	for _, key := range ignoreKeysHTTPSecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = HTTPSecurityScheme(vv)

	return nil
}

var (
	// constHTTPSecurityScheme is unconditionally added to JSON.
	constHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)
)

// MarshalJSON encodes JSON.
func (v HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHTTPSecurityScheme, marshalHTTPSecurityScheme(v), v.MapOfAnything, v.OneOf0, v.OneOf1)
}

// HTTPSecuritySchemeOneOf0 structure is generated from "#/definitions/HTTPSecurityScheme/oneOf/0".
//
// Bearer.
//
// Bearer.
type HTTPSecuritySchemeOneOf0 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *HTTPSecuritySchemeOneOf0) WithAdditionalProperties(val map[string]interface{}) *HTTPSecuritySchemeOneOf0 {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *HTTPSecuritySchemeOneOf0) WithAdditionalPropertiesItem(key string, val interface{}) *HTTPSecuritySchemeOneOf0 {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *HTTPSecuritySchemeOneOf0) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["scheme"]; !ok || string(v) != `"bearer"` {
		return fmt.Errorf(`bad or missing const value for "scheme" ("bearer" expected, %s received)`, v)
	}

	delete(m, "scheme")

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constHTTPSecuritySchemeOneOf0 is unconditionally added to JSON.
	constHTTPSecuritySchemeOneOf0 = json.RawMessage(`{"scheme":"bearer"}`)
)

// MarshalJSON encodes JSON.
func (v HTTPSecuritySchemeOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHTTPSecuritySchemeOneOf0, v.AdditionalProperties)
}

// HTTPSecuritySchemeOneOf1 structure is generated from "#/definitions/HTTPSecurityScheme/oneOf/1".
//
// Non Bearer.
//
// Non Bearer.
type HTTPSecuritySchemeOneOf1 struct {
	Scheme               *interface{}           `json:"scheme,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                // All unmatched properties.
}

// WithScheme sets Scheme value.
func (v *HTTPSecuritySchemeOneOf1) WithScheme(val interface{}) *HTTPSecuritySchemeOneOf1 {
	v.Scheme = &val
	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *HTTPSecuritySchemeOneOf1) WithAdditionalProperties(val map[string]interface{}) *HTTPSecuritySchemeOneOf1 {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *HTTPSecuritySchemeOneOf1) WithAdditionalPropertiesItem(key string, val interface{}) *HTTPSecuritySchemeOneOf1 {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

type marshalHTTPSecuritySchemeOneOf1 HTTPSecuritySchemeOneOf1

var ignoreKeysHTTPSecuritySchemeOneOf1 = []string{
	"scheme",
}

// UnmarshalJSON decodes JSON.
func (v *HTTPSecuritySchemeOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalHTTPSecuritySchemeOneOf1(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if vv.Scheme == nil {
		if _, ok := m["scheme"]; ok {
			var v interface{}
			vv.Scheme = &v
		}
	}

	for _, key := range ignoreKeysHTTPSecuritySchemeOneOf1 {
		delete(m, key)
	}

	for key, rawValue := range m {
		if vv.AdditionalProperties == nil {
			vv.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		vv.AdditionalProperties[key] = val
	}

	*v = HTTPSecuritySchemeOneOf1(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v HTTPSecuritySchemeOneOf1) MarshalJSON() ([]byte, error) {
	if len(v.AdditionalProperties) == 0 {
		return json.Marshal(marshalHTTPSecuritySchemeOneOf1(v))
	}

	return marshalUnion(marshalHTTPSecuritySchemeOneOf1(v), v.AdditionalProperties)
}

// OAuth2SecurityScheme structure is generated from "#/definitions/OAuth2SecurityScheme".
type OAuth2SecurityScheme struct {
	Flows         OAuthFlows             `json:"flows"`                 // Required.
	Description   *string                `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithFlows sets Flows value.
func (v *OAuth2SecurityScheme) WithFlows(val OAuthFlows) *OAuth2SecurityScheme {
	v.Flows = val
	return v
}

// WithDescription sets Description value.
func (v *OAuth2SecurityScheme) WithDescription(val string) *OAuth2SecurityScheme {
	v.Description = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *OAuth2SecurityScheme) WithMapOfAnything(val map[string]interface{}) *OAuth2SecurityScheme {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *OAuth2SecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *OAuth2SecurityScheme {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalOAuth2SecurityScheme OAuth2SecurityScheme

var ignoreKeysOAuth2SecurityScheme = []string{
	"flows",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *OAuth2SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOAuth2SecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"oauth2"` {
		return fmt.Errorf(`bad or missing const value for "type" ("oauth2" expected, %s received)`, v)
	}

	delete(m, "type")

	for _, key := range ignoreKeysOAuth2SecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = OAuth2SecurityScheme(vv)

	return nil
}

var (
	// constOAuth2SecurityScheme is unconditionally added to JSON.
	constOAuth2SecurityScheme = json.RawMessage(`{"type":"oauth2"}`)
)

// MarshalJSON encodes JSON.
func (v OAuth2SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOAuth2SecurityScheme, marshalOAuth2SecurityScheme(v), v.MapOfAnything)
}

// OAuthFlows structure is generated from "#/definitions/OAuthFlows".
type OAuthFlows struct {
	Implicit          *ImplicitOAuthFlow          `json:"implicit,omitempty"`
	Password          *PasswordOAuthFlow          `json:"password,omitempty"`
	ClientCredentials *ClientCredentialsFlow      `json:"clientCredentials,omitempty"`
	AuthorizationCode *AuthorizationCodeOAuthFlow `json:"authorizationCode,omitempty"`
	MapOfAnything     map[string]interface{}      `json:"-"`                           // Key must match pattern: `^x-`.
}

// WithImplicit sets Implicit value.
func (v *OAuthFlows) WithImplicit(val ImplicitOAuthFlow) *OAuthFlows {
	v.Implicit = &val
	return v
}

// WithPassword sets Password value.
func (v *OAuthFlows) WithPassword(val PasswordOAuthFlow) *OAuthFlows {
	v.Password = &val
	return v
}

// WithClientCredentials sets ClientCredentials value.
func (v *OAuthFlows) WithClientCredentials(val ClientCredentialsFlow) *OAuthFlows {
	v.ClientCredentials = &val
	return v
}

// WithAuthorizationCode sets AuthorizationCode value.
func (v *OAuthFlows) WithAuthorizationCode(val AuthorizationCodeOAuthFlow) *OAuthFlows {
	v.AuthorizationCode = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *OAuthFlows) WithMapOfAnything(val map[string]interface{}) *OAuthFlows {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *OAuthFlows) WithMapOfAnythingItem(key string, val interface{}) *OAuthFlows {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalOAuthFlows OAuthFlows

var ignoreKeysOAuthFlows = []string{
	"implicit",
	"password",
	"clientCredentials",
	"authorizationCode",
}

// UnmarshalJSON decodes JSON.
func (v *OAuthFlows) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOAuthFlows(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysOAuthFlows {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = OAuthFlows(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v OAuthFlows) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOAuthFlows(v), v.MapOfAnything)
}

// ImplicitOAuthFlow structure is generated from "#/definitions/ImplicitOAuthFlow".
type ImplicitOAuthFlow struct {
	// Format: uri-reference.
	// Required.
	AuthorizationURL string                 `json:"authorizationUrl"`
	RefreshURL       *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes           map[string]string      `json:"scopes"`               // Required.
	MapOfAnything    map[string]interface{} `json:"-"`                    // Key must match pattern: `^x-`.
}

// WithAuthorizationURL sets AuthorizationURL value.
func (v *ImplicitOAuthFlow) WithAuthorizationURL(val string) *ImplicitOAuthFlow {
	v.AuthorizationURL = val
	return v
}

// WithRefreshURL sets RefreshURL value.
func (v *ImplicitOAuthFlow) WithRefreshURL(val string) *ImplicitOAuthFlow {
	v.RefreshURL = &val
	return v
}

// WithScopes sets Scopes value.
func (v *ImplicitOAuthFlow) WithScopes(val map[string]string) *ImplicitOAuthFlow {
	v.Scopes = val
	return v
}

// WithScopesItem sets Scopes item value.
func (v *ImplicitOAuthFlow) WithScopesItem(key string, val string) *ImplicitOAuthFlow {
	if v.Scopes == nil {
		v.Scopes = make(map[string]string, 1)
	}

	v.Scopes[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *ImplicitOAuthFlow) WithMapOfAnything(val map[string]interface{}) *ImplicitOAuthFlow {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *ImplicitOAuthFlow) WithMapOfAnythingItem(key string, val interface{}) *ImplicitOAuthFlow {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalImplicitOAuthFlow ImplicitOAuthFlow

var ignoreKeysImplicitOAuthFlow = []string{
	"authorizationUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (v *ImplicitOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalImplicitOAuthFlow(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysImplicitOAuthFlow {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ImplicitOAuthFlow(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ImplicitOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalImplicitOAuthFlow(v), v.MapOfAnything)
}

// PasswordOAuthFlow structure is generated from "#/definitions/PasswordOAuthFlow".
type PasswordOAuthFlow struct {
	// Format: uri-reference.
	// Required.
	TokenURL      string                 `json:"tokenUrl"`
	RefreshURL    *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                    // Key must match pattern: `^x-`.
}

// WithTokenURL sets TokenURL value.
func (v *PasswordOAuthFlow) WithTokenURL(val string) *PasswordOAuthFlow {
	v.TokenURL = val
	return v
}

// WithRefreshURL sets RefreshURL value.
func (v *PasswordOAuthFlow) WithRefreshURL(val string) *PasswordOAuthFlow {
	v.RefreshURL = &val
	return v
}

// WithScopes sets Scopes value.
func (v *PasswordOAuthFlow) WithScopes(val map[string]string) *PasswordOAuthFlow {
	v.Scopes = val
	return v
}

// WithScopesItem sets Scopes item value.
func (v *PasswordOAuthFlow) WithScopesItem(key string, val string) *PasswordOAuthFlow {
	if v.Scopes == nil {
		v.Scopes = make(map[string]string, 1)
	}

	v.Scopes[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *PasswordOAuthFlow) WithMapOfAnything(val map[string]interface{}) *PasswordOAuthFlow {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *PasswordOAuthFlow) WithMapOfAnythingItem(key string, val interface{}) *PasswordOAuthFlow {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalPasswordOAuthFlow PasswordOAuthFlow

var ignoreKeysPasswordOAuthFlow = []string{
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (v *PasswordOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalPasswordOAuthFlow(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysPasswordOAuthFlow {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = PasswordOAuthFlow(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v PasswordOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPasswordOAuthFlow(v), v.MapOfAnything)
}

// ClientCredentialsFlow structure is generated from "#/definitions/ClientCredentialsFlow".
type ClientCredentialsFlow struct {
	// Format: uri-reference.
	// Required.
	TokenURL      string                 `json:"tokenUrl"`
	RefreshURL    *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                    // Key must match pattern: `^x-`.
}

// WithTokenURL sets TokenURL value.
func (v *ClientCredentialsFlow) WithTokenURL(val string) *ClientCredentialsFlow {
	v.TokenURL = val
	return v
}

// WithRefreshURL sets RefreshURL value.
func (v *ClientCredentialsFlow) WithRefreshURL(val string) *ClientCredentialsFlow {
	v.RefreshURL = &val
	return v
}

// WithScopes sets Scopes value.
func (v *ClientCredentialsFlow) WithScopes(val map[string]string) *ClientCredentialsFlow {
	v.Scopes = val
	return v
}

// WithScopesItem sets Scopes item value.
func (v *ClientCredentialsFlow) WithScopesItem(key string, val string) *ClientCredentialsFlow {
	if v.Scopes == nil {
		v.Scopes = make(map[string]string, 1)
	}

	v.Scopes[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *ClientCredentialsFlow) WithMapOfAnything(val map[string]interface{}) *ClientCredentialsFlow {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *ClientCredentialsFlow) WithMapOfAnythingItem(key string, val interface{}) *ClientCredentialsFlow {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalClientCredentialsFlow ClientCredentialsFlow

var ignoreKeysClientCredentialsFlow = []string{
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (v *ClientCredentialsFlow) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalClientCredentialsFlow(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysClientCredentialsFlow {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = ClientCredentialsFlow(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v ClientCredentialsFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalClientCredentialsFlow(v), v.MapOfAnything)
}

// AuthorizationCodeOAuthFlow structure is generated from "#/definitions/AuthorizationCodeOAuthFlow".
type AuthorizationCodeOAuthFlow struct {
	// Format: uri-reference.
	// Required.
	AuthorizationURL string                 `json:"authorizationUrl"`
	// Format: uri-reference.
	// Required.
	TokenURL         string                 `json:"tokenUrl"`
	RefreshURL       *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes           map[string]string      `json:"scopes,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                    // Key must match pattern: `^x-`.
}

// WithAuthorizationURL sets AuthorizationURL value.
func (v *AuthorizationCodeOAuthFlow) WithAuthorizationURL(val string) *AuthorizationCodeOAuthFlow {
	v.AuthorizationURL = val
	return v
}

// WithTokenURL sets TokenURL value.
func (v *AuthorizationCodeOAuthFlow) WithTokenURL(val string) *AuthorizationCodeOAuthFlow {
	v.TokenURL = val
	return v
}

// WithRefreshURL sets RefreshURL value.
func (v *AuthorizationCodeOAuthFlow) WithRefreshURL(val string) *AuthorizationCodeOAuthFlow {
	v.RefreshURL = &val
	return v
}

// WithScopes sets Scopes value.
func (v *AuthorizationCodeOAuthFlow) WithScopes(val map[string]string) *AuthorizationCodeOAuthFlow {
	v.Scopes = val
	return v
}

// WithScopesItem sets Scopes item value.
func (v *AuthorizationCodeOAuthFlow) WithScopesItem(key string, val string) *AuthorizationCodeOAuthFlow {
	if v.Scopes == nil {
		v.Scopes = make(map[string]string, 1)
	}

	v.Scopes[key] = val

	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *AuthorizationCodeOAuthFlow) WithMapOfAnything(val map[string]interface{}) *AuthorizationCodeOAuthFlow {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *AuthorizationCodeOAuthFlow) WithMapOfAnythingItem(key string, val interface{}) *AuthorizationCodeOAuthFlow {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalAuthorizationCodeOAuthFlow AuthorizationCodeOAuthFlow

var ignoreKeysAuthorizationCodeOAuthFlow = []string{
	"authorizationUrl",
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (v *AuthorizationCodeOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalAuthorizationCodeOAuthFlow(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysAuthorizationCodeOAuthFlow {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = AuthorizationCodeOAuthFlow(vv)

	return nil
}

// MarshalJSON encodes JSON.
func (v AuthorizationCodeOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAuthorizationCodeOAuthFlow(v), v.MapOfAnything)
}

// OpenIDConnectSecurityScheme structure is generated from "#/definitions/OpenIdConnectSecurityScheme".
type OpenIDConnectSecurityScheme struct {
	// Format: uri-reference.
	// Required.
	OpenIDConnectURL string                 `json:"openIdConnectUrl"`
	Description      *string                `json:"description,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                     // Key must match pattern: `^x-`.
}

// WithOpenIDConnectURL sets OpenIDConnectURL value.
func (v *OpenIDConnectSecurityScheme) WithOpenIDConnectURL(val string) *OpenIDConnectSecurityScheme {
	v.OpenIDConnectURL = val
	return v
}

// WithDescription sets Description value.
func (v *OpenIDConnectSecurityScheme) WithDescription(val string) *OpenIDConnectSecurityScheme {
	v.Description = &val
	return v
}

// WithMapOfAnything sets MapOfAnything value.
func (v *OpenIDConnectSecurityScheme) WithMapOfAnything(val map[string]interface{}) *OpenIDConnectSecurityScheme {
	v.MapOfAnything = val
	return v
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (v *OpenIDConnectSecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *OpenIDConnectSecurityScheme {
	if v.MapOfAnything == nil {
		v.MapOfAnything = make(map[string]interface{}, 1)
	}

	v.MapOfAnything[key] = val

	return v
}

type marshalOpenIDConnectSecurityScheme OpenIDConnectSecurityScheme

var ignoreKeysOpenIDConnectSecurityScheme = []string{
	"openIdConnectUrl",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (v *OpenIDConnectSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	vv := marshalOpenIDConnectSecurityScheme(*v)

	err = json.Unmarshal(data, &vv)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if v, ok := m["type"]; !ok || string(v) != `"openIdConnect"` {
		return fmt.Errorf(`bad or missing const value for "type" ("openIdConnect" expected, %s received)`, v)
	}

	delete(m, "type")

	for _, key := range ignoreKeysOpenIDConnectSecurityScheme {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if vv.MapOfAnything == nil {
				vv.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			vv.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*v = OpenIDConnectSecurityScheme(vv)

	return nil
}

var (
	// constOpenIDConnectSecurityScheme is unconditionally added to JSON.
	constOpenIDConnectSecurityScheme = json.RawMessage(`{"type":"openIdConnect"}`)
)

// MarshalJSON encodes JSON.
func (v OpenIDConnectSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOpenIDConnectSecurityScheme, marshalOpenIDConnectSecurityScheme(v), v.MapOfAnything)
}

// SecurityScheme structure is generated from "#/definitions/SecurityScheme".
type SecurityScheme struct {
	APIKeySecurityScheme        *APIKeySecurityScheme        `json:"-"`
	HTTPSecurityScheme          *HTTPSecurityScheme          `json:"-"`
	OAuth2SecurityScheme        *OAuth2SecurityScheme        `json:"-"`
	OpenIDConnectSecurityScheme *OpenIDConnectSecurityScheme `json:"-"`
}

// WithAPIKeySecurityScheme sets APIKeySecurityScheme value.
func (v *SecurityScheme) WithAPIKeySecurityScheme(val APIKeySecurityScheme) *SecurityScheme {
	v.APIKeySecurityScheme = &val
	return v
}

// WithHTTPSecurityScheme sets HTTPSecurityScheme value.
func (v *SecurityScheme) WithHTTPSecurityScheme(val HTTPSecurityScheme) *SecurityScheme {
	v.HTTPSecurityScheme = &val
	return v
}

// WithOAuth2SecurityScheme sets OAuth2SecurityScheme value.
func (v *SecurityScheme) WithOAuth2SecurityScheme(val OAuth2SecurityScheme) *SecurityScheme {
	v.OAuth2SecurityScheme = &val
	return v
}

// WithOpenIDConnectSecurityScheme sets OpenIDConnectSecurityScheme value.
func (v *SecurityScheme) WithOpenIDConnectSecurityScheme(val OpenIDConnectSecurityScheme) *SecurityScheme {
	v.OpenIDConnectSecurityScheme = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.APIKeySecurityScheme)
	if err != nil {
		v.APIKeySecurityScheme = nil
	}

	err = json.Unmarshal(data, &v.HTTPSecurityScheme)
	if err != nil {
		v.HTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &v.OAuth2SecurityScheme)
	if err != nil {
		v.OAuth2SecurityScheme = nil
	}

	err = json.Unmarshal(data, &v.OpenIDConnectSecurityScheme)
	if err != nil {
		v.OpenIDConnectSecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.APIKeySecurityScheme, v.HTTPSecurityScheme, v.OAuth2SecurityScheme, v.OpenIDConnectSecurityScheme)
}

// SecuritySchemeOrRef structure is generated from "#/definitions/SecuritySchemeOrRef".
type SecuritySchemeOrRef struct {
	SecuritySchemeReference *SecuritySchemeReference `json:"-"`
	SecurityScheme          *SecurityScheme          `json:"-"`
}

// WithSecuritySchemeReference sets SecuritySchemeReference value.
func (v *SecuritySchemeOrRef) WithSecuritySchemeReference(val SecuritySchemeReference) *SecuritySchemeOrRef {
	v.SecuritySchemeReference = &val
	return v
}

// WithSecurityScheme sets SecurityScheme value.
func (v *SecuritySchemeOrRef) WithSecurityScheme(val SecurityScheme) *SecuritySchemeOrRef {
	v.SecurityScheme = &val
	return v
}

// UnmarshalJSON decodes JSON.
func (v *SecuritySchemeOrRef) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &v.SecuritySchemeReference)
	if err != nil {
		v.SecuritySchemeReference = nil
	}

	err = json.Unmarshal(data, &v.SecurityScheme)
	if err != nil {
		v.SecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v SecuritySchemeOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.SecuritySchemeReference, v.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/Components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfSecuritySchemeOrRefValues map[string]SecuritySchemeOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties           map[string]interface{}         `json:"-"` // All unmatched properties.
}

// WithMapOfSecuritySchemeOrRefValues sets MapOfSecuritySchemeOrRefValues value.
func (v *ComponentsSecuritySchemes) WithMapOfSecuritySchemeOrRefValues(val map[string]SecuritySchemeOrRef) *ComponentsSecuritySchemes {
	v.MapOfSecuritySchemeOrRefValues = val
	return v
}

// WithMapOfSecuritySchemeOrRefValuesItem sets MapOfSecuritySchemeOrRefValues item value.
func (v *ComponentsSecuritySchemes) WithMapOfSecuritySchemeOrRefValuesItem(key string, val SecuritySchemeOrRef) *ComponentsSecuritySchemes {
	if v.MapOfSecuritySchemeOrRefValues == nil {
		v.MapOfSecuritySchemeOrRefValues = make(map[string]SecuritySchemeOrRef, 1)
	}

	v.MapOfSecuritySchemeOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsSecuritySchemes) WithAdditionalProperties(val map[string]interface{}) *ComponentsSecuritySchemes {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsSecuritySchemes) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsSecuritySchemes {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfSecuritySchemeOrRefValues == nil {
				v.MapOfSecuritySchemeOrRefValues = make(map[string]SecuritySchemeOrRef, 1)
			}

			var val SecuritySchemeOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfSecuritySchemeOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfSecuritySchemeOrRefValues, v.AdditionalProperties)
}

// ComponentsLinks structure is generated from "#/definitions/Components->links".
type ComponentsLinks struct {
	MapOfLinkOrRefValues map[string]LinkOrRef   `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties.
}

// WithMapOfLinkOrRefValues sets MapOfLinkOrRefValues value.
func (v *ComponentsLinks) WithMapOfLinkOrRefValues(val map[string]LinkOrRef) *ComponentsLinks {
	v.MapOfLinkOrRefValues = val
	return v
}

// WithMapOfLinkOrRefValuesItem sets MapOfLinkOrRefValues item value.
func (v *ComponentsLinks) WithMapOfLinkOrRefValuesItem(key string, val LinkOrRef) *ComponentsLinks {
	if v.MapOfLinkOrRefValues == nil {
		v.MapOfLinkOrRefValues = make(map[string]LinkOrRef, 1)
	}

	v.MapOfLinkOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsLinks) WithAdditionalProperties(val map[string]interface{}) *ComponentsLinks {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsLinks) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsLinks {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsLinks) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfLinkOrRefValues == nil {
				v.MapOfLinkOrRefValues = make(map[string]LinkOrRef, 1)
			}

			var val LinkOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfLinkOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsLinks) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfLinkOrRefValues, v.AdditionalProperties)
}

// ComponentsCallbacks structure is generated from "#/definitions/Components->callbacks".
type ComponentsCallbacks struct {
	MapOfCallbackOrRefValues map[string]CallbackOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
	AdditionalProperties     map[string]interface{}   `json:"-"` // All unmatched properties.
}

// WithMapOfCallbackOrRefValues sets MapOfCallbackOrRefValues value.
func (v *ComponentsCallbacks) WithMapOfCallbackOrRefValues(val map[string]CallbackOrRef) *ComponentsCallbacks {
	v.MapOfCallbackOrRefValues = val
	return v
}

// WithMapOfCallbackOrRefValuesItem sets MapOfCallbackOrRefValues item value.
func (v *ComponentsCallbacks) WithMapOfCallbackOrRefValuesItem(key string, val CallbackOrRef) *ComponentsCallbacks {
	if v.MapOfCallbackOrRefValues == nil {
		v.MapOfCallbackOrRefValues = make(map[string]CallbackOrRef, 1)
	}

	v.MapOfCallbackOrRefValues[key] = val

	return v
}

// WithAdditionalProperties sets AdditionalProperties value.
func (v *ComponentsCallbacks) WithAdditionalProperties(val map[string]interface{}) *ComponentsCallbacks {
	v.AdditionalProperties = val
	return v
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (v *ComponentsCallbacks) WithAdditionalPropertiesItem(key string, val interface{}) *ComponentsCallbacks {
	if v.AdditionalProperties == nil {
		v.AdditionalProperties = make(map[string]interface{}, 1)
	}

	v.AdditionalProperties[key] = val

	return v
}

// UnmarshalJSON decodes JSON.
func (v *ComponentsCallbacks) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if v.MapOfCallbackOrRefValues == nil {
				v.MapOfCallbackOrRefValues = make(map[string]CallbackOrRef, 1)
			}

			var val CallbackOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			v.MapOfCallbackOrRefValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if v.AdditionalProperties == nil {
			v.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		v.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (v ComponentsCallbacks) MarshalJSON() ([]byte, error) {
	return marshalUnion(v.MapOfCallbackOrRefValues, v.AdditionalProperties)
}

// ParameterIn is an enum type.
type ParameterIn string

// ParameterIn values enumeration.
const (
	ParameterInPath = ParameterIn("path")
	ParameterInQuery = ParameterIn("query")
	ParameterInHeader = ParameterIn("header")
	ParameterInCookie = ParameterIn("cookie")
)

// MarshalJSON encodes JSON.
func (i ParameterIn) MarshalJSON() ([]byte, error) {
	switch i {
	case ParameterInPath:
	case ParameterInQuery:
	case ParameterInHeader:
	case ParameterInCookie:

	default:
		return nil, fmt.Errorf("unexpected ParameterIn value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ParameterIn) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := ParameterIn(ii)

	switch v {
	case ParameterInPath:
	case ParameterInQuery:
	case ParameterInHeader:
	case ParameterInCookie:

	default:
		return fmt.Errorf("unexpected ParameterIn value: %v", v)
	}

	*i = v

	return nil
}

// SchemaType is an enum type.
type SchemaType string

// SchemaType values enumeration.
const (
	SchemaTypeArray = SchemaType("array")
	SchemaTypeBoolean = SchemaType("boolean")
	SchemaTypeInteger = SchemaType("integer")
	SchemaTypeNumber = SchemaType("number")
	SchemaTypeObject = SchemaType("object")
	SchemaTypeString = SchemaType("string")
)

// MarshalJSON encodes JSON.
func (i SchemaType) MarshalJSON() ([]byte, error) {
	switch i {
	case SchemaTypeArray:
	case SchemaTypeBoolean:
	case SchemaTypeInteger:
	case SchemaTypeNumber:
	case SchemaTypeObject:
	case SchemaTypeString:

	default:
		return nil, fmt.Errorf("unexpected SchemaType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *SchemaType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := SchemaType(ii)

	switch v {
	case SchemaTypeArray:
	case SchemaTypeBoolean:
	case SchemaTypeInteger:
	case SchemaTypeNumber:
	case SchemaTypeObject:
	case SchemaTypeString:

	default:
		return fmt.Errorf("unexpected SchemaType value: %v", v)
	}

	*i = v

	return nil
}

// EncodingStyle is an enum type.
type EncodingStyle string

// EncodingStyle values enumeration.
const (
	EncodingStyleForm = EncodingStyle("form")
	EncodingStyleSpaceDelimited = EncodingStyle("spaceDelimited")
	EncodingStylePipeDelimited = EncodingStyle("pipeDelimited")
	EncodingStyleDeepObject = EncodingStyle("deepObject")
)

// MarshalJSON encodes JSON.
func (i EncodingStyle) MarshalJSON() ([]byte, error) {
	switch i {
	case EncodingStyleForm:
	case EncodingStyleSpaceDelimited:
	case EncodingStylePipeDelimited:
	case EncodingStyleDeepObject:

	default:
		return nil, fmt.Errorf("unexpected EncodingStyle value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *EncodingStyle) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := EncodingStyle(ii)

	switch v {
	case EncodingStyleForm:
	case EncodingStyleSpaceDelimited:
	case EncodingStylePipeDelimited:
	case EncodingStyleDeepObject:

	default:
		return fmt.Errorf("unexpected EncodingStyle value: %v", v)
	}

	*i = v

	return nil
}

// ParameterLocationOneOf0Style is an enum type.
type ParameterLocationOneOf0Style string

// ParameterLocationOneOf0Style values enumeration.
const (
	ParameterLocationOneOf0StyleMatrix = ParameterLocationOneOf0Style("matrix")
	ParameterLocationOneOf0StyleLabel = ParameterLocationOneOf0Style("label")
	ParameterLocationOneOf0StyleSimple = ParameterLocationOneOf0Style("simple")
)

// MarshalJSON encodes JSON.
func (i ParameterLocationOneOf0Style) MarshalJSON() ([]byte, error) {
	switch i {
	case ParameterLocationOneOf0StyleMatrix:
	case ParameterLocationOneOf0StyleLabel:
	case ParameterLocationOneOf0StyleSimple:

	default:
		return nil, fmt.Errorf("unexpected ParameterLocationOneOf0Style value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf0Style) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := ParameterLocationOneOf0Style(ii)

	switch v {
	case ParameterLocationOneOf0StyleMatrix:
	case ParameterLocationOneOf0StyleLabel:
	case ParameterLocationOneOf0StyleSimple:

	default:
		return fmt.Errorf("unexpected ParameterLocationOneOf0Style value: %v", v)
	}

	*i = v

	return nil
}

// ParameterLocationOneOf1Style is an enum type.
type ParameterLocationOneOf1Style string

// ParameterLocationOneOf1Style values enumeration.
const (
	ParameterLocationOneOf1StyleForm = ParameterLocationOneOf1Style("form")
	ParameterLocationOneOf1StyleSpaceDelimited = ParameterLocationOneOf1Style("spaceDelimited")
	ParameterLocationOneOf1StylePipeDelimited = ParameterLocationOneOf1Style("pipeDelimited")
	ParameterLocationOneOf1StyleDeepObject = ParameterLocationOneOf1Style("deepObject")
)

// MarshalJSON encodes JSON.
func (i ParameterLocationOneOf1Style) MarshalJSON() ([]byte, error) {
	switch i {
	case ParameterLocationOneOf1StyleForm:
	case ParameterLocationOneOf1StyleSpaceDelimited:
	case ParameterLocationOneOf1StylePipeDelimited:
	case ParameterLocationOneOf1StyleDeepObject:

	default:
		return nil, fmt.Errorf("unexpected ParameterLocationOneOf1Style value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf1Style) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := ParameterLocationOneOf1Style(ii)

	switch v {
	case ParameterLocationOneOf1StyleForm:
	case ParameterLocationOneOf1StyleSpaceDelimited:
	case ParameterLocationOneOf1StylePipeDelimited:
	case ParameterLocationOneOf1StyleDeepObject:

	default:
		return fmt.Errorf("unexpected ParameterLocationOneOf1Style value: %v", v)
	}

	*i = v

	return nil
}

// APIKeySecuritySchemeIn is an enum type.
type APIKeySecuritySchemeIn string

// APIKeySecuritySchemeIn values enumeration.
const (
	APIKeySecuritySchemeInHeader = APIKeySecuritySchemeIn("header")
	APIKeySecuritySchemeInQuery = APIKeySecuritySchemeIn("query")
	APIKeySecuritySchemeInCookie = APIKeySecuritySchemeIn("cookie")
)

// MarshalJSON encodes JSON.
func (i APIKeySecuritySchemeIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeySecuritySchemeInHeader:
	case APIKeySecuritySchemeInQuery:
	case APIKeySecuritySchemeInCookie:

	default:
		return nil, fmt.Errorf("unexpected APIKeySecuritySchemeIn value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *APIKeySecuritySchemeIn) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := APIKeySecuritySchemeIn(ii)

	switch v {
	case APIKeySecuritySchemeInHeader:
	case APIKeySecuritySchemeInQuery:
	case APIKeySecuritySchemeInCookie:

	default:
		return fmt.Errorf("unexpected APIKeySecuritySchemeIn value: %v", v)
	}

	*i = v

	return nil
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	isObject := true

	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		if string(j) == "{}" {
			continue
		}

		if string(j) == "null" {
			continue
		}

		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false

				continue
			}

			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}

		result = append(result, j[1:]...)
	}

	// Close empty result.
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
// Regular expressions for pattern properties.
var (
	regexX = regexp.MustCompile("^x-")
	regexGetPutPostDeleteOptionsHeadPatchTrace = regexp.MustCompile("^(get|put|post|delete|options|head|patch|trace)$")
	regex15D2XX = regexp.MustCompile(`^[1-5](?:\d{2}|XX)$`)
	regex = regexp.MustCompile(`^\/`)
	regexAZAZ09 = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
)
