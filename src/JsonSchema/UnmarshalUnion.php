<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class UnmarshalUnion extends GoTemplate
{
    public $withRegex;

    protected function toString()
    {
        $code = new Code();
        $code->imports()
            ->addByName('encoding/json')
            ->addByName('bytes')
            ->addByName('errors');

        if ($this->withRegex) {
            $code->imports()
                ->addByName('regexp')
                ->addByName('strings');
        }

        $code->addSnippet(
            <<<GO
func unmarshalUnion(
	mustUnmarshal []interface{},
	mayUnmarshal []interface{},
	ignoreKeys []string,
{$this->renderRegexMapsArg()}	j []byte,
) error {
	for _, item := range mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			return err
		}
	}

	for i, item := range mayUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			mayUnmarshal[i] = nil
		}
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// removing ignored keys (defined in struct)
	for _, i := range ignoreKeys {
		delete(m, i)
	}

	// returning early on empty map
	if len(m) == 0 {
		return nil
	}
{$this->renderRegexMapsLoop()}
	return nil
}

GO

        );
        return $code;

    }

    private function renderRegexMapsArg()
    {
        if (!$this->withRegex) {
            return <<<'GO'
	_ map[string]interface{}, // unused regexMaps

GO;
        }
        return <<<'GO'
	regexMaps map[string]interface{},

GO;
    }

    private function renderRegexMapsLoop()
    {
        if (!$this->withRegex) {
            return '';
        }
        return <<<'GO'
	// preparing regexp matchers
	var reg = make(map[string]*regexp.Regexp, len(regexMaps))
	for regex := range regexMaps {
		if regex != "" {
			reg[regex], err = regexp.Compile(regex)
			if err != nil {
				return err
			}
		}
	}

	subMapsRaw := make(map[string][]byte, len(regexMaps))
	// iterating map and feeding subMaps
	for key, val := range m {
		matched := false
		var ok bool
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex, exp := range reg {
			if exp.MatchString(key) {
				matched = true
				var subMap []byte
				if subMap, ok = subMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)
				subMap = append(subMap, []byte(*val)...)
				subMap = append(subMap, '}')

				subMapsRaw[regex] = subMap
			}
		}

		// empty regex for additionalProperties
		if !matched {
			var subMap []byte
			if subMap, ok = subMapsRaw[""]; !ok {
				subMap = make([]byte, 1, 100)
				subMap[0] = '{'
			} else {
				subMap = append(subMap[:len(subMap)-1], ',')
			}
			subMap = append(subMap, []byte(keyEscaped)...)
			subMap = append(subMap, []byte(*val)...)
			subMap = append(subMap, '}')

			subMapsRaw[""] = subMap
		}
	}

	for regex := range regexMaps {
		if subMap, ok := subMapsRaw[regex]; !ok {
			continue
		} else {
			err = json.Unmarshal(subMap, regexMaps[regex])
			if err != nil {
				return err
			}
		}
	}

GO;
    }
}