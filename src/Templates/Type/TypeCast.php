<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\TypeCast\Registry;

class TypeCast extends GoTemplate
{
    /** @var  AnyType */
    private $fromType;

    /** @var AnyType */
    private $toType;

    private $fromVarName;

    private $toVarName;

    private $assignOp = ' = ';

    /** @var Registry */
    private $typeRegistry;

    /**
     * TypeCast constructor.
     * @param AnyType $toType
     * @param AnyType $fromType
     * @param $toVarName
     * @param $fromVarName
     * @param Registry $registry
     */
    public function __construct(AnyType $toType, AnyType $fromType, $toVarName, $fromVarName, Registry $registry = null)
    {
        $this->fromType = $fromType;
        $this->toType = $toType;
        $this->fromVarName = $fromVarName;
        $this->toVarName = $toVarName;
        $this->typeRegistry = $registry;
    }

    private function processFromPointer(Pointer $fromType)
    {
        if ($this->toType instanceof Pointer) {
            if (TypeUtil::equals($this->toResolved, $this->fromResolved)) {
                $this->fromType = $fromType->getType();
                $this->toType = $this->toType->getType();
                return $this->toString();
            }
        }

        if ($fromType->getType() instanceof Pointer) {
            $tmpName = 'tmp' . substr(md5($this->fromVarName), 0, 4);
            $res = new TypeCast($this->toType, $fromType->getType(), $this->toVarName, $tmpName, $this->typeRegistry);
            $res->assignOp = ' = *';
            return <<<GO
if $this->fromVarName != nil { // $this->toType <- $this->fromType
	$tmpName := *{$this->fromVarName}
{$this->indentLines($res->toString())}
}
GO;

        } else {
            $res = new TypeCast($this->toType, $fromType->getType(), $this->toVarName, $this->fromVarName, $this->typeRegistry);
            $res->assignOp = ' = *';
            $res = $res->render();
            return <<<GO
if $this->fromVarName != nil {
{$this->indentLines($res)}
}
GO;
        }
    }

    private function processToPointer(Pointer $toType)
    {
        if ($toType->getType() instanceof Pointer) {
            $tmpName = 'tmp' . substr(md5($this->fromVarName), 0, 4);
            $res = new TypeCast($toType->getType(), $this->fromType, $this->toVarName, $tmpName, $this->typeRegistry);
            $res->assignOp = ' = &';
            return <<<GO
$tmpName := &{$this->fromVarName}
{$res->toString()}
GO;

        } else {
            if ($this->assignOp === ' = *') {
                $tmpName = 'tmp' . substr(md5($this->fromType), 0, 4);
                $tmpCast = <<<GO
{$tmpName} := *$this->fromVarName

GO;
                $res = new TypeCast($toType->getType(), $this->fromType, $this->toVarName, $tmpName, $this->typeRegistry);
                $res->assignOp = ' = &';
                return $tmpCast . $res->toString();

            } else {
                $res = new TypeCast($toType->getType(), $this->fromType, $this->toVarName, $this->fromVarName, $this->typeRegistry);
                $res->assignOp = ' = &';
                return $res->toString();
            }
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
            $valName,
            $this->typeRegistry
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
        $toKeyName = 'toKey' . $postfix;

        if (TypeUtil::equals($toType->getKeyType(), $fromType->getKeyType())) {
            $castValue = new TypeCast(
                $toType->getValueType(),
                $fromType->getValueType(),
                $this->toVarName . '[' . $keyName . ']',
                $valName,
                $this->typeRegistry
            );
            return <<<GO
{$this->toVarName} = make({$this->toType->render()}, len({$this->fromVarName}))
for {$keyName}, {$valName} := range {$this->fromVarName} {
{$this->indentLines($castValue->toString())}
}
GO;
        } else {
            $castKey = new TypeCast(
                $toType->getKeyType(),
                $fromType->getKeyType(),
                $toKeyName,
                $keyName,
                $this->typeRegistry
            );
            $castKey->assignOp = ' := ';
            $castValue = new TypeCast(
                $toType->getValueType(),
                $fromType->getValueType(),
                $this->toVarName . '[' . $toKeyName . ']',
                $this->fromVarName . '[' . $keyName . ']',
                $this->typeRegistry
            );
            return <<<GO
{$this->toVarName} = make({$this->toType->render()}, len({$this->fromVarName}))
for {$keyName} := range {$this->fromVarName} {
	var {$toKeyName} {$toType->getKeyType()->render()}
{$this->indentLines($castKey->toString())}
{$this->indentLines($castValue->toString())}
}
GO;
        }
    }

    /** @var AnyType */
    private $toResolved;
    /** @var AnyType */
    private $fromResolved;

    protected function toString()
    {
        $toType = $this->toType;
        $fromType = $this->fromType;

        if (TypeUtil::equals($toType, $fromType)) {
            return $this->toVarName . $this->assignOp . $this->fromVarName;
        }

        $this->toResolved = TypeUtil::resolvePointer($toType);
        $this->fromResolved = TypeUtil::resolvePointer($fromType);

        if ($toType instanceof Pointer) {
            if (!TypeUtil::equals($this->toResolved, $this->fromResolved)) {
                if (TypeUtil::isCastable($this->toResolved, $this->fromResolved)) {
                    $tmpName = 'tmp' . substr(md5($this->toVarName), 0, 4);
                    $castBack = new TypeCast($this->toResolved, $fromType, $tmpName, $this->fromVarName, $this->typeRegistry);
                    $castForth = new TypeCast($toType, $this->toResolved, $this->toVarName, $tmpName, $this->typeRegistry);

                    return <<<GO
var $tmpName {$this->toResolved->render()}
{$castBack->toString()}
{$castForth->toString()}
GO;
                }
            }
        }


        if ($fromType instanceof Pointer) {
            return $this->processFromPointer($fromType);
        } elseif ($toType instanceof Pointer) {
            return $this->processToPointer($toType);
        } elseif (($fromType instanceof Slice) && ($toType instanceof Slice)) {
            return $this->processSlices($toType, $fromType);
        } elseif (($toType instanceof Map) && ($fromType instanceof Map)) {
            return $this->processMaps($toType, $fromType);
        }


        if (($fromType instanceof Type) && ($toType instanceof Type)) {
            if (TypeUtil::isCastable($fromType, $toType)) {
                if ($this->assignOp === ' = *') {
                    return <<<GO
{$this->toVarName} = {$toType->render()}(*{$this->fromVarName})
GO;
                } elseif ($this->assignOp === ' = &') {
                    throw new TypeCastException('Could not compute');
                } else {
                    return <<<GO
{$this->toVarName} = {$toType->render()}({$this->fromVarName})
GO;
                }
            }
        }

        $fromTypeString = $fromType->getTypeString();
        $toTypeString = $toType->getTypeString();

        if ($this->typeRegistry
            && $this->typeRegistry->canProcess($toTypeString, $fromTypeString)
        ) {
            return $this->typeRegistry->process($toTypeString, $fromTypeString, $this->toVarName, $this->fromVarName, $this->assignOp);
        }

        //return '//' . 'Could not cast ' . $fromTypeString . ' to ' . $toTypeString;
        throw new TypeCastException('Could not cast ' . $fromTypeString . ' to ' . $toTypeString);
    }
}