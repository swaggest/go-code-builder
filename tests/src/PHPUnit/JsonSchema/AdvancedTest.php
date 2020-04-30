<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
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

        $builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder) {
            if ('#' === $path) {
                $structDef->setName('Properties');
            } elseif (0 === strpos($path, '#/definitions/')) {
                $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
                $structDef->setName($name);
            }
        });
        $builder->getType($schema);

        $goFile = new GoFile('entities');
        $goFile->fileComment = '';
        $goFile->setComment('Package entities contains generated structures.');
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());


        $filePath = __DIR__ . '/../../../resources/go/advanced/entities.go';
        file_put_contents($filePath, $goFile->render());

        exec('git diff ' . $filePath, $out);
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

    public function testObjectOrNull()
    {
        $schemaData = <<<'JSON'
{
    "type": ["null", "object"],
    "required": ["a"],
    "properties": {
        "a": {"type": "string"}
    }
}
JSON;
        $schema = Schema::import(json_decode($schemaData));
        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->validateRequired = false;

        $path = __DIR__ . '/../../../resources/go/advanced/object-or-null';
        Helper::buildEntities($builder, $schema, $path, 'ObjectOrNull');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }


    public function testRequiredNullable()
    {
        $schemaData = <<<'JSON'
{
    "type": ["null", "object"],
    "required": ["a", "an", "axn", "b"],
    "properties": {
        "a": {"$ref": "#/definitions/def"},
        "an": {"$ref": "#/definitions/defNullable"},
        "axn": {"$ref": "#/definitions/defXNullable"},
        "b": {"type": ["null", "string"]}
    },
    "definitions": {
        "def": {
            "type": ["object", "null"],
            "properties": {
                "b": {"type": "string"}
            }
        },
        "defNullable": {
            "type": "object",
            "nullable": true,
            "properties": {
                "b": {"type": "string"}
            }
        },
        "defXNullable": {
            "type": "object",
            "x-nullable": true,
            "properties": {
                "b": {"type": "string"}
            }
        }
    }
}
JSON;
        $schema = Schema::import(json_decode($schemaData));
        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->validateRequired = false;
        $builder->options->ignoreRequired = false;
        $builder->options->enableXNullable = true;

        $path = __DIR__ . '/../../../resources/go/advanced/required-nullable';
        Helper::buildEntities($builder, $schema, $path, 'RequiredNullable');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    public function testNullableOmitemptyMapSlice()
    {
        $schemaData = <<<'JSON'
{
    "type": "object",
    "properties": {
        "slice": {
            "type": "array",
            "items": {"type":"string"},
            "x-nullable": true,
            "x-omitempty": true
        },
        "map": {
            "type": "object",
            "additionalProperties": {"type":"string"},
            "x-nullable": true,
            "x-omitempty": true
        }
    }
}
JSON;
        $schema = Schema::import(json_decode($schemaData));
        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->validateRequired = false;
        $builder->options->ignoreRequired = false;
        $builder->options->enableXNullable = true;

        $path = __DIR__ . '/../../../resources/go/advanced/nullable-omitempty-map-slice';
        Helper::buildEntities($builder, $schema, $path, 'RequiredNullable');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }
}