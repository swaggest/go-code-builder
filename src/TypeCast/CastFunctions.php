<?php

namespace Swaggest\GoCodeBuilder\TypeCast;


use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;

interface CastFunctions
{
    /**
     * @return FuncDef
     */
    public function getMapTo();

    /**
     * @return FuncDef
     */
    public function getLoadFrom();

    /**
     * @return string
     */
    public function getBaseTypeString();

    /**
     * @return string
     */
    public function getDerivedTypeString();

}