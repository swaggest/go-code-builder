<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Schema;

class AsyncApi2Test extends \PHPUnit_Framework_TestCase
{
    public function testGen()
    {
        chdir(__DIR__ . '/../../../resources/');
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi-2.0.0.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->options->hideConstProperties = true;
        $builder->options->trimParentFromPropertyNames = true;
        $builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder) {
            if ('#' === $path) {
                $structDef->setName('AsyncAPI');
            } elseif (0 === strpos($path, '#/definitions/')) {
                $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
                $structDef->setName($name);
            }
        });
        $builder->getType($schema);

        $goFile = new GoFile('entities');
        $goFile->fileComment = '';
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());


        $filePath = __DIR__ . '/../../../resources/go/asyncapi2/entities.go';
        file_put_contents($filePath,$goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    public function testGen21()
    {
        chdir(__DIR__ . '/../../../resources/');
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi-2.1.0.json'));
        $schema = Schema::import($schemaData);


        $builder = new GoBuilder();
        $builder->options->hideConstProperties = true;
        $builder->options->trimParentFromPropertyNames = true;
        $builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder) {
            if ('#' === $path) {
                $structDef->setName('AsyncAPI');
            } elseif (0 === strpos($path, '#/definitions/')) {
                $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
                $structDef->setName($name);
            }
        });
        $builder->getType($schema);

        $goFile = new GoFile('entities');
        $goFile->fileComment = '';
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());


        $filePath = __DIR__ . '/../../../resources/go/asyncapi21/entities.go';
        file_put_contents($filePath,$goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }
}