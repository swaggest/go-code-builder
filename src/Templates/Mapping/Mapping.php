<?php

namespace Swaggest\GoCodeBuilder\Templates\Mapping;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;

class Mapping extends GoTemplate
{
    /** @var AnyType */
    public $baseType;

    /** @var AnyType */
    public $derivedType;

    public $baseName;

    public $derivedName;

    /** @var Mapping[] */
    public $properties;

    protected function toString()
    {
        // TODO: Implement toString() method.
    }

    public function mapToDerived(Mapping $parentStruct = null)
    {
        $code = new Code();
        if ($this->baseType)

        $code->addSnippet(<<<GO
{$this->baseName} = {$this->derivedName}

GO
);

    }

    public function loadFromDerived()
    {

    }
}