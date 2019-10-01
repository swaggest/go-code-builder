<?php

namespace Swaggest\GoCodeBuilder\Templates\Mapping;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;

class Mapping
{
    /** @var AnyType */
    public $baseType;

    /** @var AnyType */
    public $derivedType;

    public $baseName;

    public $derivedName;

    /** @var Mapping[] */
    public $properties;
}