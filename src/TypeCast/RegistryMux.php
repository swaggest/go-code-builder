<?php

namespace Swaggest\GoCodeBuilder\TypeCast;


use Swaggest\GoCodeBuilder\Exception;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Type\TypeCastException;

class RegistryMux implements Registry
{
    /**
     * @var Registry[]
     */
    private $registries = array();
    public function addRegistry(Registry $registry)
    {
        $this->registries[] = $registry;
        return $this;
    }

    /**
     * @return static
     */
    public static function getStd()
    {
        $mux = new static;
        $mux->addRegistry(new Time());
        return $mux;
    }

    /**
     * @param string $toTypeString
     * @param string $fromTypeString
     * @return bool
     */
    public function canProcess($toTypeString, $fromTypeString)
    {
        foreach ($this->registries as $registry) {
            if ($registry->canProcess($toTypeString, $fromTypeString)) {
                return true;
            }
        }
        return false;
    }

    /**
     * @param $toTypeString
     * @param $fromTypeString
     * @param $toVarName
     * @param $fromVarName
     * @param $assignOp
     * @return Code|string
     * @throws TypeCastException
     */
    public function process($toTypeString, $fromTypeString, $toVarName, $fromVarName, $assignOp)
    {
        foreach ($this->registries as $registry) {
            if ($registry->canProcess($toTypeString, $fromTypeString)) {
                return $registry->process($toTypeString, $fromTypeString, $toVarName, $fromVarName, $assignOp);
            }
        }
        throw new TypeCastException('No registry can process');
    }
}