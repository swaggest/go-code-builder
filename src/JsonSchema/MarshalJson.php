<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\Map;

class MarshalJson extends GoTemplate
{
    private $type;
    private $builder;

    /** @var string[] */
    private $propertyNames;

    /** @var string[] */
    public $distinctNullNames = [];

    private $additionalPropertiesEnabled = false;

    /** @var StructProperty */
    private $additionalProperties;

    /** @var StructProperty[] */
    private $patternProperties;

    private $someOf;

    private $code;

    /** @var array */
    public $constValues;

    public function __construct(GoBuilder $builder, StructDef $type)
    {
        $this->type = $type;
        $this->builder = $builder;
        $this->code = new Code();
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
        return $this->propertyNames === null ? 'i' : 'ii';
    }

    protected function toString()
    {
        if ($this->patternProperties === null
            && $this->additionalPropertiesEnabled === false
            && $this->someOf === null
            && $this->constValues === null) {
            return '';
        }

        $result = '';

        if ($this->propertyNames !== null) {
            $result .= <<<'GO'
type marshal:type :type


GO;
        }

        $result .= <<<GO
{$this->padLines('',
            $this->renderUnmarshal()
            . $this->renderMarshal())}

GO;
        if (!$this->builder->options->skipMarshal) {
            if (null === $this->builder->marshalUnion) {
                $this->builder->marshalUnion = new MarshalUnion();
                $this->builder->getCode()->addSnippet($this->builder->marshalUnion, false, 'marshal_union');
            }
        }
        if (!$this->builder->options->skipUnmarshal) {
            $this->builder->getCode()->addSnippet($this->builder->unmarshalUnion, false, 'unmarshal_union');
            if ($this->patternProperties !== null) {
                $this->builder->unmarshalUnion->withPatternProperties = true;
            }
        }

        $this->code->addSnippet(new PlaceholderString($result, [
            ':type' => $this->type->getType(),
        ]));

        return $this->code;
    }

    private function renderConstRawMessage()
    {
        if ($this->constValues !== null) {
            $result = <<<GO
var (
	// const:type is unconditionally added to JSON.
	const:type = json.RawMessage({$this->escapeValue(json_encode($this->constValues))})
)


GO;

            return new PlaceholderString($result, [
                ':type' => $this->type->getType(),
            ]);

        }
        return '';
    }


    private function renderIgnoreKeys()
    {

        if (
            $this->propertyNames !== null &&
            ($this->patternProperties || $this->additionalPropertiesEnabled)
        ) {
            // All known property names (including constant values) has to be removed from
            // pattern and additional properties matching.
            $ignoreKeys = $this->propertyNames;
            if (!empty($this->constValues)) {
                foreach ($this->constValues as $propertyName => $value) {
                    $ignoreKeys[] = $propertyName;
                }
            }
            return <<<GO
var ignoreKeys:type = []string{
	"{$this->padLines("\t", implode("\",\n\"", $ignoreKeys))}",
}


