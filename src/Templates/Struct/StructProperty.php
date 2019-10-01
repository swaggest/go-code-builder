<?php

namespace Swaggest\GoCodeBuilder\Templates\Struct;

use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Type\AnyType;

class StructProperty extends GoTemplate
{
    /** @var string|null */
    private $name;
    private $type;
    private $tags;

    public function __construct($name, AnyType $type, Tags $tags = null)
    {
        $this->name = $name;
        $this->type = $type;
        if ($tags === null) {
            $tags = new Tags();
        }
        $this->tags = $tags;
    }

    /**
     * @return string|null
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * @return AnyType
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * @return Tags
     */
    public function getTags()
    {
        return $this->tags;
    }

    protected function toString()
    {
        // no op
        return '';
    }


}