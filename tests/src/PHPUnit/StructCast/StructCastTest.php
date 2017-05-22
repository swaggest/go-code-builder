<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\StructCast;


use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Struct\StructCast;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;
use Swaggest\GoCodeBuilder\Templates\Type\Type;
use Swaggest\GoCodeBuilder\TypeCast\CastRegistry;
use Swaggest\GoCodeBuilder\TypeCast\PropertyCast;

class StructCastTest extends \PHPUnit_Framework_TestCase
{
    public function testStructCast()
    {
        $base = new StructDef('Base');
        $base->addProperty(new StructProperty('Id', new Type('int64')));

        $derived = new StructDef('Derived');
        $derived->addProperty(new StructProperty('ID', new Pointer(new Type('uint64'))));

        $cast = new StructCast($base, $derived, array('Id' => 'ID'));

        $this->assertSame(<<<GO
func (base Base) MapTo() Derived {
	result := Derived{}
	var tmpb01a uint64
	tmpb01a = uint64(base.Id)
	result.ID = &tmpb01a
	return result
}


GO
            , $cast->getMapTo()->render());

        $this->assertSame(<<<GO
func (base Base) LoadFrom(derived Derived)  {
	if derived.ID != nil {
		base.Id = int64(*derived.ID)
	}

}


GO
            , $cast->getLoadFrom()->render());

    }

    public function testSubStruct()
    {
        $base = new StructDef('Base');
        $base->addProperty(new StructProperty('Id', new Type('int64')));
        $baseBody = new StructDef('BaseBody');
        $baseBody->addProperty(new StructProperty('Name', new Type('string')));
        $base->addProperty(new StructProperty('Body', new Slice(new Pointer($baseBody->getType()))));
        $base->addProperty(new StructProperty('BodyMap', new Map(new Pointer(new Type('string')), new Pointer($baseBody->getType()))));

        $derived = new StructDef('Derived');
        $derived->addProperty(new StructProperty('ID', new Pointer(new Type('uint64'))));
        $derivedBody = new StructDef('DerivedBody');
        $derivedBody->addProperty(new StructProperty('Name', new Pointer(new Type('string'))));
        $derived->addProperty(new StructProperty('Body', new Slice($derivedBody->getType())));
        $derived->addProperty(new StructProperty('BodyMap', new Map(new Type('string'), $derivedBody->getType())));


        $veryBase = new StructDef('VeryBase');
        $veryBase->addProperty(new StructProperty('Inner', new Pointer($base->getType())));


        $castRegistry = new CastRegistry();
        $castRegistry->addStructCast(new StructCast($baseBody, $derivedBody, array('Name' => 'Name'), $castRegistry));

        $directCast = new StructCast($base, $derived, array('Id' => 'ID', 'Body' => 'Body', 'BodyMap' => 'BodyMap'), $castRegistry);
        $castRegistry->addStructCast($directCast);

        $propertyCast = new PropertyCast($veryBase, 'Inner', $derived->getType(), $castRegistry);

        $file = new GoFile('main');
        $fileCode = $file->getCode();
        $fileCode->addSnippet(<<<GO
import "encoding/json"

func main() {
	println("BITCH")
	s := "LaLa"
	b := Base{
		Id: 123,
		Body: []*BaseBody{
			{Name: "John"},
		},
		BodyMap: map[*string]*BaseBody{
			&s: {Name: "LoLo"},
		},
	}
	vb := VeryBase{
		Inner: &b,
	}
	d := vb.MapTo()
	dd, _ := json.MarshalIndent(d, "", "  ")
	println(string(dd))
	vb.LoadFrom(d)
}


GO
        );
        $fileCode->addSnippet($base);
        $fileCode->addSnippet($derived);
        $fileCode->addSnippet($veryBase);
        $fileCode->addSnippet($baseBody);
        $fileCode->addSnippet($derivedBody);


        $fileCode->addSnippet($propertyCast->getMapTo());
        $fileCode->addSnippet($propertyCast->getLoadFrom());

        //$fileCode->addSnippet($directCast->getMapTo());
        //$fileCode->addSnippet($directCast->getLoadFrom());

        foreach ($castRegistry->resetUsedCastFuncs() as $funcDef) {
            $fileCode->addSnippet($funcDef);
        }


        $expected = <<<GO
package main



import "encoding/json"

func main() {
	println("BITCH")
	s := "LaLa"
	b := Base{
		Id: 123,
		Body: []*BaseBody{
			{Name: "John"},
		},
		BodyMap: map[*string]*BaseBody{
			&s: {Name: "LoLo"},
		},
	}
	vb := VeryBase{
		Inner: &b,
	}
	d := vb.MapTo()
	dd, _ := json.MarshalIndent(d, "", "  ")
	println(string(dd))
	vb.LoadFrom(d)
}

type Base struct {
	Id      int64
	Body    []*BaseBody
	BodyMap map[*string]*BaseBody
}

type Derived struct {
	ID      *uint64
	Body    []DerivedBody
	BodyMap map[string]DerivedBody
}

type VeryBase struct {
	Inner *Base
}

type BaseBody struct {
	Name string
}

type DerivedBody struct {
	Name *string
}

func (base VeryBase) MapTo() Derived {
	var result Derived
	if base.Inner != nil {
		tmp91cc := *base.Inner
		result = tmp91cc.MapTo()
	}
	return result
}

func (base VeryBase) LoadFrom(derived Derived)  {
	if base.Inner == nil {
	    base.Inner = new(Base)
	}
	base.Inner.LoadFrom(derived)
}

func (base BaseBody) MapTo() DerivedBody {
	result := DerivedBody{}
	result.Name = &base.Name
	return result
}

func (base Base) MapTo() Derived {
	result := Derived{}
	var tmpb01a uint64
	tmpb01a = uint64(base.Id)
	result.ID = &tmpb01a
	result.Body = make([]DerivedBody, len(base.Body))
	for index61c6, val61c6 := range base.Body {
		if val61c6 != nil {
			tmp43ba := *val61c6
			result.Body[index61c6] = tmp43ba.MapTo()
		}
	}
	result.BodyMap = make(map[string]DerivedBody, len(base.BodyMap))
	for key1cdb := range base.BodyMap {
		var toKey1cdb string
		if key1cdb != nil {
			toKey1cdb = *key1cdb
		}
		if base.BodyMap[key1cdb] != nil {
			tmpc029 := *base.BodyMap[key1cdb]
			result.BodyMap[toKey1cdb] = tmpc029.MapTo()
		}
	}
	return result
}

func (base BaseBody) LoadFrom(derived DerivedBody)  {
	if derived.Name != nil {
		base.Name = *derived.Name
	}

}

func (base Base) LoadFrom(derived Derived)  {
	if derived.ID != nil {
		base.Id = int64(*derived.ID)
	}
	base.Body = make([]*BaseBody, len(derived.Body))
	for index0922, val0922 := range derived.Body {
		if base.Body[index0922] == nil {
		    base.Body[index0922] = new(BaseBody)
		}
		base.Body[index0922].LoadFrom(val0922)
	}
	base.BodyMap = make(map[*string]*BaseBody, len(derived.BodyMap))
	for keyf30f := range derived.BodyMap {
		var toKeyf30f *string
		toKeyf30f = &keyf30f
		if base.BodyMap[toKeyf30f] == nil {
		    base.BodyMap[toKeyf30f] = new(BaseBody)
		}
		base.BodyMap[toKeyf30f].LoadFrom(derived.BodyMap[keyf30f])
	}

}
GO;

        $this->assertSame($expected, $file->render());
    }

}