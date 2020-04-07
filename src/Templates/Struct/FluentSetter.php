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
        $ensurer = self::makeEnsurer($structDef, $goProperty);
        if ($ensurer !== null) {
            $structDef->addFunc($ensurer);
        }
    }

    public static function makeEnsurer(StructDef $structDef, StructProperty $goProperty) {
        $type = $goProperty->getType();

        if (!$type instanceof Pointer || !$type->getType() instanceof StructType) {
            return null;
        }

        $receiver = strtolower($structDef->getType()->getName()[0]);

        $ensurer = new FuncDef(
            $goProperty->getName() . 'Ens',
            $goProperty->getName() . 'Ens ensures returned ' . $goProperty->getName() . ' is not nil.'
        );
        $ensurer->setSelf(new Argument($receiver, new Pointer($structDef->getType())));

        $ensurer->setResult((new Result())->add(null, $type));

        $ensurer->setBody(new Code(<<<GO
if {$receiver}.{$goProperty->getName()} == nil {
    {$receiver}.{$goProperty->getName()} = new({$type->getType()->getTypeString()})
}

return {$receiver}.{$goProperty->getName()}
GO
        ));

        return $ensurer;
    }

    public static function make(StructDef $structDef, StructProperty $goProperty)
    {
        $receiver = strtolower($structDef->getType()->getName()[0]);

        $setter = new FuncDef(
            'With' . $goProperty->getName(),
            'With' . $goProperty->getName() . ' sets ' . $goProperty->getName() . ' value.'
        );
        $setter->setSelf(new Argument($receiver, new Pointer($structDef->getType())));

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
{$receiver}.{$goProperty->getName()} = {$valRef}val
return {$receiver}
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

        $receiver = strtolower($structDef->getType()->getName()[0]);

        $setter = new FuncDef(
            'With' . $goProperty->getName() . 'Item',
            'With' . $goProperty->getName() . 'Item sets ' . $goProperty->getName() . ' item value.'
        );
        $setter->setSelf(new Argument($receiver, new Pointer($structDef->getType())));

        $setter->setArguments((new Arguments())
            ->add('key', $valType->getKeyType())
            ->add('val', $valType->getValueType())
        );
        $setter->setResult((new Result())->add(null, new Pointer($structDef->getType())));

        $setter->setBody(new Code(new PlaceholderString(<<<GO
if :receiver.{$goProperty->getName()} == nil {
    :receiver.{$goProperty->getName()} = make(:type, 1)
}

:receiver.{$goProperty->getName()}[key] = val

return :receiver
GO
            , [':type' => $valType, ':receiver' => new Code($receiver)])));

        return $setter;
    }
}