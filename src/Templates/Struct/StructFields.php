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
            $commentLines = explode("\n", trim($property->getComment()));
            $commentLine = count($commentLines) > 1 ? '' : $commentLines[0];
            if ($commentLine) {
                $commentLine = '// ' . $commentLine;
            }

            if (count($commentLines) > 1) {
                $rows [] = array();
                foreach ($commentLines as $line) {
                    $rows [] = array(
                        '1' => '',
                        '2' => '',
                        '3' => '',
                        '4' => '// ' . $line,
                    );
                }
            }

            if ($property->isEmbedded()) {
                $rows [] = array(
                    '1' => $property->getType()->render(),
                    '3' => $property->getTags()->render(),
                    '4' => $commentLine
                );
            } else {
                $rows [] = array(
                    '1' => $property->getName(),
                    '2' => $property->getType()->render(),
                    '3' => $property->getTags()->render(),
                    '4' => $commentLine
                );
            }

        }
        $result = TableRenderer::create(new \ArrayIterator($rows))->setColDelimiter(' ')->__toString();
        if ($result) {
            $result = $this->padLines("\t", $this->trimLines(substr($result, 0, -1)), false);
        }
        return $result;
    }


}