package swagger

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// SwaggerSchema structure is generated from #
// A JSON Schema for Swagger 2.0 API.
type SwaggerSchema struct {
	Info                *Info                                              `json:"info,omitempty"`                // General information about the API.
	Host                string                                             `json:"host,omitempty"`                // The host (name or ip) of the API. Example: 'swagger.io'
	BasePath            string                                             `json:"basePath,omitempty"`            // The base path to the API. Example: '/api'.
	Schemes             []SchemesListItems                                 `json:"schemes,omitempty"`             // The transfer protocol of the API.
	Consumes            *Consumes                                          `json:"consumes,omitempty"`            // A list of MIME types accepted by the API.
	Produces            *Produces                                          `json:"produces,omitempty"`            // A list of MIME types the API can produce.
	Paths               *Paths                                             `json:"paths,omitempty"`               // Relative paths to the individual endpoints. They must be relative to the 'basePath'.
	Definitions         map[string]Schema                                  `json:"definitions,omitempty"`         // One or more JSON objects describing the schemas being consumed and produced by the API.
	Parameters          map[string]Parameter                               `json:"parameters,omitempty"`          // One or more JSON representations for parameters
	Responses           map[string]Response                                `json:"responses,omitempty"`           // One or more JSON representations for parameters
	Security            []map[string][]string                              `json:"security,omitempty"`
	SecurityDefinitions map[string]SecurityDefinitionsAdditionalProperties `json:"securityDefinitions,omitempty"`
	Tags                []Tag                                              `json:"tags,omitempty"`
	ExternalDocs        *ExternalDocs                                      `json:"externalDocs,omitempty"`        // information about external documentation
	MapOfAnythingValues map[string]interface{}                             `json:"-"`                             // Key must match pattern: ^x-
}

type marshalSwaggerSchema SwaggerSchema

// UnmarshalJSON decodes JSON
func (i *SwaggerSchema) UnmarshalJSON(data []byte) error {
	ii := marshalSwaggerSchema(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"info",
			"host",
			"basePath",
			"schemes",
			"consumes",
			"produces",
			"paths",
			"definitions",
			"parameters",
			"responses",
			"security",
			"securityDefinitions",
			"tags",
			"externalDocs",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["swagger"];!ok || string(v) != `"2.0"` {
	    return errors.New(`bad or missing const value for "swagger" ("2.0" expected)`)
	}
	if err != nil {
		return err
	}
	*i = SwaggerSchema(ii)
	return err
}

var (
	// constSwaggerSchema is unconditionally added to JSON
	constSwaggerSchema = json.RawMessage(`{"swagger":"2.0"}`)
)

// MarshalJSON encodes JSON
func (i SwaggerSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSwaggerSchema(i), i.MapOfAnythingValues, constSwaggerSchema)
}

// Info structure is generated from #/definitions/info
// General information about the API.
type Info struct {
	Title               string                 `json:"title,omitempty"`          // A unique and precise title of the API.
	Version             string                 `json:"version,omitempty"`        // A semantic version number of the API.
	Description         string                 `json:"description,omitempty"`    // A longer description of the API. Should be different from the title.  GitHub Flavored Markdown is allowed.
	TermsOfService      string                 `json:"termsOfService,omitempty"` // The terms of service for the API.
	Contact             *Contact               `json:"contact,omitempty"`        // Contact information for the owners of the API.
	License             *License               `json:"license,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                        // Key must match pattern: ^x-
}

type marshalInfo Info

// UnmarshalJSON decodes JSON
func (i *Info) UnmarshalJSON(data []byte) error {
	ii := marshalInfo(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"title",
			"version",
			"description",
			"termsOfService",
			"contact",
			"license",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Info(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(i), i.MapOfAnythingValues)
}

// Contact structure is generated from #/definitions/contact
// Contact information for the owners of the API.
type Contact struct {
	Name                string                 `json:"name,omitempty"`  // The identifying name of the contact person/organization.
	URL                 string                 `json:"url,omitempty"`   // The URL pointing to the contact information.
	Email               string                 `json:"email,omitempty"` // The email address of the contact person/organization.
	MapOfAnythingValues map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalContact Contact

// UnmarshalJSON decodes JSON
func (i *Contact) UnmarshalJSON(data []byte) error {
	ii := marshalContact(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"url",
			"email",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Contact(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(i), i.MapOfAnythingValues)
}

// License structure is generated from #/definitions/license
type License struct {
	Name                string                 `json:"name,omitempty"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL                 string                 `json:"url,omitempty"`  // The URL pointing to the license.
	MapOfAnythingValues map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

type marshalLicense License

// UnmarshalJSON decodes JSON
func (i *License) UnmarshalJSON(data []byte) error {
	ii := marshalLicense(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"url",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = License(ii)
	return err
}

// MarshalJSON encodes JSON
func (i License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(i), i.MapOfAnythingValues)
}

// Consumes structure is generated from #->consumes
// A list of MIME types accepted by the API.
type Consumes struct {
	AllOf0 []string `json:"-"`
}

type marshalConsumes Consumes

// UnmarshalJSON decodes JSON
func (i *Consumes) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		[]interface{}{&i.AllOf0},
		nil,
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Consumes) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalConsumes(i), i.AllOf0)
}

// Produces structure is generated from #->produces
// A list of MIME types the API can produce.
type Produces struct {
	AllOf0 []string `json:"-"`
}

type marshalProduces Produces

// UnmarshalJSON decodes JSON
func (i *Produces) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		[]interface{}{&i.AllOf0},
		nil,
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Produces) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProduces(i), i.AllOf0)
}

// Paths structure is generated from #/definitions/paths
// Relative paths to the individual endpoints. They must be relative to the 'basePath'.
type Paths struct {
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	MapOfPathItemValues map[string]PathItem    `json:"-"` // Key must match pattern: ^/
}

type marshalPaths Paths

// UnmarshalJSON decodes JSON
func (i *Paths) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			"^x-": &i.MapOfAnythingValues,
			"^/": &i.MapOfPathItemValues,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Paths) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPaths(i), i.MapOfAnythingValues, i.MapOfPathItemValues)
}

// PathItem structure is generated from #/definitions/pathItem
type PathItem struct {
	Ref                 string                 `json:"$ref,omitempty"`
	Get                 *Operation             `json:"get,omitempty"`
	Put                 *Operation             `json:"put,omitempty"`
	Post                *Operation             `json:"post,omitempty"`
	Delete              *Operation             `json:"delete,omitempty"`
	Options             *Operation             `json:"options,omitempty"`
	Head                *Operation             `json:"head,omitempty"`
	Patch               *Operation             `json:"patch,omitempty"`
	Parameters          []ParametersListItems  `json:"parameters,omitempty"` // The parameters needed to send a valid API call.
	MapOfAnythingValues map[string]interface{} `json:"-"`                    // Key must match pattern: ^x-
}

type marshalPathItem PathItem

// UnmarshalJSON decodes JSON
func (i *PathItem) UnmarshalJSON(data []byte) error {
	ii := marshalPathItem(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"$ref",
			"get",
			"put",
			"post",
			"delete",
			"options",
			"head",
			"patch",
			"parameters",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = PathItem(ii)
	return err
}

// MarshalJSON encodes JSON
func (i PathItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPathItem(i), i.MapOfAnythingValues)
}

