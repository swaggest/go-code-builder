<?php

namespace Swaggest\GoCodeBuilder\Templates;

use Swaggest\CodeBuilder\AbstractTemplate;
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

    private $uniqueKeys = array();

    public function addSnippet($code, $prepend = false, $uniqueKey = null)
    {
        if ($uniqueKey) {
            if (isset($this->uniqueKeys[$uniqueKey])) {
                return $this;
            } else {
                $this->uniqueKeys[$uniqueKey] = $uniqueKey;
            }
        }

        if (null === $this->snippets) {
            $this->snippets = array();
        }
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
            } elseif ($code instanceof AbstractTemplate) {
                $result .= $code->render();
            } else {
                $result .= $code;
            }
        }
        return $result;
    }
}