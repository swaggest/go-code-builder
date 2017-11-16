<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\JsonSchema;

class JsonSchemaGenerateTest extends \PHPUnit_Framework_TestCase
{
    public function testJsonSchemaGenerate()
    {
        //$this->markTestSkipped();

        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../../../json-schema/spec/json-schema.json'));
        $schema = JsonSchema::import($schemaData);


        $builder = new GoBuilder();
        $builder->getType($schema);
        $expectedGen = <<<'GO'
// #
type Untitled1 struct {
	Id                   string                              `json:"id"`
	Schema               string                              `json:"$schema"`
	Title                string                              `json:"title"`
	Description          string                              `json:"description"`
	Default              interface{}                         `json:"default"`
	MultipleOf           float64                             `json:"multipleOf"`
	Maximum              float64                             `json:"maximum"`
	ExclusiveMaximum     bool                                `json:"exclusiveMaximum"`
	Minimum              float64                             `json:"minimum"`
	ExclusiveMinimum     bool                                `json:"exclusiveMinimum"`
	MaxLength            int64                               `json:"maxLength"`
	MinLength            *DefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern              string                              `json:"pattern"`
	AdditionalItems      *AdditionalItems                    `json:"additionalItems"`
	Items                *Items                              `json:"items"`
	MaxItems             int64                               `json:"maxItems"`
	MinItems             *DefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems          bool                                `json:"uniqueItems"`
	MaxProperties        int64                               `json:"maxProperties"`
	MinProperties        *DefinitionsPositiveIntegerDefault0 `json:"minProperties"`
	Required             []string                            `json:"required"`
	AdditionalProperties *AdditionalProperties               `json:"additionalProperties"`
	Definitions          *Definitions                        `json:"definitions"`
	Properties           *Properties                         `json:"properties"`
	PatternProperties    *PatternProperties                  `json:"patternProperties"`
	Dependencies         *Dependencies                       `json:"dependencies"`
	Enum                 []interface{}                       `json:"enum"`
	Type                 *Type                               `json:"type"`
	Format               string                              `json:"format"`
	Ref                  string                              `json:"$ref"`
	AllOf                []*Untitled2                        `json:"allOf"`
	AnyOf                []*Untitled2                        `json:"anyOf"`
	OneOf                []*Untitled2                        `json:"oneOf"`
	Not                  *Untitled2                          `json:"not"`                  // Core schema meta-schema
}

// #/definitions/positiveIntegerDefault0
type DefinitionsPositiveIntegerDefault0 struct {
	interface{}
}

// #->additionalItems
type AdditionalItems struct {
	*Untitled2
}

// #
type Untitled2 struct {
	Id                   string                              `json:"id"`
	Schema               string                              `json:"$schema"`
	Title                string                              `json:"title"`
	Description          string                              `json:"description"`
	Default              interface{}                         `json:"default"`
	MultipleOf           float64                             `json:"multipleOf"`
	Maximum              float64                             `json:"maximum"`
	ExclusiveMaximum     bool                                `json:"exclusiveMaximum"`
	Minimum              float64                             `json:"minimum"`
	ExclusiveMinimum     bool                                `json:"exclusiveMinimum"`
	MaxLength            int64                               `json:"maxLength"`
	MinLength            *DefinitionsPositiveIntegerDefault0 `json:"minLength"`
	Pattern              string                              `json:"pattern"`
	AdditionalItems      *AdditionalItems                    `json:"additionalItems"`
	Items                *Items                              `json:"items"`
	MaxItems             int64                               `json:"maxItems"`
	MinItems             *DefinitionsPositiveIntegerDefault0 `json:"minItems"`
	UniqueItems          bool                                `json:"uniqueItems"`
	MaxProperties        int64                               `json:"maxProperties"`
	MinProperties        *DefinitionsPositiveIntegerDefault0 `json:"minProperties"`
	Required             []string                            `json:"required"`
	AdditionalProperties *AdditionalProperties               `json:"additionalProperties"`
	Definitions          *Definitions                        `json:"definitions"`
	Properties           *Properties                         `json:"properties"`
	PatternProperties    *PatternProperties                  `json:"patternProperties"`
	Dependencies         *Dependencies                       `json:"dependencies"`
	Enum                 []interface{}                       `json:"enum"`
	Type                 *Type                               `json:"type"`
	Format               string                              `json:"format"`
	Ref                  string                              `json:"$ref"`
	AllOf                []*Untitled2                        `json:"allOf"`
	AnyOf                []*Untitled2                        `json:"anyOf"`
	OneOf                []*Untitled2                        `json:"oneOf"`
	Not                  *Untitled2                          `json:"not"`                  // Core schema meta-schema
}

// #->additionalItems
type AdditionalItems struct {
	*Untitled2
}

// #->items
type Items struct {
	[]*Untitled2
}

// #->additionalProperties
type AdditionalProperties struct {
	*Untitled2
}

// #->definitions
type Definitions struct {
	additionalProperties map[string]*Untitled2
}

// #->properties
type Properties struct {
	additionalProperties map[string]*Untitled2
}

// #->patternProperties
type PatternProperties struct {
	additionalProperties map[string]*Untitled2
}

// #->dependencies->additionalProperties
type DependenciesAdditionalProperties struct {
	[]string
}

// #->dependencies
type Dependencies struct {
	additionalProperties map[string]*DependenciesAdditionalProperties
}

// #->type
type Type struct {
	[]interface{}
}

// #->items
type Items struct {
	[]*Untitled2
}

// #->additionalProperties
type AdditionalProperties struct {
	*Untitled2
}

// #->definitions
type Definitions struct {
	additionalProperties map[string]*Untitled2
}

// #->properties
type Properties struct {
	additionalProperties map[string]*Untitled2
}

// #->patternProperties
type PatternProperties struct {
	additionalProperties map[string]*Untitled2
}

// #->dependencies->additionalProperties
type DependenciesAdditionalProperties struct {
	[]string
}

// #->dependencies
type Dependencies struct {
	additionalProperties map[string]*DependenciesAdditionalProperties
}

// #->type
type Type struct {
	[]interface{}
}


GO;

        $actualGen = '';
        foreach ($builder->getGeneratedClasses() as $generatedStruct) {
            $actualGen .= $generatedStruct->structDef;
        }

        $this->assertSame($expectedGen, $actualGen);
    }

}