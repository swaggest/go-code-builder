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
            $propertyNames = '[]string{"' . implode('", "', $this->propertyNames) . '"}';
        }

        if ($this->patternProperties === null && $this->additionalPropertiesEnabled === null && $this->someOf === null) {
            return '';
        }

        $maps = '';
        $patternProperties = 'nil';

        if ($this->patternProperties !== null || $this->additionalPropertiesEnabled) {
            $patternProperties = 'map[string]interface{}{' . "\n";
            if ($this->patternProperties !== null) {
                foreach ($this->patternProperties as $regex => $patternPropertiesName) {
                    $patternProperties .= "\t`$regex`: &{$this->receiver()}.$patternPropertiesName,\n";
                    $maps .= ', i.' . $patternPropertiesName;
                }
            }

            if ($this->additionalPropertiesEnabled) {
                $patternProperties .= "\t``: &{$this->receiver()}.AdditionalProperties,\n";

                $maps .= ', i.AdditionalProperties';
            }

            $patternProperties .= '}';

            $patternProperties = $this->padLines("\t\t", $patternProperties);
        }

        if ($patternProperties === 'nil' && $this->someOf === null) {
            return '';
        }

        if ($this->someOf !== null) {
            foreach ($this->someOf as $kind => $unionPropertyNames) {
                foreach ($unionPropertyNames as $propertyName) {
                    $maps .= ', i.' . $propertyName;
                }
            }
        }


        $result = <<<GO
// UnmarshalJSON decodes JSON
func (i *:type) UnmarshalJSON(data []byte) error {
	{$this->padLines("\t", $this->renderMainStructStart())}
	{$this->padLines("\t", $this->renderMayUnmarshalHead())}
	err := unmarshalUnion(
		{$this->renderMustUnmarshal()},
		{$this->renderMayUnmarshal()},
		$propertyNames,
		$patternProperties,
		data,
	)
	{$this->padLines("\t", $this->renderMayUnmarshalTail())}
	{$this->padLines("\t", $this->renderMainStructEnd())}
	return err
}
		
// MarshalJSON encodes JSON
func (i :type) MarshalJSON() ([]byte, error) {
	type p :type

	return marshalUnion(p(i)$maps)
}


GO;
        $this->builder->getCode()->addSnippet(MarshalMapHelper::make(), false, 'marshal_map_helper');

        return new PlaceholderString($result, [
            ':type' => $this->type->getType(),
        ]);
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
        $items = $this->mayUnmarshalItems();
        if ($items) {
            $result = '';
            foreach ($items as $i => $item) {
                $item = ltrim($item, '&');
                $result .= <<<GO
if mayUnmarshal[$i] == nil {
    $item = nil
}

GO;

            }
            return $result;
        }
        return '';
    }

    private function renderMayUnmarshalHead()
    {
        $items = $this->mayUnmarshalItems();
        if ($items) {
            $itemsString = implode(', ', $items);
            return <<<GO
mayUnmarshal := []interface{}{{$itemsString}}
GO;

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
type p :type

ii := p(*i)

GO;

        }
        return '';

    }

    private function renderMainStructEnd()
    {
        if ($this->propertyNames !== null) {
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