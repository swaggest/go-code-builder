<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\Schema;

class JsonSchemaGenerateTest extends \PHPUnit_Framework_TestCase
{
    public function testJsonSchemaGenerate()
    {
        //$this->markTestSkipped();

        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../../vendor/swaggest/json-schema/spec/json-schema.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->getType($schema);
        $expectedGen = <<<'GO'
// Untitled1 structure is generated from #
// Core schema meta-schema
type Untitled1 struct {
	ID                   string                                      `json:"id,omitempty"`
	Schema               string                                      `json:"$schema,omitempty"`
	Title                string                                      `json:"title,omitempty"`
	Description          string                                      `json:"description,omitempty"`
	Default              interface{}                                 `json:"default,omitempty"`
	MultipleOf           float64                                     `json:"multipleOf,omitempty"`
	Maximum              float64                                     `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                                        `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                                     `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                                        `json:"exclusiveMinimum,omitempty"`
	MaxLength            int64                                       `json:"maxLength,omitempty"`
	MinLength            *DefinitionsPositiveIntegerDefault0         `json:"minLength,omitempty"`
	Pattern              string                                      `json:"pattern,omitempty"`
	AdditionalItems      *AdditionalItems                            `json:"additionalItems,omitempty"`
	Items                *Items                                      `json:"items,omitempty"`
	MaxItems             int64                                       `json:"maxItems,omitempty"`
	MinItems             *DefinitionsPositiveIntegerDefault0         `json:"minItems,omitempty"`
	UniqueItems          bool                                        `json:"uniqueItems,omitempty"`
	MaxProperties        int64                                       `json:"maxProperties,omitempty"`
	MinProperties        *DefinitionsPositiveIntegerDefault0         `json:"minProperties,omitempty"`
	Required             []string                                    `json:"required,omitempty"`
	AdditionalProperties *AdditionalProperties                       `json:"additionalProperties,omitempty"`
	Definitions          map[string]Untitled1                        `json:"definitions,omitempty"`
	Properties           map[string]Untitled1                        `json:"properties,omitempty"`
	PatternProperties    map[string]Untitled1                        `json:"patternProperties,omitempty"`
	Dependencies         map[string]DependenciesAdditionalProperties `json:"dependencies,omitempty"`
	Enum                 []interface{}                               `json:"enum,omitempty"`
	Type                 *Type                                       `json:"type,omitempty"`
	Format               string                                      `json:"format,omitempty"`
	Ref                  string                                      `json:"$ref,omitempty"`
	AllOf                []Untitled1                                 `json:"allOf,omitempty"`
	AnyOf                []Untitled1                                 `json:"anyOf,omitempty"`
	OneOf                []Untitled1                                 `json:"oneOf,omitempty"`
	Not                  *Untitled1                                  `json:"not,omitempty"`                  // Core schema meta-schema
}

// DefinitionsPositiveIntegerDefault0 structure is generated from #/definitions/positiveIntegerDefault0
type DefinitionsPositiveIntegerDefault0 struct {
	Int64 *int64 `json:"-"`
}

type marshalDefinitionsPositiveIntegerDefault0 DefinitionsPositiveIntegerDefault0

// UnmarshalJSON decodes JSON
func (i *DefinitionsPositiveIntegerDefault0) UnmarshalJSON(data []byte) error {

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
func (i DefinitionsPositiveIntegerDefault0) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalDefinitionsPositiveIntegerDefault0(i), i.Int64)
}

// AdditionalItems structure is generated from #->additionalItems
type AdditionalItems struct {
	Bool      *bool      `json:"-"`
	Untitled1 *Untitled1 `json:"-"`
}

type marshalAdditionalItems AdditionalItems

// UnmarshalJSON decodes JSON
func (i *AdditionalItems) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Bool, &i.Untitled1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Bool = nil
	}
	if mayUnmarshal[1] == nil {
		i.Untitled1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i AdditionalItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAdditionalItems(i), i.Bool, i.Untitled1)
}

// Items structure is generated from #->items
type Items struct {
	Untitled1 *Untitled1  `json:"-"`
	AnyOf1    []Untitled1 `json:"-"`
}

type marshalItems Items

// UnmarshalJSON decodes JSON
func (i *Items) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Untitled1, &i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Untitled1 = nil
	}
	if mayUnmarshal[1] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Items) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalItems(i), i.Untitled1, i.AnyOf1)
}

// AdditionalProperties structure is generated from #->additionalProperties
type AdditionalProperties struct {
	Bool      *bool      `json:"-"`
	Untitled1 *Untitled1 `json:"-"`
}

type marshalAdditionalProperties AdditionalProperties

// UnmarshalJSON decodes JSON
func (i *AdditionalProperties) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Bool, &i.Untitled1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Bool = nil
	}
	if mayUnmarshal[1] == nil {
		i.Untitled1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i AdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAdditionalProperties(i), i.Bool, i.Untitled1)
}

// DependenciesAdditionalProperties structure is generated from #->dependencies->additionalProperties
type DependenciesAdditionalProperties struct {
	Untitled1 *Untitled1 `json:"-"`
	AnyOf1    []string   `json:"-"`
}

type marshalDependenciesAdditionalProperties DependenciesAdditionalProperties

// UnmarshalJSON decodes JSON
func (i *DependenciesAdditionalProperties) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Untitled1, &i.AnyOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Untitled1 = nil
	}
	if mayUnmarshal[1] == nil {
		i.AnyOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i DependenciesAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalDependenciesAdditionalProperties(i), i.Untitled1, i.AnyOf1)
}

// Type structure is generated from #->type
type Type struct {
	AnyOf1 []interface{} `json:"-"`
}

type marshalType Type

// UnmarshalJSON decodes JSON
func (i *Type) UnmarshalJSON(data []byte) error {
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
func (i Type) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalType(i), i.AnyOf1)
}


GO;

        $actualGen = '';
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $actualGen .= $generatedStruct->structDef;
        }

        $this->assertSame($expectedGen, $actualGen);
    }

}