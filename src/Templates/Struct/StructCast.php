<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func\Argument;
use Swaggest\GoCodeBuilder\Templates\Func\Arguments;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Func\Result;
use Swaggest\GoCodeBuilder\Templates\Type\TypeCast;
use Swaggest\GoCodeBuilder\TypeCast\Registry;

class StructCast
{
    /** @var StructDef */
    private $baseStruct;

    /** @var StructDef */
    private $derivedStruct;

    /**
     * @var string[] base to derived names map
     */
    private $propNamesMap = array();

    /** @var Registry */
    private $typeRegistry;

    /**
     * StructCast constructor.
     * @param StructDef $baseStruct
     * @param StructDef $derivedStruct
     * @param \string[] $propNamesMap
     * @param Registry $registry
     */
    public function __construct(StructDef $baseStruct, StructDef $derivedStruct, $propNamesMap = array(),
                                Registry $registry = null)
    {
        $this->baseStruct = $baseStruct;
        $this->derivedStruct = $derivedStruct;
        $this->propNamesMap = $propNamesMap;
        $this->typeRegistry = $registry;
    }

    public function setPropMap($baseName, $derivedName)
    {
        $this->propNamesMap[$baseName] = $derivedName;
        return $this;
    }


    public function getMapTo()
    {
        $mapTo = new FuncDef('MapTo');
        $mapTo->setSelf(new Argument('base', $this->baseStruct->getType()));
        $mapTo->setResult((new Result())->add(null, $this->derivedStruct->getType()));


        $code = new Code(<<<GO
result := {$this->derivedStruct->getType()->render()}{}

GO
        );

        $baseProperties = $this->baseStruct->getProperties();

        $derivedProperties = $this->derivedStruct->getProperties();


        foreach ($this->propNamesMap as $baseName => $derivedName) {
            $cast = new TypeCast(
                $derivedProperties[$derivedName]->getType(),
                $baseProperties[$baseName]->getType(),
                'result.' . $derivedName,
                'base.' . $baseName,
                $this->typeRegistry
            );

            $code->addSnippet($cast->render() . "\n");
        }

        $code->addSnippet('return result');

        $mapTo->setBody($code);

        return $mapTo;
    }


    public function getLoadFrom()
    {
        $loadFrom = new FuncDef('LoadFrom');
        $loadFrom->setSelf(new Argument('base', $this->baseStruct->getType()));
        $loadFrom->setArguments((new Arguments())->add('derived', $this->derivedStruct->getType()));

        $code = new Code();

        $baseProperties = $this->baseStruct->getProperties();

        $derivedProperties = $this->derivedStruct->getProperties();
        foreach ($this->propNamesMap as $baseName => $derivedName) {
            $cast = new TypeCast(
                $baseProperties[$baseName]->getType(),
                $derivedProperties[$derivedName]->getType(),
                'base.' . $baseName,
                'derived.' . $derivedName,
                $this->typeRegistry
            );

            $code->addSnippet($cast->render() . "\n");
        }

        $loadFrom->setBody($code);

        return $loadFrom;
    }

    /**
     * @return StructDef
     */
    public function getBaseStruct()
    {
        return $this->baseStruct;
    }

    /**
     * @return StructDef
     */
    public function getDerivedStruct()
    {
        return $this->derivedStruct;
    }
}