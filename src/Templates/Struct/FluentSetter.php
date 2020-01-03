<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func\Argument;
use Swaggest\GoCodeBuilder\Templates\Func\Arguments;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Func\Result;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;

class FluentSetter
{
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

}