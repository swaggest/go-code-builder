<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;

use Swaggest\GoCodeBuilder\Templates\Code;

class GoTemplateTest extends \PHPUnit_Framework_TestCase
{
    public function testTabIntents()
    {
        $t = new Code();

        $tabbed = $t->tabIndents(<<<'GO'
a := []string{}{
    "a", // a is an "a       a"
    "b",
    "c",
}
GO
        );

        $this->assertEquals(<<<'GO'
a := []string{}{
	"a", // a is an "a	   a"
	"b",
	"c",
}
GO
, $tabbed);
    }
}