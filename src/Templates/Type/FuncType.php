<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Exception;
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
        $res = 'func(';
        if ($this->func->getArguments() !== null) {
            $res .= $this->func->getArguments()->toTypesString();
        }
        $res .= ') ';
        if ($this->func->getResult() !== null) {
            $res .= $this->func->getResult()->toTypesString();
        }
        return trim($res);
    }

    public function getTypeString()
    {
        throw new Exception('TypeString rendering for func is not implemented');
    }


}