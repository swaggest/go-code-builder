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
            return rtrim($this->padLines('// ', $this->comment, false)) . "\n";
        } else {
            return '';
        }
    }

    public function padLines($with, $text, $skipFirst = true, $forcePad = false)
    {
        $lines = explode("\n", $text);
        foreach ($lines as $index => $line) {
            if ($skipFirst && !$index) {
                continue;
            }
            if ($line || $forcePad) {
                $l = $with . $line;
                if (trim($l) === '') {
                    $l = '';
                }
                $lines[$index] = $l;
            }
        }
        return implode("\n", $lines);

    }


}