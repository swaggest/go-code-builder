<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class Tags extends GoTemplate
{
    private $items;

    public function setTag($key, $value)
    {
        $this->items[$key] = $value;
        return $this;
    }

    public function toString()
    {
        $result = '';
        if ($this->items === null) {
            return $result;
        }
        foreach ($this->items as $key => $value) {
            $result .= "$key:\"$value\" ";
        }
        if ($result) {
            $result = '`' . substr($result, 0, -1) . '`';
        }
        return $result;
    }

}