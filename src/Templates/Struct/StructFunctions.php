<?php
namespace Swaggest\GoCodeBuilder\Templates\Struct;


use Swaggest\CodeBuilder\AbstractTemplate;

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


    protected function toString()
    {
        $result = '';
        foreach ($this->struct->getFuncs() as $func) {
            $result .= $func->render();
        }
        return $result;
    }


}