<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\Type;


use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;
use Swaggest\GoCodeBuilder\Templates\Type\Type;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;

class EqualsTest extends \PHPUnit_Framework_TestCase
{
    public function testType()
    {
        $this->assertTrue(TypeUtil::equals(new Type('string'), new Type('string')));
        $this->assertFalse(TypeUtil::equals(new Type('string'), new Type('int64')));

        $this->assertFalse(TypeUtil::equals(new Type('String', new Import('mypkg')), new Type('String', new Import('hispkg'))));
        $this->assertTrue(TypeUtil::equals(new Type('String', new Import('mypkg')), new Type('String', new Import('mypkg'))));
    }

    public function testSlice()
    {
        $this->assertTrue(TypeUtil::equals(new Slice(new Type('string')), new Slice(new Type('string'))));

        $this->assertFalse(TypeUtil::equals(new Slice(new Type('string')), new Slice(new Slice(new Type('string')))));
    }
}