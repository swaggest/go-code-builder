<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\JsonSchema\Schema;

class GeneratedStruct
{
    /** @var StructDef */
    public $structDef;

    /** @var Schema */
    public $schema;

    /** @var string */
    public $path;

    /** @var MarshalJson */
    public $marshalJson;
}