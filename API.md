## Table of contents

- [\Swaggest\GoCodeBuilder\GoCodeBuilder](#class-swaggestgocodebuildergocodebuilder)
- [\Swaggest\GoCodeBuilder\Import](#class-swaggestgocodebuilderimport)
- [\Swaggest\GoCodeBuilder\Exception](#class-swaggestgocodebuilderexception)
- [\Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback](#class-swaggestgocodebuilderjsonschemastructhookcallback)
- [\Swaggest\GoCodeBuilder\JsonSchema\GoBuilderPathToNameHook (interface)](#interface-swaggestgocodebuilderjsonschemagobuilderpathtonamehook)
- [\Swaggest\GoCodeBuilder\JsonSchema\MarshalEnum](#class-swaggestgocodebuilderjsonschemamarshalenum)
- [\Swaggest\GoCodeBuilder\JsonSchema\GeneratedStruct](#class-swaggestgocodebuilderjsonschemageneratedstruct)
- [\Swaggest\GoCodeBuilder\JsonSchema\Options](#class-swaggestgocodebuilderjsonschemaoptions)
- [\Swaggest\GoCodeBuilder\JsonSchema\UnmarshalUnion](#class-swaggestgocodebuilderjsonschemaunmarshalunion)
- [\Swaggest\GoCodeBuilder\JsonSchema\TypeBuilder](#class-swaggestgocodebuilderjsonschematypebuilder)
- [\Swaggest\GoCodeBuilder\JsonSchema\StripPrefixPathToNameHook](#class-swaggestgocodebuilderjsonschemastripprefixpathtonamehook)
- [\Swaggest\GoCodeBuilder\JsonSchema\GoBuilderStructHook (interface)](#interface-swaggestgocodebuilderjsonschemagobuilderstructhook)
- [\Swaggest\GoCodeBuilder\JsonSchema\MarshalingTestFunc](#class-swaggestgocodebuilderjsonschemamarshalingtestfunc)
- [\Swaggest\GoCodeBuilder\JsonSchema\Exception](#class-swaggestgocodebuilderjsonschemaexception)
- [\Swaggest\GoCodeBuilder\JsonSchema\GoBuilder](#class-swaggestgocodebuilderjsonschemagobuilder)
- [\Swaggest\GoCodeBuilder\JsonSchema\MarshalJson](#class-swaggestgocodebuilderjsonschemamarshaljson)
- [\Swaggest\GoCodeBuilder\JsonSchema\MarshalUnion](#class-swaggestgocodebuilderjsonschemamarshalunion)
- [\Swaggest\GoCodeBuilder\Style\Comment](#class-swaggestgocodebuilderstylecomment)
- [\Swaggest\GoCodeBuilder\Style\Initialisms](#class-swaggestgocodebuilderstyleinitialisms)
- [\Swaggest\GoCodeBuilder\Templates\Imports](#class-swaggestgocodebuildertemplatesimports)
- [\Swaggest\GoCodeBuilder\Templates\GoTemplate (abstract)](#class-swaggestgocodebuildertemplatesgotemplate-abstract)
- [\Swaggest\GoCodeBuilder\Templates\Code](#class-swaggestgocodebuildertemplatescode)
- [\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)
- [\Swaggest\GoCodeBuilder\Templates\Constant\TypeConstBlock](#class-swaggestgocodebuildertemplatesconstanttypeconstblock)
- [\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)
- [\Swaggest\GoCodeBuilder\Templates\Func\FuncIface](#class-swaggestgocodebuildertemplatesfuncfunciface)
- [\Swaggest\GoCodeBuilder\Templates\Func\Arguments](#class-swaggestgocodebuildertemplatesfuncarguments)
- [\Swaggest\GoCodeBuilder\Templates\Func\Result](#class-swaggestgocodebuildertemplatesfuncresult)
- [\Swaggest\GoCodeBuilder\Templates\Func\Argument](#class-swaggestgocodebuildertemplatesfuncargument)
- [\Swaggest\GoCodeBuilder\Templates\Iface\IfaceDef](#class-swaggestgocodebuildertemplatesifaceifacedef)
- [\Swaggest\GoCodeBuilder\Templates\Mapping\Mapping](#class-swaggestgocodebuildertemplatesmappingmapping)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructType](#class-swaggestgocodebuildertemplatesstructstructtype)
- [\Swaggest\GoCodeBuilder\Templates\Struct\FluentSetter](#class-swaggestgocodebuildertemplatesstructfluentsetter)
- [\Swaggest\GoCodeBuilder\Templates\Struct\Tags](#class-swaggestgocodebuildertemplatesstructtags)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructFunctions](#class-swaggestgocodebuildertemplatesstructstructfunctions)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructFields](#class-swaggestgocodebuildertemplatesstructstructfields)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructCast](#class-swaggestgocodebuildertemplatesstructstructcast)
- [\Swaggest\GoCodeBuilder\Templates\Struct\StructIface](#class-swaggestgocodebuildertemplatesstructstructiface)
- [\Swaggest\GoCodeBuilder\Templates\Type\NoOmitEmpty (interface)](#interface-swaggestgocodebuildertemplatestypenoomitempty)
- [\Swaggest\GoCodeBuilder\Templates\Type\AnyType (interface)](#interface-swaggestgocodebuildertemplatestypeanytype)
- [\Swaggest\GoCodeBuilder\Templates\Type\FuncType](#class-swaggestgocodebuildertemplatestypefunctype)
- [\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)
- [\Swaggest\GoCodeBuilder\Templates\Type\NamedType (interface)](#interface-swaggestgocodebuildertemplatestypenamedtype)
- [\Swaggest\GoCodeBuilder\Templates\Type\Map](#class-swaggestgocodebuildertemplatestypemap)
- [\Swaggest\GoCodeBuilder\Templates\Type\TypeUtil](#class-swaggestgocodebuildertemplatestypetypeutil)
- [\Swaggest\GoCodeBuilder\Templates\Type\Slice](#class-swaggestgocodebuildertemplatestypeslice)
- [\Swaggest\GoCodeBuilder\Templates\Type\TypeCast](#class-swaggestgocodebuildertemplatestypetypecast)
- [\Swaggest\GoCodeBuilder\Templates\Type\TypeCastException](#class-swaggestgocodebuildertemplatestypetypecastexception)
- [\Swaggest\GoCodeBuilder\Templates\Type\Pointer](#class-swaggestgocodebuildertemplatestypepointer)
- [\Swaggest\GoCodeBuilder\TypeCast\CastRegistry](#class-swaggestgocodebuildertypecastcastregistry)
- [\Swaggest\GoCodeBuilder\TypeCast\RegistryMux](#class-swaggestgocodebuildertypecastregistrymux)
- [\Swaggest\GoCodeBuilder\TypeCast\Registry (interface)](#interface-swaggestgocodebuildertypecastregistry)
- [\Swaggest\GoCodeBuilder\TypeCast\PropertyCast](#class-swaggestgocodebuildertypecastpropertycast)
- [\Swaggest\GoCodeBuilder\TypeCast\CastFunctions (interface)](#interface-swaggestgocodebuildertypecastcastfunctions)
- [\Swaggest\GoCodeBuilder\TypeCast\Time](#class-swaggestgocodebuildertypecasttime)

<hr />

### Class: \Swaggest\GoCodeBuilder\GoCodeBuilder

> GoCodeBuilder provides file manager and names processor.

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct()</strong> : <em>void</em> |
| public | <strong>exportableName(</strong><em>mixed</em> <strong>$name</strong>, <em>bool</em> <strong>$requireBase=false</strong>)</strong> : <em>void</em> |
| public | <strong>privateName(</strong><em>mixed</em> <strong>$name</strong>)</strong> : <em>void</em> |
| public | <strong>setBuilderVersion(</strong><em>mixed</em> <strong>$builderVersion</strong>)</strong> : <em>void</em> |
| public | <strong>storeToDisk(</strong><em>mixed</em> <strong>$srcPath</strong>)</strong> : <em>void</em> |
| protected | <strong>toCamelCase(</strong><em>string</em> <strong>$string</strong>, <em>bool</em> <strong>$lowerFirst=false</strong>)</strong> : <em>string</em> |

*This class extends \Swaggest\CodeBuilder\CodeBuilder*

<hr />

### Class: \Swaggest\GoCodeBuilder\Import

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>mixed</em> <strong>$name</strong>, <em>mixed</em> <strong>$alias=null</strong>, <em>mixed</em> <strong>$defaultPackageName=null</strong>)</strong> : <em>void</em> |
| public | <strong>getPackage()</strong> : <em>string</em> |
| public | <strong>getReferencePrefix()</strong> : <em>mixed</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\Exception

| Visibility | Function |
|:-----------|:---------|

*This class extends \Exception*

*This class implements \Throwable*

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\StructHookCallback

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Closure](http://php.net/manual/en/class.closure.php)</em> <strong>$closure</strong>)</strong> : <em>void</em><br /><em>StructHookCallback constructor.</em> |
| public | <strong>process(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$structDef</strong>, <em>mixed</em> <strong>$path</strong>, <em>mixed</em> <strong>$schema</strong>)</strong> : <em>void</em> |

*This class implements [\Swaggest\GoCodeBuilder\JsonSchema\GoBuilderStructHook](#interface-swaggestgocodebuilderjsonschemagobuilderstructhook)*

<hr />

### Interface: \Swaggest\GoCodeBuilder\JsonSchema\GoBuilderPathToNameHook

| Visibility | Function |
|:-----------|:---------|
| public | <strong>pathToName(</strong><em>mixed</em> <strong>$path</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\MarshalEnum

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)</em> <strong>$type</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\NamedType](#interface-swaggestgocodebuildertemplatestypenamedtype)</em> <strong>$base</strong>, <em>array</em> <strong>$enum</strong>, <em>[\Swaggest\GoCodeBuilder\JsonSchema\GoBuilder](#class-swaggestgocodebuilderjsonschemagobuilder)</em> <strong>$builder</strong>)</strong> : <em>void</em><br /><em>MarshalEnum constructor.</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\GeneratedStruct

| Visibility | Function |
|:-----------|:---------|

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\Options

| Visibility | Function |
|:-----------|:---------|
| public static | <strong>setUpProperties(</strong><em>\Swaggest\GoCodeBuilder\JsonSchema\Properties/\Swaggest\GoCodeBuilder\JsonSchema\static</em> <strong>$properties</strong>, <em>\Swaggest\JsonSchema\Schema</em> <strong>$ownerSchema</strong>)</strong> : <em>void</em> |

*This class extends \Swaggest\JsonSchema\Structure\ClassStructure*

*This class implements \Swaggest\JsonSchema\Structure\WithResolvedValue, \Swaggest\JsonSchema\Structure\ObjectItemContract, \Traversable, \Iterator, \JsonSerializable, \ArrayAccess, \Swaggest\JsonSchema\Structure\ClassStructureContract*

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\UnmarshalUnion

| Visibility | Function |
|:-----------|:---------|
| public | <strong>patternVarName(</strong><em>mixed</em> <strong>$pattern</strong>)</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\TypeBuilder

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>\Swaggest\JsonSchema\Schema</em> <strong>$schema</strong>, <em>string</em> <strong>$path</strong>, <em>[\Swaggest\GoCodeBuilder\JsonSchema\GoBuilder](#class-swaggestgocodebuilderjsonschemagobuilder)</em> <strong>$goBuilder</strong>, <em>\Swaggest\GoCodeBuilder\JsonSchema\StructDef/null/[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$parentStruct=null</strong>, <em>bool</em> <strong>$isRequired=false</strong>)</strong> : <em>void</em><br /><em>TypeBuilder constructor.</em> |
| public | <strong>build()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\StripPrefixPathToNameHook

| Visibility | Function |
|:-----------|:---------|
| public | <strong>pathToName(</strong><em>mixed</em> <strong>$path</strong>)</strong> : <em>void</em> |

*This class implements [\Swaggest\GoCodeBuilder\JsonSchema\GoBuilderPathToNameHook](#interface-swaggestgocodebuilderjsonschemagobuilderpathtonamehook)*

<hr />

### Interface: \Swaggest\GoCodeBuilder\JsonSchema\GoBuilderStructHook

| Visibility | Function |
|:-----------|:---------|
| public | <strong>process(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$structDef</strong>, <em>string</em> <strong>$path</strong>, <em>\Swaggest\GoCodeBuilder\JsonSchema\Schema</em> <strong>$schema</strong>)</strong> : <em>null</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\MarshalingTestFunc

| Visibility | Function |
|:-----------|:---------|
| public static | <strong>make(</strong><em>[\Swaggest\GoCodeBuilder\JsonSchema\GeneratedStruct](#class-swaggestgocodebuilderjsonschemageneratedstruct)</em> <strong>$struct</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\Exception

| Visibility | Function |
|:-----------|:---------|

*This class extends \Exception*

*This class implements \Throwable*

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\GoBuilder

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct()</strong> : <em>void</em> |
| public | <strong>getClass(</strong><em>\Swaggest\GoCodeBuilder\JsonSchema\Schema</em> <strong>$schema</strong>, <em>string</em> <strong>$path</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> |
| public | <strong>getCode()</strong> : <em>mixed</em> |
| public | <strong>getGeneratedStruct(</strong><em>\Swaggest\GoCodeBuilder\JsonSchema\Schema</em> <strong>$schema</strong>, <em>string</em> <strong>$path</strong>)</strong> : <em>mixed/[\Swaggest\GoCodeBuilder\JsonSchema\GeneratedStruct](#class-swaggestgocodebuilderjsonschemageneratedstruct)</em> |
| public | <strong>getGeneratedStructs()</strong> : <em>[\Swaggest\GoCodeBuilder\JsonSchema\GeneratedStruct](#class-swaggestgocodebuilderjsonschemageneratedstruct)[]</em> |
| public | <strong>getType(</strong><em>\Swaggest\GoCodeBuilder\JsonSchema\Schema</em> <strong>$schema</strong>, <em>string</em> <strong>$path=`'#'`</strong>, <em>\Swaggest\GoCodeBuilder\JsonSchema\StructDef/null/[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$parentStruct=null</strong>, <em>bool</em> <strong>$isRequired=false</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| public | <strong>pathToName(</strong><em>mixed</em> <strong>$path</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\MarshalJson

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\JsonSchema\GoBuilder](#class-swaggestgocodebuilderjsonschemagobuilder)</em> <strong>$builder</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$type</strong>)</strong> : <em>void</em> |
| public | <strong>addNamedProperty(</strong><em>mixed</em> <strong>$name</strong>)</strong> : <em>void</em> |
| public | <strong>addPatternProperty(</strong><em>string</em> <strong>$regex</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)</em> <strong>$structProperty</strong>)</strong> : <em>\Swaggest\GoCodeBuilder\JsonSchema\$this</em> |
| public | <strong>addSomeOf(</strong><em>mixed</em> <strong>$kind</strong>, <em>mixed</em> <strong>$name</strong>)</strong> : <em>void</em> |
| public | <strong>enableAdditionalProperties(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)</em> <strong>$structProperty</strong>)</strong> : <em>void</em> |
| public | <strong>forbidAdditionalProperties()</strong> : <em>void</em> |
| public | <strong>isAdditionalPropertiesEnabled()</strong> : <em>bool</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\JsonSchema\MarshalUnion

| Visibility | Function |
|:-----------|:---------|
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Style\Comment

| Visibility | Function |
|:-----------|:---------|
| public static | <strong>sentence(</strong><em>mixed</em> <strong>$s</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\Style\Initialisms

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct()</strong> : <em>void</em> |
| public | <strong>process(</strong><em>mixed</em> <strong>$goName</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Imports

| Visibility | Function |
|:-----------|:---------|
| public | <strong>add(</strong><em>[\Swaggest\GoCodeBuilder\Import](#class-swaggestgocodebuilderimport)</em> <strong>$import</strong>)</strong> : <em>\Swaggest\GoCodeBuilder\Templates\$this</em> |
| public | <strong>addByName(</strong><em>mixed</em> <strong>$name</strong>, <em>mixed</em> <strong>$alias=null</strong>)</strong> : <em>void</em> |
| public | <strong>demand()</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\GoTemplate (abstract)

| Visibility | Function |
|:-----------|:---------|
| public | <strong>escapeValue(</strong><em>mixed</em> <strong>$value</strong>)</strong> : <em>void</em> |
| public | <strong>getComment()</strong> : <em>string</em> |
| public | <strong>ifThenElse(</strong><em>mixed</em> <strong>$condition</strong>, <em>mixed</em> <strong>$then</strong>, <em>string</em> <strong>$else=`''`</strong>)</strong> : <em>void</em> |
| public | <strong>padLines(</strong><em>mixed</em> <strong>$with</strong>, <em>mixed</em> <strong>$text</strong>, <em>bool</em> <strong>$skipFirst=true</strong>, <em>bool</em> <strong>$forcePad=false</strong>)</strong> : <em>void</em> |
| public | <strong>setComment(</strong><em>mixed</em> <strong>$comment</strong>)</strong> : <em>void</em> |
| public | <strong>stripEmptyLines(</strong><em>mixed</em> <strong>$text</strong>)</strong> : <em>void</em> |
| public | <strong>tabIndents(</strong><em>mixed</em> <strong>$text</strong>, <em>mixed</em> <strong>$spaces=4</strong>)</strong> : <em>void</em> |
| public | <strong>trim(</strong><em>mixed</em> <strong>$s</strong>)</strong> : <em>void</em> |
| protected | <strong>renderComment()</strong> : <em>void</em> |

*This class extends \Swaggest\CodeBuilder\AbstractTemplate*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Code

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>mixed</em> <strong>$body=null</strong>)</strong> : <em>void</em> |
| public | <strong>addSnippet(</strong><em>mixed</em> <strong>$code</strong>, <em>bool</em> <strong>$prepend=false</strong>, <em>mixed</em> <strong>$uniqueKey=null</strong>)</strong> : <em>void</em> |
| public | <strong>imports()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Imports](#class-swaggestgocodebuildertemplatesimports)</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\GoFile

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>string</em> <strong>$package</strong>, <em>string/null</em> <strong>$importPath=null</strong>)</strong> : <em>void</em><br /><em>GoFile constructor.</em> |
| public | <strong>commitTransaction()</strong> : <em>void</em> |
| public | <strong>dropTransaction(</strong><em>bool</em> <strong>$ignoreMissing=true</strong>)</strong> : <em>void</em> |
| public | <strong>getCode()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Code](#class-swaggestgocodebuildertemplatescode)</em> |
| public | <strong>getComment()</strong> : <em>mixed</em> |
| public static | <strong>getCurrentGoFile()</strong> : <em>mixed</em> |
| public | <strong>getImportPath()</strong> : <em>null</em> |
| public | <strong>getImports()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Imports](#class-swaggestgocodebuildertemplatesimports)</em> |
| public | <strong>getPackage()</strong> : <em>mixed</em> |
| public | <strong>setCode(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Code](#class-swaggestgocodebuildertemplatescode)</em> <strong>$code</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)</em> |
| public static | <strong>setCurrentGoFile(</strong><em>[\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)</em> <strong>$currentGoFile=null</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)/null previous go file</em> |
| public | <strong>setDependentCode(</strong><em>mixed</em> <strong>$uniqueKey</strong>, <em>mixed</em> <strong>$value</strong>)</strong> : <em>void</em> |
| public | <strong>setImportPath(</strong><em>null</em> <strong>$importPath</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)</em> |
| public | <strong>setPackage(</strong><em>mixed</em> <strong>$package</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)</em> |
| public | <strong>setSkipImportComment(</strong><em>boolean</em> <strong>$skipImportComment</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\GoFile](#class-swaggestgocodebuildertemplatesgofile)</em> |
| public | <strong>startTransaction()</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Constant\TypeConstBlock

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)</em> <strong>$type</strong>)</strong> : <em>void</em> |
| public | <strong>addValue(</strong><em>mixed</em> <strong>$name</strong>, <em>mixed</em> <strong>$value</strong>, <em>mixed</em> <strong>$comment=null</strong>)</strong> : <em>void</em> |
| public | <strong>getValues()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Func\FuncDef

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>string</em> <strong>$name</strong>, <em>string</em> <strong>$comment=`''`</strong>)</strong> : <em>void</em><br /><em>FuncDef constructor.</em> |
| public | <strong>getArguments()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\Arguments](#class-swaggestgocodebuildertemplatesfuncarguments)</em> |
| public | <strong>getName()</strong> : <em>string</em> |
| public | <strong>getRenderMode()</strong> : <em>string</em> |
| public | <strong>getResult()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\Result](#class-swaggestgocodebuildertemplatesfuncresult)/null</em> |
| public | <strong>getSelf()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\Argument](#class-swaggestgocodebuildertemplatesfuncargument)/null</em> |
| public | <strong>setArguments(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\Arguments](#class-swaggestgocodebuildertemplatesfuncarguments)</em> <strong>$arguments</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| public | <strong>setBody(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Code](#class-swaggestgocodebuildertemplatescode)</em> <strong>$body</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| public | <strong>setName(</strong><em>string</em> <strong>$name</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| public | <strong>setRenderMode(</strong><em>string</em> <strong>$renderMode</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| public | <strong>setResult(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\Result](#class-swaggestgocodebuildertemplatesfuncresult)</em> <strong>$result=null</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| public | <strong>setSelf(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\Argument](#class-swaggestgocodebuildertemplatesfuncargument)</em> <strong>$self</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Func\FuncIface

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> <strong>$funcDef</strong>)</strong> : <em>void</em><br /><em>FuncIface constructor.</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Func\Arguments

| Visibility | Function |
|:-----------|:---------|
| public | <strong>add(</strong><em>mixed</em> <strong>$name</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>, <em>bool</em> <strong>$isVariadic=false</strong>)</strong> : <em>void</em> |
| public | <strong>count()</strong> : <em>void</em> |
| public | <strong>toTypesString()</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Func\Result

| Visibility | Function |
|:-----------|:---------|
| public | <strong>toTypesString()</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\Func\Arguments](#class-swaggestgocodebuildertemplatesfuncarguments)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Func\Argument

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>string</em> <strong>$name</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>, <em>bool</em> <strong>$isVariadic=false</strong>)</strong> : <em>void</em><br /><em>Argument constructor.</em> |
| public | <strong>getType()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Iface\IfaceDef

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>mixed</em> <strong>$name</strong>, <em>mixed</em> <strong>$comment=null</strong>)</strong> : <em>void</em> |
| public | <strong>addFunc(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> <strong>$func</strong>, <em>bool</em> <strong>$prepend=false</strong>)</strong> : <em>void</em> |
| public | <strong>addType(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)</em> <strong>$func</strong>, <em>bool</em> <strong>$prepend=false</strong>)</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Mapping\Mapping

| Visibility | Function |
|:-----------|:---------|

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructDef

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>string</em> <strong>$name</strong>, <em>string</em> <strong>$comment=`''`</strong>)</strong> : <em>void</em><br /><em>StructDef constructor.</em> |
| public | <strong>addFunc(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> <strong>$func</strong>, <em>bool</em> <strong>$prepend=false</strong>)</strong> : <em>void</em> |
| public | <strong>addProperty(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)</em> <strong>$property</strong>, <em>bool</em> <strong>$prepend=false</strong>)</strong> : <em>void</em> |
| public | <strong>getCode()</strong> : <em>mixed</em> |
| public | <strong>getFuncs()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)[]</em> |
| public | <strong>getImport()</strong> : <em>\Swaggest\GoCodeBuilder\Templates\Struct\Import/null</em> |
| public | <strong>getName()</strong> : <em>mixed</em> |
| public | <strong>getProperties()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)[]</em> |
| public | <strong>getType()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructType](#class-swaggestgocodebuildertemplatesstructstructtype)</em> |
| public | <strong>renderFields()</strong> : <em>void</em> |
| public | <strong>renderFuncs()</strong> : <em>void</em> |
| public | <strong>setImport(</strong><em>[\Swaggest\GoCodeBuilder\Import](#class-swaggestgocodebuilderimport)</em> <strong>$import=null</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> |
| public | <strong>setName(</strong><em>mixed</em> <strong>$name</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> |
| public | <strong>setType(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)</em> <strong>$type</strong>)</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructType

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$structDef</strong>)</strong> : <em>void</em><br /><em>StructType constructor.</em> |
| public | <strong>getName()</strong> : <em>mixed</em> |
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\NamedType](#interface-swaggestgocodebuildertemplatestypenamedtype), [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\FluentSetter

| Visibility | Function |
|:-----------|:---------|
| public static | <strong>addToStruct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$structDef</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)</em> <strong>$goProperty</strong>)</strong> : <em>void</em> |
| public static | <strong>make(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$structDef</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)</em> <strong>$goProperty</strong>)</strong> : <em>void</em> |
| public static | <strong>makeMap(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$structDef</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructProperty](#class-swaggestgocodebuildertemplatesstructstructproperty)</em> <strong>$goProperty</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\Tags

| Visibility | Function |
|:-----------|:---------|
| public | <strong>setTag(</strong><em>mixed</em> <strong>$key</strong>, <em>mixed</em> <strong>$value</strong>)</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructFunctions

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$struct</strong>)</strong> : <em>void</em><br /><em>StructFunctions constructor.</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends \Swaggest\CodeBuilder\AbstractTemplate*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructProperty

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>mixed</em> <strong>$name</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\Tags](#class-swaggestgocodebuildertemplatesstructtags)</em> <strong>$tags=null</strong>)</strong> : <em>void</em> |
| public | <strong>getName()</strong> : <em>string/null</em> |
| public | <strong>getTags()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\Tags](#class-swaggestgocodebuildertemplatesstructtags)</em> |
| public | <strong>getType()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| public | <strong>setType(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructFields

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$struct</strong>)</strong> : <em>void</em><br /><em>StructType constructor.</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructCast

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$baseStruct</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$derivedStruct</strong>, <em>array/string[]</em> <strong>$propNamesMap=array()</strong>, <em>[\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)</em> <strong>$registry=null</strong>)</strong> : <em>void</em><br /><em>StructCast constructor.</em> |
| public | <strong>getBaseStruct()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> |
| public | <strong>getBaseTypeString()</strong> : <em>string</em> |
| public | <strong>getDerivedStruct()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> |
| public | <strong>getDerivedTypeString()</strong> : <em>string</em> |
| public | <strong>getLoadFrom()</strong> : <em>mixed</em> |
| public | <strong>getMapTo()</strong> : <em>mixed</em> |
| public | <strong>setPropMap(</strong><em>mixed</em> <strong>$baseName</strong>, <em>mixed</em> <strong>$derivedName</strong>)</strong> : <em>void</em> |

*This class implements [\Swaggest\GoCodeBuilder\TypeCast\CastFunctions](#interface-swaggestgocodebuildertypecastcastfunctions)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Struct\StructIface

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$struct</strong>, <em>mixed</em> <strong>$name</strong>, <em>string</em> <strong>$comment=`''`</strong>)</strong> : <em>void</em> |
| public | <strong>getIface()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Interface: \Swaggest\GoCodeBuilder\Templates\Type\NoOmitEmpty

| Visibility | Function |
|:-----------|:---------|
| public | <strong>isNoOmitEmpty()</strong> : <em>bool</em> |

<hr />

### Interface: \Swaggest\GoCodeBuilder\Templates\Type\AnyType

| Visibility | Function |
|:-----------|:---------|
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| public | <strong>render()</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\FuncType

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> <strong>$func</strong>)</strong> : <em>void</em> |
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\Type

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>mixed</em> <strong>$type</strong>, <em>[\Swaggest\GoCodeBuilder\Import](#class-swaggestgocodebuilderimport)</em> <strong>$import=null</strong>)</strong> : <em>void</em> |
| public | <strong>equals(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)</em> <strong>$type</strong>)</strong> : <em>void</em> |
| public | <strong>getImport()</strong> : <em>\Swaggest\GoCodeBuilder\Templates\Type\Import/null</em> |
| public | <strong>getName()</strong> : <em>mixed</em> |
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\NamedType](#interface-swaggestgocodebuildertemplatestypenamedtype), [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)*

<hr />

### Interface: \Swaggest\GoCodeBuilder\Templates\Type\NamedType

| Visibility | Function |
|:-----------|:---------|
| public | <strong>getName()</strong> : <em>string</em> |

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\Map

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$keyType</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$valueType</strong>)</strong> : <em>void</em><br /><em>Map constructor.</em> |
| public | <strong>getKeyType()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| public | <strong>getValueType()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| public | <strong>renderName()</strong> : <em>void</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\TypeUtil

| Visibility | Function |
|:-----------|:---------|
| public static | <strong>equals(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$one</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$two</strong>)</strong> : <em>void</em> |
| public static | <strong>fromString(</strong><em>string</em> <strong>$typeString</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)/[\Swaggest\GoCodeBuilder\Templates\Type\Type](#class-swaggestgocodebuildertemplatestypetype)</em> |
| public static | <strong>getBasicType(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>mixed</em> |
| public static | <strong>isCastable(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$to</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$from</strong>)</strong> : <em>bool</em> |
| public static | <strong>isFloat(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>bool</em> |
| public static | <strong>isInt(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>bool</em> |
| public static | <strong>isNumber(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>bool</em> |
| public static | <strong>resolvePointer(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>void</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\Slice

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>void</em><br /><em>Slice constructor.</em> |
| public | <strong>getType()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\TypeCast

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$toType</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$fromType</strong>, <em>string</em> <strong>$toVarName</strong>, <em>string</em> <strong>$fromVarName</strong>, <em>[\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)</em> <strong>$registry=null</strong>)</strong> : <em>void</em><br /><em>TypeCast constructor.</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\TypeCastException

| Visibility | Function |
|:-----------|:---------|

*This class extends \Exception*

*This class implements \Throwable*

<hr />

### Class: \Swaggest\GoCodeBuilder\Templates\Type\Pointer

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>void</em><br /><em>Pointer constructor.</em> |
| public | <strong>getType()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| public | <strong>getTypeString()</strong> : <em>mixed</em> |
| public | <strong>isNoOmitEmpty()</strong> : <em>bool</em> |
| public | <strong>setNoOmitEmpty(</strong><em>bool</em> <strong>$noOmitEmpty</strong>)</strong> : <em>void</em> |
| public static | <strong>tryDereferenceOnce(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$type</strong>)</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> |
| protected | <strong>toString()</strong> : <em>void</em> |

*This class extends [\Swaggest\GoCodeBuilder\Templates\GoTemplate](#class-swaggestgocodebuildertemplatesgotemplate-abstract)*

*This class implements [\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype), [\Swaggest\GoCodeBuilder\Templates\Type\NoOmitEmpty](#interface-swaggestgocodebuildertemplatestypenoomitempty)*

<hr />

### Class: \Swaggest\GoCodeBuilder\TypeCast\CastRegistry

| Visibility | Function |
|:-----------|:---------|
| public | <strong>addStructCast(</strong><em>[\Swaggest\GoCodeBuilder\TypeCast\CastFunctions](#interface-swaggestgocodebuildertypecastcastfunctions)</em> <strong>$cast</strong>)</strong> : <em>void</em> |
| public | <strong>canProcess(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>)</strong> : <em>bool</em> |
| public | <strong>process(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>, <em>string</em> <strong>$toVarName</strong>, <em>string</em> <strong>$fromVarName</strong>, <em>string</em> <strong>$assignOp</strong>)</strong> : <em>string</em> |
| public | <strong>resetUsedCastFuncs()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)[]</em> |

*This class implements [\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)*

<hr />

### Class: \Swaggest\GoCodeBuilder\TypeCast\RegistryMux

| Visibility | Function |
|:-----------|:---------|
| public | <strong>addRegistry(</strong><em>[\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)</em> <strong>$registry</strong>)</strong> : <em>void</em> |
| public | <strong>canProcess(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>)</strong> : <em>bool</em> |
| public static | <strong>getStd()</strong> : <em>\Swaggest\GoCodeBuilder\TypeCast\static</em> |
| public | <strong>process(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>, <em>string</em> <strong>$toVarName</strong>, <em>string</em> <strong>$fromVarName</strong>, <em>string</em> <strong>$assignOp</strong>)</strong> : <em>\Swaggest\GoCodeBuilder\TypeCast\Code/string</em> |

*This class implements [\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)*

<hr />

### Interface: \Swaggest\GoCodeBuilder\TypeCast\Registry

| Visibility | Function |
|:-----------|:---------|
| public | <strong>canProcess(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>)</strong> : <em>bool</em> |
| public | <strong>process(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>, <em>string</em> <strong>$toVarName</strong>, <em>string</em> <strong>$fromVarName</strong>, <em>string</em> <strong>$assignOp</strong>)</strong> : <em>\Swaggest\GoCodeBuilder\TypeCast\Code/string</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\TypeCast\PropertyCast

| Visibility | Function |
|:-----------|:---------|
| public | <strong>__construct(</strong><em>[\Swaggest\GoCodeBuilder\Templates\Struct\StructDef](#class-swaggestgocodebuildertemplatesstructstructdef)</em> <strong>$baseStruct</strong>, <em>string</em> <strong>$propertyName</strong>, <em>[\Swaggest\GoCodeBuilder\Templates\Type\AnyType](#interface-swaggestgocodebuildertemplatestypeanytype)</em> <strong>$derivedType</strong>, <em>[\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)</em> <strong>$typeRegistry</strong>)</strong> : <em>void</em><br /><em>PropertyCast constructor.</em> |
| public | <strong>getBaseTypeString()</strong> : <em>string</em> |
| public | <strong>getDerivedTypeString()</strong> : <em>string</em> |
| public | <strong>getLoadFrom()</strong> : <em>mixed</em> |
| public | <strong>getMapTo()</strong> : <em>mixed</em> |

*This class implements [\Swaggest\GoCodeBuilder\TypeCast\CastFunctions](#interface-swaggestgocodebuildertypecastcastfunctions)*

<hr />

### Interface: \Swaggest\GoCodeBuilder\TypeCast\CastFunctions

| Visibility | Function |
|:-----------|:---------|
| public | <strong>getBaseTypeString()</strong> : <em>string</em> |
| public | <strong>getDerivedTypeString()</strong> : <em>string</em> |
| public | <strong>getLoadFrom()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |
| public | <strong>getMapTo()</strong> : <em>[\Swaggest\GoCodeBuilder\Templates\Func\FuncDef](#class-swaggestgocodebuildertemplatesfuncfuncdef)</em> |

<hr />

### Class: \Swaggest\GoCodeBuilder\TypeCast\Time

| Visibility | Function |
|:-----------|:---------|
| public | <strong>canProcess(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>)</strong> : <em>bool</em> |
| public | <strong>process(</strong><em>string</em> <strong>$toTypeString</strong>, <em>string</em> <strong>$fromTypeString</strong>, <em>string</em> <strong>$toVarName</strong>, <em>string</em> <strong>$fromVarName</strong>, <em>string</em> <strong>$assignOp</strong>)</strong> : <em>\Swaggest\GoCodeBuilder\TypeCast\Code/string</em> |

*This class implements [\Swaggest\GoCodeBuilder\TypeCast\Registry](#interface-swaggestgocodebuildertypecastregistry)*

