<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;

use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\Schema;

class AdvancedTest extends \PHPUnit_Framework_TestCase
{
    public function testGenConst()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/advanced.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->options->hideConstProperties = true;
        $builder->options->trimParentFromPropertyNames = true;
        $builder->options->inheritSchemaFromExamples = true;

        $path = __DIR__ . '/../../../resources/go/advanced';
        Helper::buildEntities($builder, $schema, $path, 'Properties');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    public function testStringOrObject()
    {
        $schemaData = <<<'JSON'
{
    "items": {
        "$ref": "#/definitions/element"
    },
    "type": "array",
    "definitions": {
        "element": {
            "properties": {
                "id": {
                    "type": "string"
                },
                "val": {
                    "type": "integer"
                }
            },
            "type": [
                "string",
                "object"
            ]
        }
    }
}
JSON;
        $schema = Schema::import(json_decode($schemaData));
        $builder = new GoBuilder();

        $path = __DIR__ . '/../../../resources/go/advanced/string-or-object';
        Helper::buildEntities($builder, $schema, $path, 'StringOrObject');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    public function testObjectOrString()
    {
        $schemaData = <<<'JSON'
{
    "items": {
        "$ref": "#/definitions/element"
    },
    "type": "array",
    "definitions": {
        "element": {
            "properties": {
                "id": {
                    "type": "string"
                },
                "val": {
                    "type": "integer"
                }
            },
            "type": [
                "object",
                "string"
            ]
        }
    }
}
JSON;
        $schema = Schema::import(json_decode($schemaData));
        $builder = new GoBuilder();


        $path = __DIR__ . '/../../../resources/go/advanced/object-or-string';
        Helper::buildEntities($builder, $schema, $path, 'ObjectOrString');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

}