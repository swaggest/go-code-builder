<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use PhpLang\ScopeExit;
use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\GoFile;

class TypeUtil
{
    /**
     * @param $typeString
     * @return AnyType|Type
     */
    public static function fromString($typeString)
    {
        if ('*' === $typeString[0]) {
            return new Pointer(self::fromString(substr($typeString, 1)));
        }

        if ('map[' === substr($typeString, 0, 4)) {
            list($key, $val) = explode(']', substr($typeString, 4));
            return new Map(self::fromString($key), self::fromString($val));
        }

        if ('[]' === substr($typeString, 0, 2)) {
            return new Slice(self::fromString(substr($typeString, 2)));
        }

        $parts = explode('::', $typeString);
        $import = null;
        if (false !== $dotPos = strrpos($parts[0], '.')) {
            if (isset($parts[1])) {
                list($packageName) = explode('.', $parts[1]);
                $import = new Import(substr($parts[0], 0, $dotPos), null, $packageName);
            } else {
                $import = new Import(substr($parts[0], 0, $dotPos));
            }
            $typeName = substr($parts[0], $dotPos + 1);
        } else {
            $typeName = $parts[0];
        }

        return new Type($typeName, $import);
    }

    public static function equals(AnyType $one, AnyType $two)
    {
        if (!$one instanceof $two) {
            return false;
        }
        if ($one instanceof Type && $two instanceof Type) {
            return $one->equals($two);
        }
        if ($one instanceof Pointer && $two instanceof Pointer) {
            return self::equals($one->getType(), $two->getType());
        }
        if ($one instanceof Map && $two instanceof Map) {
            return self::equals($one->getKeyType(), $two->getKeyType())
            && self::equals($one->getValueType(), $two->getValueType());
        }
        if ($one instanceof Slice && $two instanceof Slice) {
            return self::equals($one->getType(), $two->getType());
        }
        return false;
    }

    public static function castExpr(AnyType $to, AnyType $from, $fromVar, $toVarAssign)
    {
        if (self::equals($to, $from)) {
            return $toVarAssign . $fromVar;
        }
        if ($from instanceof Pointer) {
            if ($to instanceof Pointer) {
                return self::castExpr($to->getType(), $from->getType(), $fromVar, $toVarAssign);
            } else {
                if ($from->getType() instanceof Pointer) {
                    $toVarDeeper = 'tmp' . substr(md5($toVarAssign), 0, 4);
                    $toVarDeeperAssign = $toVarDeeper . ' := *';
                    $res = self::castExpr($to, $from->getType(), $fromVar, $toVarDeeperAssign);
                    return <<<GO
if {$fromVar} != nil {
    {$res}
    {$toVarAssign}{$toVarDeeper}
}

GO;
                } else {
                    $res = self::castExpr($to, $from->getType(), $fromVar, $toVarAssign . '*');

                    return <<<GO
if {$fromVar} != nil {
    {$res}
}

GO;
                }
            }
        }

        return '';
    }

}