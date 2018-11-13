<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;

class MarshalJson extends GoTemplate
{
    private $type;
    private $builder;
    private $propertyNames;
    private $additionalPropertiesEnabled = false;
    private $patternProperties;
    private $someOf;

    /** @var array */
    public $constValues;

    public function __construct(GoBuilder $builder, StructDef $type)
    {
        $this->type = $type;
        $this->builder = $builder;
    }

    public function enableAdditionalProperties()
    {
        $this->additionalPropertiesEnabled = true;
        return $this;
    }

    public function addPatternProperty($regex, $name)
    {
        $this->patternProperties[$regex] = $name;
        return $this;
    }

    public function addSomeOf($kind, $name)
    {
        $this->someOf[$kind][$name] = $name;
        return $this;
    }

    public function addNamedProperty($name)
    {
        $this->propertyNames [] = $name;
        return $this;
    }

    private function receiver()
    {
        return $this->propertyNames === null ? 'i' : 'ii';
    }

    protected function toString()
    {
        $propertyNames = 'nil';
        if ($this->propertyNames !== null) {
            $propertyNames = <<<GO
[]string{
	"{$this->padLines("\t", implode("\",\n\"", $this->propertyNames))}",
}
GO;


        }

        if ($this->patternProperties === null
            && $this->additionalPropertiesEnabled === null
            && $this->someOf === null
            && $this->constValues === null) {
            return '';
        }

        $maps = '';
        $patternProperties = 'nil';

        if ($this->patternProperties !== null || $this->additionalPropertiesEnabled) {
            $patternProperties = 'map[string]interface{}{' . "\n";
            if ($this->patternProperties !== null) {
                foreach ($this->patternProperties as $regex => $patternPropertiesName) {
                    $patternProperties .= "\t{$this->escapeValue($regex)}: &{$this->receiver()}.$patternPropertiesName,\n";
                    $maps .= ', i.' . $patternPropertiesName;
                }
            }

            if ($this->additionalPropertiesEnabled) {
                $patternProperties .= "\t\"\": &{$this->receiver()}.AdditionalProperties,\n";

                $maps .= ', i.AdditionalProperties';
            }

            $patternProperties .= '}';

            $patternProperties = $this->padLines("\t\t", $patternProperties);
        }

        if ($patternProperties === 'nil'
            && $this->someOf === null
            && $this->constValues === null) {
            return '';
        }

        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                foreach ($unionPropertyNames as $propertyName) {
                    $maps .= ', i.' . $propertyName;
                }
            }
        }

        if ($this->constValues !== null) {
            $maps .= ', ptr:type';
        }


        $result = <<<GO
type marshal:type :type

{$this->padLines('',
            $this->renderUnmarshal($propertyNames, $patternProperties)
            . $this->renderMarshal($maps))}

GO;
        if (!$this->builder->options->skipMarshal) {
            if (null === $this->builder->marshalUnion) {
                $this->builder->marshalUnion = new MarshalUnion();
                $this->builder->getCode()->addSnippet($this->builder->marshalUnion, false, 'marshal_union');
            }
        }
        if (!$this->builder->options->skipUnmarshal) {
            if (null === $this->builder->unmarshalUnion) {
                $this->builder->unmarshalUnion = new UnmarshalUnion();
                $this->builder->getCode()->addSnippet($this->builder->unmarshalUnion, false, 'unmarshal_union');
            }
            if ($this->patternProperties !== null) {
                $this->builder->unmarshalUnion->withRegex = true;
            }
        }

        return new PlaceholderString($result, [
            ':type' => $this->type->getType(),
        ]);
    }

    private function renderConstRawMessage()
    {
        if ($this->constValues !== null) {
            $result = <<<GO
var (
	// const:type is unconditionally added to JSON
	const:type = json.RawMessage([]byte({$this->escapeValue(json_encode($this->constValues))}))
	ptr:type = &const:type
)


GO;

            return new PlaceholderString($result, [
                ':type' => $this->type->getType(),
            ]);

        }
        return '';
    }


    private function renderUnmarshal($propertyNames, $patternProperties)
    {
        if ($this->builder->options->skipUnmarshal) {
            return '';
        }

        return <<<GO
// UnmarshalJSON decodes JSON
func (i *:type) UnmarshalJSON(data []byte) error {
	{$this->padLines("\t",
            $this->renderMainStructStart()
            . $this->renderMayUnmarshalHead()
        )}
	err := unmarshalUnion(
		{$this->renderMustUnmarshal()},
		{$this->renderMayUnmarshal()},
		{$this->padLines("\t\t", $propertyNames)},
		$patternProperties,
		data,
	)
	{$this->padLines("\t", $this->renderMayUnmarshalTail()
            . $this->renderMainStructEnd())}
	return err
}


GO;

    }

    private function renderMarshal($maps)
    {
        if ($this->builder->options->skipMarshal) {
            return '';
        }

        return <<<GO
{$this->renderConstRawMessage()}// MarshalJSON encodes JSON
func (i :type) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshal:type(i)$maps)
}

GO;


    }

    private function renderMustUnmarshal()
    {
        $items = array();
        if ($this->propertyNames) {
            $items[] = '&ii';
        }
        if (isset($this->someOf['allOf'])) {
            foreach ($this->someOf['allOf'] as $propertyName) {
                $items [] = "&{$this->receiver()}." . $propertyName;
            }
        }
        if ($items) {
            $itemsString = implode(', ', $items);
            return <<<GO
[]interface{}{{$itemsString}}
GO;

        }
        return 'nil';
    }

    private function renderMayUnmarshalTail()
    {
        $result = '';
        if ($this->constValues !== null) {
            foreach ($this->constValues as $name => $value) {
                $result .= <<<GO
if v, ok := constValues[{$this->escapeValue($name)}];!ok || string(v) != {$this->escapeValue(json_encode($value))} {
    return errors.New({$this->escapeValue('bad or missing const value for "' . $name . '" (' . json_encode($value) . ' expected)')})
}

GO;

            }
        }

        $items = $this->mayUnmarshalItems();
        if ($items) {
            foreach ($items as $i => $item) {
                $item = ltrim($item, '&');
                if ($item === 'constValues') {
                    continue;
                }
                $result .= <<<GO
if mayUnmarshal[$i] == nil {
	$item = nil
}

GO;

            }
            return $result;
        }
        return $result;
    }

    private function renderMayUnmarshalHead()
    {
        $items = $this->mayUnmarshalItems();
        if ($items) {
            $itemsString = implode(', ', $items);
            $result = '';
            if ($this->constValues !== null) {
                $result .= <<<'GO'
constValues := make(map[string]json.RawMessage)

GO;

            }
            $result .= <<<GO
mayUnmarshal := []interface{}{{$itemsString}}
GO;
            return $result;

        }
        return '';
    }

    private function mayUnmarshalItems()
    {
        $items = array();
        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                if ($kind === 'allOf') {
                    continue;
                }
                foreach ($unionPropertyNames as $propertyName) {
                    $items [] = "&{$this->receiver()}." . $propertyName;
                }
            }
        }
        if ($this->constValues !== null) {
            $items [] = "&constValues";
        }
        return $items;
    }


    private function renderMayUnmarshal()
    {
        if ($this->mayUnmarshalItems()) {
            return <<<GO
mayUnmarshal
GO;

        }
        return 'nil';
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
if err != nil {
	return err
}
*i = :type(ii)
GO;

        }
        return '';
    }
}