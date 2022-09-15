<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\JsonSchema\Constraint\Properties;
use Swaggest\JsonSchema\Schema;
use Swaggest\JsonSchema\Structure\ClassStructure;

class Options extends ClassStructure
{
    /**
     * Remove parent prefix from property name.
     * Example:
     *  type Property struct { PropertySimple *PropertySimple `json:"-"`}
     * would become
     *  type Property struct { Simple *PropertySimple `json:"-"`}
     * @var bool
     */
    public $trimParentFromPropertyNames = true;

    /**
     * Hide properties with constant values.
     * @var bool
     */
    public $hideConstProperties = true;

    /**
     * Use integer types based on minimum/maximum.
     * @var bool
     */
    public $optimizeIntegers = true;

    /**
     * Skip Unmarshal generation.
     * @var bool
     */
    public $skipUnmarshal = false;

    /**
     * Skip Marshal generation.
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
     * Map of regexps to replaces, e.g. {"/Definitions([\w\d]+)JSON([\w\d]+)?/i":"${1}${2}"}.
     * @var array
     */
    public $rewrites = [];

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
        $properties->trimParentFromPropertyNames = Schema::boolean()->setDefault(true)
            ->setDescription('Remove parent prefix from property name.');
        $properties->hideConstProperties = Schema::boolean()->setDefault(true)
            ->setDescription('Hide properties with constant values');
        $properties->optimizeIntegers = Schema::boolean()->setDefault(true)
            ->setDescription('Use integer types based on minimum/maximum');
        $properties->skipMarshal = Schema::boolean()
            ->setDescription('Skip Marshal generation');
        $properties->skipUnmarshal = Schema::boolean()
            ->setDescription('Skip Unmarshal generation');
        $properties->ignoreXGoType = Schema::boolean()
            ->setDescription('Generate structure for schema with `x-go-type` available.');
        $properties->enableXNullable = Schema::boolean()
            ->setDescription('Add `null` to types if `x-nullable` or `nullable` is available.');
        $properties->withZeroValues = Schema::boolean()
            ->setDescription('Use pointer types to avoid zero value ambiguity.');
        $properties->ignoreRequired = Schema::boolean()
            ->setDescription('Ignore if property is required when deciding on pointer type or omitempty.');
        $properties->validateRequired = Schema::boolean()->setDefault(true)
            ->setDescription('Validate that required properties are present when unmarshaling.');
        $properties->ignoreNullable = Schema::boolean()
            ->setDescription('Add omitempty to nullable values.');
        $properties->distinctNull = Schema::boolean()->setDefault(true)
            ->setDescription('Separate `null` from non-existent key by using `*interface{}` type in property.');
        $properties->defaultAdditionalProperties = Schema::boolean()->setDefault(true)
            ->setDescription('Enable `additionalProperties` if they are missing (null) in schema.');
        $properties->inheritSchemaFromExamples = Schema::boolean()
            ->setDescription('Inherit schema from schema examples where available.');
        $properties->fluentSetters = Schema::boolean()
            ->setDescription('Generate fluent setters for struct fields.');
        $properties->renames = Schema::object()->setAdditionalProperties(Schema::string())
            ->setDescription('Map of default exported symbol names to new ones, e.g. {"PoorlyGenerated":"NicelyReadable"}.');
        $properties->rewrites = Schema::object()->setAdditionalProperties(Schema::string())
            ->setDescription('Map of regexps to replaces, e.g. {"/Definitions([\w\d]+)JSON([\w\d]+)?/i":"${1}${2}"}.');
        $properties->requireXGenerate = Schema::boolean()
            ->setDescription('Only generate schemas that have `x-generate: true`.');
        $properties->nameTags = Schema::arr()->setItems(Schema::string())
            ->setDescription('Set additional field tags with property name.');
    }
}