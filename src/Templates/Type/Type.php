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

    public function toString()
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


}