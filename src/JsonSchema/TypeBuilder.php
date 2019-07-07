<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\Constant\TypeConstBlock;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Struct\StructType;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;
use Swaggest\JsonSchema\Constraint\Format;
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
     * @param string $path
     * @param GoBuilder $goBuilder
     */
    public function __construct(Schema $schema, $path, GoBuilder $goBuilder)
    {
        $this->schema = $schema;
        $this->path = $path;
        $this->goBuilder = $goBuilder;
    }

    private function makeName(AnyType $type)
    {
        $type = Pointer::tryDereferenceOnce($type);
        if ($type instanceof StructDef) {
            return $type->getName();
        }
        if ($type instanceof StructType) {
            return $type->getName();
        }
        if ($type instanceof GoType) {
            if ($type->getName() === 'interface{}') {
                return 'anything';
            }
            return $type->getName();
        }
        return null;
    }

    private function processOr($orSchemas, $kind)
    {
        foreach ($orSchemas as $i => $item) {
            $name = $this->goBuilder->codeBuilder->exportableName($kind . '/' . $i);
            $itemType = Pointer::tryDereferenceOnce($this->goBuilder->getType($item, $this->path . '/' . $kind . '/' . $i));
            if (null !== $betterName = $this->makeName($itemType)) {
                $name = $this->goBuilder->codeBuilder->exportableName($betterName);
            }

            $resultStruct = $this->makeResultStruct();

            if ($this->goBuilder->options->trimParentFromPropertyNames) {
                if (strpos($name, $resultStruct->getName()) === 0) {
                    $name = substr($name, strlen($resultStruct->getName()));
                }
            }

            if (!$itemType instanceof Map && !$itemType instanceof Slice) {
                if ($itemType->getTypeString() === 'interface{}') {
                    continue;
                }
                $itemType = new Pointer($itemType);
            }

            $structProperty = new StructProperty($name, $itemType);
            $structProperty->getTags()->setTag('json', '-');
            $resultStruct->addProperty($structProperty);
            $generatedStruct = $this->getGeneratedStruct();
            $generatedStruct->marshalJson->addSomeOf($kind, $name);
        }
    }

    private function processLogicType()
    {
        if (in_array($this->schema->type, [JsonSchema::STRING, JsonSchema::INTEGER, JsonSchema::NUMBER, JsonSchema::BOOLEAN], true)) {
            // Skip logical keywords if scalar type provided
            return;
        }
        if ($this->schema->allOf !== null) {
            $this->processOr($this->schema->allOf, 'allOf');
        } elseif ($this->schema->anyOf !== null) {
            $this->processOr($this->schema->anyOf, 'anyOf');
        } elseif ($this->schema->oneOf !== null) {
            $this->processOr($this->schema->oneOf, 'oneOf');
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
                    $this->result[] = new Slice(Pointer::tryDereferenceOnce(
                        $this->goBuilder->getType($additionalItems, $this->path . '->' . $pathItems))
                    );
                }
            }
        }
    }

    private function processObjectType()
    {
        if ($this->schema->patternProperties !== null) {
            foreach ($this->schema->patternProperties as $pattern => $schema) {
                $itemType = Pointer::tryDereferenceOnce($this->goBuilder->getType($schema, $this->path . '->' . $pattern));
                $name = $this->goBuilder->codeBuilder->exportableName('patternProperties_' . $pattern);
                if (null !== $betterName = $this->makeName($itemType)) {
                    $name = $this->goBuilder->codeBuilder->exportableName('MapOf_' . $betterName . '_Values');
                }
                $structProperty = new StructProperty(
                    $name,
                    new Map(
                        new GoType("string"),
                        $itemType
                    )
                );
                $structProperty->getTags()->setTag('json', '-');
                $structProperty->setComment('Key must match pattern: ' . $pattern);
                $this->makeResultStruct()->addProperty($structProperty);
                $this->getGeneratedStruct()->marshalJson->addPatternProperty($pattern, $structProperty->getName());
            }
        }

        if ($this->schema->additionalProperties instanceof Schema) {
            $goType = Pointer::tryDereferenceOnce(
                $this->goBuilder->getType(
                    $this->schema->additionalProperties,
                    $this->path . '->' . (string)Schema::names()->additionalProperties)
            );

            if ($this->schema->properties !== null || $this->schema->patternProperties !== null) {
                $structProperty = new StructProperty(
                    $this->goBuilder->codeBuilder->exportableName('additionalProperties'),
                    new Map(new GoType("string"), $goType)
                );
                $structProperty->getTags()->setTag('json', '-');
                $this->makeResultStruct()->addProperty($structProperty);
                $this->getGeneratedStruct()->marshalJson->enableAdditionalProperties();
            } else {
                $this->result[] = new Map(new GoType('string'), $goType);
            }
        }
    }

    private function typeSwitch($type, $minimum = null, $maximum = null, $format = null)
    {
        switch ($type) {
            case Type::INTEGER:
                if ($this->goBuilder->options->optimizeIntegers && $minimum !== null && $maximum !== null) {
                    if ($minimum >= 0 && $maximum <= 255) {
                        return new GoType('uint8');
                    }
                    if ($minimum >= 0 && $maximum <= 65535) {
                        return new GoType('uint16');
                    }
                    if ($minimum >= 0 && $maximum <= 4294967295) {
                        return new GoType('uint32');
                    }
                    if ($minimum >= 0) {
                        return new GoType('uint64');
                    }

                    if ($minimum >= -128 && $maximum <= 127) {
                        return new GoType('int8');
                    }

                    if ($minimum >= -32768 && $maximum <= 32767) {
                        return new GoType('int16');
                    }

                    if ($minimum >= -2147483648 && $maximum <= 2147483647) {
                        return new GoType('int32');
                    }
                }
                return new GoType('int64');

            case Type::NUMBER:
                return new GoType('float64');

            case TYPE::BOOLEAN:
                return new GoType('bool');

            case Type::STRING:
                if ($format === Format::DATE_TIME) {
                    return new Pointer(new GoType('Time', new Import('time')));
                }
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
            $this->result[] = $this->makeResultStruct()->getType();
        }
    }

    private function processConst()
    {
        if ($this->schema->const !== null) { // todo properly process null const
            $path = $this->goBuilder->pathToName($this->path);
            $typeName = $this->goBuilder->codeBuilder->exportableName($path);
            $type = new GoType($typeName);

            $baseType = new GoType('interface{}');
            if (is_string($this->schema->const)) {
                $baseType = new GoType('string');
            } elseif (is_int($this->schema->const)) {
                $baseType = new GoType('int64');
            } elseif (is_float($this->schema->const)) {
                $baseType = new GoType('float64');
            } elseif (is_bool($this->schema->const)) {
                $baseType = new GoType('bool');
            }

            if ($this->goBuilder->options->hideConstProperties) {
                return $baseType;
            }

            $typeConstBlock = new TypeConstBlock($type);


            $this->goBuilder->getCode()->addSnippet(<<<GO
// $typeName is a constant type.
type $typeName {$baseType->getName()}


GO
            );

            $typeConstBlock->addValue(
                $this->goBuilder->codeBuilder->exportableName($path . '_' . $this->schema->const),
                $this->schema->const
            );

            $this->goBuilder->getCode()->addSnippet($typeConstBlock);
            $this->goBuilder->getCode()->addSnippet(new MarshalEnum($type, $baseType, $typeConstBlock->getValues(), $this->goBuilder));
            return $type;
        }
        return null;
    }

    private function processEnum(AnyType $baseType)
    {
        $oneOfEnum = true;
        $enum = null;
        /** @var Schema[] $enumSchemas */
        $enumSchemas = null;
        if ($this->schema->oneOf !== null) {
            foreach ($this->schema->oneOf as $schema) {
                if ($schema->const !== null) {
                    $enum[] = $schema->const;
                    $enumSchemas[] = $schema;
                } else {
                    $oneOfEnum = false;
                }
            }
        } else {
            $oneOfEnum = false;
        }

        if (!$oneOfEnum) {
            $enum = $this->schema->enum;
            $enumSchemas = null;
        }

        if ($enum !== null) {
            if ($this->goBuilder->options->hideConstProperties && count($enum) === 1) {
                return $baseType;
            }

            $path = $this->goBuilder->pathToName($this->path);
            $typeName = $this->goBuilder->codeBuilder->exportableName($path);
            $type = new GoType($typeName);

            $typeConstBlock = new TypeConstBlock($type);

            $this->goBuilder->getCode()->addSnippet(<<<GO
// $typeName is an enum type.
type $typeName {$baseType->getName()}


GO
            );


            foreach ($enum as $index => $item) {
                $itemName = $this->goBuilder->codeBuilder->exportableName($path . '_' . $item);
                $comment = null;
                if (isset($enumSchemas[$index])) {
                    $schema = $enumSchemas[$index];
                    if (isset($schema->title)) {
                        $itemName = $this->goBuilder->codeBuilder->exportableName($path . '_' . $schema->title);
                    }
                    if (isset($schema->description)) {
                        $comment = $schema->description;
                    }
                }
                $typeConstBlock->addValue($itemName, $item, $comment);
            }

            $this->goBuilder->getCode()->addSnippet($typeConstBlock);
            $this->goBuilder->getCode()->addSnippet(new MarshalEnum($type, $baseType, $typeConstBlock->getValues(), $this->goBuilder));

            return $type;
        }


        return $baseType;
    }

    private function getGeneratedStruct()
    {
        return $this->goBuilder->getGeneratedStruct($this->schema, $this->path);
    }

    private function makeResultStruct()
    {
        if ($this->resultStruct === null) {
            $this->resultStruct = $this->goBuilder->getClass($this->schema, $this->path);
        }
        return $this->resultStruct;
    }

    /** @var StructDef */
    private $resultStruct;

    /**
     * @return AnyType
     * @throws Exception
     */
    public function build()
    {
        if ($this->schema->{'x-go-type'}) {
            return TypeUtil::fromString($this->schema->{'x-go-type'});
        }

        $this->result = array();

        if (null !== $path = $this->schema->getFromRef()) {
            $this->path = $path;
        }

        $this->processNamedClass();
        $this->processLogicType();
        $this->processArrayType();
        $this->processObjectType();

        if (is_array($this->schema->type)) {
            $or = [];
            foreach ($this->schema->type as $type) {
                $schema = clone $this->schema;
                $schema->type = $type;
                $or [] = $schema;
            }
            $this->processOr($or, 'type');
        } elseif ($this->schema->type) {
            $type = $this->typeSwitch($this->schema->type, $this->schema->minimum, $this->schema->maximum, $this->schema->format);
            if (($type !== null)) { // todo properly process const = null
                $type = $this->processEnum($type);
            }
            $this->result[] = $type;
        }

        $hasConst = $this->schema->const !== null;
        if ($hasConst) {
            $type = $this->processConst();
            if (null !== $type) {
                $this->result[] = $type;
            }
        }


        if ($this->resultStruct !== null) {
            return new Pointer($this->resultStruct->getType());
        }

        if (empty($this->result)) {
            return new GoType('interface{}');
        } else {
            if (1 === count($this->result)) {
                return $this->result[0];
            } else {
                return $this->result[0];
            }
        }


        return new GoType('interface{}');
    }
}