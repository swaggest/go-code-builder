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


    public function testFromSlicePointers()
    {
        $expected = <<<GO
toVar = make([]string, len(fromVar))
for index4c54, val4c54 := range fromVar {
	if val4c54 != nil {
		toVar[index4c54] = *val4c54
	}
}
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('[]string'),
            TypeUtil::fromString('[]*string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }

    public function testToSlicePointers()
    {
        $expected = <<<GO
toVar = make([]*string, len(fromVar))
for index4c54, val4c54 := range fromVar {
	toVar[index4c54] = &val4c54
}
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('[]*string'),
            TypeUtil::fromString('[]string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }

    public function testMap()
    {
        $expected = <<<GO
toVar = make(map[string]*string, len(fromVar))
for key, val := range fromVar {
	if key != nil {
		toKey = *key
	}
	toVar[toKey] = &fromVar[key]
}
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('map[string]*string'),
            TypeUtil::fromString('map[*string]string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);
    }


    public function testCastNumber()
    {
        $expected = <<<GO
toVar = int64(fromVar)
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('int64'),
            TypeUtil::fromString('float32'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);

    }


    public function testCastNumberFromPointer()
    {
        $expected = <<<GO
if fromVar != nil {
    tmp4c54 := *fromVar
	if tmp4c54 != nil {
	    tmpe220 := *tmp4c54
		if tmpe220 != nil {
			toVar = int64(*tmpe220)
		}
	}
}
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('int64'),
            TypeUtil::fromString('***float32'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);

    }


    public function testCastNumberFromToPointer()
    {
        $expected = <<<GO
var tmp4c6b int64
if fromVar != nil {
    tmp4c54 := *fromVar
	if tmp4c54 != nil {
	    tmpe220 := *tmp4c54
		if tmpe220 != nil {
		    tmp30c0 := *tmpe220
			if tmp30c0 != nil {
				tmp4c6b = int64(*tmp30c0)
			}
		}
	}
}
tmpd42e := &tmp4c6b
tmpe093 := &tmpd42e
tmp680c := &tmpe093
toVar = &tmp680c
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('****int64'),
            TypeUtil::fromString('****float32'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);

    }

    public function testCastComplex()
    {
        $expected = <<<GO
toVar = make(map[int64][][]*string, len(fromVar))
for key, val := range fromVar {
	if key != nil {
		toKey = int64(*key)
	}
	toVar[toKey] = make([][]*string, len(fromVar[key]))
	for index4ed4, val4ed4 := range fromVar[key] {
		toVar[toKey][index4ed4] = make([]*string, len(val4ed4))
		for index306a, val306a := range val4ed4 {
			toVar[toKey][index4ed4][index306a] = &val306a
		}
	}
}
GO;

        $tc = new TypeCast(
            TypeUtil::fromString('map[int64][][]*string'),
            TypeUtil::fromString('map[*float32][][]string'),
            'toVar', 'fromVar');

        $obtained = $tc->render();

        $this->assertSame($expected, $obtained);

    }


}