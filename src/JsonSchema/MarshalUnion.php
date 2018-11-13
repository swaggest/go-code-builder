<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class MarshalUnion extends GoTemplate
{
    public $withRegex;

    protected function toString()
    {
        $code = new Code();
        $code->imports()
            ->addByName('encoding/json')
            ->addByName('errors');

        $code->addSnippet(
            <<<'GO'
func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	for _, m := range maps {
		if m == nil {
			continue
		}
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		if string(j) == "{}" {
			continue
		}
		if string(j) == "null" {
			continue
		}
		if j[0] != '{' {
			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if len(result) > 1 {
			result = append(result[:len(result)-1], ',')
		}
		result = append(result, j[1:]...)
	}
	// closing empty result
	if len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}

GO

        );
        return $code;

    }
}