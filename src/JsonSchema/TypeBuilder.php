<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
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
     * @param GoBuilder $phpBuilder
     */
    public function __construct(JsonSchema $schema, $path, GoBuilder $phpBuilder)
    {
        $this->schema = $schema;
        $this->path = $path;
        $this->goBuilder = $phpBuilder;
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
                $this->result[] = $this->goBuilder->getType($item, $this->path);
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
                $this->result[] = new ArrayOf($this->goBuilder->getType($schema, $this->path . '->' . $pattern));
            }
        }

        if ($this->schema->additionalProperties instanceof Schema) {
            $this->result[] = new ArrayOf($this->goBuilder->getType(
                $this->schema->additionalProperties,
                $this->path . '->' . (string)Schema::names()->additionalProperties)
            );
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

    private function processNamedClass(JsonSchema $schema, $path)
    {
        if ($schema->properties !== null) {
            $class = $this->goBuilder->getClass($schema, $path);
            $this->result[] = $class;
        }
    }

    /**
     * @return AnyType
     */
    public function build()
    {
        $this->result = array();

        if (null !== $path = $this->schema->getFromRef()) {
            $this->path = $this->schema->getFromRef();
        }


        $this->processNamedClass($this->schema, $this->path);
        $this->processLogicType();
        $this->processArrayType();
        $this->processObjectType();

        if (is_array($this->schema->type)) {
            foreach ($this->schema->type as $type) {
                $this->result[] = $this->typeSwitch($type);
            }
        } elseif ($this->schema->type) {
            $this->result[] = $this->typeSwitch($this->schema->type);
        }

        return $this->result[0];

    }
}