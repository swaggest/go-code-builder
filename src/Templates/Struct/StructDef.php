<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class StructDef extends GoTemplate
{
    private $name;
    /** @var StructProperty[] */
    private $properties;

    /** @var FuncDef[] */
    private $funcs = array();

    /** @var Code */
    private $code;

    /** @var Import|null */
    private $import;

    /**
     * StructDef constructor.
     * @param string $name
     * @param string $comment
     */
    public function __construct($name, $comment = '')
    {
        $this->name = $name;
        $this->comment = $comment;
        $this->code = new Code();
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
     * @param mixed $name
     * @return StructDef
     */
    public function setName($name)
    {
        $this->name = $name;
        return $this;
    }

    /**
     * @return Import|null
     */
    public function getImport()
    {
        return $this->import;
    }

    /**
     * @param Import $import
     * @return StructDef
     */
    public function setImport(Import $import = null)
    {
        $this->import = $import;
        return $this;
    }

    public function setType(Type $type)
    {
        $this->import = $type->getImport();
        $this->name = $type->getName();
        return $this;
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

    public function getCode()
    {
        return $this->code;
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

    protected function toString()
    {
        return <<<GO
{$this->renderComment()}{$this->renderFields()}{$this->renderFuncs()}{$this->code}
GO;

    }

    public function renderFields()
    {
        return new StructFields($this);
    }

    public function renderFuncs()
    {
        return new StructFunctions($this);
    }

    /**
     * @return StructType
     */
    public function getType()
    {
        return new StructType($this);
    }
}