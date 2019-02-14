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