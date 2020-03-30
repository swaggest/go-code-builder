<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;

use Swaggest\CodeBuilder\PlaceholderString;
use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Func\Arguments;
use Swaggest\GoCodeBuilder\Templates\Func\FuncDef;
use Swaggest\GoCodeBuilder\Templates\Type\TypeUtil;
use Swaggest\JsonSchemaMaker\InstanceFaker;

class MarshalingTestFunc
{
    public static function make(GeneratedStruct $struct, Options $options = null)
    {
        $f = new FuncDef('Test' . $struct->structDef->getName() . '_MarshalJSON_roundtrip');
        $f->setArguments((new Arguments())->add('t', TypeUtil::fromString('*testing.T')));

        $fakerOptions = new \Swaggest\JsonSchemaMaker\Options();
        if ($options !== null) {
            $fakerOptions->defaultAdditionalProperties = $options->defaultAdditionalProperties;
        }

        $instanceFaker = new InstanceFaker($struct->schema, $fakerOptions);
        $jsonValue = json_encode($instanceFaker->makeValue(), JSON_UNESCAPED_SLASHES);
        $c = new Code(new PlaceholderString(<<<GO
var (
    jsonValue = []byte(`$jsonValue`)
    v :type
)

require.NoError(t, json.Unmarshal(jsonValue, &v))

marshaled, err := json.Marshal(v)
require.NoError(t, err)
require.NoError(t, json.Unmarshal(marshaled, &v))
assertjson.Equal(t, jsonValue, marshaled)
GO
            , [
                ':type' => $struct->structDef->getType(),
            ]));

        $c->imports()
            ->addByName('encoding/json')
            ->addByName('github.com/stretchr/testify/require')
            ->addByName('github.com/swaggest/assertjson');

        $f->setBody($c);

        return $f;
    }

}