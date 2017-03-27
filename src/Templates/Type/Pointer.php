<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

class Pointer extends Type implements AnyType
{
    public function toString()
    {
        return '*' . parent::toString();
    }
}