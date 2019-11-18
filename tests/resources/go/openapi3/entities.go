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
	Openapi       string                 `json:"openapi,omitempty"`
	Info          *Info                  `json:"info,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	Servers       []Server               `json:"servers,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	Tags          []Tag                  `json:"tags,omitempty"`
	Paths         *Paths                 `json:"paths,omitempty"`
	Components    *Components            `json:"components,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
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
func (i *OpenAPI) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOpenAPI(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = OpenAPI(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i OpenAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOpenAPI(i), i.MapOfAnything)
}

// Info structure is generated from "#/definitions/Info".
type Info struct {
	Title          string                 `json:"title,omitempty"`
	Description    string                 `json:"description,omitempty"`
	TermsOfService string                 `json:"termsOfService,omitempty"`
	Contact        *Contact               `json:"contact,omitempty"`
	License        *License               `json:"license,omitempty"`
	Version        string                 `json:"version,omitempty"`
	MapOfAnything  map[string]interface{} `json:"-"`                        // Key must match pattern: ^x-
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
func (i *Info) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalInfo(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Info(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(i), i.MapOfAnything)
}

// Contact structure is generated from "#/definitions/Contact".
type Contact struct {
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url,omitempty"`
	Email         string                 `json:"email,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalContact Contact

var ignoreKeysContact = []string{
	"name",
	"url",
	"email",
}

// UnmarshalJSON decodes JSON.
func (i *Contact) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalContact(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Contact(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(i), i.MapOfAnything)
}

// License structure is generated from "#/definitions/License".
type License struct {
	Name          string                 `json:"name,omitempty"`
	URL           string                 `json:"url,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

type marshalLicense License

var ignoreKeysLicense = []string{
	"name",
	"url",
}

// UnmarshalJSON decodes JSON.
func (i *License) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalLicense(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = License(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(i), i.MapOfAnything)
}

