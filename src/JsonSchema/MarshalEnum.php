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
     * MarshalEnum constructor.
     * @param Type $type
     * @param array $enum
     */
    public function __construct(Type $type, Type $base, array $enum)
    {
        $this->type = $type;
        $this->enum = $enum;
        $this->base = $base;
    }


    protected function toString()
    {
        $result = <<<GO
// UnmarshalJSON decodes JSON
func (i *:type) UnmarshalJSON(data []byte) error {
    var ii :base
    err := json.Unmarshal(data, &ii)
    if err != nil {
        return err
    }
    v := :type(ii)
    for {
    {$this->padLines("\t", $this->renderChecks('v'))}
        return errors.New("unexpected value")
    }
    *i = v
	return nil
}
        
// MarshalJSON encodes JSON
func (i :type) MarshalJSON() ([]byte, error) {
    for {
    {$this->padLines("\t", $this->renderChecks())}
        return nil, errors.New("unexpected value")
    }
    return json.Marshal(:base(i))
}


GO;

        return new PlaceholderString($result, [
            ':type' => $this->type,
            ':base' => $this->base,
        ]);
    }

    private function renderChecks($var = 'i')
    {
        $checks = '';
        foreach ($this->enum as $name => $value) {
            $checks .= <<<GO
    if $var == $name {
        break
    }

GO;
        }
        return $checks;
    }

    private function renderEnumMarshal()
    {

        return <<<GO
for {
{$this->renderChecks()}
    return nil, errors.New("unexpected value")
}

GO;

    }

    private function renderEnumUnmarshal()
    {

        return <<<GO
for {
{$this->renderChecks('v')}
    return errors.New("unexpected value")
}

GO;

    }

}