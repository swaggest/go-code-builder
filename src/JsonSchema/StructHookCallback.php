<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;

class StructHookCallback implements GoBuilderStructHook
{
    /** @var \Closure */
    private $closure;

    /**
     * StructHookCallback constructor.
     * @param \Closure $closure
     */
    public function __construct(\Closure $closure)
    {
        $this->closure = $closure;
    }

    public function process(StructDef $structDef, $path, $schema)
    {
        $this->closure->__invoke($structDef, $path, $schema);
    }


}