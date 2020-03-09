<?php


namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\MarshalingTestFunc;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\SchemaContract;

class Helper
{
    public static function buildEntities(GoBuilder $builder, SchemaContract $schema, $path, $rootName)
    {
        mt_srand(1);

        $builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder, $rootName) {
            if ('#' === $path) {
                $structDef->setName($rootName);
            } elseif (0 === strpos($path, '#/definitions/')) {
                $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
                $structDef->setName($name);
            }
        });
        $builder->getType($schema);

        $goFile = new GoFile('entities');
        $goFile->setPackage('entities');
        $goFile->fileComment = '';
        $goFile->setComment('Package entities contains generated structures.');

        $goTestFile = new GoFile('entities_test');
        $goTestFile->setPackage('entities');
        $goTestFile->fileComment = '';

        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $goFile->getCode()->addSnippet($generatedStruct->structDef);

            $goTestFile->getCode()->addSnippet(MarshalingTestFunc::make($generatedStruct));
        }
        $goFile->getCode()->addSnippet($builder->getCode());

        file_put_contents( $path . '/entities.go', $goFile->render());
        file_put_contents( $path . '/entities_test.go', $goTestFile->render());
    }

}