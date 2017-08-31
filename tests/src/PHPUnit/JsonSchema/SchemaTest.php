<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;

use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\JsonSchema\Schema;

class SchemaTest extends \PHPUnit_Framework_TestCase
{
    public function testProperties()
    {
        $anotherSchema = Schema::object()
            ->setProperty('hello', Schema::boolean())
            ->setProperty('world', Schema::string());


        $schema = Schema::object()
            ->setProperty('sampleInt', Schema::integer())
            ->setProperty('sampleBool', Schema::boolean())
            ->setProperty('sampleString', Schema::string())
            ->setProperty('sampleNumber', Schema::number())
        ;
        $schema
            ->setProperty('sampleSelf', $schema)
            ->setProperty('another', $anotherSchema)
        ;


        $builder = new GoBuilder();
        $type = $builder->getType($schema);

        $index = 0;

        foreach ($builder->getGeneratedClasses() as $class) {
            $class->class->setName('Beech' . ++$index);
        }

        foreach ($builder->getGeneratedClasses() as $class) {
            echo $class->class;
        }


        var_dump($type->getTypeString());
    }


    public function testSimple()
    {
        $builder = new GoBuilder();
        $this->assertSame('int64', $builder->getType(Schema::integer())->getTypeString());
        $this->assertSame('float64', $builder->getType(Schema::number())->getTypeString());
        $this->assertSame('string', $builder->getType(Schema::string())->getTypeString());
        $this->assertSame('bool', $builder->getType(Schema::boolean())->getTypeString());
        $this->assertSame('[]interface{}', $builder->getType(Schema::arr())->getTypeString());
        $this->assertSame('map[string]interface{}', $builder->getType(Schema::object())->getTypeString());
        $this->assertSame('interface{}', $builder->getType(Schema::null())->getTypeString());
    }

}