<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\JsonSchema\Constraint\Properties;
use Swaggest\JsonSchema\Schema;
use Swaggest\JsonSchema\Structure\ClassStructure;

class Options extends ClassStructure
{
    /**
     * Remove parent prefix from property name
     * Example:
     *  type Property struct { PropertySimple *PropertySimple `json:"-"`}
     * would become
     *  type Property struct { Simple *PropertySimple `json:"-"`}
     * @var bool
     */
    public $trimParentFromPropertyNames = true;

    /**
     * Hide properties with constant values
     * @var bool
     */
    public $hideConstProperties = true;

    /**
     * Use integer types based on minimum/maximum
     * @var bool
     */
    public $optimizeIntegers = true;

    /**
     * Skip Unmarshal generation
     * @var bool
     */
    public $skipUnmarshal = false;

    /**
     * Skip Marshal generation
     * @var bool
     */
    public $skipMarshal = false;

    /**
     * Generate structure for schema with `x-go-type` available.
     * @var bool
     */
    public $ignoreXGoType = false;

    /**
     * Add `null` to types if `x-nullable` or `nullable` is available.
     * @var bool
     */
    public $enableXNullable = false;

    /**
     * Use pointer types to avoid zero value ambiguity.
     * @var bool
     */
    public $withZeroValues = false;

    /**
     * Ignore if property is required when deciding on pointer type or omitempty.
     * @var bool
     */
    public $ignoreRequired = false;

    /**
     * Validate that required properties are present when unmarshaling.
     * @var bool
     */
    public $validateRequired = true;

    /**
     * Add omitempty to nullable values.
     * @var bool
     */
    public $ignoreNullable = false;

    /**
     * Separate `null` from non-existent key by using `*interface{}` type in property.
     * @var bool
     */
    public $distinctNull = true;

    /**
     * Enable `additionalProperties` if they are missing (null) in schema.
     * @var bool
     */
    public $defaultAdditionalProperties = true;

    /**
     * Inherit schema from schema examples where available.
     * @var bool
     */
    public $inheritSchemaFromExamples = false;

    /**
     * Generate fluent setters for struct fields.
     * @var bool
     */
    public $fluentSetters = false;

    /**
     * Map of default exported symbol names to new ones, e.g. PoorlyGenerated:NicelyReadable.
     * @var array
     */
    public $renames = [];

    /**
     * Only generate schemas that have `x-generate: true`.
     * @var bool
     */
    public $requireXGenerate = false;

    /**
     * Set additional field tags with property name.
     * @var string[]
     */
    public $nameTags = [];

    /**
     * @param Properties|static $properties
     * @param Schema $ownerSchema
     */
    public static function setUpProperties($properties, Schema $ownerSchema)
    {
        $properties->trimParentFromPropertyNames = Schema::boolean()->setDefault(true);
        $properties->hideConstProperties = Schema::boolean()->setDefault(true);
        $properties->skipMarshal = Schema::boolean();
        $properties->skipUnmarshal = Schema::boolean();
    }


}