// Operation structure is generated from #/definitions/operation
type Operation struct {
	Tags                []string               `json:"tags,omitempty"`
	Summary             string                 `json:"summary,omitempty"`      // A brief summary of the operation.
	Description         string                 `json:"description,omitempty"`  // A longer description of the operation, GitHub Flavored Markdown is allowed.
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	ID                  string                 `json:"operationId,omitempty"`  // A unique identifier of the operation.
	Produces            *OperationProduces     `json:"produces,omitempty"`     // A list of MIME types the API can produce.
	Consumes            *OperationConsumes     `json:"consumes,omitempty"`     // A list of MIME types the API can consume.
	Parameters          []ParametersListItems  `json:"parameters,omitempty"`   // The parameters needed to send a valid API call.
	Responses           *Responses             `json:"responses,omitempty"`    // Response objects names can either be any valid HTTP status code or 'default'.
	Schemes             []SchemesListItems     `json:"schemes,omitempty"`      // The transfer protocol of the API.
	Deprecated          bool                   `json:"deprecated,omitempty"`
	Security            []map[string][]string  `json:"security,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalOperation Operation

// UnmarshalJSON decodes JSON
func (i *Operation) UnmarshalJSON(data []byte) error {
	ii := marshalOperation(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"tags",
			"summary",
			"description",
			"externalDocs",
			"operationId",
			"produces",
			"consumes",
			"parameters",
			"responses",
			"schemes",
			"deprecated",
			"security",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Operation(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(i), i.MapOfAnythingValues)
}

// ExternalDocs structure is generated from #/definitions/externalDocs
// information about external documentation
type ExternalDocs struct {
	Description         string                 `json:"description,omitempty"`
	URL                 string                 `json:"url,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalExternalDocs ExternalDocs

// UnmarshalJSON decodes JSON
func (i *ExternalDocs) UnmarshalJSON(data []byte) error {
	ii := marshalExternalDocs(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"description",
			"url",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = ExternalDocs(ii)
	return err
}

// MarshalJSON encodes JSON
func (i ExternalDocs) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocs(i), i.MapOfAnythingValues)
}

// OperationProduces structure is generated from #/definitions/operation->produces
// A list of MIME types the API can produce.
type OperationProduces struct {
	AllOf0 []string `json:"-"`
}

type marshalOperationProduces OperationProduces

// UnmarshalJSON decodes JSON
func (i *OperationProduces) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		[]interface{}{&i.AllOf0},
		nil,
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i OperationProduces) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationProduces(i), i.AllOf0)
}

// OperationConsumes structure is generated from #/definitions/operation->consumes
// A list of MIME types the API can consume.
type OperationConsumes struct {
	AllOf0 []string `json:"-"`
}

type marshalOperationConsumes OperationConsumes

// UnmarshalJSON decodes JSON
func (i *OperationConsumes) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		[]interface{}{&i.AllOf0},
		nil,
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i OperationConsumes) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationConsumes(i), i.AllOf0)
}

// BodyParameter structure is generated from #/definitions/bodyParameter
type BodyParameter struct {
	Description         string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                 `json:"name,omitempty"`        // The name of the parameter.
	Required            bool                   `json:"required,omitempty"`    // Determines whether or not this parameter is required or optional.
	Schema              *Schema                `json:"schema,omitempty"`      // A deterministic version of a JSON Schema object.
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalBodyParameter BodyParameter

// UnmarshalJSON decodes JSON
func (i *BodyParameter) UnmarshalJSON(data []byte) error {
	ii := marshalBodyParameter(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
			"name",
			"required",
			"schema",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["in"];!ok || string(v) != `"body"` {
	    return errors.New(`bad or missing const value for "in" ("body" expected)`)
	}
	if err != nil {
		return err
	}
	*i = BodyParameter(ii)
	return err
}

var (
	// constBodyParameter is unconditionally added to JSON
	constBodyParameter = json.RawMessage(`{"in":"body"}`)
)

// MarshalJSON encodes JSON
func (i BodyParameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalBodyParameter(i), i.MapOfAnythingValues, constBodyParameter)
}

// Schema structure is generated from #/definitions/schema
// A deterministic version of a JSON Schema object.
type Schema struct {
	Ref                  string                                                            `json:"$ref,omitempty"`
	Format               string                                                            `json:"format,omitempty"`
	Title                string                                                            `json:"title,omitempty"`
	Description          string                                                            `json:"description,omitempty"`
	Default              interface{}                                                       `json:"default,omitempty"`
	MultipleOf           float64                                                           `json:"multipleOf,omitempty"`
	Maximum              float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                                             `json:"maxLength,omitempty"`
	MinLength            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern              string                                                            `json:"pattern,omitempty"`
	MaxItems             int64                                                             `json:"maxItems,omitempty"`
	MinItems             *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems          bool                                                              `json:"uniqueItems,omitempty"`
	MaxProperties        int64                                                             `json:"maxProperties,omitempty"`
	MinProperties        *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minProperties,omitempty"`
	Required             []string                                                          `json:"required,omitempty"`
	Enum                 []interface{}                                                     `json:"enum,omitempty"`
	MapOfAnythingValues  map[string]interface{}                                            `json:"-"`                              // Key must match pattern: ^x-
	AdditionalProperties *SchemaAdditionalProperties                                       `json:"additionalProperties,omitempty"`
	Type                 *HTTPJSONSchemaOrgDraft04SchemaPropertiesType                     `json:"type,omitempty"`
	Items                *SchemaItems                                                      `json:"items,omitempty"`
	AllOf                []Schema                                                          `json:"allOf,omitempty"`
	Properties           map[string]Schema                                                 `json:"properties,omitempty"`
	Discriminator        string                                                            `json:"discriminator,omitempty"`
	ReadOnly             bool                                                              `json:"readOnly,omitempty"`
	XML                  *XML                                                              `json:"xml,omitempty"`
	ExternalDocs         *ExternalDocs                                                     `json:"externalDocs,omitempty"`         // information about external documentation
	Example              interface{}                                                       `json:"example,omitempty"`
}

type marshalSchema Schema

// UnmarshalJSON decodes JSON
func (i *Schema) UnmarshalJSON(data []byte) error {
	ii := marshalSchema(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"$ref",
			"format",
			"title",
			"description",
			"default",
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
			"additionalProperties",
			"type",
			"items",
			"allOf",
			"properties",
			"discriminator",
			"readOnly",
			"xml",
			"externalDocs",
			"example",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Schema(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchema(i), i.MapOfAnythingValues)
}

// HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 structure is generated from http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0
type HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 struct {
	Int64 *int64 `json:"-"`
}

type marshalHTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0

// UnmarshalJSON decodes JSON
func (i *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		[]interface{}{&i.Int64},
		nil,
		nil,
		nil,
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalHTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0(i), i.Int64)
}

// SchemaAdditionalProperties structure is generated from #/definitions/schema->additionalProperties
type SchemaAdditionalProperties struct {
	Schema *Schema `json:"-"`
	Bool   *bool   `json:"-"`
}

type marshalSchemaAdditionalProperties SchemaAdditionalProperties

// UnmarshalJSON decodes JSON
func (i *SchemaAdditionalProperties) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.Bool}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}
	if mayUnmarshal[1] == nil {
		i.Bool = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i SchemaAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchemaAdditionalProperties(i), i.Schema, i.Bool)
}

// HTTPJSONSchemaOrgDraft04SchemaPropertiesType structure is generated from http://json-schema.org/draft-04/schema#/properties/type
type HTTPJSONSchemaOrgDraft04SchemaPropertiesType struct {
	AnyOf1 []interface{} `json:"-"`
}

type marshalHTTPJSONSchemaOrgDraft04SchemaPropertiesType HTTPJSONSchemaOrgDraft04SchemaPropertiesType

// UnmarshalJSON decodes JSON
func (i *HTTPJSONSchemaOrgDraft04SchemaPropertiesType) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i HTTPJSONSchemaOrgDraft04SchemaPropertiesType) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalHTTPJSONSchemaOrgDraft04SchemaPropertiesType(i), i.AnyOf1)
}

