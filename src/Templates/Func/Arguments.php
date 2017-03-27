<?php

namespace Swaggest\GoCodeBuilder\Templates\Func;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;

class Arguments extends GoTemplate
{
    /** @var Argument[] */
    protected $items;

    public function add($name, AnyType $type, $isVariadic = false)
    {
        $this->items[] = new Argument($name, $type, $isVariadic);
        return $this;
    }

    public function count()
    {
        return count($this->items);
    }

    public function toString()
    {
        $result = '';
        if ($this->items === null) {
            return $result;
        }
        /** @todo combine similar types (`a, b string`) */
        foreach ($this->items as $argument) {
            $result .= $argument->toString() . ', ';
        }
        if ($result) {
            $result = substr($result, 0, -2);
        }
        return $result;
    }

    public function toTypesString()
    {
        $result = '';
        foreach ($this->items as $argument) {
            $result .= $argument->getType() . ', ';
        }
        if ($result) {
            $result = substr($result, 0, -2);
        }
        return $result;
    }
}