// ExternalDocumentation structure is generated from "#/definitions/ExternalDocumentation".
type ExternalDocumentation struct {
	Description   string                 `json:"description,omitempty"`
	URL           string                 `json:"url,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalExternalDocumentation ExternalDocumentation

var ignoreKeysExternalDocumentation = []string{
	"description",
	"url",
}

// UnmarshalJSON decodes JSON.
func (i *ExternalDocumentation) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalExternalDocumentation(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = ExternalDocumentation(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ExternalDocumentation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocumentation(i), i.MapOfAnything)
}

// Server structure is generated from "#/definitions/Server".
type Server struct {
	URL           string                    `json:"url,omitempty"`
	Description   string                    `json:"description,omitempty"`
	Variables     map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnything map[string]interface{}    `json:"-"`                     // Key must match pattern: ^x-
}

type marshalServer Server

var ignoreKeysServer = []string{
	"url",
	"description",
	"variables",
}

// UnmarshalJSON decodes JSON.
func (i *Server) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalServer(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Server(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(i), i.MapOfAnything)
}

// ServerVariable structure is generated from "#/definitions/ServerVariable".
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalServerVariable ServerVariable

var ignoreKeysServerVariable = []string{
	"enum",
	"default",
	"description",
}

// UnmarshalJSON decodes JSON.
func (i *ServerVariable) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalServerVariable(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = ServerVariable(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(i), i.MapOfAnything)
}

// Tag structure is generated from "#/definitions/Tag".
type Tag struct {
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalTag Tag

var ignoreKeysTag = []string{
	"name",
	"description",
	"externalDocs",
}

// UnmarshalJSON decodes JSON.
func (i *Tag) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalTag(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Tag(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(i), i.MapOfAnything)
}

// PathItem structure is generated from "#/definitions/PathItem".
type PathItem struct {
	Ref                  string                    `json:"$ref,omitempty"`
	Summary              string                    `json:"summary,omitempty"`
	Description          string                    `json:"description,omitempty"`
	Servers              []Server                  `json:"servers,omitempty"`
	Parameters           []PathItemParametersItems `json:"parameters,omitempty"`
	MapOfOperationValues map[string]Operation      `json:"-"`                     // Key must match pattern: ^(get|put|post|delete|options|head|patch|trace)$
	MapOfAnything        map[string]interface{}    `json:"-"`                     // Key must match pattern: ^x-
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
func (i *PathItem) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalPathItem(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfOperationValues == nil {
				ii.MapOfOperationValues = make(map[string]Operation, 1)
			}

			var val Operation

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfOperationValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = PathItem(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i PathItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPathItem(i), i.MapOfOperationValues, i.MapOfAnything)
}

// Parameter structure is generated from "#/definitions/Parameter".
type Parameter struct {
	Name             string                                           `json:"name,omitempty"`
	In               string                                           `json:"in,omitempty"`
	Description      string                                           `json:"description,omitempty"`
	Required         bool                                             `json:"required,omitempty"`
	Deprecated       bool                                             `json:"deprecated,omitempty"`
	AllowEmptyValue  bool                                             `json:"allowEmptyValue,omitempty"`
	Style            string                                           `json:"style,omitempty"`
	Explode          bool                                             `json:"explode,omitempty"`
	AllowReserved    bool                                             `json:"allowReserved,omitempty"`
	Schema           *ParameterSchema                                 `json:"schema,omitempty"`
	Content          map[string]MediaType                             `json:"content,omitempty"`
	Example          *interface{}                                     `json:"example,omitempty"`
	Examples         map[string]ParameterExamplesAdditionalProperties `json:"examples,omitempty"`
	SchemaXORContent *SchemaXORContentOneOf1                          `json:"-"`
	Location         *ParameterLocation                               `json:"-"`
	MapOfAnything    map[string]interface{}                           `json:"-"`                         // Key must match pattern: ^x-
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
func (i *Parameter) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalParameter(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &ii.SchemaXORContent)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &ii.Location)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			ii.Example = &v
		}
	}

	for _, key := range ignoreKeysParameter {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Parameter(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(i), i.MapOfAnything, i.SchemaXORContent, i.Location)
}

// Schema structure is generated from "#/definitions/Schema".
type Schema struct {
	Title                string                                          `json:"title,omitempty"`
	MultipleOf           float64                                         `json:"multipleOf,omitempty"`
	Maximum              float64                                         `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                                            `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                         `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                                            `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                           `json:"maxLength,omitempty"`
	MinLength            int64                                           `json:"minLength,omitempty"`
	Pattern              string                                          `json:"pattern,omitempty"`
	MaxItems             int64                                           `json:"maxItems,omitempty"`
	MinItems             int64                                           `json:"minItems,omitempty"`
	UniqueItems          bool                                            `json:"uniqueItems,omitempty"`
	MaxProperties        int64                                           `json:"maxProperties,omitempty"`
	MinProperties        int64                                           `json:"minProperties,omitempty"`
	Required             []string                                        `json:"required,omitempty"`
	Enum                 []interface{}                                   `json:"enum,omitempty"`
	Type                 SchemaType                                      `json:"type,omitempty"`
	MapOfAnything        map[string]interface{}                          `json:"-"`                              // Key must match pattern: ^x-
	Not                  *SchemaNot                                      `json:"not,omitempty"`
	AllOf                []SchemaAllOfItems                              `json:"allOf,omitempty"`
	OneOf                []SchemaOneOfItems                              `json:"oneOf,omitempty"`
	AnyOf                []SchemaAnyOfItems                              `json:"anyOf,omitempty"`
	Items                *SchemaItems                                    `json:"items,omitempty"`
	Properties           map[string]SchemaPropertiesAdditionalProperties `json:"properties,omitempty"`
	AdditionalProperties *SchemaAdditionalProperties                     `json:"additionalProperties,omitempty"`
	Description          string                                          `json:"description,omitempty"`
	Format               string                                          `json:"format,omitempty"`
	Default              *interface{}                                    `json:"default,omitempty"`
	Nullable             bool                                            `json:"nullable,omitempty"`
	Discriminator        *Discriminator                                  `json:"discriminator,omitempty"`
	ReadOnly             bool                                            `json:"readOnly,omitempty"`
	WriteOnly            bool                                            `json:"writeOnly,omitempty"`
	Example              *interface{}                                    `json:"example,omitempty"`
	ExternalDocs         *ExternalDocumentation                          `json:"externalDocs,omitempty"`
	Deprecated           bool                                            `json:"deprecated,omitempty"`
	XML                  *XML                                            `json:"xml,omitempty"`
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
func (i *Schema) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalSchema(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Default == nil {
		if _, ok := m["default"]; ok {
			var v interface{}
			ii.Default = &v
		}
	}

	if ii.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			ii.Example = &v
		}
	}

	for _, key := range ignoreKeysSchema {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Schema(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchema(i), i.MapOfAnything)
}

// Reference structure is generated from "#/definitions/Reference".
type Reference struct {
	MapOfStringValues    map[string]string      `json:"-"` // Key must match pattern: ^\$ref$
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *Reference) UnmarshalJSON(data []byte) error {
	var err error

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for key, rawValue := range m {
		matched := false

		if regexRef.MatchString(key) {
			matched = true

			if i.MapOfStringValues == nil {
				i.MapOfStringValues = make(map[string]string, 1)
			}

			var val string

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfStringValues[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Reference) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfStringValues, i.AdditionalProperties)
}

// SchemaNot structure is generated from "#/definitions/Schema->not".
type SchemaNot struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaNot) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaNot) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// SchemaAllOfItems structure is generated from "#/definitions/Schema->allOf->items".
type SchemaAllOfItems struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaAllOfItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaAllOfItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// SchemaOneOfItems structure is generated from "#/definitions/Schema->oneOf->items".
type SchemaOneOfItems struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaOneOfItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaOneOfItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// SchemaAnyOfItems structure is generated from "#/definitions/Schema->anyOf->items".
type SchemaAnyOfItems struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaAnyOfItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaAnyOfItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// SchemaItems structure is generated from "#/definitions/Schema->items".
type SchemaItems struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// SchemaPropertiesAdditionalProperties structure is generated from "#/definitions/Schema->properties->additionalProperties".
type SchemaPropertiesAdditionalProperties struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaPropertiesAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaPropertiesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// SchemaAdditionalProperties structure is generated from "#/definitions/Schema->additionalProperties".
type SchemaAdditionalProperties struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
	Bool      *bool      `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SchemaAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Bool)
	if err != nil {
		i.Bool = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SchemaAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference, i.Bool)
}

// Discriminator structure is generated from "#/definitions/Discriminator".
type Discriminator struct {
	PropertyName         string                 `json:"propertyName,omitempty"`
	Mapping              map[string]string      `json:"mapping,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                      // All unmatched properties
}

type marshalDiscriminator Discriminator

var ignoreKeysDiscriminator = []string{
	"propertyName",
	"mapping",
}

// UnmarshalJSON decodes JSON.
func (i *Discriminator) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalDiscriminator(*i)

	err = json.Unmarshal(data, &ii)
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
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = Discriminator(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Discriminator) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalDiscriminator(i), i.AdditionalProperties)
}

// XML structure is generated from "#/definitions/XML".
type XML struct {
	Name          string                 `json:"name,omitempty"`
	Namespace     string                 `json:"namespace,omitempty"`
	Prefix        string                 `json:"prefix,omitempty"`
	Attribute     bool                   `json:"attribute,omitempty"`
	Wrapped       bool                   `json:"wrapped,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                   // Key must match pattern: ^x-
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
func (i *XML) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalXML(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = XML(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i XML) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalXML(i), i.MapOfAnything)
}

// ParameterSchema structure is generated from "#/definitions/Parameter->schema".
type ParameterSchema struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ParameterSchema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ParameterSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// MediaType structure is generated from "#/definitions/MediaType".
type MediaType struct {
	Schema        *MediaTypeSchema                                 `json:"schema,omitempty"`
	Example       *interface{}                                     `json:"example,omitempty"`
	Examples      map[string]MediaTypeExamplesAdditionalProperties `json:"examples,omitempty"`
	MapOfAnything map[string]interface{}                           `json:"-"`                  // Key must match pattern: ^x-
	Encoding      map[string]Encoding                              `json:"encoding,omitempty"`
}

