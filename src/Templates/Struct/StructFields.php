<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\CodeBuilder\TableRenderer;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class StructFields extends GoTemplate
{
    /** @var StructDef */
    private $struct;

    /**
     * StructType constructor.
     * @param StructDef $struct
     */
    public function __construct(StructDef $struct)
    {
        $this->struct = $struct;
    }

    protected function toString()
    {
        return <<<GO
type {$this->struct->getName()} struct {
{$this->renderProperties()}
}


GO;
    }

    private function renderProperties()
    {
        $rows = array();
        $properties = $this->struct->getProperties();
        if (empty($properties)) {
            return '';
        }
        foreach ($properties as $property) {
            if (null === $property->getName()) {
                $rows [] = array(
                    '1' => $property->getType()->render(),
                    '3' => $property->getTags()->render(),
                    '4' => $property->getComment() ? '// ' . $property->getComment() : ''
                );
            } else {
                $rows [] = array(
                    '1' => $property->getName(),
                    '2' => $property->getType()->render(),
                    '3' => $property->getTags()->render(),
                    '4' => $property->getComment() ? '// ' . $property->getComment() : ''
                );
                //$result .= "\t{$property->getName()}\t{$property->getType()->toString()}\t{$property->getTags()->toString()}\n";
            }
        }
        $result = TableRenderer::create(new \ArrayIterator($rows))->setColDelimiter(' ')->__toString();
        if ($result) {
            $result = $this->padLines("\t", $this->trimLines(substr($result, 0, -1)), false);
        }
        return $result;
    }


}