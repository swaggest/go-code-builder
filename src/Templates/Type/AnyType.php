<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

interface AnyType
{
    public function render();

    public function getTypeString();
}