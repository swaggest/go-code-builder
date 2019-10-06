<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\TypeBuilder;
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
            "additionalProperties": false,
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

        $this->assertSame('Header', $type->getTypeString());
        $expectedGen = <<<'GO'
// Header structure is generated from "#/definitions/header".
type Header struct {
	Maximum float64 `json:"maximum,omitempty"`
}


GO;

        $actualGen = '';
        foreach ($goBuilder->getGeneratedStructs() as $struct) {
            $actualGen .= $struct->structDef;
        }
        $this->assertSame($expectedGen, $actualGen);
    }

    function testXGoTypeIgnore()
    {
        $builder = new GoBuilder();
        $builder->options->ignoreXGoType = true;
        $schema = new Schema();
        $schema->{'x-go-type'} = 'my-package/domain/orders.Order';
        $schema->type = Schema::STRING;
        $tb = new TypeBuilder($schema, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('string', $type->getTypeString());
    }


    function testXGoTypeString()
    {
        $builder = new GoBuilder();
        $schema = new Schema();
        $schema->{'x-go-type'} = 'my-package/domain/orders.Order';
        $tb = new TypeBuilder($schema, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('my-package/domain/orders.Order', $type->getTypeString());
    }

    function testXGoTypeGoSwaggerObject()
    {
        $builder = new GoBuilder();
        $schema = new Schema();
        $schema->{'x-go-type'} = json_decode('{
                  "import": {
                    "package": "my-package/domain/orders"
                  },
                  "type": "Order"
                }');
        $tb = new TypeBuilder($schema, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('my-package/domain/orders.Order', $type->getTypeString());
    }


    function testNullable()
    {
        $prop = new Schema();
        $prop->type = [Schema::STRING, Schema::NULL];

        $obj = Schema::object();
        $obj->setProperty('nullable', $prop);
        $obj->setProperty('regular', Schema::string());
        $obj->additionalProperties = false;


        $builder = new GoBuilder();
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('*Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	Nullable *string `json:"nullable"`
	Regular  string  `json:"regular,omitempty"`
}


GO
            , $struct->structDef->render());


        $builder = new GoBuilder();
        $builder->options->withZeroValues = true;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('*Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	Nullable *string `json:"nullable"`
	Regular  *string `json:"regular,omitempty"`
}


GO
            , $struct->structDef->render());


        $builder = new GoBuilder();
        $builder->options->ignoreNullable = true;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('*Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	Nullable *string `json:"nullable,omitempty"`
	Regular  string  `json:"regular,omitempty"`
}


GO
            , $struct->structDef->render());


    }

}