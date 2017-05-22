<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class Slice extends GoTemplate implements AnyType
{
    /** @var AnyType */
    private $type;

    /**
     * Slice constructor.
     * @param AnyType $type
     */
    public function __construct(AnyType $type)
    {
        $this->type = $type;
    }

    protected function toString()
    {
        return '[]' . $this->type->render();
    }

    /**
     * @return AnyType
     */
    public function getType()
    {
        return $this->type;
    }

    public function getTypeString()
    {
        return '[]' . $this->type->getTypeString();
    }


}