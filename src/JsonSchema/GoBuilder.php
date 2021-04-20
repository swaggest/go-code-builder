<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\GoCodeBuilder;
use Swaggest\GoCodeBuilder\Style\Comment;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Struct\FluentSetter;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Struct\StructType;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\NoOmitEmpty;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Type;
use Swaggest\JsonSchema\Schema;
use Swaggest\JsonSchema\Wrapper;

class GoBuilder
{
    private $code;

    /** @var Options */
    public $options;

    /** @var GeneratedStruct[] */
    private $generatedStructs;

    /** @var \SplObjectStorage contains generated structs by schema object */
    private $generatedStructsBySchema;

    private $untitledIndex = 0;
    /** @var GoCodeBuilder */
    public $codeBuilder;

    /** @var GoBuilderStructHook */
    public $structPreparedHook;

    /** @var GoBuilderStructHook */
    public $structCreatedHook;

    /** @var GoBuilderPathToNameHook */
    public $pathToNameHook;

    /** @var MarshalUnion */
    public $marshalUnion;

    /** @var UnmarshalUnion */
    public $unmarshalUnion;

    /** @var Type[] */
    public $pathTypesDefined = [];

    /** @var string[] */
    public $pathByTypeName = [];

    /** @var string[] */
    private $typeNameByPath = [];

    public function __construct()
    {
        $this->code = new Code();
        $this->options = new Options();
        $this->generatedStructs = [];
        $this->generatedStructsBySchema = new \SplObjectStorage();
        $this->codeBuilder = new GoCodeBuilder();
        $this->pathToNameHook = new StripPrefixPathToNameHook();
    }

    public function getCode()
    {
        return $this->code;
    }

    /**
     * @param Schema $schema
     * @param string $path
     * @param StructDef|null $parentStruct
     * @param bool $isRequired
     * @return AnyType
     * @throws Exception
     * @throws \Swaggest\JsonSchema\Exception
     * @throws \Swaggest\JsonSchema\InvalidValue
     */
    public function getType($schema, $path = '#', StructDef $parentStruct = null, $isRequired = false)
    {
        if (isset($this->generatedStructs[$path])) {
            return $this->generatedStructs[$path]->structDef->getType();
        }

        if (!is_bool($schema) && (null !== $refs = $schema->getFromRefs()) && !in_array(false, $refs)) {
            foreach ($refs as $ref) {
                if (isset($this->generatedStructs[$ref])) {
                    return $this->generatedStructs[$ref]->structDef->getType();
                }
            }
        }

        $s = self::unboolSchema($schema);
        if ($s instanceof Wrapper) {
            $path = $s->getObjectItemClass();
            $typeBuilder = new TypeBuilder($s->exportSchema(), $path, $this, $parentStruct);
            $result = $typeBuilder->build();
            return $result;
        }
        if ($s instanceof \stdClass) {
            $s = Schema::import($s);
        }
        $typeBuilder = new TypeBuilder($s, $path, $this, $parentStruct, $isRequired);
        $result = $typeBuilder->build();
        return $result;
    }

    /**
     * @param Schema $schema
     * @param string $path
     * @return StructDef
     * @throws Exception
     * @throws \Swaggest\JsonSchema\Exception
     * @throws \Swaggest\JsonSchema\InvalidValue
     */
    public function getClass($schema, $path)
    {
        $generatedStruct = $this->getGeneratedStruct($schema, $path);
        return $generatedStruct->structDef;
    }

    /**
     * @param Schema $schema
     * @param string $path
     * @return mixed|GeneratedStruct
     * @throws Exception
     * @throws \Swaggest\JsonSchema\Exception
     * @throws \Swaggest\JsonSchema\InvalidValue
     */
    public function getGeneratedStruct($schema, $path)
    {
        if (isset($this->generatedStructs[$path])) {
            return $this->generatedStructs[$path];
        }

        if ($this->generatedStructsBySchema->contains($schema)) {
            return $this->generatedStructsBySchema[$schema];
        }

        return $this->makeStruct($schema, $path);
    }

    private $namesGenerated = [];

