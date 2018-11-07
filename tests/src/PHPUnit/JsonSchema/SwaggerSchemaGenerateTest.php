<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Schema;

class SwaggerSchemaGenerateTest extends \PHPUnit_Framework_TestCase
{
    public function testGen()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/swagger-schema.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder) {
            if ('#' === $path) {
                $structDef->setName('Properties');
            } elseif (0 === strpos($path, '#/definitions/')) {
                $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
                $structDef->setName($name);
            }
        });
        $builder->getType($schema);

        $goFile = new GoFile('entites');
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());


        $expectedGen = <<<'GO'
package entites

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// Properties structure is generated from #
// A JSON Schema for Swagger 2.0 API.
type Properties struct {
	Swagger             Swagger                                            `json:"swagger,omitempty"`             // The Swagger version of this document.
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

// UnmarshalJSON decodes JSON
func (i *Properties) UnmarshalJSON(data []byte) error {
	type p Properties

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"swagger", "info", "host", "basePath", "schemes", "consumes", "produces", "paths", "definitions", "parameters", "responses", "security", "securityDefinitions", "tags", "externalDocs"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = Properties(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i Properties) MarshalJSON() ([]byte, error) {
	type p Properties

	return marshalUnion(p(i), i.MapOfAnythingValues)
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

// UnmarshalJSON decodes JSON
func (i *Info) UnmarshalJSON(data []byte) error {
	type p Info

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"title", "version", "description", "termsOfService", "contact", "license"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Info

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Contact structure is generated from #/definitions/contact
// Contact information for the owners of the API.
type Contact struct {
	Name                string                 `json:"name,omitempty"`  // The identifying name of the contact person/organization.
	URL                 string                 `json:"url,omitempty"`   // The URL pointing to the contact information.
	Email               string                 `json:"email,omitempty"` // The email address of the contact person/organization.
	MapOfAnythingValues map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Contact) UnmarshalJSON(data []byte) error {
	type p Contact

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "url", "email"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Contact

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// License structure is generated from #/definitions/license
type License struct {
	Name                string                 `json:"name,omitempty"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL                 string                 `json:"url,omitempty"`  // The URL pointing to the license.
	MapOfAnythingValues map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *License) UnmarshalJSON(data []byte) error {
	type p License

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "url"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p License

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Consumes structure is generated from #->consumes
// A list of MIME types accepted by the API.
type Consumes struct {
	AllOf0 []string `json:"-"`
}

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
	type p Consumes

	return marshalUnion(p(i), i.AllOf0)
}

// Produces structure is generated from #->produces
// A list of MIME types the API can produce.
type Produces struct {
	AllOf0 []string `json:"-"`
}

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
	type p Produces

	return marshalUnion(p(i), i.AllOf0)
}

// Paths structure is generated from #/definitions/paths
// Relative paths to the individual endpoints. They must be relative to the 'basePath'.
type Paths struct {
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	MapOfPathItemValues map[string]PathItem    `json:"-"` // Key must match pattern: ^/
}

// UnmarshalJSON decodes JSON
func (i *Paths) UnmarshalJSON(data []byte) error {
	
	
	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			`^x-`: &i.MapOfAnythingValues,
			`^/`: &i.MapOfPathItemValues,
		},
		data,
	)
	
	
	return err
}
		
// MarshalJSON encodes JSON
func (i Paths) MarshalJSON() ([]byte, error) {
	type p Paths

	return marshalUnion(p(i), i.MapOfAnythingValues, i.MapOfPathItemValues)
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

// UnmarshalJSON decodes JSON
func (i *PathItem) UnmarshalJSON(data []byte) error {
	type p PathItem

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"$ref", "get", "put", "post", "delete", "options", "head", "patch", "parameters"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p PathItem

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Operation structure is generated from #/definitions/operation
type Operation struct {
	Tags                []string               `json:"tags,omitempty"`
	Summary             string                 `json:"summary,omitempty"`      // A brief summary of the operation.
	Description         string                 `json:"description,omitempty"`  // A longer description of the operation, GitHub Flavored Markdown is allowed.
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	OperationID         string                 `json:"operationId,omitempty"`  // A unique identifier of the operation.
	Produces            *OperationProduces     `json:"produces,omitempty"`     // A list of MIME types the API can produce.
	Consumes            *OperationConsumes     `json:"consumes,omitempty"`     // A list of MIME types the API can consume.
	Parameters          []ParametersListItems  `json:"parameters,omitempty"`   // The parameters needed to send a valid API call.
	Responses           *Responses             `json:"responses,omitempty"`    // Response objects names can either be any valid HTTP status code or 'default'.
	Schemes             []SchemesListItems     `json:"schemes,omitempty"`      // The transfer protocol of the API.
	Deprecated          bool                   `json:"deprecated,omitempty"`
	Security            []map[string][]string  `json:"security,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Operation) UnmarshalJSON(data []byte) error {
	type p Operation

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"tags", "summary", "description", "externalDocs", "operationId", "produces", "consumes", "parameters", "responses", "schemes", "deprecated", "security"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Operation

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ExternalDocs structure is generated from #/definitions/externalDocs
// information about external documentation
type ExternalDocs struct {
	Description         string                 `json:"description,omitempty"`
	URL                 string                 `json:"url,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *ExternalDocs) UnmarshalJSON(data []byte) error {
	type p ExternalDocs

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"description", "url"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p ExternalDocs

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// OperationProduces structure is generated from #/definitions/operation->produces
// A list of MIME types the API can produce.
type OperationProduces struct {
	AllOf0 []string `json:"-"`
}

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
	type p OperationProduces

	return marshalUnion(p(i), i.AllOf0)
}

// OperationConsumes structure is generated from #/definitions/operation->consumes
// A list of MIME types the API can consume.
type OperationConsumes struct {
	AllOf0 []string `json:"-"`
}

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
	type p OperationConsumes

	return marshalUnion(p(i), i.AllOf0)
}

// BodyParameter structure is generated from #/definitions/bodyParameter
type BodyParameter struct {
	Description         string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                 `json:"name,omitempty"`        // The name of the parameter.
	In                  BodyParameterIn        `json:"in,omitempty"`          // Determines the location of the parameter.
	Required            bool                   `json:"required,omitempty"`    // Determines whether or not this parameter is required or optional.
	Schema              *Schema                `json:"schema,omitempty"`      // A deterministic version of a JSON Schema object.
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *BodyParameter) UnmarshalJSON(data []byte) error {
	type p BodyParameter

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"description", "name", "in", "required", "schema"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = BodyParameter(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i BodyParameter) MarshalJSON() ([]byte, error) {
	type p BodyParameter

	return marshalUnion(p(i), i.MapOfAnythingValues)
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

// UnmarshalJSON decodes JSON
func (i *Schema) UnmarshalJSON(data []byte) error {
	type p Schema

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"$ref", "format", "title", "description", "default", "multipleOf", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "maxProperties", "minProperties", "required", "enum", "additionalProperties", "type", "items", "allOf", "properties", "discriminator", "readOnly", "xml", "externalDocs", "example"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Schema

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 structure is generated from http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0
type HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 struct {
	Int64 *int64 `json:"-"`
}

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
	type p HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0

	return marshalUnion(p(i), i.Int64)
}

// SchemaAdditionalProperties structure is generated from #/definitions/schema->additionalProperties
type SchemaAdditionalProperties struct {
	Schema *Schema `json:"-"`
	Bool   *bool   `json:"-"`
}

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
	type p SchemaAdditionalProperties

	return marshalUnion(p(i), i.Schema, i.Bool)
}

// HTTPJSONSchemaOrgDraft04SchemaPropertiesType structure is generated from http://json-schema.org/draft-04/schema#/properties/type
type HTTPJSONSchemaOrgDraft04SchemaPropertiesType struct {
	AnyOf1 []interface{} `json:"-"`
}

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
	type p HTTPJSONSchemaOrgDraft04SchemaPropertiesType

	return marshalUnion(p(i), i.AnyOf1)
}

// SchemaItems structure is generated from #/definitions/schema->items
type SchemaItems struct {
	Schema *Schema  `json:"-"`
	AnyOf1 []Schema `json:"-"`
}

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
	type p SchemaItems

	return marshalUnion(p(i), i.Schema, i.AnyOf1)
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

// UnmarshalJSON decodes JSON
func (i *XML) UnmarshalJSON(data []byte) error {
	type p XML

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "namespace", "prefix", "attribute", "wrapped"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p XML

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Parameter structure is generated from #/definitions/parameter
type Parameter struct {
	BodyParameter    *BodyParameter    `json:"-"`
	NonBodyParameter *NonBodyParameter `json:"-"`
}

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
	type p Parameter

	return marshalUnion(p(i), i.BodyParameter, i.NonBodyParameter)
}

// HeaderParameterSubSchema structure is generated from #/definitions/headerParameterSubSchema
type HeaderParameterSubSchema struct {
	Required            bool                                                              `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	In                  HeaderParameterSubSchemaIn                                        `json:"in,omitempty"`               // Determines the location of the parameter.
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

// UnmarshalJSON decodes JSON
func (i *HeaderParameterSubSchema) UnmarshalJSON(data []byte) error {
	type p HeaderParameterSubSchema

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"required", "in", "description", "name", "type", "format", "items", "collectionFormat", "default", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "enum", "multipleOf"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = HeaderParameterSubSchema(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i HeaderParameterSubSchema) MarshalJSON() ([]byte, error) {
	type p HeaderParameterSubSchema

	return marshalUnion(p(i), i.MapOfAnythingValues)
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

// UnmarshalJSON decodes JSON
func (i *PrimitivesItems) UnmarshalJSON(data []byte) error {
	type p PrimitivesItems

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "format", "items", "collectionFormat", "default", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "enum", "multipleOf"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p PrimitivesItems

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// NonBodyParameter structure is generated from #/definitions/nonBodyParameter
type NonBodyParameter struct {
	HeaderParameterSubSchema   *HeaderParameterSubSchema   `json:"-"`
	FormDataParameterSubSchema *FormDataParameterSubSchema `json:"-"`
	QueryParameterSubSchema    *QueryParameterSubSchema    `json:"-"`
	PathParameterSubSchema     *PathParameterSubSchema     `json:"-"`
}

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
	type p NonBodyParameter

	return marshalUnion(p(i), i.HeaderParameterSubSchema, i.FormDataParameterSubSchema, i.QueryParameterSubSchema, i.PathParameterSubSchema)
}

// FormDataParameterSubSchema structure is generated from #/definitions/formDataParameterSubSchema
type FormDataParameterSubSchema struct {
	Required            bool                                                              `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	In                  FormDataParameterSubSchemaIn                                      `json:"in,omitempty"`               // Determines the location of the parameter.
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

// UnmarshalJSON decodes JSON
func (i *FormDataParameterSubSchema) UnmarshalJSON(data []byte) error {
	type p FormDataParameterSubSchema

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"required", "in", "description", "name", "allowEmptyValue", "type", "format", "items", "collectionFormat", "default", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "enum", "multipleOf"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = FormDataParameterSubSchema(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i FormDataParameterSubSchema) MarshalJSON() ([]byte, error) {
	type p FormDataParameterSubSchema

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// QueryParameterSubSchema structure is generated from #/definitions/queryParameterSubSchema
type QueryParameterSubSchema struct {
	Required            bool                                                              `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	In                  QueryParameterSubSchemaIn                                         `json:"in,omitempty"`               // Determines the location of the parameter.
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

// UnmarshalJSON decodes JSON
func (i *QueryParameterSubSchema) UnmarshalJSON(data []byte) error {
	type p QueryParameterSubSchema

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"required", "in", "description", "name", "allowEmptyValue", "type", "format", "items", "collectionFormat", "default", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "enum", "multipleOf"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = QueryParameterSubSchema(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i QueryParameterSubSchema) MarshalJSON() ([]byte, error) {
	type p QueryParameterSubSchema

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// PathParameterSubSchema structure is generated from #/definitions/pathParameterSubSchema
type PathParameterSubSchema struct {
	Required            PathParameterSubSchemaRequired                                    `json:"required,omitempty"`         // Determines whether or not this parameter is required or optional.
	In                  PathParameterSubSchemaIn                                          `json:"in,omitempty"`               // Determines the location of the parameter.
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

// UnmarshalJSON decodes JSON
func (i *PathParameterSubSchema) UnmarshalJSON(data []byte) error {
	type p PathParameterSubSchema

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"required", "in", "description", "name", "type", "format", "items", "collectionFormat", "default", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "enum", "multipleOf"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = PathParameterSubSchema(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i PathParameterSubSchema) MarshalJSON() ([]byte, error) {
	type p PathParameterSubSchema

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ParametersListItems structure is generated from #/definitions/parametersList->items
type ParametersListItems struct {
	Parameter     *Parameter     `json:"-"`
	JSONReference *JSONReference `json:"-"`
}

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
	type p ParametersListItems

	return marshalUnion(p(i), i.Parameter, i.JSONReference)
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

// UnmarshalJSON decodes JSON
func (i *Response) UnmarshalJSON(data []byte) error {
	type p Response

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"description", "schema", "headers", "examples"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Response

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ResponseSchema structure is generated from #/definitions/response->schema
type ResponseSchema struct {
	Schema     *Schema     `json:"-"`
	FileSchema *FileSchema `json:"-"`
}

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
	type p ResponseSchema

	return marshalUnion(p(i), i.Schema, i.FileSchema)
}

// FileSchema structure is generated from #/definitions/fileSchema
// A deterministic version of a JSON Schema object.
type FileSchema struct {
	Format              string                 `json:"format,omitempty"`
	Title               string                 `json:"title,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Default             interface{}            `json:"default,omitempty"`
	Required            []string               `json:"required,omitempty"`
	Type                FileSchemaType         `json:"type,omitempty"`
	ReadOnly            bool                   `json:"readOnly,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	Example             interface{}            `json:"example,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *FileSchema) UnmarshalJSON(data []byte) error {
	type p FileSchema

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"format", "title", "description", "default", "required", "type", "readOnly", "externalDocs", "example"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = FileSchema(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i FileSchema) MarshalJSON() ([]byte, error) {
	type p FileSchema

	return marshalUnion(p(i), i.MapOfAnythingValues)
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

// UnmarshalJSON decodes JSON
func (i *Header) UnmarshalJSON(data []byte) error {
	type p Header

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "format", "items", "collectionFormat", "default", "maximum", "exclusiveMaximum", "minimum", "exclusiveMinimum", "maxLength", "minLength", "pattern", "maxItems", "minItems", "uniqueItems", "enum", "multipleOf", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Header

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// ResponseValue structure is generated from #/definitions/responseValue
type ResponseValue struct {
	Response      *Response      `json:"-"`
	JSONReference *JSONReference `json:"-"`
}

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
	type p ResponseValue

	return marshalUnion(p(i), i.Response, i.JSONReference)
}

// Responses structure is generated from #/definitions/responses
// Response objects names can either be any valid HTTP status code or 'default'.
type Responses struct {
	MapOfResponseValueValues map[string]ResponseValue `json:"-"` // Key must match pattern: ^([0-9]{3})$|^(default)$
	MapOfAnythingValues      map[string]interface{}   `json:"-"` // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Responses) UnmarshalJSON(data []byte) error {
	
	
	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			`^([0-9]{3})$|^(default)$`: &i.MapOfResponseValueValues,
			`^x-`: &i.MapOfAnythingValues,
		},
		data,
	)
	
	
	return err
}
		
// MarshalJSON encodes JSON
func (i Responses) MarshalJSON() ([]byte, error) {
	type p Responses

	return marshalUnion(p(i), i.MapOfResponseValueValues, i.MapOfAnythingValues)
}

// BasicAuthenticationSecurity structure is generated from #/definitions/basicAuthenticationSecurity
type BasicAuthenticationSecurity struct {
	Type                BasicAuthenticationSecurityType `json:"type,omitempty"`
	Description         string                          `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}          `json:"-"`                     // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *BasicAuthenticationSecurity) UnmarshalJSON(data []byte) error {
	type p BasicAuthenticationSecurity

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = BasicAuthenticationSecurity(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i BasicAuthenticationSecurity) MarshalJSON() ([]byte, error) {
	type p BasicAuthenticationSecurity

	return marshalUnion(p(i), i.MapOfAnythingValues)
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
	type p SecurityDefinitionsAdditionalProperties

	return marshalUnion(p(i), i.BasicAuthenticationSecurity, i.APIKeySecurity, i.Oauth2ImplicitSecurity, i.Oauth2PasswordSecurity, i.Oauth2ApplicationSecurity, i.Oauth2AccessCodeSecurity)
}

// APIKeySecurity structure is generated from #/definitions/apiKeySecurity
type APIKeySecurity struct {
	Type                APIKeySecurityType     `json:"type,omitempty"`
	Name                string                 `json:"name,omitempty"`
	In                  APIKeySecurityIn       `json:"in,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *APIKeySecurity) UnmarshalJSON(data []byte) error {
	type p APIKeySecurity

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "name", "in", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = APIKeySecurity(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i APIKeySecurity) MarshalJSON() ([]byte, error) {
	type p APIKeySecurity

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Oauth2ImplicitSecurity structure is generated from #/definitions/oauth2ImplicitSecurity
type Oauth2ImplicitSecurity struct {
	Type                Oauth2ImplicitSecurityType `json:"type,omitempty"`
	Flow                Oauth2ImplicitSecurityFlow `json:"flow,omitempty"`
	Scopes              map[string]string          `json:"scopes,omitempty"`
	AuthorizationURL    string                     `json:"authorizationUrl,omitempty"`
	Description         string                     `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}     `json:"-"`                          // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Oauth2ImplicitSecurity) UnmarshalJSON(data []byte) error {
	type p Oauth2ImplicitSecurity

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "flow", "scopes", "authorizationUrl", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = Oauth2ImplicitSecurity(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i Oauth2ImplicitSecurity) MarshalJSON() ([]byte, error) {
	type p Oauth2ImplicitSecurity

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Oauth2PasswordSecurity structure is generated from #/definitions/oauth2PasswordSecurity
type Oauth2PasswordSecurity struct {
	Type                Oauth2PasswordSecurityType `json:"type,omitempty"`
	Flow                Oauth2PasswordSecurityFlow `json:"flow,omitempty"`
	Scopes              map[string]string          `json:"scopes,omitempty"`
	TokenURL            string                     `json:"tokenUrl,omitempty"`
	Description         string                     `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}     `json:"-"`                     // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Oauth2PasswordSecurity) UnmarshalJSON(data []byte) error {
	type p Oauth2PasswordSecurity

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "flow", "scopes", "tokenUrl", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = Oauth2PasswordSecurity(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i Oauth2PasswordSecurity) MarshalJSON() ([]byte, error) {
	type p Oauth2PasswordSecurity

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Oauth2ApplicationSecurity structure is generated from #/definitions/oauth2ApplicationSecurity
type Oauth2ApplicationSecurity struct {
	Type                Oauth2ApplicationSecurityType `json:"type,omitempty"`
	Flow                Oauth2ApplicationSecurityFlow `json:"flow,omitempty"`
	Scopes              map[string]string             `json:"scopes,omitempty"`
	TokenURL            string                        `json:"tokenUrl,omitempty"`
	Description         string                        `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}        `json:"-"`                     // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Oauth2ApplicationSecurity) UnmarshalJSON(data []byte) error {
	type p Oauth2ApplicationSecurity

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "flow", "scopes", "tokenUrl", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = Oauth2ApplicationSecurity(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i Oauth2ApplicationSecurity) MarshalJSON() ([]byte, error) {
	type p Oauth2ApplicationSecurity

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Oauth2AccessCodeSecurity structure is generated from #/definitions/oauth2AccessCodeSecurity
type Oauth2AccessCodeSecurity struct {
	Type                Oauth2AccessCodeSecurityType `json:"type,omitempty"`
	Flow                Oauth2AccessCodeSecurityFlow `json:"flow,omitempty"`
	Scopes              map[string]string            `json:"scopes,omitempty"`
	AuthorizationURL    string                       `json:"authorizationUrl,omitempty"`
	TokenURL            string                       `json:"tokenUrl,omitempty"`
	Description         string                       `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}       `json:"-"`                          // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Oauth2AccessCodeSecurity) UnmarshalJSON(data []byte) error {
	type p Oauth2AccessCodeSecurity

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"type", "flow", "scopes", "authorizationUrl", "tokenUrl", "description"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
		},
		data,
	)
	
	if err != nil {
		return err
	}
	*i = Oauth2AccessCodeSecurity(ii)

	return err
}
		
// MarshalJSON encodes JSON
func (i Oauth2AccessCodeSecurity) MarshalJSON() ([]byte, error) {
	type p Oauth2AccessCodeSecurity

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

// Tag structure is generated from #/definitions/tag
type Tag struct {
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

// UnmarshalJSON decodes JSON
func (i *Tag) UnmarshalJSON(data []byte) error {
	type p Tag

	ii := p(*i)

	
	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{"name", "description", "externalDocs"},
		map[string]interface{}{
			`^x-`: &ii.MapOfAnythingValues,
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
	type p Tag

	return marshalUnion(p(i), i.MapOfAnythingValues)
}

type Swagger string

// Swagger values enumeration
const (
	Swagger20 = Swagger(`2.0`)
)

// UnmarshalJSON decodes JSON
func (i *Swagger) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Swagger(ii)
    for {
        if v == Swagger20 {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Swagger) MarshalJSON() ([]byte, error) {
    for {
        if i == Swagger20 {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type SchemesListItems string

// SchemesListItems values enumeration
const (
	SchemesListItemsHTTP = SchemesListItems(`http`)
	SchemesListItemsHTTPS = SchemesListItems(`https`)
	SchemesListItemsWs = SchemesListItems(`ws`)
	SchemesListItemsWss = SchemesListItems(`wss`)
)

// UnmarshalJSON decodes JSON
func (i *SchemesListItems) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := SchemesListItems(ii)
    for {
        if v == SchemesListItemsHTTP {
	        break
	    }
	    if v == SchemesListItemsHTTPS {
	        break
	    }
	    if v == SchemesListItemsWs {
	        break
	    }
	    if v == SchemesListItemsWss {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i SchemesListItems) MarshalJSON() ([]byte, error) {
    for {
        if i == SchemesListItemsHTTP {
	        break
	    }
	    if i == SchemesListItemsHTTPS {
	        break
	    }
	    if i == SchemesListItemsWs {
	        break
	    }
	    if i == SchemesListItemsWss {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type BodyParameterIn string

// BodyParameterIn values enumeration
const (
	BodyParameterInBody = BodyParameterIn(`body`)
)

// UnmarshalJSON decodes JSON
func (i *BodyParameterIn) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := BodyParameterIn(ii)
    for {
        if v == BodyParameterInBody {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i BodyParameterIn) MarshalJSON() ([]byte, error) {
    for {
        if i == BodyParameterInBody {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type HeaderParameterSubSchemaIn string

// HeaderParameterSubSchemaIn values enumeration
const (
	HeaderParameterSubSchemaInHeader = HeaderParameterSubSchemaIn(`header`)
)

// UnmarshalJSON decodes JSON
func (i *HeaderParameterSubSchemaIn) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := HeaderParameterSubSchemaIn(ii)
    for {
        if v == HeaderParameterSubSchemaInHeader {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i HeaderParameterSubSchemaIn) MarshalJSON() ([]byte, error) {
    for {
        if i == HeaderParameterSubSchemaInHeader {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type HeaderParameterSubSchemaType string

// HeaderParameterSubSchemaType values enumeration
const (
	HeaderParameterSubSchemaTypeString = HeaderParameterSubSchemaType(`string`)
	HeaderParameterSubSchemaTypeNumber = HeaderParameterSubSchemaType(`number`)
	HeaderParameterSubSchemaTypeBoolean = HeaderParameterSubSchemaType(`boolean`)
	HeaderParameterSubSchemaTypeInteger = HeaderParameterSubSchemaType(`integer`)
	HeaderParameterSubSchemaTypeArray = HeaderParameterSubSchemaType(`array`)
)

// UnmarshalJSON decodes JSON
func (i *HeaderParameterSubSchemaType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := HeaderParameterSubSchemaType(ii)
    for {
        if v == HeaderParameterSubSchemaTypeString {
	        break
	    }
	    if v == HeaderParameterSubSchemaTypeNumber {
	        break
	    }
	    if v == HeaderParameterSubSchemaTypeBoolean {
	        break
	    }
	    if v == HeaderParameterSubSchemaTypeInteger {
	        break
	    }
	    if v == HeaderParameterSubSchemaTypeArray {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i HeaderParameterSubSchemaType) MarshalJSON() ([]byte, error) {
    for {
        if i == HeaderParameterSubSchemaTypeString {
	        break
	    }
	    if i == HeaderParameterSubSchemaTypeNumber {
	        break
	    }
	    if i == HeaderParameterSubSchemaTypeBoolean {
	        break
	    }
	    if i == HeaderParameterSubSchemaTypeInteger {
	        break
	    }
	    if i == HeaderParameterSubSchemaTypeArray {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type PrimitivesItemsType string

// PrimitivesItemsType values enumeration
const (
	PrimitivesItemsTypeString = PrimitivesItemsType(`string`)
	PrimitivesItemsTypeNumber = PrimitivesItemsType(`number`)
	PrimitivesItemsTypeInteger = PrimitivesItemsType(`integer`)
	PrimitivesItemsTypeBoolean = PrimitivesItemsType(`boolean`)
	PrimitivesItemsTypeArray = PrimitivesItemsType(`array`)
)

// UnmarshalJSON decodes JSON
func (i *PrimitivesItemsType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := PrimitivesItemsType(ii)
    for {
        if v == PrimitivesItemsTypeString {
	        break
	    }
	    if v == PrimitivesItemsTypeNumber {
	        break
	    }
	    if v == PrimitivesItemsTypeInteger {
	        break
	    }
	    if v == PrimitivesItemsTypeBoolean {
	        break
	    }
	    if v == PrimitivesItemsTypeArray {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i PrimitivesItemsType) MarshalJSON() ([]byte, error) {
    for {
        if i == PrimitivesItemsTypeString {
	        break
	    }
	    if i == PrimitivesItemsTypeNumber {
	        break
	    }
	    if i == PrimitivesItemsTypeInteger {
	        break
	    }
	    if i == PrimitivesItemsTypeBoolean {
	        break
	    }
	    if i == PrimitivesItemsTypeArray {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat(`csv`)
	CollectionFormatSsv = CollectionFormat(`ssv`)
	CollectionFormatTsv = CollectionFormat(`tsv`)
	CollectionFormatPipes = CollectionFormat(`pipes`)
)

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := CollectionFormat(ii)
    for {
        if v == CollectionFormatCsv {
	        break
	    }
	    if v == CollectionFormatSsv {
	        break
	    }
	    if v == CollectionFormatTsv {
	        break
	    }
	    if v == CollectionFormatPipes {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
    for {
        if i == CollectionFormatCsv {
	        break
	    }
	    if i == CollectionFormatSsv {
	        break
	    }
	    if i == CollectionFormatTsv {
	        break
	    }
	    if i == CollectionFormatPipes {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat(`csv`)
	CollectionFormatSsv = CollectionFormat(`ssv`)
	CollectionFormatTsv = CollectionFormat(`tsv`)
	CollectionFormatPipes = CollectionFormat(`pipes`)
)

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := CollectionFormat(ii)
    for {
        if v == CollectionFormatCsv {
	        break
	    }
	    if v == CollectionFormatSsv {
	        break
	    }
	    if v == CollectionFormatTsv {
	        break
	    }
	    if v == CollectionFormatPipes {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
    for {
        if i == CollectionFormatCsv {
	        break
	    }
	    if i == CollectionFormatSsv {
	        break
	    }
	    if i == CollectionFormatTsv {
	        break
	    }
	    if i == CollectionFormatPipes {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type FormDataParameterSubSchemaIn string

// FormDataParameterSubSchemaIn values enumeration
const (
	FormDataParameterSubSchemaInFormData = FormDataParameterSubSchemaIn(`formData`)
)

// UnmarshalJSON decodes JSON
func (i *FormDataParameterSubSchemaIn) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := FormDataParameterSubSchemaIn(ii)
    for {
        if v == FormDataParameterSubSchemaInFormData {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i FormDataParameterSubSchemaIn) MarshalJSON() ([]byte, error) {
    for {
        if i == FormDataParameterSubSchemaInFormData {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type FormDataParameterSubSchemaType string

// FormDataParameterSubSchemaType values enumeration
const (
	FormDataParameterSubSchemaTypeString = FormDataParameterSubSchemaType(`string`)
	FormDataParameterSubSchemaTypeNumber = FormDataParameterSubSchemaType(`number`)
	FormDataParameterSubSchemaTypeBoolean = FormDataParameterSubSchemaType(`boolean`)
	FormDataParameterSubSchemaTypeInteger = FormDataParameterSubSchemaType(`integer`)
	FormDataParameterSubSchemaTypeArray = FormDataParameterSubSchemaType(`array`)
	FormDataParameterSubSchemaTypeFile = FormDataParameterSubSchemaType(`file`)
)

// UnmarshalJSON decodes JSON
func (i *FormDataParameterSubSchemaType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := FormDataParameterSubSchemaType(ii)
    for {
        if v == FormDataParameterSubSchemaTypeString {
	        break
	    }
	    if v == FormDataParameterSubSchemaTypeNumber {
	        break
	    }
	    if v == FormDataParameterSubSchemaTypeBoolean {
	        break
	    }
	    if v == FormDataParameterSubSchemaTypeInteger {
	        break
	    }
	    if v == FormDataParameterSubSchemaTypeArray {
	        break
	    }
	    if v == FormDataParameterSubSchemaTypeFile {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i FormDataParameterSubSchemaType) MarshalJSON() ([]byte, error) {
    for {
        if i == FormDataParameterSubSchemaTypeString {
	        break
	    }
	    if i == FormDataParameterSubSchemaTypeNumber {
	        break
	    }
	    if i == FormDataParameterSubSchemaTypeBoolean {
	        break
	    }
	    if i == FormDataParameterSubSchemaTypeInteger {
	        break
	    }
	    if i == FormDataParameterSubSchemaTypeArray {
	        break
	    }
	    if i == FormDataParameterSubSchemaTypeFile {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type CollectionFormatWithMulti string

// CollectionFormatWithMulti values enumeration
const (
	CollectionFormatWithMultiCsv = CollectionFormatWithMulti(`csv`)
	CollectionFormatWithMultiSsv = CollectionFormatWithMulti(`ssv`)
	CollectionFormatWithMultiTsv = CollectionFormatWithMulti(`tsv`)
	CollectionFormatWithMultiPipes = CollectionFormatWithMulti(`pipes`)
	CollectionFormatWithMultiMulti = CollectionFormatWithMulti(`multi`)
)

// UnmarshalJSON decodes JSON
func (i *CollectionFormatWithMulti) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := CollectionFormatWithMulti(ii)
    for {
        if v == CollectionFormatWithMultiCsv {
	        break
	    }
	    if v == CollectionFormatWithMultiSsv {
	        break
	    }
	    if v == CollectionFormatWithMultiTsv {
	        break
	    }
	    if v == CollectionFormatWithMultiPipes {
	        break
	    }
	    if v == CollectionFormatWithMultiMulti {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i CollectionFormatWithMulti) MarshalJSON() ([]byte, error) {
    for {
        if i == CollectionFormatWithMultiCsv {
	        break
	    }
	    if i == CollectionFormatWithMultiSsv {
	        break
	    }
	    if i == CollectionFormatWithMultiTsv {
	        break
	    }
	    if i == CollectionFormatWithMultiPipes {
	        break
	    }
	    if i == CollectionFormatWithMultiMulti {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type QueryParameterSubSchemaIn string

// QueryParameterSubSchemaIn values enumeration
const (
	QueryParameterSubSchemaInQuery = QueryParameterSubSchemaIn(`query`)
)

// UnmarshalJSON decodes JSON
func (i *QueryParameterSubSchemaIn) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := QueryParameterSubSchemaIn(ii)
    for {
        if v == QueryParameterSubSchemaInQuery {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i QueryParameterSubSchemaIn) MarshalJSON() ([]byte, error) {
    for {
        if i == QueryParameterSubSchemaInQuery {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type QueryParameterSubSchemaType string

// QueryParameterSubSchemaType values enumeration
const (
	QueryParameterSubSchemaTypeString = QueryParameterSubSchemaType(`string`)
	QueryParameterSubSchemaTypeNumber = QueryParameterSubSchemaType(`number`)
	QueryParameterSubSchemaTypeBoolean = QueryParameterSubSchemaType(`boolean`)
	QueryParameterSubSchemaTypeInteger = QueryParameterSubSchemaType(`integer`)
	QueryParameterSubSchemaTypeArray = QueryParameterSubSchemaType(`array`)
)

// UnmarshalJSON decodes JSON
func (i *QueryParameterSubSchemaType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := QueryParameterSubSchemaType(ii)
    for {
        if v == QueryParameterSubSchemaTypeString {
	        break
	    }
	    if v == QueryParameterSubSchemaTypeNumber {
	        break
	    }
	    if v == QueryParameterSubSchemaTypeBoolean {
	        break
	    }
	    if v == QueryParameterSubSchemaTypeInteger {
	        break
	    }
	    if v == QueryParameterSubSchemaTypeArray {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i QueryParameterSubSchemaType) MarshalJSON() ([]byte, error) {
    for {
        if i == QueryParameterSubSchemaTypeString {
	        break
	    }
	    if i == QueryParameterSubSchemaTypeNumber {
	        break
	    }
	    if i == QueryParameterSubSchemaTypeBoolean {
	        break
	    }
	    if i == QueryParameterSubSchemaTypeInteger {
	        break
	    }
	    if i == QueryParameterSubSchemaTypeArray {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type CollectionFormatWithMulti string

// CollectionFormatWithMulti values enumeration
const (
	CollectionFormatWithMultiCsv = CollectionFormatWithMulti(`csv`)
	CollectionFormatWithMultiSsv = CollectionFormatWithMulti(`ssv`)
	CollectionFormatWithMultiTsv = CollectionFormatWithMulti(`tsv`)
	CollectionFormatWithMultiPipes = CollectionFormatWithMulti(`pipes`)
	CollectionFormatWithMultiMulti = CollectionFormatWithMulti(`multi`)
)

// UnmarshalJSON decodes JSON
func (i *CollectionFormatWithMulti) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := CollectionFormatWithMulti(ii)
    for {
        if v == CollectionFormatWithMultiCsv {
	        break
	    }
	    if v == CollectionFormatWithMultiSsv {
	        break
	    }
	    if v == CollectionFormatWithMultiTsv {
	        break
	    }
	    if v == CollectionFormatWithMultiPipes {
	        break
	    }
	    if v == CollectionFormatWithMultiMulti {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i CollectionFormatWithMulti) MarshalJSON() ([]byte, error) {
    for {
        if i == CollectionFormatWithMultiCsv {
	        break
	    }
	    if i == CollectionFormatWithMultiSsv {
	        break
	    }
	    if i == CollectionFormatWithMultiTsv {
	        break
	    }
	    if i == CollectionFormatWithMultiPipes {
	        break
	    }
	    if i == CollectionFormatWithMultiMulti {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type PathParameterSubSchemaRequired bool

// PathParameterSubSchemaRequired values enumeration
const (
	PathParameterSubSchemaRequired1 = PathParameterSubSchemaRequired(1)
)

// UnmarshalJSON decodes JSON
func (i *PathParameterSubSchemaRequired) UnmarshalJSON(data []byte) error {
    var ii bool
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := PathParameterSubSchemaRequired(ii)
    for {
        if v == PathParameterSubSchemaRequired1 {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i PathParameterSubSchemaRequired) MarshalJSON() ([]byte, error) {
    for {
        if i == PathParameterSubSchemaRequired1 {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(bool(i))
}

type PathParameterSubSchemaIn string

// PathParameterSubSchemaIn values enumeration
const (
	PathParameterSubSchemaInPath = PathParameterSubSchemaIn(`path`)
)

// UnmarshalJSON decodes JSON
func (i *PathParameterSubSchemaIn) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := PathParameterSubSchemaIn(ii)
    for {
        if v == PathParameterSubSchemaInPath {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i PathParameterSubSchemaIn) MarshalJSON() ([]byte, error) {
    for {
        if i == PathParameterSubSchemaInPath {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type PathParameterSubSchemaType string

// PathParameterSubSchemaType values enumeration
const (
	PathParameterSubSchemaTypeString = PathParameterSubSchemaType(`string`)
	PathParameterSubSchemaTypeNumber = PathParameterSubSchemaType(`number`)
	PathParameterSubSchemaTypeBoolean = PathParameterSubSchemaType(`boolean`)
	PathParameterSubSchemaTypeInteger = PathParameterSubSchemaType(`integer`)
	PathParameterSubSchemaTypeArray = PathParameterSubSchemaType(`array`)
)

// UnmarshalJSON decodes JSON
func (i *PathParameterSubSchemaType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := PathParameterSubSchemaType(ii)
    for {
        if v == PathParameterSubSchemaTypeString {
	        break
	    }
	    if v == PathParameterSubSchemaTypeNumber {
	        break
	    }
	    if v == PathParameterSubSchemaTypeBoolean {
	        break
	    }
	    if v == PathParameterSubSchemaTypeInteger {
	        break
	    }
	    if v == PathParameterSubSchemaTypeArray {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i PathParameterSubSchemaType) MarshalJSON() ([]byte, error) {
    for {
        if i == PathParameterSubSchemaTypeString {
	        break
	    }
	    if i == PathParameterSubSchemaTypeNumber {
	        break
	    }
	    if i == PathParameterSubSchemaTypeBoolean {
	        break
	    }
	    if i == PathParameterSubSchemaTypeInteger {
	        break
	    }
	    if i == PathParameterSubSchemaTypeArray {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat(`csv`)
	CollectionFormatSsv = CollectionFormat(`ssv`)
	CollectionFormatTsv = CollectionFormat(`tsv`)
	CollectionFormatPipes = CollectionFormat(`pipes`)
)

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := CollectionFormat(ii)
    for {
        if v == CollectionFormatCsv {
	        break
	    }
	    if v == CollectionFormatSsv {
	        break
	    }
	    if v == CollectionFormatTsv {
	        break
	    }
	    if v == CollectionFormatPipes {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
    for {
        if i == CollectionFormatCsv {
	        break
	    }
	    if i == CollectionFormatSsv {
	        break
	    }
	    if i == CollectionFormatTsv {
	        break
	    }
	    if i == CollectionFormatPipes {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type FileSchemaType string

// FileSchemaType values enumeration
const (
	FileSchemaTypeFile = FileSchemaType(`file`)
)

// UnmarshalJSON decodes JSON
func (i *FileSchemaType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := FileSchemaType(ii)
    for {
        if v == FileSchemaTypeFile {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i FileSchemaType) MarshalJSON() ([]byte, error) {
    for {
        if i == FileSchemaTypeFile {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type HeaderType string

// HeaderType values enumeration
const (
	HeaderTypeString = HeaderType(`string`)
	HeaderTypeNumber = HeaderType(`number`)
	HeaderTypeInteger = HeaderType(`integer`)
	HeaderTypeBoolean = HeaderType(`boolean`)
	HeaderTypeArray = HeaderType(`array`)
)

// UnmarshalJSON decodes JSON
func (i *HeaderType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := HeaderType(ii)
    for {
        if v == HeaderTypeString {
	        break
	    }
	    if v == HeaderTypeNumber {
	        break
	    }
	    if v == HeaderTypeInteger {
	        break
	    }
	    if v == HeaderTypeBoolean {
	        break
	    }
	    if v == HeaderTypeArray {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i HeaderType) MarshalJSON() ([]byte, error) {
    for {
        if i == HeaderTypeString {
	        break
	    }
	    if i == HeaderTypeNumber {
	        break
	    }
	    if i == HeaderTypeInteger {
	        break
	    }
	    if i == HeaderTypeBoolean {
	        break
	    }
	    if i == HeaderTypeArray {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type CollectionFormat string

// CollectionFormat values enumeration
const (
	CollectionFormatCsv = CollectionFormat(`csv`)
	CollectionFormatSsv = CollectionFormat(`ssv`)
	CollectionFormatTsv = CollectionFormat(`tsv`)
	CollectionFormatPipes = CollectionFormat(`pipes`)
)

// UnmarshalJSON decodes JSON
func (i *CollectionFormat) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := CollectionFormat(ii)
    for {
        if v == CollectionFormatCsv {
	        break
	    }
	    if v == CollectionFormatSsv {
	        break
	    }
	    if v == CollectionFormatTsv {
	        break
	    }
	    if v == CollectionFormatPipes {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i CollectionFormat) MarshalJSON() ([]byte, error) {
    for {
        if i == CollectionFormatCsv {
	        break
	    }
	    if i == CollectionFormatSsv {
	        break
	    }
	    if i == CollectionFormatTsv {
	        break
	    }
	    if i == CollectionFormatPipes {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type SchemesListItems string

// SchemesListItems values enumeration
const (
	SchemesListItemsHTTP = SchemesListItems(`http`)
	SchemesListItemsHTTPS = SchemesListItems(`https`)
	SchemesListItemsWs = SchemesListItems(`ws`)
	SchemesListItemsWss = SchemesListItems(`wss`)
)

// UnmarshalJSON decodes JSON
func (i *SchemesListItems) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := SchemesListItems(ii)
    for {
        if v == SchemesListItemsHTTP {
	        break
	    }
	    if v == SchemesListItemsHTTPS {
	        break
	    }
	    if v == SchemesListItemsWs {
	        break
	    }
	    if v == SchemesListItemsWss {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i SchemesListItems) MarshalJSON() ([]byte, error) {
    for {
        if i == SchemesListItemsHTTP {
	        break
	    }
	    if i == SchemesListItemsHTTPS {
	        break
	    }
	    if i == SchemesListItemsWs {
	        break
	    }
	    if i == SchemesListItemsWss {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type BasicAuthenticationSecurityType string

// BasicAuthenticationSecurityType values enumeration
const (
	BasicAuthenticationSecurityTypeBasic = BasicAuthenticationSecurityType(`basic`)
)

// UnmarshalJSON decodes JSON
func (i *BasicAuthenticationSecurityType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := BasicAuthenticationSecurityType(ii)
    for {
        if v == BasicAuthenticationSecurityTypeBasic {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i BasicAuthenticationSecurityType) MarshalJSON() ([]byte, error) {
    for {
        if i == BasicAuthenticationSecurityTypeBasic {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type APIKeySecurityType string

// APIKeySecurityType values enumeration
const (
	APIKeySecurityTypeAPIKey = APIKeySecurityType(`apiKey`)
)

// UnmarshalJSON decodes JSON
func (i *APIKeySecurityType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := APIKeySecurityType(ii)
    for {
        if v == APIKeySecurityTypeAPIKey {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i APIKeySecurityType) MarshalJSON() ([]byte, error) {
    for {
        if i == APIKeySecurityTypeAPIKey {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type APIKeySecurityIn string

// APIKeySecurityIn values enumeration
const (
	APIKeySecurityInHeader = APIKeySecurityIn(`header`)
	APIKeySecurityInQuery = APIKeySecurityIn(`query`)
)

// UnmarshalJSON decodes JSON
func (i *APIKeySecurityIn) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := APIKeySecurityIn(ii)
    for {
        if v == APIKeySecurityInHeader {
	        break
	    }
	    if v == APIKeySecurityInQuery {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i APIKeySecurityIn) MarshalJSON() ([]byte, error) {
    for {
        if i == APIKeySecurityInHeader {
	        break
	    }
	    if i == APIKeySecurityInQuery {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2ImplicitSecurityType string

// Oauth2ImplicitSecurityType values enumeration
const (
	Oauth2ImplicitSecurityTypeOauth2 = Oauth2ImplicitSecurityType(`oauth2`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2ImplicitSecurityType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2ImplicitSecurityType(ii)
    for {
        if v == Oauth2ImplicitSecurityTypeOauth2 {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2ImplicitSecurityType) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2ImplicitSecurityTypeOauth2 {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2ImplicitSecurityFlow string

// Oauth2ImplicitSecurityFlow values enumeration
const (
	Oauth2ImplicitSecurityFlowImplicit = Oauth2ImplicitSecurityFlow(`implicit`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2ImplicitSecurityFlow) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2ImplicitSecurityFlow(ii)
    for {
        if v == Oauth2ImplicitSecurityFlowImplicit {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2ImplicitSecurityFlow) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2ImplicitSecurityFlowImplicit {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2PasswordSecurityType string

// Oauth2PasswordSecurityType values enumeration
const (
	Oauth2PasswordSecurityTypeOauth2 = Oauth2PasswordSecurityType(`oauth2`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2PasswordSecurityType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2PasswordSecurityType(ii)
    for {
        if v == Oauth2PasswordSecurityTypeOauth2 {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2PasswordSecurityType) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2PasswordSecurityTypeOauth2 {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2PasswordSecurityFlow string

// Oauth2PasswordSecurityFlow values enumeration
const (
	Oauth2PasswordSecurityFlowPassword = Oauth2PasswordSecurityFlow(`password`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2PasswordSecurityFlow) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2PasswordSecurityFlow(ii)
    for {
        if v == Oauth2PasswordSecurityFlowPassword {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2PasswordSecurityFlow) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2PasswordSecurityFlowPassword {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2ApplicationSecurityType string

// Oauth2ApplicationSecurityType values enumeration
const (
	Oauth2ApplicationSecurityTypeOauth2 = Oauth2ApplicationSecurityType(`oauth2`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2ApplicationSecurityType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2ApplicationSecurityType(ii)
    for {
        if v == Oauth2ApplicationSecurityTypeOauth2 {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2ApplicationSecurityType) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2ApplicationSecurityTypeOauth2 {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2ApplicationSecurityFlow string

// Oauth2ApplicationSecurityFlow values enumeration
const (
	Oauth2ApplicationSecurityFlowApplication = Oauth2ApplicationSecurityFlow(`application`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2ApplicationSecurityFlow) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2ApplicationSecurityFlow(ii)
    for {
        if v == Oauth2ApplicationSecurityFlowApplication {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2ApplicationSecurityFlow) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2ApplicationSecurityFlowApplication {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2AccessCodeSecurityType string

// Oauth2AccessCodeSecurityType values enumeration
const (
	Oauth2AccessCodeSecurityTypeOauth2 = Oauth2AccessCodeSecurityType(`oauth2`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2AccessCodeSecurityType) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2AccessCodeSecurityType(ii)
    for {
        if v == Oauth2AccessCodeSecurityTypeOauth2 {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2AccessCodeSecurityType) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2AccessCodeSecurityTypeOauth2 {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

type Oauth2AccessCodeSecurityFlow string

// Oauth2AccessCodeSecurityFlow values enumeration
const (
	Oauth2AccessCodeSecurityFlowAccessCode = Oauth2AccessCodeSecurityFlow(`accessCode`)
)

// UnmarshalJSON decodes JSON
func (i *Oauth2AccessCodeSecurityFlow) UnmarshalJSON(data []byte) error {
    var ii string
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := Oauth2AccessCodeSecurityFlow(ii)
    for {
        if v == Oauth2AccessCodeSecurityFlowAccessCode {
	        break
	    }

        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i Oauth2AccessCodeSecurityFlow) MarshalJSON() ([]byte, error) {
    for {
        if i == Oauth2AccessCodeSecurityFlowAccessCode {
	        break
	    }

        return nil, errors.New("unexpected value")
    }
    return json.Marshal(string(i))
}

func unmarshalUnion(mustUnmarshal []interface{}, mayUnmarshal []interface{}, ignoreKeys []string, regexMaps map[string]interface{}, j []byte) error {
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
	for regex, _ := range regexMaps {
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
			if exp.Match([]byte(key)) {
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

	for regex, _ := range regexMaps {
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

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	for _, m := range maps {
		if m == nil {
			continue
		}
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
			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if len(result) > 1 {
			result = append(result[:len(result)-1], ',')
		}
		result = append(result, j[1:]...)
	}
	// closing empty result
	if len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}

GO;

        $this->assertSame($expectedGen, $goFile->render());

    }

}