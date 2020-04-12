<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\JsonSchema\Schema;

class MarshalJson extends GoTemplate
{
    private $type;
    private $builder;

    /** @var string[] */
    private $propertyNames;

    /** @var string[] */
    public $required;

    /** @var string[] */
    public $distinctNullNames = [];

    private $additionalPropertiesEnabled = null;

    /** @var StructProperty */
    private $additionalProperties;

    /** @var StructProperty[] */
    private $patternProperties;

    private $someOf;

    /** @var AnyType */
    public $not;

    private $code;

    /** @var array */
    public $constValues;

    public function __construct(GoBuilder $builder, StructDef $type)
    {
        $this->type = $type;
        $this->builder = $builder;
        $this->code = new Code();
    }

    public function forbidAdditionalProperties()
    {
        $this->additionalPropertiesEnabled = false;
        return $this;
    }

    public function enableAdditionalProperties(StructProperty $structProperty)
    {
        $this->additionalPropertiesEnabled = true;
        $this->additionalProperties = $structProperty;
        return $this;
    }

    public function isAdditionalPropertiesEnabled()
    {
        return $this->additionalPropertiesEnabled;
    }

    /**
     * @param string $regex
     * @param StructProperty $structProperty
     * @return $this
     */
    public function addPatternProperty($regex, StructProperty $structProperty)
    {
        $this->patternProperties[$regex] = $structProperty;
        return $this;
    }

    public function addSomeOf($kind, $name)
    {
        $this->someOf[$kind][$name] = $name;
        return $this;
    }

    public function addNamedProperty($name)
    {
        $this->propertyNames[] = $name;
        return $this;
    }

    private function receiver()
    {
        return $this->propertyNames === null ? ':receiver' : 'm:receiver';
    }

    protected function toString()
    {
        if ($this->patternProperties === null
            && $this->additionalPropertiesEnabled === null
            && $this->someOf === null
            && $this->constValues === null
            && empty($this->required)) {
            return '';
        }

        $result = '';

        $renderUnmarshal = $this->renderUnmarshal();
        $renderMarshal = $this->renderMarshal();

        if ($this->propertyNames !== null && ($renderUnmarshal !== '' || $renderMarshal !== '')) {
            $result .= <<<'GO'
type marshal:type :type


GO;
        }

        $result .= <<<GO
{$this->padLines('',
            $renderUnmarshal
            . $renderMarshal)}

GO;
        if (!$this->builder->options->skipUnmarshal) {
            $this->builder->getCode()->addSnippet($this->builder->unmarshalUnion, false, 'unmarshal_union');
            if ($this->patternProperties !== null) {
                $this->builder->unmarshalUnion->withPatternProperties = true;
            }
        }

        $this->code->addSnippet(new PlaceholderString($result, [
            ':type' => $this->type->getType(),
            ':receiver' => new Code(strtolower($this->type->getType()->getName()[0]))
        ]));

        return $this->code;
    }

    private function renderConstRawMessage()
    {
        if ($this->constValues !== null) {
            $this->code->imports()->addByName('encoding/json');
            $result = <<<GO
var (
	// const:type is unconditionally added to JSON.
	const:type = json.RawMessage({$this->escapeValue(json_encode($this->constValues))})
)


GO;

            return new PlaceholderString($result, [
                ':type' => $this->type->getType(),
                ':receiver' => new Code(strtolower($this->type->getType()->getName()[0]))
            ]);

        }
        return '';
    }


    private function renderKeyNames()
    {
        $result = '';

        if (
            $this->propertyNames !== null &&
            ($this->patternProperties || $this->additionalPropertiesEnabled !== null)
        ) {
            // All known property names (including constant values) has to be removed from
            // pattern and additional properties matching.
            $knownKeys = $this->propertyNames;
            if (!empty($this->constValues)) {
                foreach ($this->constValues as $propertyName => $value) {
                    $knownKeys[] = $propertyName;
                }
            }
            $result .= <<<GO
var knownKeys:type = []string{
	"{$this->padLines("\t", implode("\",\n\"", $knownKeys))}",
}