    /**
     * @param Schema $schema
     * @param string $path
     * @return GeneratedStruct
     * @throws Exception
     * @throws \Swaggest\JsonSchema\Exception
     * @throws \Swaggest\JsonSchema\InvalidValue
     */
    private function makeStruct(Schema $schema, $path)
    {
        if (empty($path)) {
            throw new Exception('Empty path');
        }
        $generatedStruct = new GeneratedStruct();
        $this->generatedStructs[$path] = $generatedStruct;
        $this->generatedStructsBySchema->attach($schema, $generatedStruct);
        $generatedStruct->schema = $schema;

        $pathToName = $this->pathToName($path);
        if ($path === '#' && empty($schema->title)) {
            $pathToName = 'Untitled' . ++$this->untitledIndex;
        }

        $structName = $this->typeName($schema, $pathToName);

        if (isset($this->namesGenerated[$structName]) && $schema->getMeta(TypeBuilder::CONDITIONAL_META)) {
            $structName = $structName . 'Conditional';
        }

        $structPreferredName = $structName;
        $i = 2;
        while (isset($this->namesGenerated[$structName])) {
            $structName = $structPreferredName . $i;
            $i++;
        }

        $this->namesGenerated[$structName] = true;


        $structDef = new StructDef($structName);


        if ($this->structCreatedHook !== null) {
            $this->structCreatedHook->process($structDef, $path, $schema);
        }

        $comment = $structDef->getName() . ' structure is generated from "' . $path . '".';
        if ($schema->title) {
            $comment .= "\n\n" . rtrim($schema->title, '.') . '.';
        }
        if ($schema->description) {
            $comment .= "\n\n" . rtrim($schema->description, '.') . '.';
        }
        $structDef->setComment($comment);
        $marshalJson = new MarshalJson($this, $structDef);

        $generatedStruct->structDef = $structDef;
        $generatedStruct->path = $path;
        $generatedStruct->marshalJson = $marshalJson;

        // Properties are only processed if type has object semantic.
        // This removes properties from multi-type and non-object (e.g. boolean) structures.
        $processProperties = false;
        if ($schema->type === null
            || $schema->type === Schema::OBJECT
            || (is_array($schema->type)
                && in_array(Schema::OBJECT, $schema->type)
                && in_array(Schema::NULL, $schema->type)
                && count($schema->type) === 2)
        ) {
            $processProperties = true;
        }

        if ($processProperties && $schema->properties !== null) {
            // Iterating over a copy (toArray) to not conflict with any other iterations in nested processings.
            foreach ($schema->properties->toArray() as $name => $property) {
                $property = self::unboolSchema($property);

                if ($property->{TypeBuilder::X_GENERATE} === false ||
                    ($this->options->requireXGenerate && empty($property->{TypeBuilder::X_GENERATE}))) {
                    continue;
                }

                $fieldName = $this->codeBuilder->exportableName($name);

                if ($this->options->trimParentFromPropertyNames) {
                    $stripped = substr($fieldName, strlen($structDef->getName()));
                    if (
                        strpos($fieldName, $structDef->getName()) === 0
                        && $fieldName !== $structDef->getName()
                        && $stripped === $this->codeBuilder->exportableName($stripped)
                    ) {
                        $fieldName = $stripped;
                    }
                }

                $isRequired = (null !== $schema->required) && in_array($name, $schema->required);
                $goPropertyType = $this->getType($property, $path . '->' . $name, $structDef, $isRequired);
                if ($goPropertyType instanceof StructType) {
                    if (!$isRequired
                        || $this->options->ignoreRequired
                        || $this->isNullable($property)
                    ) {
                        $goPropertyType = new Pointer($goPropertyType);
                    }
                }
                $goProperty = new StructProperty(
                    $fieldName,
                    $goPropertyType
                );

                if ($property instanceof Wrapper) {
                    $property = $property->exportSchema();
                }

                if (!$property instanceof Schema) {
                    $property = new Schema();
                }

                $isOmitEmpty = false;
                // Nullable properties need `null` explicitly available in json payload.
                if (
                    false === $property->{TypeBuilder::X_OMIT_EMPTY} ||
                    (
                        $goPropertyType instanceof NoOmitEmpty &&
                        $goPropertyType->isNoOmitEmpty() &&
                        true !== $property->{TypeBuilder::X_OMIT_EMPTY}
                    ) ||
                    ($isRequired && !$this->options->ignoreRequired && true !== $property->{TypeBuilder::X_OMIT_EMPTY})
                ) {
                    $goProperty->getTags()->setTag('json', $name);
                } else {
                    $isOmitEmpty = true;
                    $goProperty->getTags()->setTag('json', $name . ',omitempty');
                }

                if (!empty($this->options->nameTags)) {
                    foreach ($this->options->nameTags as $nameTag) {
                        $goProperty->getTags()->setTag($nameTag, $name);
                    }
                }

                $comment = '';
                if ($property->title) {
                    $comment .= Comment::sentence($property->title) . "\n";
                }
                if ($property->description) {
                    $comment .= Comment::sentence($property->description) . "\n";
                }
                if ($property->pattern) {
                    $comment .= "Value must match pattern: `{$property->pattern}`.\n";
                }
                if ($property->format) {
                    $comment .= "Format: {$property->format}.\n";
                }
                if ($isRequired) {
                    $comment .= "Required.\n";
                }
                $comment = trim($comment);

                if ($comment !== '') {
                    $goProperty->setComment($comment);
                }
                if ($this->options->hideConstProperties) {
                    $enum = [];

                    if ($property->enum !== null) {
                        $enum = $property->enum;
                    }

                    if ($property->const !== null) {
                        $enum = [$property->const];
                    }

                    if (count($enum) === 1) {
                        $marshalJson->constValues[$name] = $enum[0];
                        continue;
                    }
                }

                if ($goPropertyType->getTypeString() === 'interface{}') {
                    if ($this->options->distinctNull && $isOmitEmpty) {
                        $goPropertyType = new Pointer($goPropertyType);
                        $goProperty->setType($goPropertyType);
                        $marshalJson->distinctNullNames[$goProperty->getName()] = $name;
                    }
                }

                $structDef->addProperty($goProperty);

                if ($this->options->fluentSetters) {
                    FluentSetter::addToStruct($structDef, $goProperty);
                }

                $marshalJson->addNamedProperty($name);
            }
        }

        if ($processProperties && !empty($schema->required) && !$this->options->ignoreRequired) {
            $marshalJson->required = $schema->required;
        }

        if (!empty($schema->not)) {
            $not = $schema->not;
            if ($not instanceof Schema) {
                $marshalJson->not = $this->getType($not, $path . '->not', $structDef);
            }
        }

        $structDef->getCode()->addSnippet($marshalJson);

        if ($this->structPreparedHook !== null) {
            $this->structPreparedHook->process($structDef, $path, $schema);
        }

        return $generatedStruct;
    }

