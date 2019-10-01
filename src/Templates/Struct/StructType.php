<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;


use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\NamedType;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class StructType extends GoTemplate implements NamedType
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

    public function getName()
    {
        return $this->structDef->getName();
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