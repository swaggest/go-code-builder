<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Context;
use Swaggest\JsonSchema\RemoteRef\Preloaded;
use Swaggest\JsonSchema\Schema;

class AsyncApiTest extends \PHPUnit_Framework_TestCase
{
    public function testGen()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi.json'));
        $refResolver = new Preloaded();
        $context = new Context($refResolver);
        $schema = Schema::import($schemaData, $context);


        $builder = new GoBuilder();
        $builder->options->hideConstProperties = false;
        $builder->options->trimParentFromPropertyNames = false;
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
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());


        $filePath = __DIR__ . '/../../../resources/go/asyncapi/entities.go';
        file_put_contents($filePath, $goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    public function testGenConst()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi.json'));
        $refResolver = new Preloaded();
        $context = new Context($refResolver);
        $schema = Schema::import($schemaData, $context);


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
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $filePath = __DIR__ . '/../../../resources/go/asyncapi-default/entities.go';
        file_put_contents($filePath, $goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    public function testGenSkipMarshal()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi.json'));
        $refResolver = new Preloaded();
        $context = new Context($refResolver);
        $schema = Schema::import($schemaData, $context);


        $builder = new GoBuilder();
        $builder->options->skipMarshal = true;
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
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $filePath = __DIR__ . '/../../../resources/go/asyncapi-skip-marshal/entities.go';
        file_put_contents($filePath, $goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }


    public function testGenSkipUnmarshal()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi.json'));
        $refResolver = new Preloaded();
        $context = new Context($refResolver);
        $schema = Schema::import($schemaData, $context);


        $builder = new GoBuilder();
        $builder->options->skipUnmarshal = true;
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
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $filePath = __DIR__ . '/../../../resources/go/asyncapi-skip-unmarshal/entities.go';
        file_put_contents($filePath,$goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }

    function testGenData()
    {
        $dataValue = json_decode(file_get_contents(__DIR__ . '/../../../resources/asyncapi-data.json'));

        $preloaded = new Preloaded();
        $preloaded->setSchemaData('asyncapi-data.json', $dataValue);

        $data = (object)[
            Schema::PROP_REF => 'asyncapi-data.json'
                . '#/components/messages/MessagingReaderReads/payload'
        ];
        $schema = Schema::import($data, new Context($preloaded));

        $builder = new GoBuilder();
        $builder->options->hideConstProperties = true;
        $builder->options->trimParentFromPropertyNames = true;
        $builder->pathToNameHook->prefixes [] = '#/components/schemas';
        $builder->getType($schema);

        $goFile = new GoFile('entities');
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $filePath = __DIR__ . '/../../../resources/go/asyncapi-data/entities.go';
        file_put_contents($filePath, $goFile->render());

        exec('git diff ' . $filePath, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }
}