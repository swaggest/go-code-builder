<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;

use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\Schema;

class SchemaTest extends \PHPUnit_Framework_TestCase
{
    public function testProperties()
    {
        $anotherSchema = Schema::object()
            ->setProperty('hello', Schema::boolean())
            ->setProperty('world', Schema::string());
        $anotherSchema->additionalProperties = false;


        $schema = Schema::object()
            ->setProperty('sampleInt', Schema::integer())
            ->setProperty('sampleBool', Schema::boolean())
            ->setProperty('sampleString', Schema::string())
            ->setProperty('sampleNumber', Schema::number())
        ;
        $schema
            ->setProperty('sampleSelf', $schema)
            ->setProperty('another', $anotherSchema)
        ;


        $builder = new GoBuilder();
        $type = $builder->getType($schema);

        $expectedStructs = <<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	SampleInt            int64                  `json:"sampleInt,omitempty"`
	SampleBool           bool                   `json:"sampleBool,omitempty"`
	SampleString         string                 `json:"sampleString,omitempty"`
	SampleNumber         float64                `json:"sampleNumber,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                      // All unmatched properties
	SampleSelf           *Untitled1             `json:"sampleSelf,omitempty"`
	Another              *Another               `json:"another,omitempty"`
}

type marshalUntitled1 Untitled1

var ignoreKeysUntitled1 = []string{
	"sampleInt",
	"sampleBool",
	"sampleString",
	"sampleNumber",
	"sampleSelf",
	"another",
}

// UnmarshalJSON decodes JSON.
func (i *Untitled1) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalUntitled1(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysUntitled1 {
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

	*i = Untitled1(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Untitled1) MarshalJSON() ([]byte, error) {
	if len(i.AdditionalProperties) == 0 {
		return json.Marshal(marshalUntitled1(i))
	}
	return marshalUnion(marshalUntitled1(i), i.AdditionalProperties)
}

// Another structure is generated from "#->another".
type Another struct {
	Hello bool   `json:"hello,omitempty"`
	World string `json:"world,omitempty"`
}


GO;

        $actualStructs = '';
        foreach ($builder->getGeneratedStructs() as $class) {
            $actualStructs .= $class->structDef;
        }

        $this->assertSame($expectedStructs, $actualStructs);
        $this->assertSame('*Untitled1', $type->getTypeString());
    }


    public function testSimple()
    {
        $builder = new GoBuilder();
        $this->assertSame('int64', $builder->getType(Schema::integer())->getTypeString());
        $this->assertSame('float64', $builder->getType(Schema::number())->getTypeString());
        $this->assertSame('string', $builder->getType(Schema::string())->getTypeString());
        $this->assertSame('bool', $builder->getType(Schema::boolean())->getTypeString());
        $this->assertSame('[]interface{}', $builder->getType(Schema::arr())->getTypeString());
        $this->assertSame('map[string]interface{}', $builder->getType(Schema::object())->getTypeString());
        $this->assertSame('interface{}', $builder->getType(Schema::null())->getTypeString());
    }

}