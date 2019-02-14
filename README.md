# Swaggest JSON-schema enabled Go code builder

[![Build Status](https://travis-ci.org/swaggest/go-code-builder.svg?branch=master)](https://travis-ci.org/swaggest/go-code-builder)
[![codecov](https://codecov.io/gh/swaggest/go-code-builder/branch/master/graph/badge.svg)](https://codecov.io/gh/swaggest/go-code-builder)

This library generates Go mapping structures defined by [JSON schema](http://json-schema.org/).

## Example

[Generated code](tests/resources/go/swagger/entities.go) for [schema](tests/resources/swagger-schema.json).

```php
<?php
namespace MyApp;

use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Schema;

require_once __DIR__ . '/vendor/autoload.php';

$schemaData = json_decode(file_get_contents(__DIR__ . '/swagger-schema.json'));
$schema = Schema::import($schemaData);

$builder = new GoBuilder();
$builder->structCreatedHook = new StructHookCallback(function (StructDef $structDef, $path, $schema) use ($builder) {
    if ('#' === $path) {
        $structDef->setName('SwaggerSchema');
    } elseif (0 === strpos($path, '#/definitions/')) {
        $name = $builder->codeBuilder->exportableName(substr($path, strlen('#/definitions/')));
        $structDef->setName($name);
    }
});
$builder->getType($schema);

$goFile = new GoFile('swagger');
foreach ($builder->getGeneratedStructs() as $generatedStruct) {
    $goFile->getCode()->addSnippet($generatedStruct->structDef);
}
$goFile->getCode()->addSnippet($builder->getCode());
```
