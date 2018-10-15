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
// #
type Untitled1 struct {
	*DefinitionsHeader
}

// #/definitions/header
type DefinitionsHeader struct {
	Maximum float64 `json:"maximum"`
}


GO;

        $actualGen = '';
        foreach ($goBuilder->getGeneratedClasses() as $struct) {
            $actualGen .= $struct->structDef;
        }
        $this->assertSame($expectedGen, $actualGen);
    }

}