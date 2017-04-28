<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class FuncType extends GoTemplate implements AnyType
{
    private $func;

    public function __construct(FuncDef $func)
    {
        $this->func = $func;
    }

    protected function toString()
    {
        // func({$data->iteratorResultGoType}, error) bool
        return 'func(' . $this->func->getArguments()->toTypesString() . ') ' . $this->func->getResult()->toTypesString();
    }


}