type marshalMediaType MediaType

var ignoreKeysMediaType = []string{
	"schema",
	"example",
	"examples",
	"encoding",
}

// UnmarshalJSON decodes JSON.
func (i *MediaType) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalMediaType(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			ii.Example = &v
		}
	}

	for _, key := range ignoreKeysMediaType {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = MediaType(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i MediaType) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMediaType(i), i.MapOfAnything)
}

// MediaTypeSchema structure is generated from "#/definitions/MediaType->schema".
type MediaTypeSchema struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MediaTypeSchema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MediaTypeSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// Example structure is generated from "#/definitions/Example".
type Example struct {
	Summary       string                 `json:"summary,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Value         *interface{}           `json:"value,omitempty"`
	ExternalValue string                 `json:"externalValue,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                       // Key must match pattern: ^x-
}

type marshalExample Example

var ignoreKeysExample = []string{
	"summary",
	"description",
	"value",
	"externalValue",
}

// UnmarshalJSON decodes JSON.
func (i *Example) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalExample(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Value == nil {
		if _, ok := m["value"]; ok {
			var v interface{}
			ii.Value = &v
		}
	}

	for _, key := range ignoreKeysExample {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Example(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Example) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExample(i), i.MapOfAnything)
}

// MediaTypeExamplesAdditionalProperties structure is generated from "#/definitions/MediaType->examples->additionalProperties".
type MediaTypeExamplesAdditionalProperties struct {
	Example   *Example   `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *MediaTypeExamplesAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Example)
	if err != nil {
		i.Example = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i MediaTypeExamplesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Example, i.Reference)
}

// Encoding structure is generated from "#/definitions/Encoding".
type Encoding struct {
	ContentType   string            `json:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty"`
	Style         EncodingStyle     `json:"style,omitempty"`
	Explode       bool              `json:"explode,omitempty"`
	AllowReserved bool              `json:"allowReserved,omitempty"`
}

// Header structure is generated from "#/definitions/Header".
type Header struct {
	Description     string                                        `json:"description,omitempty"`
	Required        bool                                          `json:"required,omitempty"`
	Deprecated      bool                                          `json:"deprecated,omitempty"`
	AllowEmptyValue bool                                          `json:"allowEmptyValue,omitempty"`
	Explode         bool                                          `json:"explode,omitempty"`
	AllowReserved   bool                                          `json:"allowReserved,omitempty"`
	Schema          *HeaderSchema                                 `json:"schema,omitempty"`
	Content         map[string]MediaType                          `json:"content,omitempty"`
	Example         *interface{}                                  `json:"example,omitempty"`
	Examples        map[string]HeaderExamplesAdditionalProperties `json:"examples,omitempty"`
	MapOfAnything   map[string]interface{}                        `json:"-"`                         // Key must match pattern: ^x-
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
func (i *Header) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalHeader(*i)

	err = json.Unmarshal(data, &ii)
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

	if ii.Example == nil {
		if _, ok := m["example"]; ok {
			var v interface{}
			ii.Example = &v
		}
	}

	for _, key := range ignoreKeysHeader {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Header(ii)

	return nil
}

var (
	// constHeader is unconditionally added to JSON.
	constHeader = json.RawMessage(`{"style":"simple"}`)
)

// MarshalJSON encodes JSON.
func (i Header) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHeader, marshalHeader(i), i.MapOfAnything)
}

// HeaderSchema structure is generated from "#/definitions/Header->schema".
type HeaderSchema struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *HeaderSchema) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i HeaderSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// HeaderExamplesAdditionalProperties structure is generated from "#/definitions/Header->examples->additionalProperties".
type HeaderExamplesAdditionalProperties struct {
	Example   *Example   `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *HeaderExamplesAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Example)
	if err != nil {
		i.Example = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i HeaderExamplesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Example, i.Reference)
}

// SchemaXORContentOneOf1 structure is generated from "#/definitions/SchemaXORContent/oneOf/1".
//
// Some properties are not allowed if content is present.
type SchemaXORContentOneOf1 struct {

}

// ParameterExamplesAdditionalProperties structure is generated from "#/definitions/Parameter->examples->additionalProperties".
type ParameterExamplesAdditionalProperties struct {
	Example   *Example   `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ParameterExamplesAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Example)
	if err != nil {
		i.Example = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ParameterExamplesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Example, i.Reference)
}

// ParameterLocationOneOf0 structure is generated from "#/definitions/ParameterLocation/oneOf/0".
//
// Parameter in path.
type ParameterLocationOneOf0 struct {
	Style                ParameterLocationOneOf0Style `json:"style,omitempty"`
	AdditionalProperties map[string]interface{}       `json:"-"`               // All unmatched properties
}

type marshalParameterLocationOneOf0 ParameterLocationOneOf0

var ignoreKeysParameterLocationOneOf0 = []string{
	"style",
	"in",
	"required",
}

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf0) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalParameterLocationOneOf0(*i)

	err = json.Unmarshal(data, &ii)
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
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = ParameterLocationOneOf0(ii)

	return nil
}

var (
	// constParameterLocationOneOf0 is unconditionally added to JSON.
	constParameterLocationOneOf0 = json.RawMessage(`{"in":"path","required":true}`)
)

// MarshalJSON encodes JSON.
func (i ParameterLocationOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf0, marshalParameterLocationOneOf0(i), i.AdditionalProperties)
}

// ParameterLocationOneOf1 structure is generated from "#/definitions/ParameterLocation/oneOf/1".
//
// Parameter in query.
type ParameterLocationOneOf1 struct {
	Style                ParameterLocationOneOf1Style `json:"style,omitempty"`
	AdditionalProperties map[string]interface{}       `json:"-"`               // All unmatched properties
}

