<?php
namespace Swaggest\GoCodeBuilder\Templates\Struct;


use Swaggest\CodeBuilder\AbstractTemplate;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;

class StructFunctions extends AbstractTemplate
{
    /** @var StructDef */
    private $struct;

    /**
     * StructFunctions constructor.
     * @param StructDef $struct
     */
    public function __construct(StructDef $struct)
    {
        $this->struct = $struct;
    }


    public function toString()
    {
        $result = '';
        foreach ($this->struct->getFuncs() as $func) {
            $result .= $func->toString();
        }
        return $result;
    }


}