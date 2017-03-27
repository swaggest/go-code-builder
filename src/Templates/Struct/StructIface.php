<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;


use Swaggest\CodeBuilder\AbstractTemplate;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Iface\IfaceDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;

class StructIface extends GoTemplate
{
    private $struct;
    private $name;

    public function __construct(StructDef $struct, $name, $comment = '')
    {
        $this->struct = $struct;
        $this->name = $name;
        $this->comment = $comment;
    }

    public function getIface()
    {
        $iface = new IfaceDef($this->name, $this->comment);
        foreach ($this->struct->getFuncs() as $func) {
            $iface->addFunc($func);
        }
        return $iface;
    }

    public function toString()
    {
        return $this->getIface()->toString();
    }


}