    /**
     * @return GeneratedStruct[]
     */
    public function getGeneratedStructs()
    {
        return $this->generatedStructs;
    }

    public function pathToName($path)
    {
        // Removing type marker, e.g. #[object]/properties => #/properties
        $path = preg_replace('/\[\w+\]/', '', $path);

        if (null !== $this->pathToNameHook) {
            return $this->pathToNameHook->pathToName($path);
        }
        return $path;
    }

    /**
     * Resolves boolean schema into Schema instance
     *
     * @param mixed $schema
     * @return mixed|Schema
     */
    private static function unboolSchema($schema)
    {
        static $trueSchema;
        static $falseSchema;

        if (null === $trueSchema) {
            $trueSchema = new Schema();
            $falseSchema = new Schema();
            $falseSchema->not = $trueSchema;
        }

        if ($schema === true) {
            return $trueSchema;
        } elseif ($schema === false) {
            return $falseSchema;
        } else {
            return $schema;
        }
    }

    public function isNullable(Schema $schema)
    {
        if ($schema->type === Schema::NULL) {
            return true;
        }

        if (is_array($schema->type) && in_array(Schema::NULL, $schema->type)) {
            return true;
        }

        if ($this->options->enableXNullable) {
            if ($schema->{TypeBuilder::X_NULLABLE} === true || $schema->{TypeBuilder::NULLABLE} === true) {
                return true;
            }
        }

        return false;
    }

    /**
     * @param Schema $schema
     * @param $path
     * @param StructDef|null $parentStruct
     */
    public function typeName($schema, $path, $parentStruct = null)
    {
        if (isset($this->typeNameByPath[$path])) {
            return $this->typeNameByPath[$path];
        }

        if (!empty($schema->title)) {
            $typeName = $this->codeBuilder->exportableName($schema->title, true);
        }

        if (empty($typeName) || (isset($this->pathByTypeName[$typeName]) && $this->pathByTypeName[$typeName] !== $path)) {
            $typeName = $this->codeBuilder->exportableName($path);
        }

        if ($parentStruct && false !== $pos = strrpos($path, '->')) {
            $propertyName = substr($path, $pos);
            $typeName = $this->codeBuilder->exportableName($parentStruct->getName() . $propertyName);
        }

        if (isset($this->options->renames[$typeName])) {
            $typeName = $this->options->renames[$typeName];
        }

        $tn = $typeName;
        $i = 2;

        while (isset($this->pathByTypeName[$typeName]) && $this->pathByTypeName[$typeName] !== $path) {
            $typeName = $tn . 'Type' + $i;
            $i++;
        }

        $this->pathByTypeName[$typeName] = $path;
        $this->typeNameByPath[$path] = $typeName;

        return $typeName;
    }
}