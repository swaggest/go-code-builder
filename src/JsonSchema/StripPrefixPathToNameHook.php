<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

class StripPrefixPathToNameHook implements GoBuilderPathToNameHook
{
    public $prefixes = [
        '#/definitions',
    ];

    function pathToName($path)
    {
        foreach ($this->prefixes as $prefix) {
            if (0 === strpos($path, $prefix)) {
                return substr($path, strlen($prefix));
            }
        }
        return $path;
    }
}