// SchemaItems structure is generated from #/definitions/schema->items
type SchemaItems struct {
	Schema *Schema  `json:"-"`
	AnyOf1 []Schema `json:"-"`
}

type marshalSchemaItems SchemaItems

// UnmarshalJSON decodes JSON
func (i *SchemaItems) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}
	if mayUnmarshal[1] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i SchemaItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchemaItems(i), i.Schema, i.AnyOf1)
}

// XML structure is generated from #/definitions/xml
type XML struct {
	Name                string                 `json:"name,omitempty"`
	Namespace           string                 `json:"namespace,omitempty"`
	Prefix              string                 `json:"prefix,omitempty"`
	Attribute           bool                   `json:"attribute,omitempty"`
	Wrapped             bool                   `json:"wrapped,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                   // Key must match pattern: ^x-
}

type marshalXML XML

// UnmarshalJSON decodes JSON
func (i *XML) UnmarshalJSON(data []byte) error {
	ii := marshalXML(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"namespace",
			"prefix",
			"attribute",
			"wrapped",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = XML(ii)
	return err
}

// MarshalJSON encodes JSON
func (i XML) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalXML(i), i.MapOfAnythingValues)
}

// Parameter structure is generated from #/definitions/parameter
type Parameter struct {
	BodyParameter    *BodyParameter    `json:"-"`
	NonBodyParameter *NonBodyParameter `json:"-"`
}

type marshalParameter Parameter

// UnmarshalJSON decodes JSON
func (i *Parameter) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.BodyParameter, &i.NonBodyParameter}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.BodyParameter = nil
	}
	if mayUnmarshal[1] == nil {
		i.NonBodyParameter = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(i), i.BodyParameter, i.NonBodyParameter)
}

// HeaderParameterSubSchema structure is generated from #/definitions/headerParameterSubSchema
type HeaderParameterSubSchema struct {
	Required            bool                                                              `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	Description         string                                                            `json:"description,omitempty"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                                                            `json:"name,omitempty"`             // The name of the parameter.
	Type                HeaderParameterSubSchemaType                                      `json:"type,omitempty"`
	Format              string                                                            `json:"format,omitempty"`
	Items               *PrimitivesItems                                                  `json:"items,omitempty"`
	CollectionFormat    CollectionFormat                                                  `json:"collectionFormat,omitempty"`
	Default             interface{}                                                       `json:"default,omitempty"`
	Maximum             float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum    bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum             float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum    bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength           int64                                                             `json:"maxLength,omitempty"`
	MinLength           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern             string                                                            `json:"pattern,omitempty"`
	MaxItems            int64                                                             `json:"maxItems,omitempty"`
	MinItems            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems         bool                                                              `json:"uniqueItems,omitempty"`
	Enum                []interface{}                                                     `json:"enum,omitempty"`
	MultipleOf          float64                                                           `json:"multipleOf,omitempty"`
	MapOfAnythingValues map[string]interface{}                                            `json:"-"`                          // Key must match pattern: ^x-
}

type marshalHeaderParameterSubSchema HeaderParameterSubSchema

// UnmarshalJSON decodes JSON
func (i *HeaderParameterSubSchema) UnmarshalJSON(data []byte) error {
	ii := marshalHeaderParameterSubSchema(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"required",
			"description",
			"name",
			"type",
			"format",
			"items",
			"collectionFormat",
			"default",
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
			"enum",
			"multipleOf",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["in"];!ok || string(v) != `"header"` {
	    return errors.New(`bad or missing const value for "in" ("header" expected)`)
	}
	if err != nil {
		return err
	}
	*i = HeaderParameterSubSchema(ii)
	return err
}

var (
	// constHeaderParameterSubSchema is unconditionally added to JSON
	constHeaderParameterSubSchema = json.RawMessage(`{"in":"header"}`)
)

// MarshalJSON encodes JSON
func (i HeaderParameterSubSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalHeaderParameterSubSchema(i), i.MapOfAnythingValues, constHeaderParameterSubSchema)
}

// PrimitivesItems structure is generated from #/definitions/primitivesItems
type PrimitivesItems struct {
	Type                PrimitivesItemsType                                               `json:"type,omitempty"`
	Format              string                                                            `json:"format,omitempty"`
	MapOfAnythingValues map[string]interface{}                                            `json:"-"`                          // Key must match pattern: ^x-
	Items               *PrimitivesItems                                                  `json:"items,omitempty"`
	CollectionFormat    CollectionFormat                                                  `json:"collectionFormat,omitempty"`
	Default             interface{}                                                       `json:"default,omitempty"`
	Maximum             float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum    bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum             float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum    bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength           int64                                                             `json:"maxLength,omitempty"`
	MinLength           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern             string                                                            `json:"pattern,omitempty"`
	MaxItems            int64                                                             `json:"maxItems,omitempty"`
	MinItems            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems         bool                                                              `json:"uniqueItems,omitempty"`
	Enum                []interface{}                                                     `json:"enum,omitempty"`
	MultipleOf          float64                                                           `json:"multipleOf,omitempty"`
}

type marshalPrimitivesItems PrimitivesItems

// UnmarshalJSON decodes JSON
func (i *PrimitivesItems) UnmarshalJSON(data []byte) error {
	ii := marshalPrimitivesItems(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"type",
			"format",
			"items",
			"collectionFormat",
			"default",
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
			"enum",
			"multipleOf",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = PrimitivesItems(ii)
	return err
}

// MarshalJSON encodes JSON
func (i PrimitivesItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPrimitivesItems(i), i.MapOfAnythingValues)
}

// NonBodyParameter structure is generated from #/definitions/nonBodyParameter
type NonBodyParameter struct {
	HeaderParameterSubSchema   *HeaderParameterSubSchema   `json:"-"`
	FormDataParameterSubSchema *FormDataParameterSubSchema `json:"-"`
	QueryParameterSubSchema    *QueryParameterSubSchema    `json:"-"`
	PathParameterSubSchema     *PathParameterSubSchema     `json:"-"`
}

type marshalNonBodyParameter NonBodyParameter

// UnmarshalJSON decodes JSON
func (i *NonBodyParameter) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.HeaderParameterSubSchema, &i.FormDataParameterSubSchema, &i.QueryParameterSubSchema, &i.PathParameterSubSchema}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.HeaderParameterSubSchema = nil
	}
	if mayUnmarshal[1] == nil {
		i.FormDataParameterSubSchema = nil
	}
	if mayUnmarshal[2] == nil {
		i.QueryParameterSubSchema = nil
	}
	if mayUnmarshal[3] == nil {
		i.PathParameterSubSchema = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i NonBodyParameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalNonBodyParameter(i), i.HeaderParameterSubSchema, i.FormDataParameterSubSchema, i.QueryParameterSubSchema, i.PathParameterSubSchema)
}

// FormDataParameterSubSchema structure is generated from #/definitions/formDataParameterSubSchema
type FormDataParameterSubSchema struct {
	Required            bool                                                              `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	Description         string                                                            `json:"description,omitempty"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                                                            `json:"name,omitempty"`             // The name of the parameter.
	AllowEmptyValue     bool                                                              `json:"allowEmptyValue,omitempty"`  // allows sending a parameter by name only or with an empty value.
	Type                FormDataParameterSubSchemaType                                    `json:"type,omitempty"`
	Format              string                                                            `json:"format,omitempty"`
	Items               *PrimitivesItems                                                  `json:"items,omitempty"`
	CollectionFormat    CollectionFormatWithMulti                                         `json:"collectionFormat,omitempty"`
	Default             interface{}                                                       `json:"default,omitempty"`
	Maximum             float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum    bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum             float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum    bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength           int64                                                             `json:"maxLength,omitempty"`
	MinLength           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern             string                                                            `json:"pattern,omitempty"`
	MaxItems            int64                                                             `json:"maxItems,omitempty"`
	MinItems            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems         bool                                                              `json:"uniqueItems,omitempty"`
	Enum                []interface{}                                                     `json:"enum,omitempty"`
	MultipleOf          float64                                                           `json:"multipleOf,omitempty"`
	MapOfAnythingValues map[string]interface{}                                            `json:"-"`                          // Key must match pattern: ^x-
}

type marshalFormDataParameterSubSchema FormDataParameterSubSchema

// UnmarshalJSON decodes JSON
func (i *FormDataParameterSubSchema) UnmarshalJSON(data []byte) error {
	ii := marshalFormDataParameterSubSchema(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"required",
			"description",
			"name",
			"allowEmptyValue",
			"type",
			"format",
			"items",
			"collectionFormat",
			"default",
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
			"enum",
			"multipleOf",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["in"];!ok || string(v) != `"formData"` {
	    return errors.New(`bad or missing const value for "in" ("formData" expected)`)
	}
	if err != nil {
		return err
	}
	*i = FormDataParameterSubSchema(ii)
	return err
}

var (
	// constFormDataParameterSubSchema is unconditionally added to JSON
	constFormDataParameterSubSchema = json.RawMessage(`{"in":"formData"}`)
)

// MarshalJSON encodes JSON
func (i FormDataParameterSubSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalFormDataParameterSubSchema(i), i.MapOfAnythingValues, constFormDataParameterSubSchema)
}

// QueryParameterSubSchema structure is generated from #/definitions/queryParameterSubSchema
type QueryParameterSubSchema struct {
	Required            bool                                                              `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	Description         string                                                            `json:"description,omitempty"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                                                            `json:"name,omitempty"`             // The name of the parameter.
	AllowEmptyValue     bool                                                              `json:"allowEmptyValue,omitempty"`  // allows sending a parameter by name only or with an empty value.
	Type                QueryParameterSubSchemaType                                       `json:"type,omitempty"`
	Format              string                                                            `json:"format,omitempty"`
	Items               *PrimitivesItems                                                  `json:"items,omitempty"`
	CollectionFormat    CollectionFormatWithMulti                                         `json:"collectionFormat,omitempty"`
	Default             interface{}                                                       `json:"default,omitempty"`
	Maximum             float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum    bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum             float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum    bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength           int64                                                             `json:"maxLength,omitempty"`
	MinLength           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern             string                                                            `json:"pattern,omitempty"`
	MaxItems            int64                                                             `json:"maxItems,omitempty"`
	MinItems            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems         bool                                                              `json:"uniqueItems,omitempty"`
	Enum                []interface{}                                                     `json:"enum,omitempty"`
	MultipleOf          float64                                                           `json:"multipleOf,omitempty"`
	MapOfAnythingValues map[string]interface{}                                            `json:"-"`                          // Key must match pattern: ^x-
}

type marshalQueryParameterSubSchema QueryParameterSubSchema

// UnmarshalJSON decodes JSON
func (i *QueryParameterSubSchema) UnmarshalJSON(data []byte) error {
	ii := marshalQueryParameterSubSchema(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"required",
			"description",
			"name",
			"allowEmptyValue",
			"type",
			"format",
			"items",
			"collectionFormat",
			"default",
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
			"enum",
			"multipleOf",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["in"];!ok || string(v) != `"query"` {
	    return errors.New(`bad or missing const value for "in" ("query" expected)`)
	}
	if err != nil {
		return err
	}
	*i = QueryParameterSubSchema(ii)
	return err
}

var (
	// constQueryParameterSubSchema is unconditionally added to JSON
	constQueryParameterSubSchema = json.RawMessage(`{"in":"query"}`)
)

// MarshalJSON encodes JSON
func (i QueryParameterSubSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalQueryParameterSubSchema(i), i.MapOfAnythingValues, constQueryParameterSubSchema)
}

// PathParameterSubSchema structure is generated from #/definitions/pathParameterSubSchema
type PathParameterSubSchema struct {
	Description         string                                                            `json:"description,omitempty"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                                                            `json:"name,omitempty"`             // The name of the parameter.
	Type                PathParameterSubSchemaType                                        `json:"type,omitempty"`
	Format              string                                                            `json:"format,omitempty"`
	Items               *PrimitivesItems                                                  `json:"items,omitempty"`
	CollectionFormat    CollectionFormat                                                  `json:"collectionFormat,omitempty"`
	Default             interface{}                                                       `json:"default,omitempty"`
	Maximum             float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum    bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum             float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum    bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength           int64                                                             `json:"maxLength,omitempty"`
	MinLength           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern             string                                                            `json:"pattern,omitempty"`
	MaxItems            int64                                                             `json:"maxItems,omitempty"`
	MinItems            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems         bool                                                              `json:"uniqueItems,omitempty"`
	Enum                []interface{}                                                     `json:"enum,omitempty"`
	MultipleOf          float64                                                           `json:"multipleOf,omitempty"`
	MapOfAnythingValues map[string]interface{}                                            `json:"-"`                          // Key must match pattern: ^x-
}

