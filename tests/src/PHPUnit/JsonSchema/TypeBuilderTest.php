<?php

namespace Swaggest\GoCodeBuilder\Tests\PHPUnit\JsonSchema;


use Swaggest\GoCodeBuilder\JsonSchema\GoBuilder;
use Swaggest\GoCodeBuilder\JsonSchema\TypeBuilder;
use Swaggest\JsonSchema\Context;
use Swaggest\JsonSchema\RemoteRef\Preloaded;
use Swaggest\JsonSchema\Schema;

class TypeBuilderTest extends \PHPUnit_Framework_TestCase
{
    public function testRef()
    {
        $schemaData = json_decode(<<<'JSON'
{
    "anyOf": [
        {},
        {"$ref":"#/definitions/header"}
    ],
    "definitions": {
        "header": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "maximum": {"$ref": "#/definitions/maximum"}
            }
        },
        "maximum": {
            "$ref": "http://json-schema.org/draft-04/schema#/properties/maximum"
        }
    }
}
JSON
        );
        $refResolver = new Preloaded();
        $context = new Context($refResolver);
        $schema = Schema::import($schemaData, $context);
        $goBuilder = new GoBuilder();
        $type = $goBuilder->getType($schema);

        $this->assertSame('Header', $type->getTypeString());
        $expectedGen = <<<'GO'
// Header structure is generated from "#/definitions/header".
type Header struct {
	Maximum float64 `json:"maximum,omitempty"`
}

type marshalHeader Header

var knownKeysHeader = []string{
	"maximum",
}

// UnmarshalJSON decodes JSON.
func (h *Header) UnmarshalJSON(data []byte) error {
	var err error

	mh := marshalHeader(*h)

	err = json.Unmarshal(data, &mh)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysHeader {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Header: %v", offendingKeys)
	}

	*h = Header(mh)

	return nil
}



GO;

        $actualGen = '';
        foreach ($goBuilder->getGeneratedStructs() as $struct) {
            $actualGen .= $struct->structDef;
        }
        $this->assertSame($expectedGen, $actualGen);
    }

    function testXGoTypeIgnore()
    {
        $builder = new GoBuilder();
        $builder->options->ignoreXGoType = true;
        $schema = new Schema();
        $schema->{'x-go-type'} = 'my-package/domain/orders.Order';
        $schema->type = Schema::STRING;
        $tb = new TypeBuilder($schema, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('string', $type->getTypeString());
    }


    function testXGoTypeString()
    {
        $builder = new GoBuilder();
        $schema = new Schema();
        $schema->{'x-go-type'} = 'my-package/domain/orders.Order';
        $tb = new TypeBuilder($schema, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('my-package/domain/orders.Order', $type->getTypeString());
    }

    function testXGoTypeGoSwaggerObject()
    {
        $builder = new GoBuilder();
        $schema = new Schema();
        $schema->{'x-go-type'} = json_decode('{
                  "import": {
                    "package": "my-package/domain/orders"
                  },
                  "type": "Order"
                }');
        $tb = new TypeBuilder($schema, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('my-package/domain/orders.Order', $type->getTypeString());
    }


    function testXFlags()
    {
        $prop = new Schema();
        $prop->type = [Schema::STRING, Schema::NULL];

        $propXNullable = new Schema();
        $propXNullable->type = Schema::STRING;
        $propXNullable->{TypeBuilder::X_NULLABLE} = true;
        $propXNullable->{TypeBuilder::X_OMIT_EMPTY} = true;
        $propXNullable->{TypeBuilder::X_GENERATE} = true;

        $propNullable = new Schema();
        $propNullable->type = Schema::STRING;
        $propNullable->{TypeBuilder::NULLABLE} = true;

        $propKeepEmpty = new Schema();
        $propKeepEmpty->type = Schema::STRING;
        $propKeepEmpty->{TypeBuilder::X_OMIT_EMPTY} = false;

        $propSkipGenerate = new Schema();
        $propSkipGenerate->type = Schema::STRING;
        $propSkipGenerate->{TypeBuilder::X_GENERATE} = false;

        $obj = Schema::object();
        $obj->setProperty('schema-nullable', $prop);
        $obj->setProperty('x-nullable', $propXNullable);
        $obj->setProperty('nullable', $propNullable);
        $obj->setProperty('regular', Schema::string());
        $obj->setProperty('keep-empty', $propKeepEmpty);
        $obj->setProperty('skip-generate', $propSkipGenerate);


        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	SchemaNullable *string `json:"schema-nullable"`
	XNullable      string  `json:"x-nullable,omitempty"`
	Nullable       string  `json:"nullable,omitempty"`
	Regular        string  `json:"regular,omitempty"`
	KeepEmpty      string  `json:"keep-empty"`
}


GO
            , $struct->structDef->render());


        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->withZeroValues = true;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	SchemaNullable *string `json:"schema-nullable"`
	XNullable      *string `json:"x-nullable,omitempty"`
	Nullable       *string `json:"nullable,omitempty"`
	Regular        *string `json:"regular,omitempty"`
	KeepEmpty      string  `json:"keep-empty"`
}


GO
            , $struct->structDef->render());


        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->ignoreNullable = true;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	SchemaNullable *string `json:"schema-nullable,omitempty"`
	XNullable      string  `json:"x-nullable,omitempty"`
	Nullable       string  `json:"nullable,omitempty"`
	Regular        string  `json:"regular,omitempty"`
	KeepEmpty      string  `json:"keep-empty"`
}


