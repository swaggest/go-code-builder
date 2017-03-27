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


}