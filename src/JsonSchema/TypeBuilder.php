<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\Templates\Constant\TypeConstBlock;
use Swaggest\GoCodeBuilder\Templates\Struct\FluentSetter;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Struct\StructType;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\GoCodeBuilder\Templates\Type\NamedType;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;
use Swaggest\JsonSchema\Constraint\Format;
use Swaggest\JsonSchema\Constraint\Type;
use Swaggest\JsonSchema\JsonSchema;
use Swaggest\JsonSchema\Schema;
use Swaggest\GoCodeBuilder\Templates\Type\Type as GoType;
use Swaggest\JsonSchema\SchemaContract;
use Swaggest\JsonSchema\SchemaExporter;
use Swaggest\JsonSchemaMaker\SchemaMaker;

class TypeBuilder
{
    const X_GO_TYPE = 'x-go-type';
    const X_OMIT_EMPTY = 'x-omitempty';
    const X_NULLABLE = 'x-nullable';
    const X_GENERATE = 'x-generate';
    const NULLABLE = 'nullable';
    const EXAMPLES = 'examples';
    const EXAMPLE = 'example';

    const CONDITIONAL_META = 'conditional';

    const NAME_ANY = 'anything';

    /** @var Schema */
    private $schema;
    /** @var string */
    private $path;
    /** @var GoBuilder */
    private $goBuilder;
    /** @var AnyType[] */
    private $result;

    /** @var null|string|string[] JSON Schema type */
    private $type;

    private $nullable = false;

    /** @var StructDef|null parent structure if available */
    private $parentStruct;

    /** @var bool indicates type of required property */
    private $isRequired;

    /**
     * TypeBuilder constructor.
     * @param Schema $schema
     * @param string $path
     * @param GoBuilder $goBuilder
     * @param StructDef|null $parentStruct
     * @param bool $isRequired
     */
    public function __construct(Schema $schema, $path, GoBuilder $goBuilder, StructDef $parentStruct = null, $isRequired = false)
    {
        $this->schema = $schema;
        $this->path = $path;
        $this->goBuilder = $goBuilder;
        $this->parentStruct = $parentStruct;
        $this->isRequired = $isRequired;
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
        if ($type instanceof NamedType) {
            if ($type->getName() === 'interface{}') {
                return self::NAME_ANY;
            }
            return $type->getName();
        }
        if ($type->getTypeString() === '[]interface{}') {
            return 'SliceOf_' . self::NAME_ANY;
        }
        if ($type instanceof Slice) {
            $itemName = $this->makeName($type->getType());
            if ($itemName !== null) {
                return 'SliceOf_' . $itemName . '_Values';
            }
        }

        return null;
    }

    /**
     * @param Schema $schema
     * @return bool
     */
    private function isSimpleObject($schema)
    {
        if ($schema->type !== null && $schema->type !== Schema::OBJECT) {
            return false;
        }

        if (empty($schema->properties)) {
            return false;
        }

        if (isset($schema->additionalProperties) && $schema->additionalProperties !== true) {
            return false;
        }

        $names = Schema::names();
        // Not a simple object if has additional logical constraints.
        foreach ([
                     $names->anyOf,
                     $names->oneOf,
                     $names->if,
                     $names->not,
                     $names->patternProperties,
                     $names->else,
                 ] as $name) {
            if (isset($schema->$name)) {
                return false;
            }
        }

        return true;
    }

    /**
     * @param Schema[]|SchemaContract[] $allOfs
     */
    private function processAllOfEmbeddableObject($allOfs)
    {
        $allProperties = array();
        foreach ($allOfs as $allOf) {
            if (!$allOf instanceof Schema) {
                return false;
            }

            if (!$this->isSimpleObject($allOf) || null === $allOf->properties) {
                return false;
            }

            foreach ($allOf->properties as $propertyName => $property) {
                if (isset($allProperties[$propertyName])) {
                    return false;
                }

                $allProperties[$propertyName] = true;
            }
        }

        $result = $this->makeResultStruct();
        foreach ($allOfs as $i => $allOf) {
            if (!$allOf instanceof Schema) {
                return false;
            }

            $type = $this->goBuilder->getType($allOf, $this->path . '/allOf/' . $i, $result);
            $result->addProperty(new StructProperty(null, $type));
        }

        return true;
    }

