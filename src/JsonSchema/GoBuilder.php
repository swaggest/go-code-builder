<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\GoCodeBuilder;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Struct\Tags;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\JsonSchema\JsonSchema;

/**
 * @todo properly process $ref, $schema property names
 */
class GoBuilder
{
    /** @var \SplObjectStorage|GeneratedStruct[] */
    private $generatedClasses;
    private $untitledIndex = 0;
    /** @var GoCodeBuilder */
    public $codeBuilder;

    public function __construct()
    {
        $this->generatedClasses = new \SplObjectStorage();
        $this->codeBuilder = new GoCodeBuilder();
    }

    /**
     * @param JsonSchema $schema
     * @param string $path
     * @return AnyType
     */
    public function getType(JsonSchema $schema, $path = '#')
    {
        $typeBuilder = new TypeBuilder($schema, $path, $this);
        return $typeBuilder->build();
    }


    public function getClass(JsonSchema $schema, $path)
    {
        if ($this->generatedClasses->contains($schema)) {
            return $this->generatedClasses[$schema]->structDef;
        } else {
            return $this->makeClass($schema, $path)->structDef;
        }
    }

    private function makeClass(JsonSchema $schema, $path)
    {
        if (empty($path)) {
            throw new Exception('Empty path');
        }
        $generatedStruct = new GeneratedStruct();
        $generatedStruct->schema = $schema;

        if ($path === '#') {
            $structDef = new StructDef('Untitled' . ++$this->untitledIndex);
        } else {
            $structDef = new StructDef($this->codeBuilder->exportableName($path));
        }

        $generatedStruct->structDef = $structDef;
        $generatedStruct->path = $path;

        $this->generatedClasses->attach($schema, $generatedStruct);

        if ($schema->properties !== null) {
            foreach ($schema->properties as $name => $property) {
                $fieldName = $this->codeBuilder->exportableName($name);
                $goProperty = new StructProperty(
                    $fieldName,
                    $this->getType($property, $path . '->' . $name),
                    (new Tags())->setTag('json', $name)
                );
                if ($property->description) {
                    $goProperty->setComment($property->description);
                }
                $structDef->addProperty($goProperty);
            }
        }

        return $generatedStruct;
    }

    /**
     * @return GeneratedStruct[]
     */
    public function getGeneratedClasses()
    {
        $result = array();
        foreach ($this->generatedClasses as $schema) {
            $result[] = $this->generatedClasses[$schema];
        }
        return $result;
    }

}