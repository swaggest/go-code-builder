<?php


namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\JsonSchema\Schema;

class AllOfEmbedTest extends \PHPUnit_Framework_TestCase
{
    public function testEmbed()
    {
        $schemaData = json_decode(file_get_contents(__DIR__ . '/../../../resources/allOf-embed.json'));
        $schema = Schema::import($schemaData);

        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;

        $builder->getType($schema);

        $goFile = new GoFile('entities');
        $goFile->fileComment = '';
        $goFile->setComment('Package entities contains generated structures.');
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        $result = $goFile->render();

        $this->assertSame(<<<'GO'
// Package entities contains generated structures.
package entities



// Untitled1 structure is generated from "#".
type Untitled1 struct {
	OwnProperty  string `json:"ownProperty,omitempty"`
	AnotherTrait
	AllOf1
}

// AnotherTrait structure is generated from "#/definitions/anotherTrait".
type AnotherTrait struct {
	AtOne int64   `json:"atOne,omitempty"`
	AtTwo float64 `json:"atTwo,omitempty"`
}

// AllOf1 structure is generated from "#/allOf/1".
type AllOf1 struct {
	AoOne string `json:"aoOne,omitempty"`
	AoTwo bool   `json:"aoTwo,omitempty"`
}

GO
, $result);

    }

}