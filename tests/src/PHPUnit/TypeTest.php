<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;


use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class TypeTest extends \PHPUnit_Framework_TestCase
{
    public function testSlice()
    {
        $this->assertSame('[]string', (new Type\Slice('string'))->toString());
        $this->assertSame('[]package.MyType',
            (new Type\Slice('MyType', new Import('package')))
                ->toString()
        );

        $this->assertSame('[][]package.MyType',
            (new Type\Slice(new Type\Slice('MyType', new Import('package'))))
                ->toString()
        );
    }

}