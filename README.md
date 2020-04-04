# Swaggest JSON-schema enabled Go code builder

[![Build Status](https://travis-ci.org/swaggest/go-code-builder.svg?branch=master)](https://travis-ci.org/swaggest/go-code-builder)
[![codecov](https://codecov.io/gh/swaggest/go-code-builder/branch/master/graph/badge.svg)](https://codecov.io/gh/swaggest/go-code-builder)
![Code lines](https://sloc.xyz/github/swaggest/go-code-builder/?category=code)
![Comments](https://sloc.xyz/github/swaggest/go-code-builder/?category=comments)

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

## API Documentation

Classes [documentation](API.md).

## Schema extensions

Magic properties (vendor extensions) defined in schema enable special handling.

### `x-go-type`

Can have a value string or an object. Contains type reference that can be used instead of generated type.

If `$ignoreXGoType` option is `true` value of vendor extension is disregarded and type is generated.

Value examples:

* `"[]myorg.com/go-null-types::nulltypes.NullString"`
* `"myorg.com/my-app/order.Entity"`
* `"float64"`
* `{"import": {"package": "my-package/domain/orders"}, "type": "Order"}`
* `{"import": {"package": "my-package", "alias": "mp"}, "type": "Order"}`

### `x-nullable`, `nullable`

If `true` schema type is converted to `[<type>, "null"]`. Requires `$enableXNullable` option.

### `x-omitempty`

A `boolean` value to control `,omitempty` presence.

### `x-generate`

A `boolean` value to control whether property should be added to generated `struct`. 
Property with `"x-generate": true` will be skipped.
If `GoBuilder` option `requireXGenerate` is set to `true` only properties with `"x-generate": true` will be generated. 

## CLI Tool

You can use [json-cli](https://github.com/swaggest/json-cli#gengo) to generate Go structures from command line.
