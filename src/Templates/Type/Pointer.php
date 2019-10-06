<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class Pointer extends GoTemplate implements AnyType, NoOmitEmpty
{
    /** @var AnyType */
    private $type;

    /**
     * Json tag `omitempty` should be avoided if this flag is true.
     * @var bool
     */
    private $noOmitEmpty = false;

    /**
     * @return bool
     */
    public function isNoOmitEmpty()
    {
        return $this->noOmitEmpty;
    }

    /**
     * @param bool $noOmitEmpty
     */
    public function setNoOmitEmpty($noOmitEmpty)
    {
        $this->noOmitEmpty = $noOmitEmpty;
    }

    /**
     * Pointer constructor.
     * @param AnyType $type
     */
    public function __construct(AnyType $type)
    {
        $this->type = $type;
    }


    protected function toString()
    {
        return '*' . $this->type->render();
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
        return '*' . $this->type->getTypeString();
    }

    /**
     * @param AnyType $type
     * @return AnyType
     */
    public static function tryDereferenceOnce(AnyType $type)
    {
        if ($type instanceof Pointer) {
            return $type->getType();
        }
        return $type;
    }


}