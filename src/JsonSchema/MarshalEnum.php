<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class MarshalEnum extends GoTemplate
{
    /**
     * @var Type
     */
    private $type;

    /**
     * @var Type
     */
    private $base;

    /**
     * @var array
     */
    private $enum;

    /**
     * @var GoBuilder
     */
    private $builder;

    /**
     * MarshalEnum constructor.
     * @param Type $type
     * @param array $enum
     */
    public function __construct(Type $type, Type $base, array $enum, GoBuilder $builder)
    {
        $this->type = $type;
        $this->enum = $enum;
        $this->base = $base;
        $this->builder = $builder;
    }


    private function renderConstMarshal()
    {
        if ($this->builder->options->skipMarshal) {
            return '';
        }

        return <<<GO
// MarshalJSON encodes JSON.
func (i :type) MarshalJSON() ([]byte, error) {
	{$this->padLines("\t", $this->renderIfCheck('i', 'return nil, fmt.Errorf("unexpected :type value: %v", i)'))}
	return json.Marshal(:base(i))
}


GO;
    }

    private function renderConstUnmarshal()
    {
        if ($this->builder->options->skipUnmarshal) {
            return '';
        }

        return <<<GO
// UnmarshalJSON decodes JSON.
func (i *:type) UnmarshalJSON(data []byte) error {
	var ii :base
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := :type(ii)
	{$this->padLines("\t", $this->renderIfCheck('v', 'return fmt.Errorf("unexpected :type value: %v", v)'))}
	*i = v
	return nil
}
		

GO;

    }

    private function constToString()
    {
        $result = <<<GO
{$this->padLines('', $this->renderConstMarshal() . $this->renderConstUnmarshal())}
GO;

        return new PlaceholderString($result, [
            ':type' => $this->type,
            ':base' => $this->base,
        ]);

    }

    private function renderMarshal()
    {
        if ($this->builder->options->skipMarshal) {
            return '';
        }

        return <<<GO
// MarshalJSON encodes JSON.
func (i :type) MarshalJSON() ([]byte, error) {
	{$this->padLines("\t", $this->renderIfCheck('i', 'return nil, fmt.Errorf("unexpected :type value: %v", i)'))}
	return json.Marshal(:base(i))
}


GO;

    }


    private function renderUnmarshal()
    {
        if ($this->builder->options->skipUnmarshal) {
            return '';
        }

		return <<<GO
// UnmarshalJSON decodes JSON.
func (i *:type) UnmarshalJSON(data []byte) error {
	var ii :base
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := :type(ii)
	{$this->padLines("\t", $this->renderIfCheck('v', 'return fmt.Errorf("unexpected :type value: %v", v)'))}
	*i = v
	return nil
}


GO;

    }

    protected function toString()
    {
        if (count($this->enum) === 1) {
            return $this->constToString();
        }

        $result = <<<GO
{$this->padLines('', $this->renderMarshal() . $this->renderUnmarshal())}
GO;

        return new PlaceholderString($result, [
            ':type' => $this->type,
            ':base' => $this->base,
        ]);
    }

    private function renderIfCheck($var, $return)
    {
        if (empty($this->enum)) {
            return '';
        }
        $checks = '';
        foreach ($this->enum as $name => $value) {
            $checks .= "case $name:\n";
        }
        return <<<GO
switch $var {
$checks
default:
	$return
}

GO;
    }

}