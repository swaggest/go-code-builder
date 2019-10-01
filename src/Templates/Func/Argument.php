<?php

namespace Swaggest\GoCodeBuilder\Templates\Func;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class Argument extends GoTemplate
{
    /** @var string */
    public $name;
    /** @var AnyType */
    public $type;
    /** @var boolean */
    public $isVariadic;

    /**
     * Argument constructor.
     * @param string $name
     * @param AnyType $type
     * @param bool $isVariadic
     */
    public function __construct($name, AnyType $type, $isVariadic = false)
    {
        $this->name = $name;
        $this->type = $type;
        $this->isVariadic = $isVariadic;
    }

    protected function toString()
    {
        if ($this->name === null) {
            return $this->type->render();
        } else {
            return $this->name . ' ' . $this->type->render() . ($this->isVariadic ? '...' : '');
        }
    }

    public function getType()
    {
        return $this->type;
    }

}