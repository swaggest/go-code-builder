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

        $path = __DIR__ . '/../../../resources/go/asyncapi2';
        Helper::buildEntities($builder, $schema, $path, 'AsyncAPI');

        exec('git diff ' . $path, $out);
        $out = implode("\n", $out);
        $this->assertSame('', $out, "Generated files changed");
    }
}