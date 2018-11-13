<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\JsonSchema;
use Swaggest\JsonSchema\Schema;

class TypeBuilderTest extends \PHPUnit_Framework_TestCase
{
    public function testRef()
    {
        $schemaData = json_decode(<<<'JSON'
{
    "anyOf": [
        {},
        {"$ref":"#/definitions/header"}
    ],
    "definitions": {
        "header": {
            "type": "object",
            "properties": {
                "maximum": {"$ref": "#/definitions/maximum"}
            }
        },
        "maximum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/maximum"
        }
    }
}
JSON
        );
        $schema = Schema::import($schemaData);
        $goBuilder = new GoBuilder();
        $type = $goBuilder->getType($schema);

        $this->assertSame('*Untitled1', $type->getTypeString());
        $expectedGen = <<<'GO'
// Untitled1 structure is generated from #
type Untitled1 struct {
	DefinitionsHeader *DefinitionsHeader `json:"-"`
}

type marshalUntitled1 Untitled1

// UnmarshalJSON decodes JSON
func (i *Untitled1) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.DefinitionsHeader}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.DefinitionsHeader = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Untitled1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalUntitled1(i), i.DefinitionsHeader)
}

// DefinitionsHeader structure is generated from #/definitions/header
type DefinitionsHeader struct {
	Maximum float64 `json:"maximum,omitempty"`
}


GO;

        $actualGen = '';
        foreach ($goBuilder->getGeneratedStructs() as $struct) {
            $actualGen .= $struct->structDef;
        }
        $this->assertSame($expectedGen, $actualGen);
    }

}