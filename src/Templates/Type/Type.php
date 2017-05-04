<?php

namespace Swaggest\GoCodeBuilder\Templates\Type;

use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

/**
 * Class Type
 * @package Swaggest\GoCodeBuilder\Templates
 */
class Type extends GoTemplate implements AnyType
{
    /** @var string */
    private $type;
    /** @var Import */
    private $import;

    public function __construct($type, Import $import = null)
    {
        $this->type = $type;
        $this->import = $import;
    }

    public function getType()
    {
        return $this->type;
    }

    /**
     * @return Import
     */
    public function getImport()
    {
        return $this->import;
    }

    protected function toString()
    {

        $prefix = '';
        if ($this->import) {
            if ($goFile = GoFile::getCurrentGoFile()) {
                if ($goFile->getImportPath() !== $this->import->name) {
                    $goFile->getImports()->add($this->import);
                    $prefix = $this->import->getReferencePrefix();
                }
            }
        }

        return $prefix . $this->type;
    }

    public function equals(Type $type) {
        if ($type->type !== $this->type) {
            return false;
        }
        if ($this->import !== null) {
            if ($type->import === null) {
                return false;
            }
            if ($type->import->name !== $this->import->name) {
                return false;
            }
            return true;
        } elseif ($type->import !== null) {
            return false;
        }
        return true;
    }
}