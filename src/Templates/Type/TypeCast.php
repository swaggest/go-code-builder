<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;

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
        $toResolved = TypeUtil::resolvePointer($toType);

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

    private function processSlices(Slice $toType, Slice $fromType)
    {
        $postfix = substr(md5($this->fromVarName), 0, 4);
        $indexName = 'index' . $postfix;
        $valName = 'val' . $postfix;
        $cast = new TypeCast(
            $toType->getType(),
            $fromType->getType(),
            $this->toVarName . '[' . $indexName . ']',
            $valName
        );
        return <<<GO
{$this->toVarName} = make({$this->toType->render()}, len({$this->fromVarName}))
for {$indexName}, {$valName} := range {$this->fromVarName} {
{$this->indentLines($cast->toString())}
}
GO;
    }

    private function processMaps(Map $toType, Map $fromType)
    {
        $postfix = substr(md5($this->fromVarName), 0, 4);
        $keyName = 'key' . $postfix;
        $valName = 'val' . $postfix;

        if (TypeUtil::equals($toType->getKeyType(), $fromType->getKeyType())) {
            $castValue = new TypeCast(
                $toType->getValueType(),
                $fromType->getValueType(),
                $this->toVarName . '[' . $keyName . ']',
                $valName
            );
            return <<<GO
{$this->toVarName} = make({$this->toType->render()}, len({$this->fromVarName}))
for {$keyName}, {$valName} := range {$this->fromVarName} {
{$this->indentLines($castValue->toString())}
}
GO;
        } else {
            $castKey = new TypeCast($toType->getKeyType(), $fromType->getKeyType(), 'toKey', 'key');
            $castKey->assignOp = ' := ';
            $castValue = new TypeCast($toType->getValueType(), $fromType->getValueType(), $this->toVarName . '[toKey]', $this->fromVarName . '[key]');
            return <<<GO
{$this->toVarName} = make({$this->toType->render()}, len({$this->fromVarName}))
for key, val := range {$this->fromVarName} {
{$this->indentLines($castKey->toString())}
{$this->indentLines($castValue->toString())}
}
GO;
        }


    }

    protected function toString()
    {
        if (TypeUtil::equals($this->toType, $this->fromType)) {
            return $this->toVarName . $this->assignOp . $this->fromVarName;
        }

        if ($this->toType instanceof Pointer) {
            $toResolved = TypeUtil::resolvePointer($this->toType);
            $fromResolved = TypeUtil::resolvePointer($this->fromType);
            if (!TypeUtil::equals($toResolved, $fromResolved)) {
                if (TypeUtil::isCastable($toResolved, $fromResolved)) {
                    $tmpName = 'tmp' . substr(md5($this->toVarName), 0, 4);
                    $castBack = new TypeCast($toResolved, $this->fromType, $tmpName, $this->fromVarName);
                    $castForth = new TypeCast($this->toType, $toResolved, $this->toVarName, $tmpName);

                    return <<<GO
var $tmpName {$toResolved->render()}
{$castBack->toString()}
{$castForth->toString()}
GO;
                }
            }
        }


        if ($this->fromType instanceof Pointer) {
            return $this->processFromPointer($this->fromType);
        } elseif ($this->toType instanceof Pointer) {
            return $this->processToPointer($this->toType);
        } elseif (($this->fromType instanceof Slice) && ($this->toType instanceof Slice)) {
            return $this->processSlices($this->toType, $this->fromType);
        } elseif (($this->toType instanceof Map) && ($this->fromType instanceof Map)) {
            return $this->processMaps($this->toType, $this->fromType);
        }


        if (($this->fromType instanceof Type) && ($this->toType instanceof Type)) {
            if (TypeUtil::isCastable($this->fromType, $this->toType)) {
                if ($this->assignOp === ' = *') {
                    return <<<GO
{$this->toVarName} = {$this->toType->render()}(*{$this->fromVarName})
GO;
                } elseif ($this->assignOp === ' = &') {
                    throw new TypeCastException('Could not compute');
                } else {
                    return <<<GO
{$this->toVarName} = {$this->toType->render()}({$this->fromVarName})
GO;
                }
            }
        }

        $fromTypeString = $this->fromType instanceof Type ? $this->fromType->getTypeString() : $this->fromType->render();
        $toTypeString = $this->toType instanceof Type ? $this->toType->getTypeString() : $this->toType->render();

        return $this->processSpecial($toTypeString, $fromTypeString);
    }

    private function processSpecial($toTypeString, $fromTypeString)
    {
        if ($fromTypeString === 'time.Time' && $toTypeString === 'string') {
            return <<<GO
{$this->toVarName}{$this->assignOp}{$this->fromVarName}.Format(time.RFC3339Nano)
GO;
        }

        if ($toTypeString === 'time.Time' && $fromTypeString === 'string') {
            return <<<GO
{$this->toVarName}, _{$this->assignOp}time.Parse(time.RFC3339Nano, {$this->fromVarName})
GO;

        }

        //return '//' . 'Could not cast ' . $fromTypeString . ' to ' . $toTypeString;
        throw new TypeCastException('Could not cast ' . $fromTypeString . ' to ' . $toTypeString);
    }
}