type marshalParameterLocationOneOf1 ParameterLocationOneOf1

var ignoreKeysParameterLocationOneOf1 = []string{
	"style",
	"in",
}

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalParameterLocationOneOf1(*i)

	err = json.Unmarshal(data, &ii)
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
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = ParameterLocationOneOf1(ii)

	return nil
}

var (
	// constParameterLocationOneOf1 is unconditionally added to JSON.
	constParameterLocationOneOf1 = json.RawMessage(`{"in":"query"}`)
)

// MarshalJSON encodes JSON.
func (i ParameterLocationOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf1, marshalParameterLocationOneOf1(i), i.AdditionalProperties)
}

// ParameterLocationOneOf2 structure is generated from "#/definitions/ParameterLocation/oneOf/2".
//
// Parameter in header.
type ParameterLocationOneOf2 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf2) UnmarshalJSON(data []byte) error {
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
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constParameterLocationOneOf2 is unconditionally added to JSON.
	constParameterLocationOneOf2 = json.RawMessage(`{"in":"header","style":"simple"}`)
)

// MarshalJSON encodes JSON.
func (i ParameterLocationOneOf2) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf2, i.AdditionalProperties)
}

// ParameterLocationOneOf3 structure is generated from "#/definitions/ParameterLocation/oneOf/3".
//
// Parameter in cookie.
type ParameterLocationOneOf3 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf3) UnmarshalJSON(data []byte) error {
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
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constParameterLocationOneOf3 is unconditionally added to JSON.
	constParameterLocationOneOf3 = json.RawMessage(`{"in":"cookie","style":"form"}`)
)

// MarshalJSON encodes JSON.
func (i ParameterLocationOneOf3) MarshalJSON() ([]byte, error) {
	return marshalUnion(constParameterLocationOneOf3, i.AdditionalProperties)
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

// UnmarshalJSON decodes JSON.
func (i *ParameterLocation) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.OneOf0)
	if err != nil {
		i.OneOf0 = nil
	}

	err = json.Unmarshal(data, &i.OneOf1)
	if err != nil {
		i.OneOf1 = nil
	}

	err = json.Unmarshal(data, &i.OneOf2)
	if err != nil {
		i.OneOf2 = nil
	}

	err = json.Unmarshal(data, &i.OneOf3)
	if err != nil {
		i.OneOf3 = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ParameterLocation) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.OneOf0, i.OneOf1, i.OneOf2, i.OneOf3)
}

// PathItemParametersItems structure is generated from "#/definitions/PathItem->parameters->items".
type PathItemParametersItems struct {
	Parameter *Parameter `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *PathItemParametersItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Parameter)
	if err != nil {
		i.Parameter = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i PathItemParametersItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Parameter, i.Reference)
}

// Operation structure is generated from "#/definitions/Operation".
type Operation struct {
	Tags          []string                                          `json:"tags,omitempty"`
	Summary       string                                            `json:"summary,omitempty"`
	Description   string                                            `json:"description,omitempty"`
	ExternalDocs  *ExternalDocumentation                            `json:"externalDocs,omitempty"`
	ID            string                                            `json:"operationId,omitempty"`
	Parameters    []OperationParametersItems                        `json:"parameters,omitempty"`
	RequestBody   *OperationRequestBody                             `json:"requestBody,omitempty"`
	Responses     *Responses                                        `json:"responses,omitempty"`
	MapOfAnything map[string]interface{}                            `json:"-"`                      // Key must match pattern: ^x-
	Callbacks     map[string]OperationCallbacksAdditionalProperties `json:"callbacks,omitempty"`
	Deprecated    bool                                              `json:"deprecated,omitempty"`
	Security      []map[string][]string                             `json:"security,omitempty"`
	Servers       []Server                                          `json:"servers,omitempty"`
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
func (i *Operation) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOperation(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Operation(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(i), i.MapOfAnything)
}

// OperationParametersItems structure is generated from "#/definitions/Operation->parameters->items".
type OperationParametersItems struct {
	Parameter *Parameter `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *OperationParametersItems) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Parameter)
	if err != nil {
		i.Parameter = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationParametersItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Parameter, i.Reference)
}

