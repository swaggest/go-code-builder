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

    /**
     * @return string
     */
    public function getPackage()
    {
        if ($this->defaultPackageName) {
            return $this->defaultPackageName;
        } else {
            $pathItems = explode('/', $this->name);
            /** @var string $path */
            $path = array_pop($pathItems);
            $pathItems = explode('.', $path);
            $path = array_shift($pathItems);
            return (string)$path;
        }
    }

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
            $pathItems = explode('/', $this->name);
            /** @var string $path */
            $path = array_pop($pathItems);
            $pathItems = explode('.', $path);
            $path = array_shift($pathItems);
            return $path . '.';
        }
    }
}