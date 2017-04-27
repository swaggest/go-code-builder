<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;


use Swaggest\GoCodeBuilder\Import;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\GoFile;

class GoFileTest extends \PHPUnit_Framework_TestCase
{
    public function testImports()
    {
        $func = new FuncDef('Sample');
        $code = new Code(<<<GO
fmt.Println("Hello world")
GO
        );
        $code->imports()->add(new Import('fmt'));
        $func->setBody($code);


        $goFile = new GoFile('my_package', 'github.com/acme/my_package');

        $goFile->setCode($func);

        $this->assertSame(<<<GO
package my_package // import "github.com/acme/my_package"

import (
	"fmt"
)

func Sample()  {
	fmt.Println("Hello world")
}
GO
            , $goFile->render());

    }

}