GO;
        }

        if (!empty($this->required)) {
            $result .= <<<GO
var requireKeys:type = []string{
	"{$this->padLines("\t", implode("\",\n\"", $this->required))}",
}


GO;

        }

        return $result;
    }

    private function renderUnmarshal()
    {
        if ($this->builder->options->skipUnmarshal) {
            return '';
        }

        if ($this->builder->unmarshalUnion === null) {
            $this->builder->unmarshalUnion = new UnmarshalUnion();
            $this->builder->unmarshalUnion->goBuilder = $this->builder;
        }

        $mustUnmarshal = $this->renderMustUnmarshal();
        $mayUnmarshal = $this->renderTypeUnmarshal() . $this->renderAnyOfUnmarshal() . $this->renderOneOfUnmarshal();
        $withKnownKeys = false; // TODO move to renderer

        if (
            $this->propertyNames !== null &&
            ($this->patternProperties || $this->additionalPropertiesEnabled !== null)
        ) {
            $withKnownKeys = true;
        }

        $mapUnmarshal = '';
        if ($this->patternProperties !== null
            || $this->additionalPropertiesEnabled !== null
            || $this->constValues !== null
            || !empty($this->distinctNullNames)
            || !empty($this->required)) {
            $this->code->imports()->addByName('encoding/json');
            $mapUnmarshal = <<<'GO'


var rawMap map[string]json.RawMessage

err = json.Unmarshal(data, &rawMap)
if err != nil {
    rawMap = nil
}

GO;

        }

        if (!empty($this->required)) {
            $this->code->imports()->addByName('errors');
            $mapUnmarshal .= <<<GO

for _, key := range requireKeys:type {
    if _, found := rawMap[key]; !found {
        return errors.New("required key missing: " + key)
    }
}


GO;
        }

        if ($this->constValues !== null) {
            $this->code->imports()->addByName('fmt');
            foreach ($this->constValues as $name => $value) {
                $mapUnmarshal .= <<<GO

if v, ok := rawMap[{$this->escapeValue($name)}]; !ok || string(v) != {$this->escapeValue(json_encode($value))} {
	return fmt.Errorf({$this->escapeValue('bad or missing const value for "' . $name . '" (' . json_encode($value) . ' expected, %s received)')}, v)
}

delete(rawMap, {$this->escapeValue($name)})


GO;
            }
        }

        foreach ($this->distinctNullNames as $goPropertyName => $name) {
            $name = $this->escapeValue($name);
            $mapUnmarshal .= <<<GO

if m:receiver.$goPropertyName == nil {
    if _, ok := rawMap[$name]; ok {
        var v interface{}
        m:receiver.$goPropertyName = &v
    }
}

GO;

        }

        if ($mapUnmarshal && $withKnownKeys) {
            $mapUnmarshal .= <<<'GO'

for _, key := range knownKeys:type {
    delete(rawMap, key)
}


GO;

        }


        if ($this->patternProperties !== null) {
            $mapUnmarshal .= <<<GO

for key, rawValue := range rawMap {
    matched := false
    
GO;

            foreach ($this->patternProperties as $regex => $patternProperty) {
                $regexName = $this->builder->unmarshalUnion->patternVarName($regex);
                $mapType = $patternProperty->getType();
                $itemType = 'interface{}';
                if ($mapType instanceof Map) {
                    $itemType = $mapType->getValueType()->render();
                }
                $this->code->imports()->addByName('encoding/json');
                $mapUnmarshal .= <<<GO

    if $regexName.MatchString(key) {
        matched = true
        
        if {$this->receiver()}.{$patternProperty->getName()} == nil {
            {$this->receiver()}.{$patternProperty->getName()} = make({$patternProperty->getType()->render()}, 1)
        }
        
        var val {$itemType}

        err = json.Unmarshal(rawValue, &val)
        if err != nil {
            return err
        }
        
        {$this->receiver()}.{$patternProperty->getName()}[key] = val
    }

GO;
            }


            $mapUnmarshal .= <<<GO

    if matched {
        delete(rawMap, key)
    }
}

GO;


        }


        if ($this->additionalPropertiesEnabled) {
            $mapType = $this->additionalProperties->getType();
            $itemType = 'interface{}';
            if ($mapType instanceof Map) {
                $itemType = $mapType->getValueType()->render();
            }

            $this->code->imports()->addByName('encoding/json');
            $mapUnmarshal .= <<<GO

for key, rawValue := range rawMap {
    if {$this->receiver()}.{$this->additionalProperties->getName()} == nil {
        {$this->receiver()}.{$this->additionalProperties->getName()} = make({$this->additionalProperties->getType()->render()}, 1)
    }

    var val {$itemType}

    err = json.Unmarshal(rawValue, &val)
    if err != nil {
        return err
    }

    {$this->receiver()}.{$this->additionalProperties->getName()}[key] = val
}

GO;
        }

        // Additional properties forbidden.
        if ($this->additionalPropertiesEnabled === false) {
            $mapUnmarshal .= <<<'GO'

if len(rawMap) != 0 {
    offendingKeys := make([]string, 0, len(rawMap))

    for key := range rawMap {
        offendingKeys = append(offendingKeys, key)
    }

    return fmt.Errorf("additional properties not allowed in :type: %v", offendingKeys)
}

GO;

        }


        $funcBody = <<<GO
var err error

{$this->renderNot()}{$this->renderMainStructStart()}{$mustUnmarshal}{$mayUnmarshal}{$mapUnmarshal}
{$this->renderMainStructEnd()}

return nil
GO;


        return <<<GO
{$this->renderKeyNames()}// UnmarshalJSON decodes JSON.
func (:receiver *:type) UnmarshalJSON(data []byte) error {
{$this->padLines("\t", $this->tabIndents($this->stripEmptyLines($funcBody)), false)}
}


GO;
    }

    private function renderMarshal()
    {
        if ($this->builder->options->skipMarshal) {
            return '';
        }

        $maps = '';
        $mapsCnt = 0;

        if ($this->constValues !== null) {
            $maps .= ', const:type';
            $mapsCnt++;
        }

        if ($this->propertyNames !== null) {
            $maps .= ', marshal:type(:receiver)';
            $mapsCnt++;
        }

        if ($this->patternProperties !== null) {
            foreach ($this->patternProperties as $regex => $patternProperty) {
                $maps .= ', :receiver.' . $patternProperty->getName();
                $mapsCnt++;
            }
        }

        if ($this->additionalPropertiesEnabled) {
            $maps .= ', :receiver.' . $this->additionalProperties->getName();
            $mapsCnt++;
        }

        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                foreach ($unionPropertyNames as $propertyName) {
                    $maps .= ', :receiver.' . $propertyName;
                    $mapsCnt++;
                }
            }
        }

        // Return if only marshaling original struct.
        if ($maps === ', marshal:type(:receiver)') {
            return '';
        }

        $maps = substr($maps, 2);

        $earlyReturn = '';
        if ($this->additionalPropertiesEnabled && $this->propertyNames !== null && $mapsCnt === 2) {
            $earlyReturn = <<<GO
	if len(:receiver.{$this->additionalProperties->getName()}) == 0 {
		return json.Marshal(marshal:type(:receiver))
	}


GO;

        }

        if (!$this->builder->options->skipMarshal) {
            if (null === $this->builder->marshalUnion) {
                $this->builder->marshalUnion = new MarshalUnion();
                $this->builder->getCode()->addSnippet($this->builder->marshalUnion, false, 'marshal_union');
            }
        }

        return <<<GO
{$this->renderConstRawMessage()}// MarshalJSON encodes JSON.
func (:receiver :type) MarshalJSON() ([]byte, error) {
{$earlyReturn}	return marshalUnion($maps)
}