// RequestBody structure is generated from "#/definitions/RequestBody".
type RequestBody struct {
	Description   string                 `json:"description,omitempty"`
	Content       map[string]MediaType   `json:"content,omitempty"`
	Required      bool                   `json:"required,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalRequestBody RequestBody

var ignoreKeysRequestBody = []string{
	"description",
	"content",
	"required",
}

// UnmarshalJSON decodes JSON.
func (i *RequestBody) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalRequestBody(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = RequestBody(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i RequestBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalRequestBody(i), i.MapOfAnything)
}

// OperationRequestBody structure is generated from "#/definitions/Operation->requestBody".
type OperationRequestBody struct {
	RequestBody *RequestBody `json:"-"`
	Reference   *Reference   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *OperationRequestBody) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.RequestBody)
	if err != nil {
		i.RequestBody = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationRequestBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.RequestBody, i.Reference)
}

// Responses structure is generated from "#/definitions/Responses".
type Responses struct {
	Default                    *ResponsesDefault          `json:"default,omitempty"`
	MapOfResponses15D2XXValues map[string]Responses15D2XX `json:"-"`                 // Key must match pattern: ^[1-5](?:\d{2}|XX)$
	MapOfAnything              map[string]interface{}     `json:"-"`                 // Key must match pattern: ^x-
}

type marshalResponses Responses

var ignoreKeysResponses = []string{
	"default",
}

// UnmarshalJSON decodes JSON.
func (i *Responses) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalResponses(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfResponses15D2XXValues == nil {
				ii.MapOfResponses15D2XXValues = make(map[string]Responses15D2XX, 1)
			}

			var val Responses15D2XX

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfResponses15D2XXValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Responses(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Responses) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponses(i), i.MapOfResponses15D2XXValues, i.MapOfAnything)
}

// Response structure is generated from "#/definitions/Response".
type Response struct {
	Description   string                                         `json:"description,omitempty"`
	Headers       map[string]ResponseHeadersAdditionalProperties `json:"headers,omitempty"`
	Content       map[string]MediaType                           `json:"content,omitempty"`
	Links         map[string]ResponseLinksAdditionalProperties   `json:"links,omitempty"`
	MapOfAnything map[string]interface{}                         `json:"-"`                     // Key must match pattern: ^x-
}

type marshalResponse Response

var ignoreKeysResponse = []string{
	"description",
	"headers",
	"content",
	"links",
}

// UnmarshalJSON decodes JSON.
func (i *Response) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalResponse(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Response(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Response) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponse(i), i.MapOfAnything)
}

// ResponseHeadersAdditionalProperties structure is generated from "#/definitions/Response->headers->additionalProperties".
type ResponseHeadersAdditionalProperties struct {
	Header    *Header    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ResponseHeadersAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Header)
	if err != nil {
		i.Header = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ResponseHeadersAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Header, i.Reference)
}

// Link structure is generated from "#/definitions/Link".
type Link struct {
	OperationID   string                 `json:"operationId,omitempty"`
	OperationRef  string                 `json:"operationRef,omitempty"`
	Parameters    map[string]interface{} `json:"parameters,omitempty"`
	RequestBody   *interface{}           `json:"requestBody,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Server        *Server                `json:"server,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
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
func (i *Link) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalLink(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.RequestBody == nil {
		if _, ok := m["requestBody"]; ok {
			var v interface{}
			ii.RequestBody = &v
		}
	}

	for _, key := range ignoreKeysLink {
		delete(m, key)
	}

	for key, rawValue := range m {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Link(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Link) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLink(i), i.MapOfAnything)
}

// ResponseLinksAdditionalProperties structure is generated from "#/definitions/Response->links->additionalProperties".
type ResponseLinksAdditionalProperties struct {
	Link      *Link      `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ResponseLinksAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Link)
	if err != nil {
		i.Link = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ResponseLinksAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Link, i.Reference)
}

// ResponsesDefault structure is generated from "#/definitions/Responses->default".
type ResponsesDefault struct {
	Response  *Response  `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ResponsesDefault) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Response)
	if err != nil {
		i.Response = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ResponsesDefault) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Response, i.Reference)
}

// Responses15D2XX structure is generated from "#/definitions/Responses->^[1-5](?:\d{2}|XX)$".
type Responses15D2XX struct {
	Response  *Response  `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Responses15D2XX) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Response)
	if err != nil {
		i.Response = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Responses15D2XX) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Response, i.Reference)
}

// Callback structure is generated from "#/definitions/Callback".
type Callback struct {
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	AdditionalProperties map[string]PathItem    `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *Callback) UnmarshalJSON(data []byte) error {
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

			if i.MapOfAnything == nil {
				i.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]PathItem, 1)
		}

		var val PathItem

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Callback) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfAnything, i.AdditionalProperties)
}

// OperationCallbacksAdditionalProperties structure is generated from "#/definitions/Operation->callbacks->additionalProperties".
type OperationCallbacksAdditionalProperties struct {
	Callback  *Callback  `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *OperationCallbacksAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Callback)
	if err != nil {
		i.Callback = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i OperationCallbacksAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Callback, i.Reference)
}

// Paths structure is generated from "#/definitions/Paths".
type Paths struct {
	MapOfPathItemValues map[string]PathItem    `json:"-"` // Key must match pattern: ^\/
	MapOfAnything       map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON.
func (i *Paths) UnmarshalJSON(data []byte) error {
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

			if i.MapOfPathItemValues == nil {
				i.MapOfPathItemValues = make(map[string]PathItem, 1)
			}

			var val PathItem

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfPathItemValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if i.MapOfAnything == nil {
				i.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Paths) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfPathItemValues, i.MapOfAnything)
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
	MapOfAnything   map[string]interface{}     `json:"-"`                         // Key must match pattern: ^x-
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
func (i *Components) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalComponents(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = Components(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Components) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponents(i), i.MapOfAnything)
}

// ComponentsSchemasAZAZ09 structure is generated from "#/definitions/Components->schemas->^[a-zA-Z0-9\.\-_]+$".
type ComponentsSchemasAZAZ09 struct {
	Schema    *Schema    `json:"-"`
	Reference *Reference `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSchemasAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Schema)
	if err != nil {
		i.Schema = nil
	}

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSchemasAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// ComponentsSchemas structure is generated from "#/definitions/Components->schemas".
type ComponentsSchemas struct {
	MapOfComponentsSchemasAZAZ09Values map[string]ComponentsSchemasAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties               map[string]interface{}             `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSchemas) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsSchemasAZAZ09Values == nil {
				i.MapOfComponentsSchemasAZAZ09Values = make(map[string]ComponentsSchemasAZAZ09, 1)
			}

			var val ComponentsSchemasAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsSchemasAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSchemas) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsSchemasAZAZ09Values, i.AdditionalProperties)
}

// ComponentsResponsesAZAZ09 structure is generated from "#/definitions/Components->responses->^[a-zA-Z0-9\.\-_]+$".
type ComponentsResponsesAZAZ09 struct {
	Reference *Reference `json:"-"`
	Response  *Response  `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsResponsesAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Response)
	if err != nil {
		i.Response = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsResponsesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Response)
}

// ComponentsResponses structure is generated from "#/definitions/Components->responses".
type ComponentsResponses struct {
	MapOfComponentsResponsesAZAZ09Values map[string]ComponentsResponsesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                 map[string]interface{}               `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsResponses) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsResponsesAZAZ09Values == nil {
				i.MapOfComponentsResponsesAZAZ09Values = make(map[string]ComponentsResponsesAZAZ09, 1)
			}

			var val ComponentsResponsesAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsResponsesAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsResponses) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsResponsesAZAZ09Values, i.AdditionalProperties)
}

// ComponentsParametersAZAZ09 structure is generated from "#/definitions/Components->parameters->^[a-zA-Z0-9\.\-_]+$".
type ComponentsParametersAZAZ09 struct {
	Reference *Reference `json:"-"`
	Parameter *Parameter `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsParametersAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Parameter)
	if err != nil {
		i.Parameter = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsParametersAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Parameter)
}

