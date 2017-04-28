<?php

namespace Swaggest\GoCodeBuilder\Templates;

use Swaggest\CodeBuilder\ClosureString;

class Code extends GoTemplate
{
    /** @var Imports */
    private $imports;

    public $snippets;

    public function __construct($body = null)
    {
        if ($body != null) {
            $this->addSnippet($body);
        }
    }

    public function addSnippet($code, $prepend = false)
    {
        if ($prepend) {
            array_unshift($this->snippets, $code);
        } else {
            $this->snippets[] = $code;
        }
        return $this;
    }

    /**
     * @return Imports
     */
    public function imports()
    {
        if (null === $this->imports) {
            $this->imports = new Imports();
        }
        return $this->imports;
    }

    protected function toString()
    {
        $this->imports()->demand();
        $result = '';
        if ($this->snippets === null) {
            return '';
        }
        foreach ($this->snippets as $code) {
            if ($code instanceof ClosureString) {
                $result .= $code->render();
            } else {
                $result .= $code;
            }
        }
        return $result;
    }
}