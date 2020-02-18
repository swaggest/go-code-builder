<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func\Argument;
use Swaggest\GoCodeBuilder\Templates\Func\Arguments;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Func\Result;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;

class FluentSetter
{
    public static function addToStruct(StructDef $structDef, StructProperty $goProperty)
    {
        $structDef->addFunc(self::make($structDef, $goProperty));
        $map = self::makeMap($structDef, $goProperty);
        if ($map !== null) {
            $structDef->addFunc($map);
        }
    }

    public static function make(StructDef $structDef, StructProperty $goProperty)
    {
        $setter = new FuncDef(
            'With' . $goProperty->getName(),
            'With' . $goProperty->getName() . ' sets ' . $goProperty->getName() . ' value.'
        );
        $setter->setSelf(new Argument('v', new Pointer($structDef->getType())));

        $valType = $goProperty->getType();
        $valRef = '';
        $valVariadic = false;

        if ($valType instanceof Pointer) {
            $valType = $valType->getType();
            $valRef = '&';
        } elseif ($valType instanceof Slice) {
            $valType = $valType->getType();
            $valVariadic = true;
        }


        $setter->setArguments((new Arguments())->add('val', $valType, $valVariadic));
        $setter->setResult((new Result())->add(null, new Pointer($structDef->getType())));

        $setter->setBody(new Code(<<<GO
v.{$goProperty->getName()} = {$valRef}val
return v
GO
        ));

        return $setter;
    }

    public static function makeMap(StructDef $structDef, StructProperty $goProperty)
    {
        $valType = $goProperty->getType();
        if (!$valType instanceof Map) {
            return null;
        }

        $setter = new FuncDef(
            'With' . $goProperty->getName() . 'Item',
            'With' . $goProperty->getName() . 'Item sets ' . $goProperty->getName() . ' item value.'
        );
        $setter->setSelf(new Argument('v', new Pointer($structDef->getType())));

        $setter->setArguments((new Arguments())
            ->add('key', $valType->getKeyType())
            ->add('val', $valType->getValueType())
        );
        $setter->setResult((new Result())->add(null, new Pointer($structDef->getType())));

        $setter->setBody(new Code(new PlaceholderString(<<<GO
if v.{$goProperty->getName()} == nil {
    v.{$goProperty->getName()} = make(:type, 1)
}

v.{$goProperty->getName()}[key] = val

return v
GO
            , [':type' => $valType])));

        return $setter;
    }
}