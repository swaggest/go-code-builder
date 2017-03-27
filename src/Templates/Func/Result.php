<?php

namespace Swaggest\GoCodeBuilder\Templates\Func;

class Result extends Arguments
{
    public function toString()
    {

        if ($this->items === null) {
            return '';
        } elseif ((count($this->items) === 1) && $this->items[0]->name === null) {
            return parent::toString();
        } else {
            return '(' . parent::toString() . ')';
        }
    }

    public function toTypesString()
    {
        if ($this->items === null) {
            return '';
        } elseif ((count($this->items) === 1) && $this->items[0]->name === null) {
            return parent::toTypesString();
        } else {
            return '(' . parent::toTypesString() . ')';
        }
    }

}