// ComponentsParameters structure is generated from "#/definitions/Components->parameters".
type ComponentsParameters struct {
	MapOfComponentsParametersAZAZ09Values map[string]ComponentsParametersAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                  map[string]interface{}                `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsParameters) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsParametersAZAZ09Values == nil {
				i.MapOfComponentsParametersAZAZ09Values = make(map[string]ComponentsParametersAZAZ09, 1)
			}

			var val ComponentsParametersAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsParametersAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsParameters) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsParametersAZAZ09Values, i.AdditionalProperties)
}

// ComponentsExamplesAZAZ09 structure is generated from "#/definitions/Components->examples->^[a-zA-Z0-9\.\-_]+$".
type ComponentsExamplesAZAZ09 struct {
	Reference *Reference `json:"-"`
	Example   *Example   `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsExamplesAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Example)
	if err != nil {
		i.Example = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsExamplesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Example)
}

// ComponentsExamples structure is generated from "#/definitions/Components->examples".
type ComponentsExamples struct {
	MapOfComponentsExamplesAZAZ09Values map[string]ComponentsExamplesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                map[string]interface{}              `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsExamples) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsExamplesAZAZ09Values == nil {
				i.MapOfComponentsExamplesAZAZ09Values = make(map[string]ComponentsExamplesAZAZ09, 1)
			}

			var val ComponentsExamplesAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsExamplesAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsExamples) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsExamplesAZAZ09Values, i.AdditionalProperties)
}

// ComponentsRequestBodiesAZAZ09 structure is generated from "#/definitions/Components->requestBodies->^[a-zA-Z0-9\.\-_]+$".
type ComponentsRequestBodiesAZAZ09 struct {
	Reference   *Reference   `json:"-"`
	RequestBody *RequestBody `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsRequestBodiesAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.RequestBody)
	if err != nil {
		i.RequestBody = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsRequestBodiesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.RequestBody)
}

// ComponentsRequestBodies structure is generated from "#/definitions/Components->requestBodies".
type ComponentsRequestBodies struct {
	MapOfComponentsRequestBodiesAZAZ09Values map[string]ComponentsRequestBodiesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                     map[string]interface{}                   `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsRequestBodies) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsRequestBodiesAZAZ09Values == nil {
				i.MapOfComponentsRequestBodiesAZAZ09Values = make(map[string]ComponentsRequestBodiesAZAZ09, 1)
			}

			var val ComponentsRequestBodiesAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsRequestBodiesAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsRequestBodies) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsRequestBodiesAZAZ09Values, i.AdditionalProperties)
}

// ComponentsHeadersAZAZ09 structure is generated from "#/definitions/Components->headers->^[a-zA-Z0-9\.\-_]+$".
type ComponentsHeadersAZAZ09 struct {
	Reference *Reference `json:"-"`
	Header    *Header    `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsHeadersAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Header)
	if err != nil {
		i.Header = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsHeadersAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Header)
}

// ComponentsHeaders structure is generated from "#/definitions/Components->headers".
type ComponentsHeaders struct {
	MapOfComponentsHeadersAZAZ09Values map[string]ComponentsHeadersAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties               map[string]interface{}             `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsHeaders) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsHeadersAZAZ09Values == nil {
				i.MapOfComponentsHeadersAZAZ09Values = make(map[string]ComponentsHeadersAZAZ09, 1)
			}

			var val ComponentsHeadersAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsHeadersAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsHeadersAZAZ09Values, i.AdditionalProperties)
}

// APIKeySecurityScheme structure is generated from "#/definitions/APIKeySecurityScheme".
type APIKeySecurityScheme struct {
	Name          string                 `json:"name,omitempty"`
	In            APIKeySecuritySchemeIn `json:"in,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeySecurityScheme APIKeySecurityScheme

var ignoreKeysAPIKeySecurityScheme = []string{
	"name",
	"in",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *APIKeySecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAPIKeySecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = APIKeySecurityScheme(ii)

	return nil
}

var (
	// constAPIKeySecurityScheme is unconditionally added to JSON.
	constAPIKeySecurityScheme = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON.
func (i APIKeySecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKeySecurityScheme, marshalAPIKeySecurityScheme(i), i.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	Scheme        string                    `json:"scheme,omitempty"`
	BearerFormat  string                    `json:"bearerFormat,omitempty"`
	Description   string                    `json:"description,omitempty"`
	OneOf0        *HTTPSecuritySchemeOneOf0 `json:"-"`
	OneOf1        *HTTPSecuritySchemeOneOf1 `json:"-"`
	MapOfAnything map[string]interface{}    `json:"-"`                      // Key must match pattern: ^x-
}

type marshalHTTPSecurityScheme HTTPSecurityScheme

var ignoreKeysHTTPSecurityScheme = []string{
	"scheme",
	"bearerFormat",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalHTTPSecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &ii.OneOf0)
	if err != nil {
		ii.OneOf0 = nil
	}

	err = json.Unmarshal(data, &ii.OneOf1)
	if err != nil {
		ii.OneOf1 = nil
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = HTTPSecurityScheme(ii)

	return nil
}

var (
	// constHTTPSecurityScheme is unconditionally added to JSON.
	constHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)
)

// MarshalJSON encodes JSON.
func (i HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHTTPSecurityScheme, marshalHTTPSecurityScheme(i), i.MapOfAnything, i.OneOf0, i.OneOf1)
}

