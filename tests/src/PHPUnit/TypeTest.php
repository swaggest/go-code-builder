<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;


use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class TypeTest extends \PHPUnit_Framework_TestCase
{
    public function testSlice()
    {
        $this->assertSame('[]string', (new Slice(new Type('string')))->getTypeString());
        $this->assertSame('[]package.MyType',
            (new Slice(new Type('MyType', new Import('package'))))
                ->getTypeString()
        );

        $this->assertSame('[][]package.MyType',
            (new Slice(new Slice(new Type('MyType', new Import('package')))))
                ->getTypeString()
        );
    }

}