GO;
        }

        return '';
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
        $mayUnmarshal = $this->renderMayUnmarshal();
        $withIgnoreKeys = false; // TODO move to renderer

        if (
            $this->propertyNames !== null &&
            ($this->patternProperties || $this->additionalPropertiesEnabled)
        ) {
            $withIgnoreKeys = true;
        }

        $mapUnmarshal = '';
        if ($this->patternProperties !== null
            || $this->additionalPropertiesEnabled
            || $this->constValues !== null
            || !empty($this->distinctNullNames)) {
            $mapUnmarshal = <<<'GO'


var m map[string]json.RawMessage

err = json.Unmarshal(data, &m)
if err != nil {
    m = nil
}

GO;

        }

        if ($this->constValues !== null) {
            $this->code->imports()->addByName('fmt');
            foreach ($this->constValues as $name => $value) {
                $mapUnmarshal .= <<<GO

if v, ok := m[{$this->escapeValue($name)}]; !ok || string(v) != {$this->escapeValue(json_encode($value))} {
	return fmt.Errorf({$this->escapeValue('bad or missing const value for "' . $name . '" (' . json_encode($value) . ' expected, %s received)')}, v)
}

delete(m, {$this->escapeValue($name)})


GO;
            }
        }

        foreach ($this->distinctNullNames as $goPropertyName => $name) {
            $name = $this->escapeValue($name);
            $mapUnmarshal .= <<<GO

if ii.$goPropertyName == nil {
    if _, ok := m[$name]; ok {
        var v interface{}
        ii.$goPropertyName = &v
    }
}

GO;

        }

        if ($mapUnmarshal && $withIgnoreKeys) {
            $mapUnmarshal .= <<<'GO'

for _, key := range ignoreKeys:type {
    delete(m, key)
}


GO;

        }


        if ($this->patternProperties !== null) {
            $mapUnmarshal .= <<<GO

for key, rawValue := range m {
    matched := false
    
GO;

            foreach ($this->patternProperties as $regex => $patternProperty) {
                $regexName = $this->builder->unmarshalUnion->patternVarName($regex);
                $mapType = $patternProperty->getType();
                $itemType = 'interface{}';
                if ($mapType instanceof Map) {
                    $itemType = $mapType->getValueType()->render();
                }
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
        delete(m, key)
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

            $mapUnmarshal .= <<<GO

for key, rawValue := range m {
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


        $funcBody = <<<GO
var err error

{$this->renderMainStructStart()}{$mustUnmarshal}{$mayUnmarshal}{$mapUnmarshal}
{$this->renderMainStructEnd()}

return nil
GO;


        return <<<GO
{$this->renderIgnoreKeys()}// UnmarshalJSON decodes JSON.
func (i *:type) UnmarshalJSON(data []byte) error {
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

        if ($this->constValues !== null) {
            $maps .= ', const:type';
        }

        if ($this->propertyNames !== null) {
            $maps .= ', marshal:type(i)';
        }

        if ($this->patternProperties !== null) {
            foreach ($this->patternProperties as $regex => $patternProperty) {
                $maps .= ', i.' . $patternProperty->getName();
            }
        }

        if ($this->additionalPropertiesEnabled) {
            $maps .= ', i.' . $this->additionalProperties->getName();
        }

        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                foreach ($unionPropertyNames as $propertyName) {
                    $maps .= ', i.' . $propertyName;
                }
            }
        }

        $maps = substr($maps, 2);

        return <<<GO
{$this->renderConstRawMessage()}// MarshalJSON encodes JSON.
func (i :type) MarshalJSON() ([]byte, error) {
	return marshalUnion($maps)
}

GO;


    }

    private function renderMustUnmarshal()
    {
        $result = '';
        if ($this->propertyNames) {
            $result .= <<<'GO'


err = json.Unmarshal(data, &ii)
if err != nil {
    return err
}

GO;
        }

        if (isset($this->someOf['allOf'])) {
            foreach ($this->someOf['allOf'] as $propertyName) {
                $result .= <<<GO


err = json.Unmarshal(data, &{$this->receiver()}.{$propertyName})
if err != nil {
    return err
}

GO;
            }
        }

        return $result;
    }

    private function renderMayUnmarshal()
    {
        $result = '';
        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                if ($kind === 'allOf') {
                    continue;
                }

                foreach ($unionPropertyNames as $propertyName) {
                    $result .= <<<GO


err = json.Unmarshal(data, &{$this->receiver()}.{$propertyName})
if err != nil {
    {$this->receiver()}.{$propertyName} = nil
}

GO;
                }
            }
        }
        return $result;
    }


    private function renderMainStructStart()
    {
        if ($this->propertyNames !== null) {
            return <<<'GO'
ii := marshal:type(*i)

GO;

        }
        return '';

    }

    private function renderMainStructEnd()
    {
        if ($this->propertyNames !== null) {
            // todo error is checked after constants, fix
            return <<<'GO'

*i = :type(ii)
GO;

        }
        return '';
    }
}