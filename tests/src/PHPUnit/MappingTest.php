<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;


use Swaggest\GoCodeBuilder\Templates\Mapping\Mapping;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;

class MappingTest extends \PHPUnit_Framework_TestCase
{
    public function testMapping()
    {
        $mapping = new Mapping();
        $mapping->baseType = TypeUtil::fromString('base.MyStruct');
        $mapping->derivedType = TypeUtil::fromString('derived.MineStruct');




    }


    public function testMapString()
    {
        $mapping = new Mapping();
        $mapping->baseName = 'MyString';
        $mapping->baseType = TypeUtil::fromString('string');

        $mapping->derivedName = 'TheString';
        $mapping->derivedType = TypeUtil::fromString('*string');


    }

}