GO;


    }

    private function renderMustUnmarshal()
    {
        $result = '';
        if ($this->propertyNames) {
            $this->code->imports()->addByName('encoding/json');
            $result .= <<<'GO'


err = json.Unmarshal(data, &m:receiver)
if err != nil {
    return err
}

GO;
        }

        $kind = Schema::names()->allOf;
        if (!isset($this->someOf[$kind])) {
            return $result;
        }

        foreach ($this->someOf[$kind] as $propertyName) {
            $this->code->imports()->addByName('encoding/json');
            $result .= <<<GO


err = json.Unmarshal(data, &{$this->receiver()}.{$propertyName})
if err != nil {
    return err
}

GO;
        }

        return $result;
    }

    private function renderOneOfUnmarshal()
    {
        $result = '';

        $kind = Schema::names()->oneOf;
        if (!isset($this->someOf[$kind])) {
            return $result;
        }

        $this->code->imports()
            ->addByName('encoding/json')
            ->addByName('fmt');

        $count = count($this->someOf[$kind]);
        $result .= <<<GO


oneOfErrors := make(map[string]error, $count)
oneOfValid := 0

GO;

        foreach ($this->someOf[$kind] as $i => $propertyName) {
            $result .= <<<GO


err = json.Unmarshal(data, &{$this->receiver()}.{$propertyName})
if err != nil {
    oneOfErrors["$i"] = err
    {$this->receiver()}.{$propertyName} = nil
} else {
    oneOfValid++
}

GO;
        }

        $result .= <<<'GO'


if oneOfValid != 1 {
    return fmt.Errorf("oneOf constraint failed for :type with %d valid results: %v", oneOfValid, oneOfErrors)
}

GO;

        return $result;
    }

    private function renderAnyOfUnmarshal()
    {
        $result = '';

        $kind = Schema::names()->anyOf;
        if (!isset($this->someOf[$kind])) {
            return $result;
        }

        $this->code->imports()
            ->addByName('encoding/json')
            ->addByName('fmt');

        $count = count($this->someOf[$kind]);
        $result .= <<<GO


anyOfErrors := make(map[string]error, $count)
anyOfValid := 0

GO;

        foreach ($this->someOf[$kind] as $i => $propertyName) {
            $result .= <<<GO

err = json.Unmarshal(data, &{$this->receiver()}.{$propertyName})
if err != nil {
    anyOfErrors["$i"] = err
    {$this->receiver()}.{$propertyName} = nil
} else {
    anyOfValid++
}

GO;
        }

        $result .= <<<'GO'


if anyOfValid == 0 {
    return fmt.Errorf("anyOf constraint for :type failed with %d valid results: %v", anyOfValid, anyOfErrors)
}

GO;

        return $result;
    }

    private function renderTypeUnmarshal()
    {
        $result = '';

        $kind = Schema::names()->type;
        if (!isset($this->someOf[$kind])) {
            return $result;
        }

        $this->code->imports()
            ->addByName('encoding/json');

        $result .= <<<'GO'

typeValid := false

GO;


        foreach ($this->someOf[$kind] as $i => $propertyName) {
            $result .= <<<GO

if !typeValid {
    err = json.Unmarshal(data, &{$this->receiver()}.{$propertyName})
    if err != nil {
        {$this->receiver()}.{$propertyName} = nil
    } else {
        typeValid = true
    }
}

GO;
        }

        $result .= <<<'GO'

if !typeValid {
    return err
}

GO;


        return $result;
    }


    private function renderNot()
    {
        if (empty($this->not)) {
            return '';
        }

        return <<<GO

var not {$this->not->render()}

if json.Unmarshal(data, &not) == nil {
    return errors.New("not constraint failed for :type")
}


GO;

    }

    private function renderMainStructStart()
    {
        if ($this->propertyNames !== null) {
            return <<<'GO'
m:receiver := marshal:type(*:receiver)

GO;

        }
        return '';

    }

    private function renderMainStructEnd()
    {
        if ($this->propertyNames !== null) {
            // todo error is checked after constants, fix
            return <<<'GO'

*:receiver = :type(m:receiver)
GO;

        }
        return '';
    }
}