<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\GoCodeBuilder;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\NoOmitEmpty;
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
     * @return AnyType
     * @throws Exception
     * @throws \Swaggest\JsonSchema\Exception
     * @throws \Swaggest\JsonSchema\InvalidValue
     */
    public function getType($schema, $path = '#')
    {
        $s = self::unboolSchema($schema);
        if ($s instanceof Wrapper) {
            $path = $s->getObjectItemClass();
            $typeBuilder = new TypeBuilder($s->exportSchema(), $path, $this);
            $result = $typeBuilder->build();
            return $result;
        }
        if ($s instanceof \stdClass) {
            $s = Schema::import($s);
        }
        $typeBuilder = new TypeBuilder($s, $path, $this);
        $result = $typeBuilder->build();
        return $result;
    }

    /**
     * @param Schema $schema
     * @param string $path
     * @return StructDef
     * @throws Exception
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

        $structName = '';
        if ($schema->title) {
            $structName = $this->codeBuilder->exportableName($schema->title);
        }

        if (empty($structName)) {
            if ($path === '#') {
                $structName = 'Untitled' . ++$this->untitledIndex;
            } else {
                $structName = $this->codeBuilder->exportableName($this->pathToName($path));
            }
        }
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
        if ($schema->type === null || $schema->type === Schema::OBJECT) {
            $processProperties = true;
        }

        if ($processProperties && $schema->properties !== null) {
            // Iterating over a copy (toArray) to not conflict with any other iterations in nested processings.
            foreach ($schema->properties->toArray() as $name => $property) {
                $fieldName = $this->codeBuilder->exportableName($name);

                if ($this->options->trimParentFromPropertyNames) {
                    if (strpos($fieldName, $structDef->getName()) === 0 && $fieldName !== $structDef->getName()) {
                        $fieldName = substr($fieldName, strlen($structDef->getName()));
                    }
                }

                $goPropertyType = $this->getType($property, $path . '->' . $name);
                $goProperty = new StructProperty(
                    $fieldName,
                    $goPropertyType
                );
                // Nullable properties need `null` explicitly available in json payload.
                if (!$goPropertyType instanceof NoOmitEmpty || !$goPropertyType->isNoOmitEmpty()) {
                    $goProperty->getTags()->setTag('json', $name . ',omitempty');
                } else {
                    $goProperty->getTags()->setTag('json', $name);
                }
                $comment = '';
                $property = self::unboolSchema($property);

                if ($property instanceof Wrapper) {
                    $property = $property->exportSchema();
                }

                if (!$property instanceof Schema) {
                    $property = new Schema();
                }

                if ($property->title) {
                    $comment = $property->title . "\n";
                }
                if ($property->description) {
                    $comment .= $property->description;
                }
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

                $structDef->addProperty($goProperty);
                $marshalJson->addNamedProperty($name);
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
        $path = preg_replace('/\[.+\]/', '', $path);

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
}