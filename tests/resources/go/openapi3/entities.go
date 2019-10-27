package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
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

// UnmarshalJSON decodes JSON.
func (i *OpenAPI) UnmarshalJSON(data []byte) error {
	ii := marshalOpenAPI(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"openapi",
			"info",
			"externalDocs",
			"servers",
			"security",
			"tags",
			"paths",
			"components",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = OpenAPI(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Info) UnmarshalJSON(data []byte) error {
	ii := marshalInfo(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"title",
			"description",
			"termsOfService",
			"contact",
			"license",
			"version",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Info(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Contact) UnmarshalJSON(data []byte) error {
	ii := marshalContact(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"url",
			"email",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Contact(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *License) UnmarshalJSON(data []byte) error {
	ii := marshalLicense(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"url",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = License(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ExternalDocumentation) UnmarshalJSON(data []byte) error {
	ii := marshalExternalDocumentation(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"description",
			"url",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = ExternalDocumentation(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Server) UnmarshalJSON(data []byte) error {
	ii := marshalServer(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"url",
			"description",
			"variables",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Server(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ServerVariable) UnmarshalJSON(data []byte) error {
	ii := marshalServerVariable(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"enum",
			"default",
			"description",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = ServerVariable(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Tag) UnmarshalJSON(data []byte) error {
	ii := marshalTag(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"description",
			"externalDocs",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Tag(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *PathItem) UnmarshalJSON(data []byte) error {
	ii := marshalPathItem(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"$ref",
			"summary",
			"description",
			"servers",
			"parameters",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexGetPutPostDeleteOptionsHeadPatchTrace: &ii.MapOfOperationValues, // ^(get|put|post|delete|options|head|patch|trace)$
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = PathItem(ii)

	return err
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
	Example          interface{}                                      `json:"example,omitempty"`
	Examples         map[string]ParameterExamplesAdditionalProperties `json:"examples,omitempty"`
	SchemaXORContent *SchemaXORContentOneOf1                          `json:"-"`
	Location         *ParameterLocation                               `json:"-"`
	MapOfAnything    map[string]interface{}                           `json:"-"`                         // Key must match pattern: ^x-
}

type marshalParameter Parameter

// UnmarshalJSON decodes JSON.
func (i *Parameter) UnmarshalJSON(data []byte) error {
	ii := marshalParameter(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii, &ii.SchemaXORContent, &ii.Location},
		ignoreKeys: []string{
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
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Parameter(ii)

	return err
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
	Default              interface{}                                     `json:"default,omitempty"`
	Nullable             bool                                            `json:"nullable,omitempty"`
	Discriminator        *Discriminator                                  `json:"discriminator,omitempty"`
	ReadOnly             bool                                            `json:"readOnly,omitempty"`
	WriteOnly            bool                                            `json:"writeOnly,omitempty"`
	Example              interface{}                                     `json:"example,omitempty"`
	ExternalDocs         *ExternalDocumentation                          `json:"externalDocs,omitempty"`
	Deprecated           bool                                            `json:"deprecated,omitempty"`
	XML                  *XML                                            `json:"xml,omitempty"`
}

type marshalSchema Schema

// UnmarshalJSON decodes JSON.
func (i *Schema) UnmarshalJSON(data []byte) error {
	ii := marshalSchema(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
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
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Schema(ii)

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexRef: &i.MapOfStringValues, // ^\$ref$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference, &i.Bool}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[2] == nil {
		i.Bool = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Discriminator) UnmarshalJSON(data []byte) error {
	ii := marshalDiscriminator(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"propertyName",
			"mapping",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Discriminator(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *XML) UnmarshalJSON(data []byte) error {
	ii := marshalXML(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"namespace",
			"prefix",
			"attribute",
			"wrapped",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = XML(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
}

// MarshalJSON encodes JSON.
func (i ParameterSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// MediaType structure is generated from "#/definitions/MediaType".
type MediaType struct {
	Schema        *MediaTypeSchema                                 `json:"schema,omitempty"`
	Example       interface{}                                      `json:"example,omitempty"`
	Examples      map[string]MediaTypeExamplesAdditionalProperties `json:"examples,omitempty"`
	MapOfAnything map[string]interface{}                           `json:"-"`                  // Key must match pattern: ^x-
	Encoding      map[string]Encoding                              `json:"encoding,omitempty"`
}

type marshalMediaType MediaType

// UnmarshalJSON decodes JSON.
func (i *MediaType) UnmarshalJSON(data []byte) error {
	ii := marshalMediaType(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"schema",
			"example",
			"examples",
			"encoding",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = MediaType(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
}

// MarshalJSON encodes JSON.
func (i MediaTypeSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.Schema, i.Reference)
}

// Example structure is generated from "#/definitions/Example".
type Example struct {
	Summary       string                 `json:"summary,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Value         interface{}            `json:"value,omitempty"`
	ExternalValue string                 `json:"externalValue,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                       // Key must match pattern: ^x-
}

type marshalExample Example

// UnmarshalJSON decodes JSON.
func (i *Example) UnmarshalJSON(data []byte) error {
	ii := marshalExample(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"summary",
			"description",
			"value",
			"externalValue",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Example(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Example, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Example = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	Example         interface{}                                   `json:"example,omitempty"`
	Examples        map[string]HeaderExamplesAdditionalProperties `json:"examples,omitempty"`
	MapOfAnything   map[string]interface{}                        `json:"-"`                         // Key must match pattern: ^x-
}

type marshalHeader Header

// UnmarshalJSON decodes JSON.
func (i *Header) UnmarshalJSON(data []byte) error {
	ii := marshalHeader(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
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
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["style"]; !ok || string(v) != `"simple"` {
		return fmt.Errorf(`bad or missing const value for "style" ("simple" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = Header(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Example, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Example = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Example, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Example = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf0) UnmarshalJSON(data []byte) error {
	ii := marshalParameterLocationOneOf0(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"style",
			"in",
			"required",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["in"]; !ok || string(v) != `"path"` {
		return fmt.Errorf(`bad or missing const value for "in" ("path" expected, %v received)`, v)
	}

	if v, ok := constValues["required"]; !ok || string(v) != "true" {
		return fmt.Errorf(`bad or missing const value for "required" (true expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = ParameterLocationOneOf0(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ParameterLocationOneOf1) UnmarshalJSON(data []byte) error {
	ii := marshalParameterLocationOneOf1(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"style",
			"in",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["in"]; !ok || string(v) != `"query"` {
		return fmt.Errorf(`bad or missing const value for "in" ("query" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = ParameterLocationOneOf1(ii)

	return err
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
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["in"]; !ok || string(v) != `"header"` {
		return fmt.Errorf(`bad or missing const value for "in" ("header" expected, %v received)`, v)
	}

	if v, ok := constValues["style"]; !ok || string(v) != `"simple"` {
		return fmt.Errorf(`bad or missing const value for "style" ("simple" expected, %v received)`, v)
	}

	return err
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
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["in"]; !ok || string(v) != `"cookie"` {
		return fmt.Errorf(`bad or missing const value for "in" ("cookie" expected, %v received)`, v)
	}

	if v, ok := constValues["style"]; !ok || string(v) != `"form"` {
		return fmt.Errorf(`bad or missing const value for "style" ("form" expected, %v received)`, v)
	}

	return err
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
	mayUnmarshal := []interface{}{&i.OneOf0, &i.OneOf1, &i.OneOf2, &i.OneOf3}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.OneOf0 = nil
	}

	if mayUnmarshal[1] == nil {
		i.OneOf1 = nil
	}

	if mayUnmarshal[2] == nil {
		i.OneOf2 = nil
	}

	if mayUnmarshal[3] == nil {
		i.OneOf3 = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Parameter, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Parameter = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Operation) UnmarshalJSON(data []byte) error {
	ii := marshalOperation(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
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
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Operation(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Parameter, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Parameter = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *RequestBody) UnmarshalJSON(data []byte) error {
	ii := marshalRequestBody(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"description",
			"content",
			"required",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = RequestBody(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.RequestBody, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.RequestBody = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Responses) UnmarshalJSON(data []byte) error {
	ii := marshalResponses(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"default",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regex15D2XX: &ii.MapOfResponses15D2XXValues, // ^[1-5](?:\d{2}|XX)$
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Responses(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Response) UnmarshalJSON(data []byte) error {
	ii := marshalResponse(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"description",
			"headers",
			"content",
			"links",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Response(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Header, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Header = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	RequestBody   interface{}            `json:"requestBody,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Server        *Server                `json:"server,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalLink Link

// UnmarshalJSON decodes JSON.
func (i *Link) UnmarshalJSON(data []byte) error {
	ii := marshalLink(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"operationId",
			"operationRef",
			"parameters",
			"requestBody",
			"description",
			"server",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Link(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Link, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Link = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Response, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Response = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Response, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Response = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &i.MapOfAnything, // ^x-
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Callback, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Callback = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regex: &i.MapOfPathItemValues, // ^\/
			regexX: &i.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	return err
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

// UnmarshalJSON decodes JSON.
func (i *Components) UnmarshalJSON(data []byte) error {
	ii := marshalComponents(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"schemas",
			"responses",
			"parameters",
			"examples",
			"requestBodies",
			"headers",
			"securitySchemes",
			"links",
			"callbacks",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = Components(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.Schema, &i.Reference}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}

	if mayUnmarshal[1] == nil {
		i.Reference = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsSchemasAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.Response}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Response = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsResponsesAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.Parameter}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Parameter = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsParametersAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.Example}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Example = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsExamplesAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.RequestBody}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.RequestBody = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsRequestBodiesAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.Header}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Header = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsHeadersAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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

// UnmarshalJSON decodes JSON.
func (i *APIKeySecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalAPIKeySecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"name",
			"in",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"apiKey"` {
		return fmt.Errorf(`bad or missing const value for "type" ("apiKey" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = APIKeySecurityScheme(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&ii.OneOf0, &ii.OneOf1, &constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"scheme",
			"bearerFormat",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"http"` {
		return fmt.Errorf(`bad or missing const value for "type" ("http" expected, %v received)`, v)
	}

	if mayUnmarshal[0] == nil {
		ii.OneOf0 = nil
	}

	if mayUnmarshal[1] == nil {
		ii.OneOf1 = nil
	}

	if err != nil {
		return err
	}

	*i = HTTPSecurityScheme(ii)

	return err
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
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["scheme"]; !ok || string(v) != `"bearer"` {
		return fmt.Errorf(`bad or missing const value for "scheme" ("bearer" expected, %v received)`, v)
	}

	return err
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
	Scheme               interface{}            `json:"scheme,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                // All unmatched properties
}

type marshalHTTPSecuritySchemeOneOf1 HTTPSecuritySchemeOneOf1

// UnmarshalJSON decodes JSON.
func (i *HTTPSecuritySchemeOneOf1) UnmarshalJSON(data []byte) error {
	ii := marshalHTTPSecuritySchemeOneOf1(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"scheme",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = HTTPSecuritySchemeOneOf1(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *OAuth2SecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalOAuth2SecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"flows",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"oauth2"` {
		return fmt.Errorf(`bad or missing const value for "type" ("oauth2" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = OAuth2SecurityScheme(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *OAuthFlows) UnmarshalJSON(data []byte) error {
	ii := marshalOAuthFlows(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"implicit",
			"password",
			"clientCredentials",
			"authorizationCode",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = OAuthFlows(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ImplicitOAuthFlow) UnmarshalJSON(data []byte) error {
	ii := marshalImplicitOAuthFlow(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"authorizationUrl",
			"refreshUrl",
			"scopes",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = ImplicitOAuthFlow(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *PasswordOAuthFlow) UnmarshalJSON(data []byte) error {
	ii := marshalPasswordOAuthFlow(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"tokenUrl",
			"refreshUrl",
			"scopes",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = PasswordOAuthFlow(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *ClientCredentialsFlow) UnmarshalJSON(data []byte) error {
	ii := marshalClientCredentialsFlow(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"tokenUrl",
			"refreshUrl",
			"scopes",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = ClientCredentialsFlow(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *AuthorizationCodeOAuthFlow) UnmarshalJSON(data []byte) error {
	ii := marshalAuthorizationCodeOAuthFlow(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"authorizationUrl",
			"tokenUrl",
			"refreshUrl",
			"scopes",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if err != nil {
		return err
	}

	*i = AuthorizationCodeOAuthFlow(ii)

	return err
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

// UnmarshalJSON decodes JSON.
func (i *OpenIDConnectSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalOpenIDConnectSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		mayUnmarshal: mayUnmarshal,
		ignoreKeys: []string{
			"openIdConnectUrl",
			"description",
			"type",
		},
		patternProperties: map[*regexp.Regexp]interface{}{
			regexX: &ii.MapOfAnything, // ^x-
		},
		jsonData: data,
	}.unmarshal()

	if v, ok := constValues["type"]; !ok || string(v) != `"openIdConnect"` {
		return fmt.Errorf(`bad or missing const value for "type" ("openIdConnect" expected, %v received)`, v)
	}

	if err != nil {
		return err
	}

	*i = OpenIDConnectSecurityScheme(ii)

	return err
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
	mayUnmarshal := []interface{}{&i.APIKeySecurityScheme, &i.HTTPSecurityScheme, &i.OAuth2SecurityScheme, &i.OpenIDConnectSecurityScheme}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.APIKeySecurityScheme = nil
	}

	if mayUnmarshal[1] == nil {
		i.HTTPSecurityScheme = nil
	}

	if mayUnmarshal[2] == nil {
		i.OAuth2SecurityScheme = nil
	}

	if mayUnmarshal[3] == nil {
		i.OpenIDConnectSecurityScheme = nil
	}

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.SecurityScheme}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.SecurityScheme = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsSecuritySchemesAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.Link}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Link = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsLinksAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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
	mayUnmarshal := []interface{}{&i.Reference, &i.Callback}
	err := unionMap{
		mayUnmarshal: mayUnmarshal,
		jsonData: data,
	}.unmarshal()

	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}

	if mayUnmarshal[1] == nil {
		i.Callback = nil
	}

	return err
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
	err := unionMap{
		patternProperties: map[*regexp.Regexp]interface{}{
			regexAZAZ09: &i.MapOfComponentsCallbacksAZAZ09Values, // ^[a-zA-Z0-9\.\-_]+$
		},
		additionalProperties: &i.AdditionalProperties,
		jsonData: data,
	}.unmarshal()

	return err
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

type unionMap struct {
	mustUnmarshal        []interface{}
	mayUnmarshal         []interface{}
	ignoreKeys           []string
	patternProperties    map[*regexp.Regexp]interface{}
	additionalProperties interface{}
	jsonData             []byte
}

func (u unionMap) unmarshal() error {
	for _, item := range u.mustUnmarshal {
		// Unmarshal to struct.
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			return err
		}
	}

	for i, item := range u.mayUnmarshal {
		// Unmarshal to struct.
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			u.mayUnmarshal[i] = nil
		}
	}

	if len(u.patternProperties) == 0 && u.additionalProperties == nil {
		return nil
	}

	// Unmarshal to a generic map.
	var m map[string]*json.RawMessage

	err := json.Unmarshal(u.jsonData, &m)
	if err != nil {
		return err
	}

	// Remove ignored keys (defined in struct).
	for _, i := range u.ignoreKeys {
		delete(m, i)
	}

	// Return early on empty map.
	if len(m) == 0 {
		return nil
	}

	if len(u.patternProperties) != 0 {
		err = u.unmarshalPatternProperties(m)
		if err != nil {
			return err
		}
	}

	// Return early on empty map.
	if len(m) == 0 {
		return nil
	}

	if u.additionalProperties != nil {
		return u.unmarshalAdditionalProperties(m)
	}

	return nil
}

func (u unionMap) unmarshalAdditionalProperties(m map[string]*json.RawMessage) error {
	var err error

	subMap := make([]byte, 1, 100)

	subMap[0] = '{'

	// Iterating map and filling additional properties.
	for key, val := range m {
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		if len(subMap) != 1 {
			subMap = append(subMap[:len(subMap)-1], ',')
		}

		subMap = append(subMap, []byte(keyEscaped)...)

		if val != nil {
			subMap = append(subMap, []byte(*val)...)
		} else {
			subMap = append(subMap, []byte("null")...)
		}

		subMap = append(subMap, '}')
	}

	if len(subMap) > 1 {
		err = json.Unmarshal(subMap, u.additionalProperties)
		if err != nil {
			return err
		}
	}

	return nil
}
func (u unionMap) unmarshalPatternProperties(m map[string]*json.RawMessage) error {
	patternMapsRaw := make(map[*regexp.Regexp][]byte, len(u.patternProperties))

	// Iterating map and filling pattern properties sub maps.
	for key, val := range m {
		matched := false
		ok := false
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex := range u.patternProperties {
			if regex.MatchString(key) {
				matched = true

				var subMap []byte

				if subMap, ok = patternMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)

				if val != nil {
					subMap = append(subMap, []byte(*val)...)
				} else {
					subMap = append(subMap, []byte("null")...)
				}

				subMap = append(subMap, '}')

				patternMapsRaw[regex] = subMap
			}
		}

		// Remove from properties map if matched to at least one regex.
		if matched {
			delete(m, key)
		}
	}

	for regex := range u.patternProperties {
		if subMap, ok := patternMapsRaw[regex]; !ok {
			continue
		} else {
			err := json.Unmarshal(subMap, u.patternProperties[regex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
