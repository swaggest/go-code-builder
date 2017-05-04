<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\Type;

use Swaggest\GoCodeBuilder\Templates\Type\TypeCast;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;

class CastTest extends \PHPUnit_Framework_TestCase
{
    public function testSame()
    {
        $this->assertSame('toVar = fromVar', (new TypeCast(
            TypeUtil::fromString('string'),
            TypeUtil::fromString('string'),
            'toVar',
            'fromVar'))
            ->render()
        );
    }


    public function testDualPointer()
    {
        $expected = <<<GO
toVar = fromVar
GO;

        $this->assertSame($expected, (new TypeCast(
            TypeUtil::fromString('*string'),
            TypeUtil::fromString('*string'),
            'toVar', 'fromVar'))
            ->render()
        );
    }


    public function testDereference()
    {
        $expected = <<<GO
if fromVar != nil {
	toVar = *fromVar
}
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('string'),
            TypeUtil::fromString('*string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }

    public function testDoubleDereference()
    {
        $expected = <<<GO
if fromVar != nil {
    tmp4c54 := *fromVar
	if tmp4c54 != nil {
		toVar = *tmp4c54
	}
}
GO;

        $this->assertSame($expected, (new TypeCast(
            TypeUtil::fromString('string'),
            TypeUtil::fromString('**string'),
            'toVar', 'fromVar'))
            ->render()
        );
    }

    public function testTripleDereference()
    {
        $expected = <<<GO
if fromVar != nil {
    tmp4c54 := *fromVar
	if tmp4c54 != nil {
	    tmpe220 := *tmp4c54
		if tmpe220 != nil {
			toVar = *tmpe220
		}
	}
}
GO;

        $this->assertSame($expected, (new TypeCast(
            TypeUtil::fromString('string'),
            TypeUtil::fromString('***string'),
            'toVar', 'fromVar'))->render()
        );
    }


    public function testToPointer()
    {
        $expected = <<<GO
toVar = &fromVar
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('*string'),
            TypeUtil::fromString('string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }

    public function testToDoublePointer()
    {
        $expected = <<<GO
tmp4c54 := &fromVar
toVar = &tmp4c54
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('**string'),
            TypeUtil::fromString('string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }

    public function testToTriplePointer()
    {
        $expected = <<<GO
tmp4c54 := &fromVar
tmpe220 := &tmp4c54
toVar = &tmpe220
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('***string'),
            TypeUtil::fromString('string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }


}