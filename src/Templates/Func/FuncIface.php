<?php

namespace Swaggest\GoCodeBuilder\Templates\Func;


use PhpLang\ScopeExit;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class FuncIface extends GoTemplate
{
    /** @var FuncDef */
    private $funcDef;

    /**
     * FuncIface constructor.
     * @param FuncDef $funcDef
     */
    public function __construct(FuncDef $funcDef)
    {
        $this->funcDef = $funcDef;
    }

    protected function toString()
    {
        $this->funcDef->setRenderMode(FuncDef::RENDER_IFACE);
        /** @noinspection PhpUnusedLocalVariableInspection */
        $_ = new ScopeExit(function(){
            $this->funcDef->setRenderMode(FuncDef::RENDER_FUNC);
        });
        return (string)$this->funcDef;
    }
}