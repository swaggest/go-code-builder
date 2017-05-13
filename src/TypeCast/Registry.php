<?php

namespace Swaggest\GoCodeBuilder\TypeCast;

use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Type\TypeCastException;

interface Registry
{
    /**
     * @param string $toTypeString
     * @param string $fromTypeString
     * @return bool
     */
    public function canProcess($toTypeString, $fromTypeString);

    /**
     * @param $toTypeString
     * @param $fromTypeString
     * @param $toVarName
     * @param $fromVarName
     * @param $assignOp
     * @return Code|string
     * @throws TypeCastException
     */
    public function process($toTypeString, $fromTypeString, $toVarName, $fromVarName, $assignOp);

}