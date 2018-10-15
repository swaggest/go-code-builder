<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\JsonSchema;
use Swaggest\JsonSchema\Schema;

class SwaggerSchemaGenerateTest extends \PHPUnit_Framework_TestCase
{
    public function testJsonSchemaGenerate()
    {
        //$this->markTestSkipped();

        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../../vendor/swaggest/json-schema/spec/swagger-schema.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->getType($schema);
        $expectedGen = <<<'GO'
// #
type Untitled1 struct {
	Swagger             string                            `json:"swagger"`             // The Swagger version of this document.
	Info                *DefinitionsInfo                  `json:"info"`                // General information about the API.
	Host                string                            `json:"host"`                // The host (name or ip) of the API. Example: 'swagger.io'
	BasePath            string                            `json:"basePath"`            // The base path to the API. Example: '/api'.
	Schemes             []string                          `json:"schemes"`             // The transfer protocol of the API.
	Consumes            *Consumes                         `json:"consumes"`            // A list of MIME types accepted by the API.
	Produces            *Produces                         `json:"produces"`            // A list of MIME types the API can produce.
	Paths               *DefinitionsPaths                 `json:"paths"`               // Relative paths to the individual endpoints. They must be relative to the 'basePath'.
	Definitions         *DefinitionsDefinitions           `json:"definitions"`         // One or more JSON objects describing the schemas being consumed and produced by the API.
	Parameters          *DefinitionsParameterDefinitions  `json:"parameters"`          // One or more JSON representations for parameters
	Responses           *DefinitionsResponseDefinitions   `json:"responses"`           // One or more JSON representations for parameters
	Security            []*DefinitionsSecurityRequirement `json:"security"`
	SecurityDefinitions *DefinitionsSecurityDefinitions   `json:"securityDefinitions"`
	Tags                []*DefinitionsTag                 `json:"tags"`
	ExternalDocs        *DefinitionsExternalDocs          `json:"externalDocs"`        // information about external documentation
	patternPropertiesX  map[string]interface{}
}

// #/definitions/info
type DefinitionsInfo struct {
	Title              string                 `json:"title"`          // A unique and precise title of the API.
	Version            string                 `json:"version"`        // A semantic version number of the API.
	Description        string                 `json:"description"`    // A longer description of the API. Should be different from the title.  GitHub Flavored Markdown is allowed.
	TermsOfService     string                 `json:"termsOfService"` // The terms of service for the API.
	Contact            *DefinitionsContact    `json:"contact"`        // Contact information for the owners of the API.
	License            *DefinitionsLicense    `json:"license"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/contact
type DefinitionsContact struct {
	Name               string                 `json:"name"`  // The identifying name of the contact person/organization.
	URL                string                 `json:"url"`   // The URL pointing to the contact information.
	Email              string                 `json:"email"` // The email address of the contact person/organization.
	patternPropertiesX map[string]interface{}
}

// #/definitions/license
type DefinitionsLicense struct {
	Name               string                 `json:"name"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL                string                 `json:"url"`  // The URL pointing to the license.
	patternPropertiesX map[string]interface{}
}

// #->consumes
type Consumes struct {
	[]string
}

// #->produces
type Produces struct {
	[]string
}

// #/definitions/paths
type DefinitionsPaths struct {
	patternPropertiesX map[string]interface{}
	patternProperties  map[string]*DefinitionsPathItem
}

// #/definitions/pathItem
type DefinitionsPathItem struct {
	Ref                string                            `json:"$ref"`
	Get                *DefinitionsOperation             `json:"get"`
	Put                *DefinitionsOperation             `json:"put"`
	Post               *DefinitionsOperation             `json:"post"`
	Delete             *DefinitionsOperation             `json:"delete"`
	Options            *DefinitionsOperation             `json:"options"`
	Head               *DefinitionsOperation             `json:"head"`
	Patch              *DefinitionsOperation             `json:"patch"`
	Parameters         []*DefinitionsParametersListItems `json:"parameters"` // The parameters needed to send a valid API call.
	patternPropertiesX map[string]interface{}
}

// #/definitions/operation
type DefinitionsOperation struct {
	Tags               []string                          `json:"tags"`
	Summary            string                            `json:"summary"`      // A brief summary of the operation.
	Description        string                            `json:"description"`  // A longer description of the operation, GitHub Flavored Markdown is allowed.
	ExternalDocs       *DefinitionsExternalDocs          `json:"externalDocs"` // information about external documentation
	OperationID        string                            `json:"operationId"`  // A unique identifier of the operation.
	Produces           *DefinitionsOperationProduces     `json:"produces"`     // A list of MIME types the API can produce.
	Consumes           *DefinitionsOperationConsumes     `json:"consumes"`     // A list of MIME types the API can consume.
	Parameters         []*DefinitionsParametersListItems `json:"parameters"`   // The parameters needed to send a valid API call.
	Responses          *DefinitionsResponses             `json:"responses"`    // Response objects names can either be any valid HTTP status code or 'default'.
	Schemes            []string                          `json:"schemes"`      // The transfer protocol of the API.
	Deprecated         bool                              `json:"deprecated"`
	Security           []*DefinitionsSecurityRequirement `json:"security"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/externalDocs
type DefinitionsExternalDocs struct {
	Description        string                 `json:"description"`
	URL                string                 `json:"url"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/operation->produces
type DefinitionsOperationProduces struct {
	[]string
}

// #/definitions/operation->consumes
type DefinitionsOperationConsumes struct {
	[]string
}

// #/definitions/bodyParameter
type DefinitionsBodyParameter struct {
	Description        string                 `json:"description"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name               string                 `json:"name"`        // The name of the parameter.
	In                 string                 `json:"in"`          // Determines the location of the parameter.
	Required           bool                   `json:"required"`    // Determines whether or not this parameter is required or optional.
	Schema             *DefinitionsSchema     `json:"schema"`      // A deterministic version of a JSON Schema object.
	patternPropertiesX map[string]interface{}
}

// #/definitions/schema
type DefinitionsSchema struct {
	Ref                  string                                                            `json:"$ref"`
	Format               string                                                            `json:"format"`
	Title                string                                                            `json:"title"`
	Description          string                                                            `json:"description"`
	Default              interface{}                                                       `json:"default"`
	MultipleOf           float64                                                           `json:"multipleOf"`
	Maximum              float64                                                           `json:"maximum"`
	ExclusiveMaximum     bool                                                              `json:"exclusiveMaximum"`
	Minimum              float64                                                           `json:"minimum"`
	ExclusiveMinimum     bool                                                              `json:"exclusiveMinimum"`
	MaxLength            int64                                                             `json:"maxLength"`
	MinLength            *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern              string                                                            `json:"pattern"`
	MaxItems             int64                                                             `json:"maxItems"`
	MinItems             *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems          bool                                                              `json:"uniqueItems"`
	MaxProperties        int64                                                             `json:"maxProperties"`
	MinProperties        *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minProperties"`
	Required             []string                                                          `json:"required"`
	Enum                 []interface{}                                                     `json:"enum"`
	patternPropertiesX   map[string]interface{}
	AdditionalProperties *DefinitionsSchemaAdditionalProperties                            `json:"additionalProperties"`
	Type                 *HTTPJSONSchemaOrgDraft04SchemaPropertiesType                     `json:"type"`
	Items                *DefinitionsSchemaItems                                           `json:"items"`
	AllOf                []*DefinitionsSchema                                              `json:"allOf"`
	Properties           *DefinitionsSchemaProperties                                      `json:"properties"`
	Discriminator        string                                                            `json:"discriminator"`
	ReadOnly             bool                                                              `json:"readOnly"`
	XML                  *DefinitionsXML                                                   `json:"xml"`
	ExternalDocs         *DefinitionsExternalDocs                                          `json:"externalDocs"`         // information about external documentation
	Example              interface{}                                                       `json:"example"`
}

// http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0
type HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 struct {
	interface{}
}

// #/definitions/schema->additionalProperties
type DefinitionsSchemaAdditionalProperties struct {
	bool
}

// http://json-schema.org/draft-04/schema#/properties/type
type HTTPJSONSchemaOrgDraft04SchemaPropertiesType struct {
	[]interface{}
}

// #/definitions/schema->items
type DefinitionsSchemaItems struct {
	[]*DefinitionsSchema
}

// #/definitions/schema->properties
type DefinitionsSchemaProperties struct {
	additionalProperties map[string]*DefinitionsSchema
}

// #/definitions/xml
type DefinitionsXML struct {
	Name               string                 `json:"name"`
	Namespace          string                 `json:"namespace"`
	Prefix             string                 `json:"prefix"`
	Attribute          bool                   `json:"attribute"`
	Wrapped            bool                   `json:"wrapped"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/parameter
type DefinitionsParameter struct {
	*DefinitionsNonBodyParameter
}

// #/definitions/headerParameterSubSchema
type DefinitionsHeaderParameterSubSchema struct {
	Required           bool                                                              `json:"required"`         // Determines whether or not this parameter is required or optional.
	In                 string                                                            `json:"in"`               // Determines the location of the parameter.
	Description        string                                                            `json:"description"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name               string                                                            `json:"name"`             // The name of the parameter.
	Type               string                                                            `json:"type"`
	Format             string                                                            `json:"format"`
	Items              *DefinitionsPrimitivesItems                                       `json:"items"`
	CollectionFormat   string                                                            `json:"collectionFormat"`
	Default            interface{}                                                       `json:"default"`
	Maximum            float64                                                           `json:"maximum"`
	ExclusiveMaximum   bool                                                              `json:"exclusiveMaximum"`
	Minimum            float64                                                           `json:"minimum"`
	ExclusiveMinimum   bool                                                              `json:"exclusiveMinimum"`
	MaxLength          int64                                                             `json:"maxLength"`
	MinLength          *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern            string                                                            `json:"pattern"`
	MaxItems           int64                                                             `json:"maxItems"`
	MinItems           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems        bool                                                              `json:"uniqueItems"`
	Enum               []interface{}                                                     `json:"enum"`
	MultipleOf         float64                                                           `json:"multipleOf"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/primitivesItems
type DefinitionsPrimitivesItems struct {
	Type               string                                                            `json:"type"`
	Format             string                                                            `json:"format"`
	patternPropertiesX map[string]interface{}
	Items              *DefinitionsPrimitivesItems                                       `json:"items"`
	CollectionFormat   string                                                            `json:"collectionFormat"`
	Default            interface{}                                                       `json:"default"`
	Maximum            float64                                                           `json:"maximum"`
	ExclusiveMaximum   bool                                                              `json:"exclusiveMaximum"`
	Minimum            float64                                                           `json:"minimum"`
	ExclusiveMinimum   bool                                                              `json:"exclusiveMinimum"`
	MaxLength          int64                                                             `json:"maxLength"`
	MinLength          *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern            string                                                            `json:"pattern"`
	MaxItems           int64                                                             `json:"maxItems"`
	MinItems           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems        bool                                                              `json:"uniqueItems"`
	Enum               []interface{}                                                     `json:"enum"`
	MultipleOf         float64                                                           `json:"multipleOf"`
}

// http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0
type HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 struct {
	interface{}
}

// http://json-schema.org/draft-04/schema#/definitions/positiveIntegerDefault0
type HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 struct {
	interface{}
}

// #/definitions/nonBodyParameter
type DefinitionsNonBodyParameter struct {
	*DefinitionsPathParameterSubSchema
}

// #/definitions/formDataParameterSubSchema
type DefinitionsFormDataParameterSubSchema struct {
	Required           bool                                                              `json:"required"`         // Determines whether or not this parameter is required or optional.
	In                 string                                                            `json:"in"`               // Determines the location of the parameter.
	Description        string                                                            `json:"description"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name               string                                                            `json:"name"`             // The name of the parameter.
	AllowEmptyValue    bool                                                              `json:"allowEmptyValue"`  // allows sending a parameter by name only or with an empty value.
	Type               string                                                            `json:"type"`
	Format             string                                                            `json:"format"`
	Items              *DefinitionsPrimitivesItems                                       `json:"items"`
	CollectionFormat   string                                                            `json:"collectionFormat"`
	Default            interface{}                                                       `json:"default"`
	Maximum            float64                                                           `json:"maximum"`
	ExclusiveMaximum   bool                                                              `json:"exclusiveMaximum"`
	Minimum            float64                                                           `json:"minimum"`
	ExclusiveMinimum   bool                                                              `json:"exclusiveMinimum"`
	MaxLength          int64                                                             `json:"maxLength"`
	MinLength          *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern            string                                                            `json:"pattern"`
	MaxItems           int64                                                             `json:"maxItems"`
	MinItems           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems        bool                                                              `json:"uniqueItems"`
	Enum               []interface{}                                                     `json:"enum"`
	MultipleOf         float64                                                           `json:"multipleOf"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/queryParameterSubSchema
type DefinitionsQueryParameterSubSchema struct {
	Required           bool                                                              `json:"required"`         // Determines whether or not this parameter is required or optional.
	In                 string                                                            `json:"in"`               // Determines the location of the parameter.
	Description        string                                                            `json:"description"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name               string                                                            `json:"name"`             // The name of the parameter.
	AllowEmptyValue    bool                                                              `json:"allowEmptyValue"`  // allows sending a parameter by name only or with an empty value.
	Type               string                                                            `json:"type"`
	Format             string                                                            `json:"format"`
	Items              *DefinitionsPrimitivesItems                                       `json:"items"`
	CollectionFormat   string                                                            `json:"collectionFormat"`
	Default            interface{}                                                       `json:"default"`
	Maximum            float64                                                           `json:"maximum"`
	ExclusiveMaximum   bool                                                              `json:"exclusiveMaximum"`
	Minimum            float64                                                           `json:"minimum"`
	ExclusiveMinimum   bool                                                              `json:"exclusiveMinimum"`
	MaxLength          int64                                                             `json:"maxLength"`
	MinLength          *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern            string                                                            `json:"pattern"`
	MaxItems           int64                                                             `json:"maxItems"`
	MinItems           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems        bool                                                              `json:"uniqueItems"`
	Enum               []interface{}                                                     `json:"enum"`
	MultipleOf         float64                                                           `json:"multipleOf"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/pathParameterSubSchema
type DefinitionsPathParameterSubSchema struct {
	Required           bool                                                              `json:"required"`         // Determines whether or not this parameter is required or optional.
	In                 string                                                            `json:"in"`               // Determines the location of the parameter.
	Description        string                                                            `json:"description"`      // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name               string                                                            `json:"name"`             // The name of the parameter.
	Type               string                                                            `json:"type"`
	Format             string                                                            `json:"format"`
	Items              *DefinitionsPrimitivesItems                                       `json:"items"`
	CollectionFormat   string                                                            `json:"collectionFormat"`
	Default            interface{}                                                       `json:"default"`
	Maximum            float64                                                           `json:"maximum"`
	ExclusiveMaximum   bool                                                              `json:"exclusiveMaximum"`
	Minimum            float64                                                           `json:"minimum"`
	ExclusiveMinimum   bool                                                              `json:"exclusiveMinimum"`
	MaxLength          int64                                                             `json:"maxLength"`
	MinLength          *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern            string                                                            `json:"pattern"`
	MaxItems           int64                                                             `json:"maxItems"`
	MinItems           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems        bool                                                              `json:"uniqueItems"`
	Enum               []interface{}                                                     `json:"enum"`
	MultipleOf         float64                                                           `json:"multipleOf"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/parametersList->items
type DefinitionsParametersListItems struct {
	*DefinitionsJSONReference
}

// #/definitions/jsonReference
type DefinitionsJSONReference struct {
	Ref string `json:"$ref"`
}

// #/definitions/response
type DefinitionsResponse struct {
	Description        string                     `json:"description"`
	Schema             *DefinitionsResponseSchema `json:"schema"`
	Headers            *DefinitionsHeaders        `json:"headers"`
	Examples           map[string]interface{}     `json:"examples"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/response->schema
type DefinitionsResponseSchema struct {
	*DefinitionsFileSchema
}

// #/definitions/fileSchema
type DefinitionsFileSchema struct {
	Format             string                   `json:"format"`
	Title              string                   `json:"title"`
	Description        string                   `json:"description"`
	Default            interface{}              `json:"default"`
	Required           []string                 `json:"required"`
	Type               string                   `json:"type"`
	ReadOnly           bool                     `json:"readOnly"`
	ExternalDocs       *DefinitionsExternalDocs `json:"externalDocs"` // information about external documentation
	Example            interface{}              `json:"example"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/header
type DefinitionsHeader struct {
	Type               string                                                            `json:"type"`
	Format             string                                                            `json:"format"`
	Items              *DefinitionsPrimitivesItems                                       `json:"items"`
	CollectionFormat   string                                                            `json:"collectionFormat"`
	Default            interface{}                                                       `json:"default"`
	Maximum            float64                                                           `json:"maximum"`
	ExclusiveMaximum   bool                                                              `json:"exclusiveMaximum"`
	Minimum            float64                                                           `json:"minimum"`
	ExclusiveMinimum   bool                                                              `json:"exclusiveMinimum"`
	MaxLength          int64                                                             `json:"maxLength"`
	MinLength          *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern            string                                                            `json:"pattern"`
	MaxItems           int64                                                             `json:"maxItems"`
	MinItems           *HTTPJSONSchemaOrgDraft04SchemaDefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems        bool                                                              `json:"uniqueItems"`
	Enum               []interface{}                                                     `json:"enum"`
	MultipleOf         float64                                                           `json:"multipleOf"`
	Description        string                                                            `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/headers
type DefinitionsHeaders struct {
	additionalProperties map[string]*DefinitionsHeader
}

// #/definitions/responseValue
type DefinitionsResponseValue struct {
	*DefinitionsJSONReference
}

// #/definitions/responses
type DefinitionsResponses struct {
	patternProperties093Default map[string]*DefinitionsResponseValue
	patternPropertiesX          map[string]interface{}
}

// #/definitions/securityRequirement
type DefinitionsSecurityRequirement struct {
	additionalProperties map[string][]string
}

// #/definitions/definitions
type DefinitionsDefinitions struct {
	additionalProperties map[string]*DefinitionsSchema
}

// #/definitions/parameterDefinitions
type DefinitionsParameterDefinitions struct {
	additionalProperties map[string]*DefinitionsParameter
}

// #/definitions/responseDefinitions
type DefinitionsResponseDefinitions struct {
	additionalProperties map[string]*DefinitionsResponse
}

// #/definitions/basicAuthenticationSecurity
type DefinitionsBasicAuthenticationSecurity struct {
	Type               string                 `json:"type"`
	Description        string                 `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/securityDefinitions->additionalProperties
type DefinitionsSecurityDefinitionsAdditionalProperties struct {
	*DefinitionsOauth2AccessCodeSecurity
}

// #/definitions/apiKeySecurity
type DefinitionsAPIKeySecurity struct {
	Type               string                 `json:"type"`
	Name               string                 `json:"name"`
	In                 string                 `json:"in"`
	Description        string                 `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/oauth2ImplicitSecurity
type DefinitionsOauth2ImplicitSecurity struct {
	Type               string                   `json:"type"`
	Flow               string                   `json:"flow"`
	Scopes             *DefinitionsOauth2Scopes `json:"scopes"`
	AuthorizationURL   string                   `json:"authorizationUrl"`
	Description        string                   `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/oauth2Scopes
type DefinitionsOauth2Scopes struct {
	additionalProperties map[string]string
}

// #/definitions/oauth2PasswordSecurity
type DefinitionsOauth2PasswordSecurity struct {
	Type               string                   `json:"type"`
	Flow               string                   `json:"flow"`
	Scopes             *DefinitionsOauth2Scopes `json:"scopes"`
	TokenURL           string                   `json:"tokenUrl"`
	Description        string                   `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/oauth2ApplicationSecurity
type DefinitionsOauth2ApplicationSecurity struct {
	Type               string                   `json:"type"`
	Flow               string                   `json:"flow"`
	Scopes             *DefinitionsOauth2Scopes `json:"scopes"`
	TokenURL           string                   `json:"tokenUrl"`
	Description        string                   `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/oauth2AccessCodeSecurity
type DefinitionsOauth2AccessCodeSecurity struct {
	Type               string                   `json:"type"`
	Flow               string                   `json:"flow"`
	Scopes             *DefinitionsOauth2Scopes `json:"scopes"`
	AuthorizationURL   string                   `json:"authorizationUrl"`
	TokenURL           string                   `json:"tokenUrl"`
	Description        string                   `json:"description"`
	patternPropertiesX map[string]interface{}
}

// #/definitions/securityDefinitions
type DefinitionsSecurityDefinitions struct {
	additionalProperties map[string]*DefinitionsSecurityDefinitionsAdditionalProperties
}

// #/definitions/tag
type DefinitionsTag struct {
	Name               string                   `json:"name"`
	Description        string                   `json:"description"`
	ExternalDocs       *DefinitionsExternalDocs `json:"externalDocs"` // information about external documentation
	patternPropertiesX map[string]interface{}
}


GO;

        $actualGen = '';
        foreach ($builder->getGeneratedClasses() as $generatedStruct) {
            $actualGen .= $generatedStruct->structDef;
        }

        $this->assertSame($expectedGen, $actualGen);
    }

}