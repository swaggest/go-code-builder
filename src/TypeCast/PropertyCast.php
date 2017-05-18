<?php

namespace Swaggest\GoCodeBuilder\TypeCast;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func\Argument;
use Swaggest\GoCodeBuilder\Templates\Func\Arguments;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Func\Result;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\TypeCast;

class PropertyCast implements CastFunctions
{
    /** @var StructDef */
    private $baseStruct;

    private $propertyName;

    /** @var AnyType */
    private $derivedType;

    /** @var Registry */
    private $typeRegistry;

    /**
     * PropertyCast constructor.
     * @param StructDef $baseStruct
     * @param $propertyName
     * @param AnyType $derivedType
     * @param Registry $typeRegistry
     */
    public function __construct(StructDef $baseStruct, $propertyName, AnyType $derivedType, Registry $typeRegistry = null)
    {
        $this->baseStruct = $baseStruct;
        $this->propertyName = $propertyName;
        $this->derivedType = $derivedType;
        $this->typeRegistry = $typeRegistry;
    }


    function getMapTo()
    {
        $mapTo = new FuncDef('MapTo');
        $mapTo->setSelf(new Argument('base', $this->baseStruct->getType()));
        $mapTo->setResult((new Result())->add(null, $this->derivedType));

        $properties = $this->baseStruct->getProperties();
        $property = $properties[$this->propertyName];

        $cast = new TypeCast($this->derivedType, $property->getType(), 'result', 'base.' . $this->propertyName, $this->typeRegistry);

        $code = <<<GO
var result {$this->derivedType->render()}
{$cast->render()}
return result
GO;

        $mapTo->setBody(new Code($code));
        return $mapTo;
    }

    function getLoadFrom()
    {
        $loadFrom = new FuncDef('LoadFrom');
        $loadFrom->setSelf(new Argument('base', $this->baseStruct->getType()));
        $loadFrom->setArguments((new Arguments())->add('derived', $this->derivedType));

        $properties = $this->baseStruct->getProperties();
        $property = $properties[$this->propertyName];

        $cast = new TypeCast($property->getType(), $this->derivedType, 'base.' . $this->propertyName, 'derived', $this->typeRegistry);

        $code = <<<GO
{$cast->render()}
GO;

        $loadFrom->setBody(new Code($code));
        return $loadFrom;

    }

    /**
     * @return string
     */
    public function getBaseTypeString()
    {
        return $this->baseStruct->getType()->getTypeString();
    }

    /**
     * @return string
     */
    public function getDerivedTypeString()
    {
        return $this->derivedType->getTypeString();
    }


}