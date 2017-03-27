<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;


use Swaggest\GoCodeBuilder\Templates\Struct\Tags;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class StructTest extends \PHPUnit_Framework_TestCase
{
    public function testStructDef()
    {
        $struct = new StructDef('Sample', 'Sample is a sample structure');
        $struct
            ->addProperty(
                new StructProperty(
                    'PropOne',
                    new Type('string'),
                    (new Tags())->setTag('json', 'prop_one')
                )
            )->addProperty(
                new StructProperty(
                    'PropTwo',
                    new Type('int64'),
                    (new Tags())
                        ->setTag('schema', 'prop_two')
                        ->setTag('description', 'The second property')
                )
            );


        $this->assertSame(<<<GO
// Sample is a sample structure
type Sample struct {
	PropOne	string	`json:"prop_one"`
	PropTwo	int64	`schema:"prop_two" description:"The second property"`
}
GO
            , $struct->toString());
    }

    public function testAnonProperty()
    {
        $struct = new StructDef('Sample', 'Sample is a sample structure');
        $struct
            ->addProperty(
                new StructProperty(
                    null,
                    new Type('MyTypeOne'),
                    (new Tags())->setTag('json', '-')
                )
            )
            ->addProperty(
                new StructProperty(
                    'PropTwo',
                    new Type('int64'),
                    (new Tags())
                        ->setTag('schema', 'prop_two')
                        ->setTag('description', 'The second property')
                )
            );

        $this->assertSame(<<<GO
// Sample is a sample structure
type Sample struct {
	MyTypeOne	`json:"-"`
	PropTwo	int64	`schema:"prop_two" description:"The second property"`
}
GO
            , $struct->toString());
    }

}