type marshalPathParameterSubSchema PathParameterSubSchema

// UnmarshalJSON decodes JSON
func (i *PathParameterSubSchema) UnmarshalJSON(data []byte) error {
	ii := marshalPathParameterSubSchema(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
			"name",
			"type",
			"format",
			"items",
			"collectionFormat",
			"default",
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
			"enum",
			"multipleOf",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["required"];!ok || string(v) != "true" {
	    return errors.New(`bad or missing const value for "required" (true expected)`)
	}
	if v, ok := constValues["in"];!ok || string(v) != `"path"` {
	    return errors.New(`bad or missing const value for "in" ("path" expected)`)
	}
	if err != nil {
		return err
	}
	*i = PathParameterSubSchema(ii)
	return err
}

var (
	// constPathParameterSubSchema is unconditionally added to JSON
	constPathParameterSubSchema = json.RawMessage(`{"required":true,"in":"path"}`)
)

// MarshalJSON encodes JSON
func (i PathParameterSubSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPathParameterSubSchema(i), i.MapOfAnythingValues, constPathParameterSubSchema)
}

// ParametersListItems structure is generated from #/definitions/parametersList->items
type ParametersListItems struct {
	Parameter     *Parameter     `json:"-"`
	JSONReference *JSONReference `json:"-"`
}

type marshalParametersListItems ParametersListItems

// UnmarshalJSON decodes JSON
func (i *ParametersListItems) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Parameter, &i.JSONReference}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Parameter = nil
	}
	if mayUnmarshal[1] == nil {
		i.JSONReference = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i ParametersListItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParametersListItems(i), i.Parameter, i.JSONReference)
}

// JSONReference structure is generated from #/definitions/jsonReference
type JSONReference struct {
	Ref string `json:"$ref,omitempty"`
}

