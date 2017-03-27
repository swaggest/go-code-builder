<?php

namespace Swaggest\GoCodeBuilder\Templates\Iface;

use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Func\FuncIface;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class IfaceDef extends GoTemplate
{
    private $name;

    public function __construct($name, $comment = null)
    {
        $this->name = $name;
        $this->comment = $comment;
    }

    /** @var FuncDef[] */
    private $funcs;

    /** @var Type[] */
    private $types;

    public function addFunc(FuncDef $func, $prepend = false)
    {
        if ($prepend) {
            array_unshift($this->funcs, $func);
        } else {
            $this->funcs[] = $func;
        }
        return $this;
    }

    public function addType(Type $func, $prepend = false)
    {
        if ($prepend) {
            array_unshift($this->types, $func);
        } else {
            $this->types[] = $func;
        }
        return $this;
    }

    private function renderTypes()
    {
        $result = '';
        if (empty($this->types)) {
            return $result;
        }
        foreach ($this->types as $type) {
            $result .= $type . "\n\n";
        }
        return $this->padLines("\t", substr($result, 0, -1), false);
    }

    private function renderFuncs()
    {
        $result = '';
        if (empty($this->funcs)) {
            return $result;
        }
        foreach ($this->funcs as $func) {
            $result .= (new FuncIface($func)) . "\n\n";
        }
        return $this->padLines("\t", substr($result, 0, -1), false);
    }

    public function toString()
    {
        $separator = !empty($this->types) && !empty($this->funcs) ? "\n" : '';
        $result = <<<GO
{$this->renderComment()}type {$this->name} interface {
{$this->renderTypes()}{$separator}{$this->renderFuncs()}}


GO;
        return $result;
    }
}