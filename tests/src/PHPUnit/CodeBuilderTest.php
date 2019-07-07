<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit;


use Swaggest\GoCodeBuilder\GoCodeBuilder;

class CodeBuilderTest extends \PHPUnit_Framework_TestCase
{
    public function testExportedName()
    {
        $codeBuilder = new GoCodeBuilder();
        $this->assertSame('B2B', $codeBuilder->exportableName('b2b'));
        $this->assertSame('Schema', $codeBuilder->exportableName('$schema'));
    }

}