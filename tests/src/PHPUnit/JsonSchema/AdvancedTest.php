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

        $builder->getType($schema);
        $goFile = new GoFile('entities');
        $goFile->fileComment = '';
        $goFile->setComment('Package entities contains generated structures.');
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $filePath = __DIR__ . '/../../../resources/go/advanced/string-or-object/entities.go';
        file_put_contents($filePath, $goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }
}