GO
            , $struct->structDef->render());


        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->enableXNullable = true;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	SchemaNullable *string `json:"schema-nullable"`
	XNullable      *string `json:"x-nullable,omitempty"`
	Nullable       *string `json:"nullable"`
	Regular        string  `json:"regular,omitempty"`
	KeepEmpty      string  `json:"keep-empty"`
}


GO
            , $struct->structDef->render());


        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->enableXNullable = true;
        $builder->options->requireXGenerate = true;
        $tb = new TypeBuilder($prop, '#', $builder);
        $type = $tb->build();
        $this->assertEquals('*string', $type->getTypeString());

        $this->assertEquals('Untitled1', $builder->getType($obj)->getTypeString());
        $struct = $builder->getGeneratedStruct($obj, '#');

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	XNullable *string `json:"x-nullable,omitempty"`
}


GO
            , $struct->structDef->render());

    }

    public function testXGenerate()
    {
        $schemaJson = <<<'JSON'
{
    "type": "object",
    "properties": {
        "id": {"type": "integer", "x-generate": true},
        "entity": {"$ref": "#/definitions/entity"}
    },
    "definitions": {
        "entity": {
            "type": "object",
            "properties": {"prop1": {"type": "string"}}
        }
    }
}
JSON;
        $schema = Schema::import(json_decode($schemaJson));

        $builder = new GoBuilder();
        $builder->options->defaultAdditionalProperties = false;
        $builder->options->requireXGenerate = true;
        $tb = new TypeBuilder($schema, '#', $builder);
        $tb->build();

        $res = '';
        foreach ($builder->getGeneratedStructs() as $generatedStruct) {
            $res .= $generatedStruct->structDef->render();
        }

        $this->assertEquals(<<<'GO'
// Untitled1 structure is generated from "#".
type Untitled1 struct {
	ID int64 `json:"id,omitempty"`
}


GO
            , $res);

    }

    public function testEscape()
    {
        $schemaJson = <<<'JSON'
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "type": "string",
    "title": "Escape Issue",
    "enum": [
        "Doesn't Generate Correctly",
        "Does Generate"
    ]
}
JSON;
        $schema = Schema::import(json_decode($schemaJson));

        $builder = new GoBuilder();
        $tb = new TypeBuilder($schema, '#', $builder);
        $tb->build();

        $res = $builder->getCode()->render();

        $this->assertEquals(<<<'GO'
// EscapeIssue is an enum type.
type EscapeIssue string

// EscapeIssue values enumeration.
const (
	EscapeIssueDoesnTGenerateCorrectly = EscapeIssue("Doesn't Generate Correctly")
	EscapeIssueDoesGenerate = EscapeIssue("Does Generate")
)

// MarshalJSON encodes JSON.
func (i EscapeIssue) MarshalJSON() ([]byte, error) {
	switch i {
	case EscapeIssueDoesnTGenerateCorrectly:
	case EscapeIssueDoesGenerate:

	default:
		return nil, fmt.Errorf("unexpected EscapeIssue value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *EscapeIssue) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := EscapeIssue(ii)

	switch v {
	case EscapeIssueDoesnTGenerateCorrectly:
	case EscapeIssueDoesGenerate:

	default:
		return fmt.Errorf("unexpected EscapeIssue value: %v", v)
	}

	*i = v

	return nil
}


GO
            , $res);

    }

}