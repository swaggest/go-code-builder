<?php

namespace Swaggest\GoCodeBuilder\Templates;

use Swaggest\GoCodeBuilder\Import;

class Imports extends GoTemplate
{
    /**
     * @var Import[]
     */
    public $imports = array();

    /**
     * @param Import $import
     * @return $this
     */
    public function add(Import $import)
    {
        $this->imports[$import->name] = $import;
        return $this;
    }

    public function addByName($name, $alias = null)
    {
        $this->add(new Import($name, $alias));
        return $this;
    }

    public function demand()
    {
        if (!$goFile = GoFile::getCurrentGoFile()) {
            return;
        }
        $imports = $goFile->getImports();

        foreach ($this->imports as $import) {
            $imports->add($import);
        }
    }

    public function toString()
    {
        $result = '';
        $builtin = array();
        $external = array();
        foreach ($this->imports as $import) {
            $path = explode('/', $import->name);
            if (strpos($path[0], '.') === false) {
                $builtin[] = $import;
            } else {
                $external[] = $import;
            }
        }

        asort($builtin);
        asort($external);

        foreach ($builtin as $import) {
            $result .= "\t" . ($import->alias ? $import->alias . ' ' : '') . '"' . $import->name . '"' . "\n";
        }

        if ($result && $external) {
            $result .= "\n";
        }
        foreach ($external as $import) {
            $result .= "\t" . ($import->alias ? $import->alias . ' ' : '') . '"' . $import->name . '"' . "\n";
        }

        if ($result) {
            $result = <<<GO
import (
$result)
GO;
        }

        return $result;
    }
}