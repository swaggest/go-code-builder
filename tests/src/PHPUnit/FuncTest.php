<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;

use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Func\Result;
use Swaggest\GoCodeBuilder\Templates\Type\Type;
use Swaggest\GoCodeBuilder\Templates\Type\Pointer;

class FuncTest extends \PHPUnit_Framework_TestCase
{
    public function testPointerResult()
    {
        $func = new FuncDef('Sample');
        $func->setResult(
            (new Result())
                ->add(null, new Pointer(new Type('MyType')))
                ->add(null, new Type('error'))
        );

        $this->assertSame(<<<GO
func Sample() (*MyType, error) {

}


GO
            , $func->render());

    }

    public function testSingleResult()
    {
        $func = new FuncDef('Sample');
        $func->setResult(
            (new Result())
                ->add(null, new Type('error'))
        );

        $this->assertSame(<<<GO
func Sample() error {

}


GO
            , $func->render());

    }


    public function testNoArgsNoResult()
    {
        $func = new FuncDef('Sample');

        $this->assertSame(<<<GO
func Sample() {

}


GO
            , $func->render());

    }

}