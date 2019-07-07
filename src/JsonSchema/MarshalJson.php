<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\Code;
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

    private $code;

    /** @var array */
    public $constValues;

    public function __construct(GoBuilder $builder, StructDef $type)
    {
        $this->type = $type;
        $this->builder = $builder;
        $this->code = new Code();
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
        if ($this->patternProperties === null
            && $this->additionalPropertiesEnabled === false
            && $this->someOf === null
            && $this->constValues === null) {
            return '';
        }

        $result = <<<GO
type marshal:type :type

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


    private function renderUnmarshal()
    {
        if ($this->builder->options->skipUnmarshal) {
            return '';
        }

        if ($this->builder->unmarshalUnion === null) {
            $this->builder->unmarshalUnion = new UnmarshalUnion();
            $this->builder->unmarshalUnion->goBuilder = $this->builder;
        }

        $unionMap = '';
        $mustUnmarshal = $this->renderMustUnmarshal();
        if ($mustUnmarshal !== 'nil') {
            $this->builder->unmarshalUnion->withMustUnmarshal = true;
            $unionMap .= 'mustUnmarshal: ' . $mustUnmarshal . ",\n";
        }

        $mayUnmarshal = $this->renderMayUnmarshal();
        if ($mayUnmarshal !== 'nil') {
            $this->builder->unmarshalUnion->withMayUnmarshal = true;
            $unionMap .= 'mayUnmarshal: ' . $mayUnmarshal . ",\n";
        }

        if ($this->propertyNames !== null && ($this->patternProperties || $this->additionalPropertiesEnabled)) {
            $propertyNames = <<<GO
[]string{
	"{$this->padLines("\t", implode("\",\n\"", $this->propertyNames))}",
}
GO;
            $this->builder->unmarshalUnion->withIgnoreKeys = true;
            $unionMap .= 'ignoreKeys: ' . $propertyNames . ",\n";
        }

        if ($this->patternProperties !== null) {
            $patternProperties = 'map[*regexp.Regexp]interface{}{' . "\n";
            foreach ($this->patternProperties as $regex => $patternPropertiesName) {
                $regexName = $this->builder->unmarshalUnion->patternVarName($regex);
                $patternProperties .= "\t$regexName: &{$this->receiver()}.$patternPropertiesName, // $regex \n";
            }
            $patternProperties .= '}';

            $this->builder->unmarshalUnion->withPatternProperties = true;
            $unionMap .= 'patternProperties: ' . $patternProperties . ",\n";
        }

        if ($this->additionalPropertiesEnabled) {
            $this->builder->unmarshalUnion->withAdditionalProperties = true;
            $unionMap .= 'additionalProperties: ' . "&{$this->receiver()}.AdditionalProperties,\n";
        }


        return <<<GO
// UnmarshalJSON decodes JSON.
func (i *:type) UnmarshalJSON(data []byte) error {
	{$this->padLines("\t",
			$this->renderMainStructStart()
			. $this->renderMayUnmarshalHead()
		)}
	err := unionMap{
{$this->padLines("\t\t", $unionMap, false)}		jsonData: data,
	}.unmarshal()
	{$this->padLines("\t", $this->renderMayUnmarshalTail()
			. $this->renderMainStructEnd())}
	return err
}


GO;
    }

    private function renderMarshal()
    {
        if ($this->builder->options->skipMarshal) {
            return '';
        }

        $maps = '';

        if ($this->patternProperties !== null) {
            foreach ($this->patternProperties as $regex => $patternPropertiesName) {
                $maps .= ', i.' . $patternPropertiesName;
            }
        }

        if ($this->additionalPropertiesEnabled) {
            $maps .= ', i.AdditionalProperties';
        }

        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                foreach ($unionPropertyNames as $propertyName) {
                    $maps .= ', i.' . $propertyName;
                }
            }
        }

        if ($this->constValues !== null) {
            $maps .= ', const:type';
        }

        return <<<GO
{$this->renderConstRawMessage()}// MarshalJSON encodes JSON.
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
            $this->code->imports()->addByName('fmt');
            foreach ($this->constValues as $name => $value) {
                $result .= <<<GO
if v, ok := constValues[{$this->escapeValue($name)}]; !ok || string(v) != {$this->escapeValue(json_encode($value))} {
	return fmt.Errorf({$this->escapeValue('bad or missing const value for "' . $name . '" (' . json_encode($value) . ' expected, %v received)')}, v)
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