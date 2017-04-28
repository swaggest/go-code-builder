<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

class Pointer extends Type implements AnyType
{
    protected function toString()
    {
        return '*' . parent::toString();
    }
}