    /**
     * @param Schema[]|SchemaContract[] $orSchemas
     * @param string|mixed $kind
     * @throws Exception
     * @throws \Swaggest\JsonSchema\Exception
     * @throws \Swaggest\JsonSchema\InvalidValue
     */
    private function processOr($orSchemas, $kind)
    {
        $types = [];
        foreach ($orSchemas as $i => $item) {
            if (!$item instanceof Schema && $item instanceof SchemaExporter) {
                $item = $item->exportSchema();
            }
            if ($item instanceof Schema) {
                $path = $this->path . '/' . $kind . '/' . $i;
                if ($kind === Schema::names()->type && is_string($item->type)) {
                    $path = $this->path . '[' . $item->type . ']';
                }
                $itemType = Pointer::tryDereferenceOnce($this->goBuilder->getType($item, $path));
                if ($itemType->getTypeString() === 'interface{}') {
                    continue;
                }
                $types [] = $itemType;
            }
        }
        if (count($types) == 1) {
            $this->result[] = $types[0];
            return;
        }

        $props = [];
        foreach ($orSchemas as $i => $item) {
            if ($item instanceof Schema) {
                $name = $this->goBuilder->codeBuilder->exportableName($kind . '/' . $i);
                $path = $this->path . '/' . $kind . '/' . $i;

                if ($kind === Schema::names()->type && is_string($item->type) && is_string($kind)) {
                    $path = $this->path . '[' . $item->type . ']';
                    $name = $this->goBuilder->codeBuilder->exportableName($kind . '/' . $item->type);
                }

                $betterName = null;
                if (($refs = $item->getFromRefs()) && (($kind !== Schema::names()->type) || $item->type === Schema::OBJECT)) {
                    $betterName = $this->goBuilder->pathToName($refs[0]);
                    $betterName = $this->goBuilder->codeBuilder->exportableName($betterName, true);
                    if (empty($betterName) || isset($props[$betterName])) {
                        $betterName = null;
                    }
                }

                $itemType = Pointer::tryDereferenceOnce($this->goBuilder->getType($item, $path));
                if (empty($betterName) &&
                    ($kind !== Schema::names()->type)) {
                    $betterName = $this->makeName($itemType);
                    $betterName = $this->goBuilder->codeBuilder->exportableName($betterName, true);
                }

                if (!empty($betterName) && !isset($props[$betterName])) {
                    $name = $betterName;
                }

                $this->schema->addMeta(true, self::CONDITIONAL_META);
                $resultStruct = $this->makeResultStruct();

                if ($this->goBuilder->options->trimParentFromPropertyNames) {
                    $stripped = substr($name, strlen($resultStruct->getName()));
                    if (
                        strpos($name, $resultStruct->getName()) === 0
                        && $name !== $resultStruct->getName()
                        && $stripped === $this->goBuilder->codeBuilder->exportableName($stripped)
                    ) {
                        $name = $stripped;
                    }
                }

                if ((!$itemType instanceof Map) && (!$itemType instanceof Slice)) {
                    if ($itemType->getTypeString() === 'interface{}') {
                        continue;
                    }
                    $itemType = new Pointer($itemType);
                }

                $props[$name] = true;
                $structProperty = new StructProperty($name, $itemType);
//                $structProperty->setComment($path);
                $structProperty->getTags()->setTag('json', '-');
                $resultStruct->addProperty($structProperty);

                if ($this->goBuilder->options->fluentSetters) {
                    FluentSetter::addToStruct($resultStruct, $structProperty);
                }

                $generatedStruct = $this->getGeneratedStruct();
                $generatedStruct->marshalJson->addSomeOf($kind, $name);
            }
        }
    }

    private function processLogicType()
    {
        if (in_array($this->schema->type, [JsonSchema::STRING, JsonSchema::INTEGER, JsonSchema::NUMBER, JsonSchema::BOOLEAN], true)) {
            // Skip logical keywords if scalar type provided
            return;
        }
        if ($this->schema->allOf !== null) {
            if (!$this->processAllOfEmbeddableObject($this->schema->allOf)) {
                $this->processOr($this->schema->allOf, Schema::names()->allOf);
            }
        } elseif ($this->schema->anyOf !== null) {
            $this->processOr($this->schema->anyOf, Schema::names()->anyOf);
        } elseif ($this->schema->oneOf !== null) {
            $this->processOr($this->schema->oneOf, Schema::names()->oneOf);
        }
    }

