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
        $builder->options->nameTags = ['msgp', 'bson'];
        $type = $builder->getType($schema);

        $expectedStructs = <<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	SampleInt            int64                  `json:"sampleInt,omitempty" msgp:"sampleInt" bson:"sampleInt"`
	SampleBool           bool                   `json:"sampleBool,omitempty" msgp:"sampleBool" bson:"sampleBool"`
	SampleString         string                 `json:"sampleString,omitempty" msgp:"sampleString" bson:"sampleString"`
	SampleNumber         float64                `json:"sampleNumber,omitempty" msgp:"sampleNumber" bson:"sampleNumber"`
	AdditionalProperties map[string]interface{} `json:"-"`                                                              // All unmatched properties.
	SampleSelf           *Untitled1             `json:"sampleSelf,omitempty" msgp:"sampleSelf" bson:"sampleSelf"`
	Another              *Another               `json:"another,omitempty" msgp:"another" bson:"another"`
}

type marshalUntitled1 Untitled1

var knownKeysUntitled1 = []string{
	"sampleInt",
	"sampleBool",
	"sampleString",
	"sampleNumber",
	"sampleSelf",
	"another",
}

// UnmarshalJSON decodes JSON.
func (u *Untitled1) UnmarshalJSON(data []byte) error {
	var err error

	mu := marshalUntitled1(*u)

	err = json.Unmarshal(data, &mu)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysUntitled1 {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		if mu.AdditionalProperties == nil {
			mu.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		mu.AdditionalProperties[key] = val
	}

	*u = Untitled1(mu)

	return nil
}

// MarshalJSON encodes JSON.
func (u Untitled1) MarshalJSON() ([]byte, error) {
	if len(u.AdditionalProperties) == 0 {
		return json.Marshal(marshalUntitled1(u))
	}

	return marshalUnion(marshalUntitled1(u), u.AdditionalProperties)
}

// Another structure is generated from "#->another".
type Another struct {
	Hello bool   `json:"hello,omitempty" msgp:"hello" bson:"hello"`
	World string `json:"world,omitempty" msgp:"world" bson:"world"`
}

type marshalAnother Another

var knownKeysAnother = []string{
	"hello",
	"world",
}

// UnmarshalJSON decodes JSON.
func (a *Another) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAnother(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysAnother {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Another: %v", offendingKeys)
	}

	*a = Another(ma)

	return nil
}



GO;

        $actualStructs = '';
        foreach ($builder->getGeneratedStructs() as $class) {
            $actualStructs .= $class->structDef;
        }

        $this->assertSame($expectedStructs, $actualStructs);
        $this->assertSame('Untitled1', $type->getTypeString());
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