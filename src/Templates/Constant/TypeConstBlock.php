<?php

namespace Swaggest\GoCodeBuilder\Templates\Constant;

use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class TypeConstBlock extends GoTemplate
{
    /** @var Type */
    private $type;

    /** @var array */
    private $values;

    private $comments;

    public function __construct(Type $type)
    {
        $this->type = $type;
    }

    public function addValue($name, $value, $comment = null)
    {
        $this->values[$name] = $value;
        if ($comment !== null) {
            $this->comments[$name] = $comment;
        }
        return $this;
    }

    protected function toString()
    {
        if ($this->values === null) {
            return '';
        }

        $result = '';
        foreach ($this->values as $name => $value) {
            if (!is_scalar($value)) {
                continue;
            }
            $value = $this->escapeValue($value);

            if (isset($this->comments[$name])) {
                $result .= "\t//" . $name . ' ' . $this->comments[$name] . "\n";
            }
            $result .= "\t" . $name . ' = ' . ':type(' . $value . ')' . "\n";
        }
        if (!$result) {
            return '';
        }
        $result =
            '// :type values enumeration.' . "\n" .
            'const (' . "\n" . $result . ')' . "\n\n";

        return new PlaceholderString($result, [
            ':type' => $this->type,
        ]);
    }

    public function getValues()
    {
        return $this->values;
    }

}