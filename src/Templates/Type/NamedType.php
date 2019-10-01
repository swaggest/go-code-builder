<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

interface NamedType extends AnyType
{
    /**
     * @return string
     */
    public function getName();
}