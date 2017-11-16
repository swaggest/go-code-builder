<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;
use Swaggest\JsonSchema\Constraint\Type;
use Swaggest\JsonSchema\JsonSchema;
use Swaggest\JsonSchema\Schema;
use Swaggest\GoCodeBuilder\Templates\Type\Type as GoType;

class TypeBuilder
{
    /** @var JsonSchema */
    private $schema;
    /** @var string */
    private $path;
    /** @var GoBuilder */
    private $goBuilder;
    /** @var AnyType[] */
    private $result;

    /**
     * TypeBuilder constructor.
     * @param Schema $schema
     * @param $path
     * @param GoBuilder $goBuilder
     */
    public function __construct(JsonSchema $schema, $path, GoBuilder $goBuilder)
    {
        $this->schema = $schema;
        $this->path = $path;
        $this->goBuilder = $goBuilder;
    }

    private function processLogicType()
    {
        $orSchemas = null;
        if ($this->schema->allOf !== null) {
            $orSchemas = $this->schema->allOf;
        } elseif ($this->schema->anyOf !== null) {
            $orSchemas = $this->schema->anyOf;
        } elseif ($this->schema->oneOf !== null) {
            $orSchemas = $this->schema->oneOf;
        }

        if ($orSchemas !== null) {
            foreach ($orSchemas as $item) {
                $structProperty = new StructProperty(null, $this->goBuilder->getType($item, $this->path));
                $this->makeResultStruct()->addProperty($structProperty);
            }
        }
    }

    private function processArrayType()
    {
        $schema = $this->schema;

        $pathItems = (string)Schema::names()->items;
        if ($schema->items instanceof Schema) {
            $items = array();
            $additionalItems = $schema->items;
        } elseif ($schema->items === null) { // items defaults to empty schema so everything is valid
            $items = array();
            $additionalItems = true;
        } else { // listed items
            $items = $schema->items;
            $additionalItems = $schema->additionalItems;
            $pathItems = (string)Schema::names()->additionalItems;
        }

        if ($items !== null || $additionalItems !== null) {
            $itemsLen = is_array($items) ? count($items) : 0;
            $index = 0;
            if ($index < $itemsLen) {
            } else {
                if ($additionalItems instanceof Schema) {
                    $this->result[] = new Slice($this->goBuilder->getType($additionalItems, $this->path . '->' . $pathItems));
                }
            }
        }
    }

    private function processObjectType()
    {
        if ($this->schema->patternProperties !== null) {
            foreach ($this->schema->patternProperties as $pattern => $schema) {
                $structProperty = new StructProperty(
                    $this->goBuilder->codeBuilder->privateName('patternProperties_' . $pattern),
                    new Map(
                        new GoType("string"),
                        $this->goBuilder->getType($schema, $this->path . '->' . $pattern)
                    )
                );
                $this->makeResultStruct()->addProperty($structProperty);
            }
        }

        if ($this->schema->additionalProperties instanceof Schema) {
            $structProperty = new StructProperty(
                $this->goBuilder->codeBuilder->privateName('additionalProperties'),
                new Map(
                    new GoType("string"),
                    $this->goBuilder->getType(
                        $this->schema->additionalProperties,
                        $this->path . '->' . (string)Schema::names()->additionalProperties)
                )
            );
            $this->makeResultStruct()->addProperty($structProperty);
        }
    }

    private function typeSwitch($type)
    {
        switch ($type) {
            case Type::INTEGER:
                return new GoType('int64');

            case Type::NUMBER:
                return new GoType('float64');

            case TYPE::BOOLEAN:
                return new GoType('bool');

            case Type::STRING:
                return new GoType('string');

            case Type::OBJECT:
                return new GoType('map[string]interface{}');

            case Type::ARR:
                return new GoType('[]interface{}');

            case Type::NULL:
                return new GoType('interface{}');

            default:
                return null;
        }
    }

    private function processNamedClass()
    {
        if ($this->schema->properties !== null) {
            $this->result[] = new Pointer($this->makeResultStruct()->getType());
        }
    }


    private function makeResultStruct()
    {
        if ($this->resultStruct === null) {
            $this->resultStruct = $this->goBuilder->getClass($this->schema, $this->path);
            $this->resultStruct->setComment($this->path);
        }
        return $this->resultStruct;
    }

    /** @var FuncDef */
    private $resultStructImportFunc;
    /** @var FuncDef */
    private $resultStructExportFunc;

    /** @var StructDef */
    private $resultStruct;

    /**
     * @return AnyType
     * @throws Exception
     */
    public function build()
    {
        $this->result = array();

        if (null !== $path = $this->schema->getFromRef()) {
            $this->path = $this->schema->getFromRef();
        }


        $this->processNamedClass();
        $this->processLogicType();
        $this->processArrayType();
        $this->processObjectType();

        if (is_array($this->schema->type)) {
            throw new Exception("Can not map multiple types in GO type");
        } elseif ($this->schema->type) {
            $this->result[] = $this->typeSwitch($this->schema->type);
        }


        if ($this->resultStruct !== null) {
            return new Pointer($this->resultStruct->getType());
        }
        if (empty($this->result)) {
            return new \Swaggest\GoCodeBuilder\Templates\Type\Type('interface{}');
        }
        return $this->result[0];

    }
}