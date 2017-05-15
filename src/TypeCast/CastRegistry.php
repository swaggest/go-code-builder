<?php

namespace Swaggest\GoCodeBuilder\TypeCast;


use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructCast;
use Swaggest\GoCodeBuilder\Templates\Type\TypeCastException;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;

class CastRegistry implements Registry
{
    /** @var array */
    private $castIndex = array();

    /** @var FuncDef[] */
    private $usedCasts = array();

    /** @var StructCast[] */
    private $casts = array();

    public function addStructCast(StructCast $cast)
    {
        $baseTypeString = $cast->getBaseStruct()->getType()->getTypeString();
        $derivedTypeString = $cast->getDerivedStruct()->getType()->getTypeString();

        $this->castIndex[$baseTypeString][$derivedTypeString] = true;
        $this->castIndex[$derivedTypeString][$baseTypeString] = false;

        $this->casts[$baseTypeString . ':' . $derivedTypeString] = $cast;
        return $this;
    }

    /**
     * @param string $toTypeString
     * @param string $fromTypeString
     * @return bool
     */
    public function canProcess($toTypeString, $fromTypeString)
    {
        return isset($this->castIndex[$toTypeString][$fromTypeString]);
    }

    /**
     * @param $toTypeString
     * @param $fromTypeString
     * @param $toVarName
     * @param $fromVarName
     * @param $assignOp
     * @return string
     * @throws TypeCastException
     */
    public function process($toTypeString, $fromTypeString, $toVarName, $fromVarName, $assignOp)
    {

        if ($this->castIndex[$toTypeString][$fromTypeString]) {
            $cast = $this->casts[$toTypeString . ':' . $fromTypeString];
        } else {
            $cast = $this->casts[$fromTypeString . ':' . $toTypeString];
        }
        $usedCastKey = $toTypeString . ':' . $fromTypeString;



        if ($this->castIndex[$toTypeString][$fromTypeString]) {
            if (!isset($this->usedCasts[$usedCastKey])) {
                $this->usedCasts[$usedCastKey] = $cast->getLoadFrom();
            }
            if (' = &' === $assignOp) {
                $toType = TypeUtil::fromString($toTypeString);
                return <<<GO
if {$toVarName} == nil {
    {$toVarName} = new({$toType->render()})
}
{$toVarName}.LoadFrom({$fromVarName})
GO;
            }
            return <<<GO
{$toVarName}.LoadFrom({$fromVarName})
GO;
        } else {
            if (!isset($this->usedCasts[$usedCastKey])) {
                $this->usedCasts[$usedCastKey] = $cast->getMapTo();
            }
            $tmpVar = 'tmp' . substr(md5($fromVarName), 0, 4);
            $code = '';

            $pDir = substr($assignOp, -1);

            if ($pDir === '*') {
                $assignOp = substr($assignOp, 0, -1);
                $code = <<<GO
{$tmpVar} := *{$fromVarName}

GO;
                $fromVarName = $tmpVar;
            } elseif ($pDir === '&') {
                $assignOp = substr($assignOp, 0, -1);
                $code = <<<GO
{$tmpVar} := &{$fromVarName}

GO;
                $fromVarName = $tmpVar;
            }


            $code .= <<<GO
{$toVarName}{$assignOp}{$fromVarName}.MapTo()
GO;

        }

        return $code;
    }

    /**
     * @return \Swaggest\GoCodeBuilder\Templates\Func\FuncDef[]
     */
    public function getUsedCastFuncs()
    {
        return array_values($this->usedCasts);
    }


}