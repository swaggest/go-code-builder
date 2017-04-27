<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class Map extends GoTemplate implements AnyType
{
    /** @var Type */
    private $keyType;
    /** @var Type */
    private $valueType;

    /**
     * Map constructor.
     * @param Type $keyType
     * @param Type $valueType
     */
    public function __construct(AnyType $keyType, AnyType $valueType)
    {
        $this->keyType = $keyType;
        $this->valueType = $valueType;
    }

    protected function toString()
    {
        return 'map[' . $this->keyType . ']' . $this->valueType;
    }
}