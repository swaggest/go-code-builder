<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\Type;


use Swaggest\GoCodeBuilder\Templates\GoFile;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;
use Swaggest\GoCodeBuilder\Templates\Type\Map;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;
use Swaggest\GoCodeBuilder\Templates\Type\Slice;

class TypeFromStringTest extends \PHPUnit_Framework_TestCase
{
    public function setUp()
    {
        GoFile::setCurrentGoFile(new GoFile('blabla'));
    }

    public function testPointer()
    {
        $typeString = '*blabla.com/some-api/internal/handler.swaggerFeaturesRequest';
        /** @var Pointer $type */
        $type = TypeUtil::fromString($typeString);
        $this->assertTrue($type instanceof Pointer);
        $this->assertSame('*handler.swaggerFeaturesRequest', $type->render());
        $this->assertSame('blabla.com/some-api/internal/handler', $type->getType()->getImport()->name);
    }

    public function testDefPackageName()
    {
        $type = TypeUtil::fromString('blabla.com/go-null-types.NullString::nulltypes.NullString');
        $this->assertSame('nulltypes.NullString', $type->render());
        $this->assertSame('blabla.com/go-null-types', $type->getImport()->name);
    }

    public function testMap()
    {
        /** @var Map $type */
        $type = TypeUtil::fromString('map[string]blabla.com/go-null-types.NullString::nulltypes.NullString');
        $this->assertSame('map[string]nulltypes.NullString', $type->render());
        $this->assertSame('string', $type->getKeyType()->render());
        $this->assertSame('nulltypes.NullString', $type->getValueType()->render());
    }

    public function testSlice()
    {
        /** @var Slice $type */
        $type = TypeUtil::fromString('[]blabla.com/go-null-types.NullString::nulltypes.NullString');
        $this->assertSame('[]nulltypes.NullString', $type->render());
        $this->assertSame('nulltypes.NullString', $type->getType()->render());
    }

    public function tearDown()
    {
        GoFile::setCurrentGoFile(null);
    }

}