    private function processArrayType()
    {
        $schema = $this->schema;

        $pathItems = Schema::names()->items;
        if ($schema->items instanceof Schema) {
            $items = array();
            $additionalItems = $schema->items;
        } elseif ($schema->items === null) { // items defaults to empty schema so everything is valid
            $items = array();
            $additionalItems = true;
        } else { // listed items
            $items = $schema->items;
            $additionalItems = $schema->additionalItems;
            $pathItems = Schema::names()->additionalItems;
        }

        $itemsLen = is_array($items) ? count($items) : 0;
        $index = 0;
        if ($index < $itemsLen) {
        } else {
            if ($additionalItems instanceof Schema) {
                $sliceType = new Slice(Pointer::tryDereferenceOnce(
                    $this->goBuilder->getType($additionalItems, $this->path . '->' . $pathItems))
                );
                if ($this->nullable) {
                    $sliceType = new Pointer($sliceType);
                }
                $this->result[] = $sliceType;
            }
        }
    }

    private function processObjectType()
    {
        $canBeObject = false;
        if ($this->type === null ||
            $this->type === Schema::OBJECT
        ) {
            $canBeObject = true;
        }

        if (!$canBeObject) {
            return;
        }

        if ($this->schema->patternProperties !== null) {
            foreach ($this->schema->patternProperties as $pattern => $schema) {
                if (!$schema instanceof Schema) {
                    continue;
                }
                $itemType = Pointer::tryDereferenceOnce($this->goBuilder->getType($schema, $this->path . '->' . $pattern));
                $name = $this->goBuilder->codeBuilder->exportableName('patternProperties_' . $pattern);
                if (null !== $betterName = $this->makeName($itemType)) {
                    if ($betterName === self::NAME_ANY) {
                        $name = $this->goBuilder->codeBuilder->exportableName('MapOf_' . $betterName);
                    } else {
                        $name = $this->goBuilder->codeBuilder->exportableName('MapOf_' . $betterName . '_Values');
                    }
                }
                $structProperty = new StructProperty(
                    $name,
                    new Map(
                        new GoType("string"),
                        $itemType
                    )
                );
                $structProperty->getTags()->setTag('json', '-');
                $structProperty->setComment('Key must match pattern: `' . $pattern . '`.');

                $resultStruct = $this->makeResultStruct();
                $resultStruct->addProperty($structProperty);

                if ($this->goBuilder->options->fluentSetters) {
                    FluentSetter::addToStruct($resultStruct, $structProperty);
                }

                $this->getGeneratedStruct()->marshalJson->addPatternProperty($pattern, $structProperty);
            }
        }

        /** @var null|boolean|Schema $additionalProperties */
        $additionalProperties = $this->schema->additionalProperties;
        if ($this->goBuilder->options->defaultAdditionalProperties && $additionalProperties === null) {
            $additionalProperties = true;
        }
        if (
            $additionalProperties instanceof Schema ||
            $additionalProperties === true
        ) {

            if ($additionalProperties instanceof Schema) {
                $goType = Pointer::tryDereferenceOnce(
                    $this->goBuilder->getType(
                        $additionalProperties,
                        $this->path . '->' . Schema::names()->additionalProperties)
                );
            } else {
                $goType = TypeUtil::fromString('interface{}');
            }

            if (
                $this->schema->properties !== null ||
                $this->schema->patternProperties !== null
            ) {
                $resultStruct = $this->makeResultStruct();
                $existingProperties = [];
                if ($this->schema->properties !== null) {
                    foreach ($this->schema->properties->toArray() as $propertyName => $property) {
                        $existingProperties[$this->goBuilder->codeBuilder->exportableName($propertyName)] = true;
                    }
                }

                $possibleNames = ['additionalProperties', 'extraProperties', 'unmatchedProperties'];
                foreach ($possibleNames as $name) {
                    $propName = $this->goBuilder->codeBuilder->exportableName($name);
                    if (!isset($existingProperties[$propName])) {
                        break;
                    }
                }
                $i = 1;
                while (isset($existingProperties[$propName])) {
                    $i++;
                    $propName = $this->goBuilder->codeBuilder->exportableName('additionalProperties' . $i);
                }

                if (!$this->getGeneratedStruct()->marshalJson->isAdditionalPropertiesEnabled()) {
                    $structProperty = new StructProperty(
                        $propName,
                        new Map(new GoType("string"), $goType)
                    );
                    $structProperty->setComment('All unmatched properties.');
                    $structProperty->getTags()->setTag('json', '-');
                    $resultStruct->addProperty($structProperty);

                    if ($this->goBuilder->options->fluentSetters) {
                        FluentSetter::addToStruct($resultStruct, $structProperty);
                    }

                    $this->getGeneratedStruct()->marshalJson->enableAdditionalProperties($structProperty);
                }
            } elseif ($additionalProperties instanceof Schema) {
                $mapType = new Map(new GoType('string'), $goType);
                if ($this->nullable) {
                    $mapType = new Pointer($mapType);
                }

                $this->result[] = $mapType;
            }
        }

        if ($additionalProperties === false) {
            $this->getGeneratedStruct()->marshalJson->forbidAdditionalProperties();
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
                    return TypeUtil::fromString('*time.Time');
                }
                return new GoType('string');

            case Type::OBJECT:
                return TypeUtil::fromString('map[string]interface{}');

            case Type::ARR:
                return TypeUtil::fromString('[]interface{}');

            case Type::NULL:
                return new GoType('interface{}');

            default:
                return null;
        }
    }

    private function processNamedClass()
    {
        $canBeObject = false;
        if ($this->type === null ||
            $this->type === Schema::OBJECT) {
            $canBeObject = true;
        }

        if (!$canBeObject) {
            return;
        }

        if ($this->schema->properties !== null) {
            $this->result[] = $this->makeResultStruct()->getType();
        }
    }

    private function processConst()
    {
        if ($this->schema->const !== null) { // todo properly process null const
            $path = $this->goBuilder->pathToName($this->path);
            $typeName = $this->goBuilder->typeName($this->schema, $path, $this->parentStruct);
            if ($this->parentStruct && false !== $pos = strrpos($path, '->')) {
                $propertyName = substr($path, $pos);
                $typeName = $this->goBuilder->codeBuilder->exportableName($this->parentStruct->getName() . $propertyName);
            }

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
                $this->goBuilder->codeBuilder->exportableName($typeName . '_' . $this->schema->const),
                $this->schema->const
            );

            $this->goBuilder->getCode()->addSnippet($typeConstBlock);
            $this->goBuilder->getCode()->addSnippet(new MarshalEnum($type, $baseType, $typeConstBlock->getValues(), $this->goBuilder));
            return $type;
        }
        return null;
    }


    private function typeFromValue($val)
    {
        if (is_string($val)) {
            return "string";
        } elseif (is_int($val)) {
            return "integer";
        } elseif (is_bool($val)) {
            return "boolean";
        } elseif (is_numeric($val)) {
            return "number";
        }
        return null;
    }

    private function inferTypeFromEnumConst()
    {
        // Return early if the type is already defined.
        if ($this->type !== null) {
            return;
        }

        // Null type means no override.
        $type = null;

        // Check const.
        if (property_exists($this->schema, Schema::CONST_PROPERTY)) {
            $type = $this->typeFromValue($this->schema->const);
            if ($type !== null) {
                $this->type = $type;
                return;
            }
        }

        // Iterate enum to find if its values share same scalar type.
        if ($this->schema->enum !== null) {
            foreach ($this->schema->enum as $enumItem) {
                $itemType = $this->typeFromValue($enumItem);
                if ($type !== null && $type !== $itemType) {
                    // Break on types mismatch between items.
                    return;
                }
                $type = $itemType;
            }
        }

        // Apply the type found in enum iteration.
        if ($type !== null) {
            $this->type = $type;
        }
    }

    private function processEnum(NamedType $baseType)
    {
        $oneOfEnum = true;
        $enum = null;
        /** @var Schema[] $enumSchemas */
        $enumSchemas = null;
        if ($this->schema->oneOf !== null) {
            foreach ($this->schema->oneOf as $schema) {
                if ($schema instanceof Schema) {
                    if ($schema->const !== null) {
                        $enum[] = $schema->const;
                        $enumSchemas[] = $schema;
                    } else {
                        $oneOfEnum = false;
                    }
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

            if (isset($this->goBuilder->pathTypesDefined[$this->path])) {
                return $this->goBuilder->pathTypesDefined[$this->path];
            }
            $path = $this->goBuilder->pathToName($this->path);

            $typeName = $this->goBuilder->typeName($this->schema, $path, $this->parentStruct);
            $type = new GoType($typeName);
            $this->goBuilder->pathTypesDefined[$this->path] = $type;

            $typeConstBlock = new TypeConstBlock($type);

            $this->goBuilder->getCode()->addSnippet(<<<GO
// $typeName is an enum type.
type $typeName {$baseType->getName()}


GO
            );

            foreach ($enum as $index => $item) {
                $itemName = $this->goBuilder->codeBuilder->exportableName($typeName . '_' . $item);
                $comment = null;
                if (isset($enumSchemas[$index])) {
                    $schema = $enumSchemas[$index];
                    if (isset($schema->title)) {
                        $itemName = $this->goBuilder->codeBuilder->exportableName($typeName . '_' . $schema->title);
                    }

                    if (isset($schema->description)) {
                        $comment = $schema->description;
                    }
                }

                if (isset($this->goBuilder->options->renames[$itemName])) {
                    $itemName = $this->goBuilder->options->renames[$itemName];
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
        if (!$this->goBuilder->options->ignoreXGoType && $this->schema->{self::X_GO_TYPE}) {
            // go-swagger formatted type.
            /*
             * {
                  "import": {
                    "package": "github.com/my-org/my-service/internal/domain/orders"
                  },
                  "type": "Order"
                }
             */
            $xGoType = $this->schema->{self::X_GO_TYPE};
            if ($xGoType instanceof \stdClass) {
                $typeString = '';
                if (isset($xGoType->import) && isset($xGoType->import->package)) {
                    $typeString .= $xGoType->import->package . '.';
                }
                if (isset($xGoType->type)) {
                    $typeString .= $xGoType->type;
                }
                return TypeUtil::fromString($typeString);
            } elseif (is_string($xGoType)) {
                return TypeUtil::fromString($xGoType);
            }
        }

        if (
            $this->goBuilder->options->inheritSchemaFromExamples
            && empty($this->schema->type) && empty($this->schema->properties)
        ) {
            $schemaMaker = new SchemaMaker($this->schema);

            if (
                isset($this->schema->{self::EXAMPLES})
                && is_array($this->schema->{self::EXAMPLES})
            ) {
                foreach ($this->schema->{self::EXAMPLES} as $example) {
                    $schemaMaker->addInstanceValue($example);
                }
            }

            if (isset($this->schema->{self::EXAMPLE})) {
                $schemaMaker->addInstanceValue($this->schema->{self::EXAMPLE});
            }
        }


        $this->result = array();

        // TODO reduce structure complexity in Swaggest\JsonSchema.
        /** @var string[]|boolean[] $path */
        $path = $this->schema->getFromRefs();
        if (!empty($path)) {
            if (false === $path[count($path) - 1]) {
                $path = null;
            } else {
                $path = $path[0];
            }
        }
        if (!empty($path) && is_string($path)) {
            $this->path = $path;
        }

        $this->type = $this->schema->type;

        if ($this->goBuilder->options->enableXNullable) {
            if ($this->schema->{self::X_NULLABLE} === true || $this->schema->{self::NULLABLE} === true) {
                if (is_array($this->type) && !in_array(Schema::NULL, $this->type)) {
                    $this->type[] = Schema::NULL;
                } elseif (is_string($this->type) && $this->type !== Schema::NULL) {
                    $this->type = [$this->type, Schema::NULL];
                }
            }
        }

        if (is_array($this->type)) {
            foreach ($this->type as $k => $item) {
                if ($item === Schema::NULL) {
                    $this->nullable = true;
                    unset($this->type[$k]);
                    $this->type = array_values($this->type);
                    break;
                }
            }

            if (empty($this->type)) {
                $this->type = null;
            } elseif (count($this->type) === 1) {
                $this->type = $this->type[0];
            }
        }

        $this->inferTypeFromEnumConst();

        $this->processNamedClass();
        $this->processLogicType();
        $this->processArrayType();
        $this->processObjectType();

        if (is_array($this->type)) {
            $or = [];
            foreach ($this->type as $type) {
                $schema = clone $this->schema;
                $schema->setFromRef(false);
                $schema->type = $type;
                $or [] = $schema;
            }
            $this->processOr($or, Schema::names()->type);
        } elseif ($this->type) {
            $type = $this->typeSwitch($this->type, $this->schema->minimum, $this->schema->maximum, $this->schema->format);
            if ($type instanceof NamedType) { // todo properly process const = null
                $type = $this->processEnum($type);
            }
            if (
                $this->nullable ||
                ($this->goBuilder->options->withZeroValues && false !== $this->schema->{self::X_OMIT_EMPTY})
            ) {
                if (
                    (!$type instanceof Pointer) &&
                    (!$type instanceof Map) &&
                    (!$type instanceof Slice) &&
                    (!$this->isRequired || $this->goBuilder->options->ignoreRequired || $this->nullable) &&
                    ($type->getTypeString() !== 'interface{}')
                ) {
                    $type = new Pointer($type);
                    if ($this->nullable && !$this->goBuilder->options->ignoreNullable) {
                        $type->setNoOmitEmpty(true);
                    }
                }
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
            return $this->resultStruct->getType();
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
    }
}