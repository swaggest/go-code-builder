<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class Map extends GoTemplate implements AnyType
{
    /** @var AnyType */
    private $keyType;
    /** @var AnyType */
    private $valueType;

    /**
     * Map constructor.
     * @param AnyType $keyType
     * @param AnyType $valueType
     */
    public function __construct(AnyType $keyType, AnyType $valueType)
    {
        $this->keyType = $keyType;
        $this->valueType = $valueType;
    }

    protected function toString()
    {
        return 'map[' . $this->keyType->render() . ']' . $this->valueType->render();
    }

    public function renderName()
    {
        return $this->toString();
    }

    /**
     * @return AnyType
     */
    public function getKeyType()
    {
        return $this->keyType;
    }

    /**
     * @return AnyType
     */
    public function getValueType()
    {
        return $this->valueType;
    }
}