// Response structure is generated from #/definitions/response
type Response struct {
	Description         string                 `json:"description,omitempty"`
	Schema              *ResponseSchema        `json:"schema,omitempty"`
	Headers             map[string]Header      `json:"headers,omitempty"`
	Examples            map[string]interface{} `json:"examples,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalResponse Response

// UnmarshalJSON decodes JSON
func (i *Response) UnmarshalJSON(data []byte) error {
	ii := marshalResponse(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"description",
			"schema",
			"headers",
			"examples",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Response(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Response) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponse(i), i.MapOfAnythingValues)
}

// ResponseSchema structure is generated from #/definitions/response->schema
type ResponseSchema struct {
	Schema     *Schema     `json:"-"`
	FileSchema *FileSchema `json:"-"`
}

type marshalResponseSchema ResponseSchema

// UnmarshalJSON decodes JSON
func (i *ResponseSchema) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Schema, &i.FileSchema}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Schema = nil
	}
	if mayUnmarshal[1] == nil {
		i.FileSchema = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i ResponseSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponseSchema(i), i.Schema, i.FileSchema)
}

// FileSchema structure is generated from #/definitions/fileSchema
// A deterministic version of a JSON Schema object.
type FileSchema struct {
	Format              string                 `json:"format,omitempty"`
	Title               string                 `json:"title,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Default             interface{}            `json:"default,omitempty"`
	Required            []string               `json:"required,omitempty"`
	ReadOnly            bool                   `json:"readOnly,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	Example             interface{}            `json:"example,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalFileSchema FileSchema

// UnmarshalJSON decodes JSON
func (i *FileSchema) UnmarshalJSON(data []byte) error {
	ii := marshalFileSchema(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"format",
			"title",
			"description",
			"default",
			"required",
			"readOnly",
			"externalDocs",
			"example",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"file"` {
	    return errors.New(`bad or missing const value for "type" ("file" expected)`)
	}
	if err != nil {
		return err
	}
	*i = FileSchema(ii)
	return err
}

var (
	// constFileSchema is unconditionally added to JSON
	constFileSchema = json.RawMessage(`{"type":"file"}`)
)

// MarshalJSON encodes JSON
func (i FileSchema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalFileSchema(i), i.MapOfAnythingValues, constFileSchema)
}

// Header structure is generated from #/definitions/header
type Header struct {
	Type                HeaderType                                                        `json:"type,omitempty"`
	Format              string                                                            `json:"format,omitempty"`
	Items               *PrimitivesItems                                                  `json:"items,omitempty"`
	CollectionFormat    CollectionFormat                                                  `json:"collectionFormat,omitempty"`
	Default             interface{}                                                       `json:"default,omitempty"`
	Maximum             float64                                                           `json:"maximum,omitempty"`
	ExclusiveMaximum    bool                                                              `json:"exclusiveMaximum,omitempty"`
	Minimum             float64                                                           `json:"minimum,omitempty"`
	ExclusiveMinimum    bool                                                              `json:"exclusiveMinimum,omitempty"`
	MaxLength           int64                                                             `json:"maxLength,omitempty"`
	MinLength           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength,omitempty"`
	Pattern             string                                                            `json:"pattern,omitempty"`
	MaxItems            int64                                                             `json:"maxItems,omitempty"`
	MinItems            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems,omitempty"`
	UniqueItems         bool                                                              `json:"uniqueItems,omitempty"`
	Enum                []interface{}                                                     `json:"enum,omitempty"`
	MultipleOf          float64                                                           `json:"multipleOf,omitempty"`
	Description         string                                                            `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}                                            `json:"-"`                          // Key must match pattern: ^x-
}

type marshalHeader Header

// UnmarshalJSON decodes JSON
func (i *Header) UnmarshalJSON(data []byte) error {
	ii := marshalHeader(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"type",
			"format",
			"items",
			"collectionFormat",
			"default",
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
			"enum",
			"multipleOf",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Header(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Header) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalHeader(i), i.MapOfAnythingValues)
}

// ResponseValue structure is generated from #/definitions/responseValue
type ResponseValue struct {
	Response      *Response      `json:"-"`
	JSONReference *JSONReference `json:"-"`
}

type marshalResponseValue ResponseValue

// UnmarshalJSON decodes JSON
func (i *ResponseValue) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Response, &i.JSONReference}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Response = nil
	}
	if mayUnmarshal[1] == nil {
		i.JSONReference = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i ResponseValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponseValue(i), i.Response, i.JSONReference)
}

// Responses structure is generated from #/definitions/responses
// Response objects names can either be any valid HTTP status code or 'default'.
type Responses struct {
	MapOfResponseValueValues map[string]ResponseValue `json:"-"` // Key must match pattern: ^([0-9]{3})$|^(default)$
	MapOfAnythingValues      map[string]interface{}   `json:"-"` // Key must match pattern: ^x-
}

type marshalResponses Responses

// UnmarshalJSON decodes JSON
func (i *Responses) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			"^([0-9]{3})$|^(default)$": &i.MapOfResponseValueValues,
			"^x-": &i.MapOfAnythingValues,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Responses) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponses(i), i.MapOfResponseValueValues, i.MapOfAnythingValues)
}

// BasicAuthenticationSecurity structure is generated from #/definitions/basicAuthenticationSecurity
type BasicAuthenticationSecurity struct {
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalBasicAuthenticationSecurity BasicAuthenticationSecurity

// UnmarshalJSON decodes JSON
func (i *BasicAuthenticationSecurity) UnmarshalJSON(data []byte) error {
	ii := marshalBasicAuthenticationSecurity(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"basic"` {
	    return errors.New(`bad or missing const value for "type" ("basic" expected)`)
	}
	if err != nil {
		return err
	}
	*i = BasicAuthenticationSecurity(ii)
	return err
}

var (
	// constBasicAuthenticationSecurity is unconditionally added to JSON
	constBasicAuthenticationSecurity = json.RawMessage(`{"type":"basic"}`)
)

// MarshalJSON encodes JSON
func (i BasicAuthenticationSecurity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalBasicAuthenticationSecurity(i), i.MapOfAnythingValues, constBasicAuthenticationSecurity)
}

// SecurityDefinitionsAdditionalProperties structure is generated from #/definitions/securityDefinitions->additionalProperties
type SecurityDefinitionsAdditionalProperties struct {
	BasicAuthenticationSecurity *BasicAuthenticationSecurity `json:"-"`
	APIKeySecurity              *APIKeySecurity              `json:"-"`
	Oauth2ImplicitSecurity      *Oauth2ImplicitSecurity      `json:"-"`
	Oauth2PasswordSecurity      *Oauth2PasswordSecurity      `json:"-"`
	Oauth2ApplicationSecurity   *Oauth2ApplicationSecurity   `json:"-"`
	Oauth2AccessCodeSecurity    *Oauth2AccessCodeSecurity    `json:"-"`
}

type marshalSecurityDefinitionsAdditionalProperties SecurityDefinitionsAdditionalProperties

// UnmarshalJSON decodes JSON
func (i *SecurityDefinitionsAdditionalProperties) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.BasicAuthenticationSecurity, &i.APIKeySecurity, &i.Oauth2ImplicitSecurity, &i.Oauth2PasswordSecurity, &i.Oauth2ApplicationSecurity, &i.Oauth2AccessCodeSecurity}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.BasicAuthenticationSecurity = nil
	}
	if mayUnmarshal[1] == nil {
		i.APIKeySecurity = nil
	}
	if mayUnmarshal[2] == nil {
		i.Oauth2ImplicitSecurity = nil
	}
	if mayUnmarshal[3] == nil {
		i.Oauth2PasswordSecurity = nil
	}
	if mayUnmarshal[4] == nil {
		i.Oauth2ApplicationSecurity = nil
	}
	if mayUnmarshal[5] == nil {
		i.Oauth2AccessCodeSecurity = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i SecurityDefinitionsAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSecurityDefinitionsAdditionalProperties(i), i.BasicAuthenticationSecurity, i.APIKeySecurity, i.Oauth2ImplicitSecurity, i.Oauth2PasswordSecurity, i.Oauth2ApplicationSecurity, i.Oauth2AccessCodeSecurity)
}

