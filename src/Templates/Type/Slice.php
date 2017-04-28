<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

class Slice extends Type implements AnyType
{
    protected function toString()
    {
        return '[]' . parent::toString();
    }
}