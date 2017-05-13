<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Import;

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
            list($key, $val) = explode(']', substr($typeString, 4), 2);
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


    public static function getBasicType(AnyType $type)
    {
        if (!$type instanceof Type) {
            return false;
        }

        if ($type->getImport() !== null) {
            return false;
        }

        return $type->getName();
    }

    public static function isInt(AnyType $type)
    {
        if (!$basicType = self::getBasicType($type)) {
            return false;
        }
        return in_array($basicType, array('int', 'int8', 'int16', 'int32', 'int64',
            'uint', 'uint8', 'uint16', 'uint32', 'uint64'), true);
    }

    public static function isFloat(AnyType $type)
    {
        if (!$basicType = self::getBasicType($type)) {
            return false;
        }
        return in_array($basicType, array('float32', 'float64'), true);
    }

    public static function isNumber(AnyType $type)
    {
        return self::isFloat($type) || self::isInt($type);
    }

    public static function isCastable(Type $to, Type $from)
    {
        if (self::isNumber($from) && self::isNumber($to)) {
            return true;
        }

        return false;
    }

    public static function resolvePointer(AnyType $type) {
        if ($type instanceof Pointer) {
            return self::resolvePointer($type->getType());
        } else {
            return $type;
        }
    }
}