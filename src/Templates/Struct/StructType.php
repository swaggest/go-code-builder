<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;


use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class StructType extends GoTemplate implements AnyType
{
    /** @var StructDef */
    private $structDef;

    /**
     * StructType constructor.
     * @param StructDef $structDef
     */
    public function __construct(StructDef $structDef)
    {
        $this->structDef = $structDef;
    }

    private function getType()
    {
        return new Type($this->structDef->getName(), $this->structDef->getImport());
    }

    protected function toString()
    {
        return $this->getType()->toString();
    }

    public function getTypeString()
    {
        return $this->getType()->getTypeString();
    }

}