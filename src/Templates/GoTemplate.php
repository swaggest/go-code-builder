<?php

namespace Swaggest\GoCodeBuilder\Templates;


use Swaggest\CodeBuilder\AbstractTemplate;

abstract class GoTemplate extends AbstractTemplate
{
    protected $comment;

    public function setComment($comment)
    {
        $this->comment = $comment;
        return $this;
    }

    /**
     * @return string
     */
    public function getComment()
    {
        return $this->comment;
    }

    protected function renderComment()
    {
        if ($this->comment) {
            return rtrim($this->padLines('// ', $this->comment, false, true)) . "\n";
        } else {
            return '';
        }
    }

    public function padLines($with, $text, $skipFirst = true, $forcePad = false)
    {
        $lines = explode("\n", $text);
        $prevLine = '';
        foreach ($lines as $index => $line) {
            if ($skipFirst && !$index) {
                continue;
            }
            if ('' === trim($line) && $prevLine === '') {
                continue;
            }
            $prevLine = $line;
            if ($line || $forcePad) {
                $l = $with . $line;
                if (trim($l) === '') {
                    $l = '';
                }
                $lines[$index] = rtrim($l);
            }
        }
        return implode("\n", $lines);

    }

    public function stripEmptyLines($text)
    {
        $lines = explode("\n", $text);
        $prevLine = '';
        foreach ($lines as $index => $line) {
            if ('' === trim($line) && trim($prevLine) === '') {
                unset($lines[$index]);
                continue;
            }
            $prevLine = $line;
            if (trim($line) === '') {
                $lines[$index] = '';
            }
        }
        return implode("\n", $lines);
    }

    public function trim($s)
    {
        return trim($s);
    }

    public function escapeValue($value)
    {
        if (is_string($value)) {
            if (strpos($value, '"') === false && strpos($value, '\\') === false) {
                $value = '"' . addslashes($value) . '"';
            } elseif (strpos($value, '`')) {
                $value = '"' . addslashes($value) . '"';
            } else {
                $value = '`' . $value . '`';
            }
        }

        return $value;
    }

    public function ifThenElse($condition, $then, $else = '')
    {
        return $condition ? $then : $else;
    }


}