// APIKeySecurity structure is generated from #/definitions/apiKeySecurity
type APIKeySecurity struct {
	Name                string                 `json:"name,omitempty"`
	In                  APIKeySecurityIn       `json:"in,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeySecurity APIKeySecurity

// UnmarshalJSON decodes JSON
func (i *APIKeySecurity) UnmarshalJSON(data []byte) error {
	ii := marshalAPIKeySecurity(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"name",
			"in",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"apiKey"` {
	    return errors.New(`bad or missing const value for "type" ("apiKey" expected)`)
	}
	if err != nil {
		return err
	}
	*i = APIKeySecurity(ii)
	return err
}

var (
	// constAPIKeySecurity is unconditionally added to JSON
	constAPIKeySecurity = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON
func (i APIKeySecurity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAPIKeySecurity(i), i.MapOfAnythingValues, constAPIKeySecurity)
}

// Oauth2ImplicitSecurity structure is generated from #/definitions/oauth2ImplicitSecurity
type Oauth2ImplicitSecurity struct {
	Scopes              map[string]string      `json:"scopes,omitempty"`
	AuthorizationURL    string                 `json:"authorizationUrl,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-
}

type marshalOauth2ImplicitSecurity Oauth2ImplicitSecurity

// UnmarshalJSON decodes JSON
func (i *Oauth2ImplicitSecurity) UnmarshalJSON(data []byte) error {
	ii := marshalOauth2ImplicitSecurity(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"scopes",
			"authorizationUrl",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"oauth2"` {
	    return errors.New(`bad or missing const value for "type" ("oauth2" expected)`)
	}
	if v, ok := constValues["flow"];!ok || string(v) != `"implicit"` {
	    return errors.New(`bad or missing const value for "flow" ("implicit" expected)`)
	}
	if err != nil {
		return err
	}
	*i = Oauth2ImplicitSecurity(ii)
	return err
}

var (
	// constOauth2ImplicitSecurity is unconditionally added to JSON
	constOauth2ImplicitSecurity = json.RawMessage(`{"type":"oauth2","flow":"implicit"}`)
)

// MarshalJSON encodes JSON
func (i Oauth2ImplicitSecurity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2ImplicitSecurity(i), i.MapOfAnythingValues, constOauth2ImplicitSecurity)
}

// Oauth2PasswordSecurity structure is generated from #/definitions/oauth2PasswordSecurity
type Oauth2PasswordSecurity struct {
	Scopes              map[string]string      `json:"scopes,omitempty"`
	TokenURL            string                 `json:"tokenUrl,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalOauth2PasswordSecurity Oauth2PasswordSecurity

// UnmarshalJSON decodes JSON
func (i *Oauth2PasswordSecurity) UnmarshalJSON(data []byte) error {
	ii := marshalOauth2PasswordSecurity(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"scopes",
			"tokenUrl",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"oauth2"` {
	    return errors.New(`bad or missing const value for "type" ("oauth2" expected)`)
	}
	if v, ok := constValues["flow"];!ok || string(v) != `"password"` {
	    return errors.New(`bad or missing const value for "flow" ("password" expected)`)
	}
	if err != nil {
		return err
	}
	*i = Oauth2PasswordSecurity(ii)
	return err
}

var (
	// constOauth2PasswordSecurity is unconditionally added to JSON
	constOauth2PasswordSecurity = json.RawMessage(`{"type":"oauth2","flow":"password"}`)
)

// MarshalJSON encodes JSON
func (i Oauth2PasswordSecurity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2PasswordSecurity(i), i.MapOfAnythingValues, constOauth2PasswordSecurity)
}

// Oauth2ApplicationSecurity structure is generated from #/definitions/oauth2ApplicationSecurity
type Oauth2ApplicationSecurity struct {
	Scopes              map[string]string      `json:"scopes,omitempty"`
	TokenURL            string                 `json:"tokenUrl,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalOauth2ApplicationSecurity Oauth2ApplicationSecurity

// UnmarshalJSON decodes JSON
func (i *Oauth2ApplicationSecurity) UnmarshalJSON(data []byte) error {
	ii := marshalOauth2ApplicationSecurity(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"scopes",
			"tokenUrl",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"oauth2"` {
	    return errors.New(`bad or missing const value for "type" ("oauth2" expected)`)
	}
	if v, ok := constValues["flow"];!ok || string(v) != `"application"` {
	    return errors.New(`bad or missing const value for "flow" ("application" expected)`)
	}
	if err != nil {
		return err
	}
	*i = Oauth2ApplicationSecurity(ii)
	return err
}

var (
	// constOauth2ApplicationSecurity is unconditionally added to JSON
	constOauth2ApplicationSecurity = json.RawMessage(`{"type":"oauth2","flow":"application"}`)
)

// MarshalJSON encodes JSON
func (i Oauth2ApplicationSecurity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2ApplicationSecurity(i), i.MapOfAnythingValues, constOauth2ApplicationSecurity)
}

// Oauth2AccessCodeSecurity structure is generated from #/definitions/oauth2AccessCodeSecurity
type Oauth2AccessCodeSecurity struct {
	Scopes              map[string]string      `json:"scopes,omitempty"`
	AuthorizationURL    string                 `json:"authorizationUrl,omitempty"`
	TokenURL            string                 `json:"tokenUrl,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                          // Key must match pattern: ^x-
}

type marshalOauth2AccessCodeSecurity Oauth2AccessCodeSecurity

// UnmarshalJSON decodes JSON
func (i *Oauth2AccessCodeSecurity) UnmarshalJSON(data []byte) error {
	ii := marshalOauth2AccessCodeSecurity(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"scopes",
			"authorizationUrl",
			"tokenUrl",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"oauth2"` {
	    return errors.New(`bad or missing const value for "type" ("oauth2" expected)`)
	}
	if v, ok := constValues["flow"];!ok || string(v) != `"accessCode"` {
	    return errors.New(`bad or missing const value for "flow" ("accessCode" expected)`)
	}
	if err != nil {
		return err
	}
	*i = Oauth2AccessCodeSecurity(ii)
	return err
}

var (
	// constOauth2AccessCodeSecurity is unconditionally added to JSON
	constOauth2AccessCodeSecurity = json.RawMessage(`{"type":"oauth2","flow":"accessCode"}`)
)

// MarshalJSON encodes JSON
func (i Oauth2AccessCodeSecurity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOauth2AccessCodeSecurity(i), i.MapOfAnythingValues, constOauth2AccessCodeSecurity)
}