// HTTPSecuritySchemeOneOf0 structure is generated from "#/definitions/HTTPSecurityScheme/oneOf/0".
//
// Bearer.
type HTTPSecuritySchemeOneOf0 struct {
	AdditionalProperties map[string]interface{} `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *HTTPSecuritySchemeOneOf0) UnmarshalJSON(data []byte) error {
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
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

var (
	// constHTTPSecuritySchemeOneOf0 is unconditionally added to JSON.
	constHTTPSecuritySchemeOneOf0 = json.RawMessage(`{"scheme":"bearer"}`)
)

// MarshalJSON encodes JSON.
func (i HTTPSecuritySchemeOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHTTPSecuritySchemeOneOf0, i.AdditionalProperties)
}

// HTTPSecuritySchemeOneOf1 structure is generated from "#/definitions/HTTPSecurityScheme/oneOf/1".
//
// Non Bearer.
type HTTPSecuritySchemeOneOf1 struct {
	Scheme               *interface{}           `json:"scheme,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                // All unmatched properties
}

type marshalHTTPSecuritySchemeOneOf1 HTTPSecuritySchemeOneOf1

var ignoreKeysHTTPSecuritySchemeOneOf1 = []string{
	"scheme",
}

// UnmarshalJSON decodes JSON.
func (i *HTTPSecuritySchemeOneOf1) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalHTTPSecuritySchemeOneOf1(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	if ii.Scheme == nil {
		if _, ok := m["scheme"]; ok {
			var v interface{}
			ii.Scheme = &v
		}
	}

	for _, key := range ignoreKeysHTTPSecuritySchemeOneOf1 {
		delete(m, key)
	}

	for key, rawValue := range m {
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = HTTPSecuritySchemeOneOf1(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i HTTPSecuritySchemeOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalHTTPSecuritySchemeOneOf1(i), i.AdditionalProperties)
}

// OAuth2SecurityScheme structure is generated from "#/definitions/OAuth2SecurityScheme".
type OAuth2SecurityScheme struct {
	Flows         *OAuthFlows            `json:"flows,omitempty"`
	Description   string                 `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalOAuth2SecurityScheme OAuth2SecurityScheme

var ignoreKeysOAuth2SecurityScheme = []string{
	"flows",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *OAuth2SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOAuth2SecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = OAuth2SecurityScheme(ii)

	return nil
}

var (
	// constOAuth2SecurityScheme is unconditionally added to JSON.
	constOAuth2SecurityScheme = json.RawMessage(`{"type":"oauth2"}`)
)

// MarshalJSON encodes JSON.
func (i OAuth2SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOAuth2SecurityScheme, marshalOAuth2SecurityScheme(i), i.MapOfAnything)
}

// OAuthFlows structure is generated from "#/definitions/OAuthFlows".
type OAuthFlows struct {
	Implicit          *ImplicitOAuthFlow          `json:"implicit,omitempty"`
	Password          *PasswordOAuthFlow          `json:"password,omitempty"`
	ClientCredentials *ClientCredentialsFlow      `json:"clientCredentials,omitempty"`
	AuthorizationCode *AuthorizationCodeOAuthFlow `json:"authorizationCode,omitempty"`
	MapOfAnything     map[string]interface{}      `json:"-"`                           // Key must match pattern: ^x-
}

type marshalOAuthFlows OAuthFlows

var ignoreKeysOAuthFlows = []string{
	"implicit",
	"password",
	"clientCredentials",
	"authorizationCode",
}

// UnmarshalJSON decodes JSON.
func (i *OAuthFlows) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOAuthFlows(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = OAuthFlows(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i OAuthFlows) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOAuthFlows(i), i.MapOfAnything)
}

// ImplicitOAuthFlow structure is generated from "#/definitions/ImplicitOAuthFlow".
type ImplicitOAuthFlow struct {
	AuthorizationURL string                 `json:"authorizationUrl,omitempty"`
	RefreshURL       string                 `json:"refreshUrl,omitempty"`
	Scopes           map[string]string      `json:"scopes,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-
}

type marshalImplicitOAuthFlow ImplicitOAuthFlow

var ignoreKeysImplicitOAuthFlow = []string{
	"authorizationUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (i *ImplicitOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalImplicitOAuthFlow(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = ImplicitOAuthFlow(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ImplicitOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalImplicitOAuthFlow(i), i.MapOfAnything)
}

// PasswordOAuthFlow structure is generated from "#/definitions/PasswordOAuthFlow".
type PasswordOAuthFlow struct {
	TokenURL      string                 `json:"tokenUrl,omitempty"`
	RefreshURL    string                 `json:"refreshUrl,omitempty"`
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                    // Key must match pattern: ^x-
}

type marshalPasswordOAuthFlow PasswordOAuthFlow

var ignoreKeysPasswordOAuthFlow = []string{
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (i *PasswordOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalPasswordOAuthFlow(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = PasswordOAuthFlow(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i PasswordOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPasswordOAuthFlow(i), i.MapOfAnything)
}

// ClientCredentialsFlow structure is generated from "#/definitions/ClientCredentialsFlow".
type ClientCredentialsFlow struct {
	TokenURL      string                 `json:"tokenUrl,omitempty"`
	RefreshURL    string                 `json:"refreshUrl,omitempty"`
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                    // Key must match pattern: ^x-
}

type marshalClientCredentialsFlow ClientCredentialsFlow

var ignoreKeysClientCredentialsFlow = []string{
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (i *ClientCredentialsFlow) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalClientCredentialsFlow(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = ClientCredentialsFlow(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ClientCredentialsFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalClientCredentialsFlow(i), i.MapOfAnything)
}

// AuthorizationCodeOAuthFlow structure is generated from "#/definitions/AuthorizationCodeOAuthFlow".
type AuthorizationCodeOAuthFlow struct {
	AuthorizationURL string                 `json:"authorizationUrl,omitempty"`
	TokenURL         string                 `json:"tokenUrl,omitempty"`
	RefreshURL       string                 `json:"refreshUrl,omitempty"`
	Scopes           map[string]string      `json:"scopes,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-
}

type marshalAuthorizationCodeOAuthFlow AuthorizationCodeOAuthFlow

var ignoreKeysAuthorizationCodeOAuthFlow = []string{
	"authorizationUrl",
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (i *AuthorizationCodeOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalAuthorizationCodeOAuthFlow(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = AuthorizationCodeOAuthFlow(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i AuthorizationCodeOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAuthorizationCodeOAuthFlow(i), i.MapOfAnything)
}

// OpenIDConnectSecurityScheme structure is generated from "#/definitions/OpenIdConnectSecurityScheme".
type OpenIDConnectSecurityScheme struct {
	OpenIDConnectURL string                 `json:"openIdConnectUrl,omitempty"`
	Description      string                 `json:"description,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-
}

type marshalOpenIDConnectSecurityScheme OpenIDConnectSecurityScheme

var ignoreKeysOpenIDConnectSecurityScheme = []string{
	"openIdConnectUrl",
	"description",
	"type",
}

// UnmarshalJSON decodes JSON.
func (i *OpenIDConnectSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalOpenIDConnectSecurityScheme(*i)

	err = json.Unmarshal(data, &ii)
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

			if ii.MapOfAnything == nil {
				ii.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ii.MapOfAnything[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	*i = OpenIDConnectSecurityScheme(ii)

	return nil
}

var (
	// constOpenIDConnectSecurityScheme is unconditionally added to JSON.
	constOpenIDConnectSecurityScheme = json.RawMessage(`{"type":"openIdConnect"}`)
)

// MarshalJSON encodes JSON.
func (i OpenIDConnectSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOpenIDConnectSecurityScheme, marshalOpenIDConnectSecurityScheme(i), i.MapOfAnything)
}

// SecurityScheme structure is generated from "#/definitions/SecurityScheme".
type SecurityScheme struct {
	APIKeySecurityScheme        *APIKeySecurityScheme        `json:"-"`
	HTTPSecurityScheme          *HTTPSecurityScheme          `json:"-"`
	OAuth2SecurityScheme        *OAuth2SecurityScheme        `json:"-"`
	OpenIDConnectSecurityScheme *OpenIDConnectSecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.APIKeySecurityScheme)
	if err != nil {
		i.APIKeySecurityScheme = nil
	}

	err = json.Unmarshal(data, &i.HTTPSecurityScheme)
	if err != nil {
		i.HTTPSecurityScheme = nil
	}

	err = json.Unmarshal(data, &i.OAuth2SecurityScheme)
	if err != nil {
		i.OAuth2SecurityScheme = nil
	}

	err = json.Unmarshal(data, &i.OpenIDConnectSecurityScheme)
	if err != nil {
		i.OpenIDConnectSecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.APIKeySecurityScheme, i.HTTPSecurityScheme, i.OAuth2SecurityScheme, i.OpenIDConnectSecurityScheme)
}

// ComponentsSecuritySchemesAZAZ09 structure is generated from "#/definitions/Components->securitySchemes->^[a-zA-Z0-9\.\-_]+$".
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemesAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.SecurityScheme)
	if err != nil {
		i.SecurityScheme = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSecuritySchemesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/Components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesAZAZ09Values map[string]ComponentsSecuritySchemesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                       map[string]interface{}                     `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsSecuritySchemesAZAZ09Values == nil {
				i.MapOfComponentsSecuritySchemesAZAZ09Values = make(map[string]ComponentsSecuritySchemesAZAZ09, 1)
			}

			var val ComponentsSecuritySchemesAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsSecuritySchemesAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsSecuritySchemesAZAZ09Values, i.AdditionalProperties)
}

// ComponentsLinksAZAZ09 structure is generated from "#/definitions/Components->links->^[a-zA-Z0-9\.\-_]+$".
type ComponentsLinksAZAZ09 struct {
	Reference *Reference `json:"-"`
	Link      *Link      `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsLinksAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Link)
	if err != nil {
		i.Link = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsLinksAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Link)
}

// ComponentsLinks structure is generated from "#/definitions/Components->links".
type ComponentsLinks struct {
	MapOfComponentsLinksAZAZ09Values map[string]ComponentsLinksAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties             map[string]interface{}           `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsLinks) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsLinksAZAZ09Values == nil {
				i.MapOfComponentsLinksAZAZ09Values = make(map[string]ComponentsLinksAZAZ09, 1)
			}

			var val ComponentsLinksAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsLinksAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsLinks) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsLinksAZAZ09Values, i.AdditionalProperties)
}

// ComponentsCallbacksAZAZ09 structure is generated from "#/definitions/Components->callbacks->^[a-zA-Z0-9\.\-_]+$".
type ComponentsCallbacksAZAZ09 struct {
	Reference *Reference `json:"-"`
	Callback  *Callback  `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsCallbacksAZAZ09) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.Reference)
	if err != nil {
		i.Reference = nil
	}

	err = json.Unmarshal(data, &i.Callback)
	if err != nil {
		i.Callback = nil
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsCallbacksAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Reference, i.Callback)
}

// ComponentsCallbacks structure is generated from "#/definitions/Components->callbacks".
type ComponentsCallbacks struct {
	MapOfComponentsCallbacksAZAZ09Values map[string]ComponentsCallbacksAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
	AdditionalProperties                 map[string]interface{}               `json:"-"` // All unmatched properties
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsCallbacks) UnmarshalJSON(data []byte) error {
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

			if i.MapOfComponentsCallbacksAZAZ09Values == nil {
				i.MapOfComponentsCallbacksAZAZ09Values = make(map[string]ComponentsCallbacksAZAZ09, 1)
			}

			var val ComponentsCallbacksAZAZ09

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			i.MapOfComponentsCallbacksAZAZ09Values[key] = val
		}

		if matched {
			delete(m, key)
		}
	}

	for key, rawValue := range m {
		if i.AdditionalProperties == nil {
			i.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		i.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsCallbacks) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.MapOfComponentsCallbacksAZAZ09Values, i.AdditionalProperties)
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
	regexRef = regexp.MustCompile(`^\$ref$`)
	regex15D2XX = regexp.MustCompile(`^[1-5](?:\d{2}|XX)$`)
	regex = regexp.MustCompile(`^\/`)
	regexAZAZ09 = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
)
