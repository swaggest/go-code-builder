<?php

namespace Swaggest\GoCodeBuilder;


class Import
{

    public function __construct($name, $alias = null, $defaultPackageName = null)
    {
        $this->name = $name;
        $this->alias = $alias;
        $this->defaultPackageName = $defaultPackageName;
    }

    public $name;
    public $alias;
    public $defaultPackageName;

    public function getReferencePrefix()
    {
        if ($this->alias) {
            if ($this->alias === '.') {
                return '';
            }
            return $this->alias . '.';
        } elseif ($this->defaultPackageName) {
            return $this->defaultPackageName . '.';
        } else {
            $path = explode('/', $this->name);
            $path = array_pop($path);
            $path = explode('.', $path);
            $path = array_shift($path);
            return $path . '.';
        }
    }
}