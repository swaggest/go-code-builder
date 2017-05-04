<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\Type;

use Swaggest\GoCodeBuilder\Templates\Struct\StructCast;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;

class StructCastTest extends \PHPUnit_Framework_TestCase
{
    public function testMapTo()
    {
        $base = new StructDef('Base');
        $base->addProperty(new StructProperty('BInt64', TypeUtil::fromString('*int64')));
        $base->addProperty(new StructProperty('BString', TypeUtil::fromString('string')));


        $derived = new StructDef('Derived');
        $derived->addProperty(new StructProperty('DUInt32', TypeUtil::fromString('uint32')));
        $derived->addProperty(new StructProperty('DString', TypeUtil::fromString('*string')));

        $cast = new StructCast($base, $derived);
        $cast->setPropMap('BInt64', 'DUInt32');
        $cast->setPropMap('BString', 'DString');

        $mapTo = $cast->getMapTo();
        $expected = <<<GO
func (base Base) MapTo() Derived {
	result := Derived{}
	if base.BInt64 != nil {
		result.DUInt32 = uint32(*base.BInt64)
	}
	result.DString = &base.BString
	return result
}


GO;

        $obtained = $mapTo->render();
        $this->assertSame($expected, $obtained);

        $loadFrom = $cast->getLoadFrom();
        $expected = <<<GO
func (base *Base) LoadFrom(derived Derived)  {
	var tmp4c90 int64
	tmp4c90 = int64(derived.DUInt32)
	base.BInt64 = &tmp4c90
	if derived.DString != nil {
		base.BString = *derived.DString
	}

}


GO;

        $obtained = $loadFrom->render();
        $this->assertSame($expected, $obtained);
    }

}