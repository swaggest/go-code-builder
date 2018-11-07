<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Schema;

interface GoBuilderStructHook
{
    /**
     * @param StructDef $structDef
     * @param $path
     * @param Schema $schema
     * @return null
     */
    public function process(StructDef $structDef, $path, $schema);
}