// Tag structure is generated from #/definitions/tag
type Tag struct {
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalTag Tag

// UnmarshalJSON decodes JSON
func (i *Tag) UnmarshalJSON(data []byte) error {
	ii := marshalTag(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"description",
			"externalDocs",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Tag(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(i), i.MapOfAnythingValues)
}

// SchemesListItems is a constant type
type SchemesListItems string

// SchemesListItems values enumeration
const (
	SchemesListItemsHTTP = SchemesListItems("http")
	SchemesListItemsHTTPS = SchemesListItems("https")
	SchemesListItemsWs = SchemesListItems("ws")
	SchemesListItemsWss = SchemesListItems("wss")
)

// MarshalJSON encodes JSON
func (i SchemesListItems) MarshalJSON() ([]byte, error) {
	switch i {
	case SchemesListItemsHTTP:
	case SchemesListItemsHTTPS:
	case SchemesListItemsWs:
	case SchemesListItemsWss:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *SchemesListItems) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := SchemesListItems(ii)
	switch v {
	case SchemesListItemsHTTP:
	case SchemesListItemsHTTPS:
	case SchemesListItemsWs:
	case SchemesListItemsWss:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// HeaderParameterSubSchemaType is a constant type
type HeaderParameterSubSchemaType string

// HeaderParameterSubSchemaType values enumeration
const (
	HeaderParameterSubSchemaTypeString = HeaderParameterSubSchemaType("string")
	HeaderParameterSubSchemaTypeNumber = HeaderParameterSubSchemaType("number")
	HeaderParameterSubSchemaTypeBoolean = HeaderParameterSubSchemaType("boolean")
	HeaderParameterSubSchemaTypeInteger = HeaderParameterSubSchemaType("integer")
	HeaderParameterSubSchemaTypeArray = HeaderParameterSubSchemaType("array")
)

// MarshalJSON encodes JSON
func (i HeaderParameterSubSchemaType) MarshalJSON() ([]byte, error) {
	switch i {
	case HeaderParameterSubSchemaTypeString:
	case HeaderParameterSubSchemaTypeNumber:
	case HeaderParameterSubSchemaTypeBoolean:
	case HeaderParameterSubSchemaTypeInteger:
	case HeaderParameterSubSchemaTypeArray:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *HeaderParameterSubSchemaType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := HeaderParameterSubSchemaType(ii)
	switch v {
	case HeaderParameterSubSchemaTypeString:
	case HeaderParameterSubSchemaTypeNumber:
	case HeaderParameterSubSchemaTypeBoolean:
	case HeaderParameterSubSchemaTypeInteger:
	case HeaderParameterSubSchemaTypeArray:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// PrimitivesItemsType is a constant type
type PrimitivesItemsType string

// PrimitivesItemsType values enumeration
const (
	PrimitivesItemsTypeString = PrimitivesItemsType("string")
	PrimitivesItemsTypeNumber = PrimitivesItemsType("number")
	PrimitivesItemsTypeInteger = PrimitivesItemsType("integer")
	PrimitivesItemsTypeBoolean = PrimitivesItemsType("boolean")
	PrimitivesItemsTypeArray = PrimitivesItemsType("array")
)

// MarshalJSON encodes JSON
func (i PrimitivesItemsType) MarshalJSON() ([]byte, error) {
	switch i {
	case PrimitivesItemsTypeString:
	case PrimitivesItemsTypeNumber:
	case PrimitivesItemsTypeInteger:
	case PrimitivesItemsTypeBoolean:
	case PrimitivesItemsTypeArray:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *PrimitivesItemsType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := PrimitivesItemsType(ii)
	switch v {
	case PrimitivesItemsTypeString:
	case PrimitivesItemsTypeNumber:
	case PrimitivesItemsTypeInteger:
	case PrimitivesItemsTypeBoolean:
	case PrimitivesItemsTypeArray:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// CollectionFormat is a constant type
type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat("csv")
	CollectionFormatSsv = CollectionFormat("ssv")
	CollectionFormatTsv = CollectionFormat("tsv")
	CollectionFormatPipes = CollectionFormat("pipes")
)

// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
	switch i {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := CollectionFormat(ii)
	switch v {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// CollectionFormat is a constant type
type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat("csv")
	CollectionFormatSsv = CollectionFormat("ssv")
	CollectionFormatTsv = CollectionFormat("tsv")
	CollectionFormatPipes = CollectionFormat("pipes")
)

// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
	switch i {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := CollectionFormat(ii)
	switch v {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// FormDataParameterSubSchemaType is a constant type
type FormDataParameterSubSchemaType string

// FormDataParameterSubSchemaType values enumeration
const (
	FormDataParameterSubSchemaTypeString = FormDataParameterSubSchemaType("string")
	FormDataParameterSubSchemaTypeNumber = FormDataParameterSubSchemaType("number")
	FormDataParameterSubSchemaTypeBoolean = FormDataParameterSubSchemaType("boolean")
	FormDataParameterSubSchemaTypeInteger = FormDataParameterSubSchemaType("integer")
	FormDataParameterSubSchemaTypeArray = FormDataParameterSubSchemaType("array")
	FormDataParameterSubSchemaTypeFile = FormDataParameterSubSchemaType("file")
)

// MarshalJSON encodes JSON
func (i FormDataParameterSubSchemaType) MarshalJSON() ([]byte, error) {
	switch i {
	case FormDataParameterSubSchemaTypeString:
	case FormDataParameterSubSchemaTypeNumber:
	case FormDataParameterSubSchemaTypeBoolean:
	case FormDataParameterSubSchemaTypeInteger:
	case FormDataParameterSubSchemaTypeArray:
	case FormDataParameterSubSchemaTypeFile:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *FormDataParameterSubSchemaType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := FormDataParameterSubSchemaType(ii)
	switch v {
	case FormDataParameterSubSchemaTypeString:
	case FormDataParameterSubSchemaTypeNumber:
	case FormDataParameterSubSchemaTypeBoolean:
	case FormDataParameterSubSchemaTypeInteger:
	case FormDataParameterSubSchemaTypeArray:
	case FormDataParameterSubSchemaTypeFile:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// CollectionFormatWithMulti is a constant type
type CollectionFormatWithMulti string

// CollectionFormatWithMulti values enumeration
const (
	CollectionFormatWithMultiCsv = CollectionFormatWithMulti("csv")
	CollectionFormatWithMultiSsv = CollectionFormatWithMulti("ssv")
	CollectionFormatWithMultiTsv = CollectionFormatWithMulti("tsv")
	CollectionFormatWithMultiPipes = CollectionFormatWithMulti("pipes")
	CollectionFormatWithMultiMulti = CollectionFormatWithMulti("multi")
)

// MarshalJSON encodes JSON
func (i CollectionFormatWithMulti) MarshalJSON() ([]byte, error) {
	switch i {
	case CollectionFormatWithMultiCsv:
	case CollectionFormatWithMultiSsv:
	case CollectionFormatWithMultiTsv:
	case CollectionFormatWithMultiPipes:
	case CollectionFormatWithMultiMulti:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *CollectionFormatWithMulti) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := CollectionFormatWithMulti(ii)
	switch v {
	case CollectionFormatWithMultiCsv:
	case CollectionFormatWithMultiSsv:
	case CollectionFormatWithMultiTsv:
	case CollectionFormatWithMultiPipes:
	case CollectionFormatWithMultiMulti:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// QueryParameterSubSchemaType is a constant type
type QueryParameterSubSchemaType string

// QueryParameterSubSchemaType values enumeration
const (
	QueryParameterSubSchemaTypeString = QueryParameterSubSchemaType("string")
	QueryParameterSubSchemaTypeNumber = QueryParameterSubSchemaType("number")
	QueryParameterSubSchemaTypeBoolean = QueryParameterSubSchemaType("boolean")
	QueryParameterSubSchemaTypeInteger = QueryParameterSubSchemaType("integer")
	QueryParameterSubSchemaTypeArray = QueryParameterSubSchemaType("array")
)

// MarshalJSON encodes JSON
func (i QueryParameterSubSchemaType) MarshalJSON() ([]byte, error) {
	switch i {
	case QueryParameterSubSchemaTypeString:
	case QueryParameterSubSchemaTypeNumber:
	case QueryParameterSubSchemaTypeBoolean:
	case QueryParameterSubSchemaTypeInteger:
	case QueryParameterSubSchemaTypeArray:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *QueryParameterSubSchemaType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := QueryParameterSubSchemaType(ii)
	switch v {
	case QueryParameterSubSchemaTypeString:
	case QueryParameterSubSchemaTypeNumber:
	case QueryParameterSubSchemaTypeBoolean:
	case QueryParameterSubSchemaTypeInteger:
	case QueryParameterSubSchemaTypeArray:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// CollectionFormatWithMulti is a constant type
type CollectionFormatWithMulti string

// CollectionFormatWithMulti values enumeration
const (
	CollectionFormatWithMultiCsv = CollectionFormatWithMulti("csv")
	CollectionFormatWithMultiSsv = CollectionFormatWithMulti("ssv")
	CollectionFormatWithMultiTsv = CollectionFormatWithMulti("tsv")
	CollectionFormatWithMultiPipes = CollectionFormatWithMulti("pipes")
	CollectionFormatWithMultiMulti = CollectionFormatWithMulti("multi")
)

// MarshalJSON encodes JSON
func (i CollectionFormatWithMulti) MarshalJSON() ([]byte, error) {
	switch i {
	case CollectionFormatWithMultiCsv:
	case CollectionFormatWithMultiSsv:
	case CollectionFormatWithMultiTsv:
	case CollectionFormatWithMultiPipes:
	case CollectionFormatWithMultiMulti:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *CollectionFormatWithMulti) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := CollectionFormatWithMulti(ii)
	switch v {
	case CollectionFormatWithMultiCsv:
	case CollectionFormatWithMultiSsv:
	case CollectionFormatWithMultiTsv:
	case CollectionFormatWithMultiPipes:
	case CollectionFormatWithMultiMulti:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// PathParameterSubSchemaType is a constant type
type PathParameterSubSchemaType string

// PathParameterSubSchemaType values enumeration
const (
	PathParameterSubSchemaTypeString = PathParameterSubSchemaType("string")
	PathParameterSubSchemaTypeNumber = PathParameterSubSchemaType("number")
	PathParameterSubSchemaTypeBoolean = PathParameterSubSchemaType("boolean")
	PathParameterSubSchemaTypeInteger = PathParameterSubSchemaType("integer")
	PathParameterSubSchemaTypeArray = PathParameterSubSchemaType("array")
)

// MarshalJSON encodes JSON
func (i PathParameterSubSchemaType) MarshalJSON() ([]byte, error) {
	switch i {
	case PathParameterSubSchemaTypeString:
	case PathParameterSubSchemaTypeNumber:
	case PathParameterSubSchemaTypeBoolean:
	case PathParameterSubSchemaTypeInteger:
	case PathParameterSubSchemaTypeArray:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *PathParameterSubSchemaType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := PathParameterSubSchemaType(ii)
	switch v {
	case PathParameterSubSchemaTypeString:
	case PathParameterSubSchemaTypeNumber:
	case PathParameterSubSchemaTypeBoolean:
	case PathParameterSubSchemaTypeInteger:
	case PathParameterSubSchemaTypeArray:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// CollectionFormat is a constant type
type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat("csv")
	CollectionFormatSsv = CollectionFormat("ssv")
	CollectionFormatTsv = CollectionFormat("tsv")
	CollectionFormatPipes = CollectionFormat("pipes")
)

// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
	switch i {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := CollectionFormat(ii)
	switch v {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// HeaderType is a constant type
type HeaderType string

// HeaderType values enumeration
const (
	HeaderTypeString = HeaderType("string")
	HeaderTypeNumber = HeaderType("number")
	HeaderTypeInteger = HeaderType("integer")
	HeaderTypeBoolean = HeaderType("boolean")
	HeaderTypeArray = HeaderType("array")
)

// MarshalJSON encodes JSON
func (i HeaderType) MarshalJSON() ([]byte, error) {
	switch i {
	case HeaderTypeString:
	case HeaderTypeNumber:
	case HeaderTypeInteger:
	case HeaderTypeBoolean:
	case HeaderTypeArray:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *HeaderType) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := HeaderType(ii)
	switch v {
	case HeaderTypeString:
	case HeaderTypeNumber:
	case HeaderTypeInteger:
	case HeaderTypeBoolean:
	case HeaderTypeArray:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// CollectionFormat is a constant type
type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat("csv")
	CollectionFormatSsv = CollectionFormat("ssv")
	CollectionFormatTsv = CollectionFormat("tsv")
	CollectionFormatPipes = CollectionFormat("pipes")
)

// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
	switch i {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := CollectionFormat(ii)
	switch v {
	case CollectionFormatCsv:
	case CollectionFormatSsv:
	case CollectionFormatTsv:
	case CollectionFormatPipes:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// SchemesListItems is a constant type
type SchemesListItems string

// SchemesListItems values enumeration
const (
	SchemesListItemsHTTP = SchemesListItems("http")
	SchemesListItemsHTTPS = SchemesListItems("https")
	SchemesListItemsWs = SchemesListItems("ws")
	SchemesListItemsWss = SchemesListItems("wss")
)

// MarshalJSON encodes JSON
func (i SchemesListItems) MarshalJSON() ([]byte, error) {
	switch i {
	case SchemesListItemsHTTP:
	case SchemesListItemsHTTPS:
	case SchemesListItemsWs:
	case SchemesListItemsWss:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *SchemesListItems) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := SchemesListItems(ii)
	switch v {
	case SchemesListItemsHTTP:
	case SchemesListItemsHTTPS:
	case SchemesListItemsWs:
	case SchemesListItemsWss:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// APIKeySecurityIn is a constant type
type APIKeySecurityIn string

// APIKeySecurityIn values enumeration
const (
	APIKeySecurityInHeader = APIKeySecurityIn("header")
	APIKeySecurityInQuery = APIKeySecurityIn("query")
)

// MarshalJSON encodes JSON
func (i APIKeySecurityIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeySecurityInHeader:
	case APIKeySecurityInQuery:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *APIKeySecurityIn) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := APIKeySecurityIn(ii)
	switch v {
	case APIKeySecurityInHeader:
	case APIKeySecurityInQuery:

	default:
		return errors.New("unexpected value")
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
	// closing empty result
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
func unmarshalUnion(
	mustUnmarshal []interface{},
	mayUnmarshal []interface{},
	ignoreKeys []string,
	regexMaps map[string]interface{},
	j []byte,
) error {
	for _, item := range mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			return err
		}
	}

	for i, item := range mayUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			mayUnmarshal[i] = nil
		}
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// removing ignored keys (defined in struct)
	for _, i := range ignoreKeys {
		delete(m, i)
	}

	// returning early on empty map
	if len(m) == 0 {
		return nil
	}

	// preparing regexp matchers
	var reg = make(map[string]*regexp.Regexp, len(regexMaps))
	for regex := range regexMaps {
		if regex != "" {
			reg[regex], err = regexp.Compile(regex)
			if err != nil {
				return err //todo use errors.Wrap
			}
		}
	}

	subMapsRaw := make(map[string][]byte, len(regexMaps))
	// iterating map and feeding subMaps
	for key, val := range m {
		matched := false
		var ok bool
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex, exp := range reg {
			if exp.MatchString(key) {
				matched = true
				var subMap []byte
				if subMap, ok = subMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)
				subMap = append(subMap, []byte(*val)...)
				subMap = append(subMap, '}')

				subMapsRaw[regex] = subMap
			}
		}

		// empty regex for additionalProperties
		if !matched {
			var subMap []byte
			if subMap, ok = subMapsRaw[""]; !ok {
				subMap = make([]byte, 1, 100)
				subMap[0] = '{'
			} else {
				subMap = append(subMap[:len(subMap)-1], ',')
			}
			subMap = append(subMap, []byte(keyEscaped)...)
			subMap = append(subMap, []byte(*val)...)
			subMap = append(subMap, '}')

			subMapsRaw[""] = subMap
		}
	}

	for regex := range regexMaps {
		if subMap, ok := subMapsRaw[regex]; !ok {
			continue
		} else {
			err = json.Unmarshal(subMap, regexMaps[regex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
