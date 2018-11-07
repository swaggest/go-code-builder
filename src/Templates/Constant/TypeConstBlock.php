<?php

namespace Swaggest\GoCodeBuilder\Templates\Constant;

use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class TypeConstBlock extends GoTemplate
{
    /** @var Type */
    private $type;

    /** @var Type */
    private $base;

    /** @var array */
    private $values;

    private $comments;

    public function __construct(Type $type, Type $base)
    {
        $this->type = $type;
        $this->base = $base;
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
            $value = self::prepareValue($value);

            if (isset($this->comments[$name])) {
                $result .= "\t//" . $name . ' ' . $this->comments[$name] . "\n";
            }
            $result .= "\t" . $name . ' = ' . ':type(' . $value . ')' . "\n";
        }
        if (!$result) {
            return '';
        }
        $result =
            '// :type values enumeration' . "\n" .
            'const (' . "\n" . $result . ')' . "\n\n";

        return new PlaceholderString($result, [
            ':type' => $this->type,
        ]);
    }

    public function getValues()
    {
        return $this->values;
    }

    public static function prepareValue($value)
    {
        if (is_string($value)) {
            if (strpos($value, '`')) {
                $value = '"' . addslashes($value) . '"';
            } else {
                $value = '`' . $value . '`';
            }
        }

        return $value;
    }
}