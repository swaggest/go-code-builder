<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\GoCodeBuilder\Templates\Func;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;

class StructDef extends GoTemplate
{
    private $name;
    /** @var StructProperty[] */
    private $properties;

    /** @var FuncDef[] */
    private $funcs = array();

    public function __construct($name, $comment = '')
    {
        $this->name = $name;
        $this->comment = $comment;
    }

    public function addProperty(StructProperty $property, $prepend = false)
    {
        if ($prepend) {
            $this->properties = array($property->getName() => $property) + $this->properties;
        } else {
            $this->properties[$property->getName()] = $property;
        }
        return $this;
    }

    /**
     * @return mixed
     */
    public function getName()
    {
        return $this->name;
    }



    /**
     * @return StructProperty[]
     */
    public function getProperties()
    {
        if (null === $this->properties) {
            return array();
        }
        return $this->properties;
    }

    /**
     * @return Func\FuncDef[]
     */
    public function getFuncs()
    {
        return $this->funcs;
    }


    public function addFunc(FuncDef $func, $prepend = false)
    {
        if ($prepend) {
            array_unshift($this->funcs, $func);
        } else {
            $this->funcs[] = $func;
        }
        return $this;
    }

    public function toString()
    {
        return <<<GO
{$this->renderComment()}{$this->renderStruct()}{$this->renderFuncs()}
GO;

    }

    public function renderStruct()
    {
        return new StructType($this);
    }

    public function renderFuncs()
    {
        return new StructFunctions($this);
    }


}