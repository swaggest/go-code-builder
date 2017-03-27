<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

class Slice extends Type implements AnyType
{
    public function toString()
    {
        return '[]' . parent::toString();
    }
}