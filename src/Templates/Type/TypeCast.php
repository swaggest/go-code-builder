<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\JsonSchema\Constraint\Type;

class TypeCast extends GoTemplate
{
    /** @var  AnyType */
    private $fromType;

    /** @var AnyType */
    private $toType;

    private $fromVarName;

    private $toVarName;

    private $assignOp = ' = ';

    /**
     * TypeCast constructor.
     * @param AnyType $toType
     * @param AnyType $fromType
     * @param $toVarName
     * @param $fromVarName
     */
    public function __construct(AnyType $toType, AnyType $fromType, $toVarName, $fromVarName)
    {
        $this->fromType = $fromType;
        $this->toType = $toType;
        $this->fromVarName = $fromVarName;
        $this->toVarName = $toVarName;
    }

    private function processFromPointer(Pointer $fromType)
    {
        if ($this->toType instanceof Pointer) {
            $this->fromType = $fromType->getType();
            $this->toType = $this->toType->getType();
            return $this->toString();
        } else {
            if ($fromType->getType() instanceof Pointer) {
                $tmpName = 'tmp' . substr(md5($this->fromVarName), 0, 4);
                $res = new TypeCast($this->toType, $fromType->getType(), $this->toVarName, $tmpName);
                $res->assignOp = ' = *';
                return <<<GO
if $this->fromVarName != nil {
    $tmpName := *{$this->fromVarName}
{$this->indentLines($res->toString())}
}
GO;

            } else {
                $res = new TypeCast($this->toType, $fromType->getType(), $this->toVarName, $this->fromVarName);
                $res->assignOp = ' = *';
                return <<<GO
if $this->fromVarName != nil {
{$this->indentLines($res)}
}
GO;
            }
        }
    }

    private function processToPointer(Pointer $toType)
    {
        //$res = new TypeCast($toType->getType(), $this->fromType, $this->toVarName, $this->fromVarName);
        //$res->assignOp = ' = &';
        //return $res->toString();


        if ($toType->getType() instanceof Pointer) {
            $tmpName = 'tmp' . substr(md5($this->fromVarName), 0, 4);
            $res = new TypeCast($toType->getType(), $this->fromType, $this->toVarName, $tmpName);
            $res->assignOp = ' = &';
            return <<<GO
$tmpName := &{$this->fromVarName}
{$res->toString()}
GO;

        } else {
            $res = new TypeCast($toType->getType(), $this->fromType, $this->toVarName, $this->fromVarName);
            $res->assignOp = ' = &';
            return $res->toString();
        }

    }

    protected function toString()
    {
        if (TypeUtil::equals($this->toType, $this->fromType)) {
            return $this->toVarName . $this->assignOp . $this->fromVarName;
        }


        if ($this->fromType instanceof Pointer) {
            return $this->processFromPointer($this->fromType);
        } elseif ($this->toType instanceof Pointer) {
            return $this->processToPointer($this->toType);
        }


        return '';
    }
}