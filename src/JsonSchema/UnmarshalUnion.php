<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;
use Swaggest\GoCodeBuilder\Templates\Struct\StructProperty;
use Swaggest\GoCodeBuilder\Templates\Type\Type;

class UnmarshalUnion extends GoTemplate
{
    /** @var GoBuilder */
    public $goBuilder;

    public $withMustUnmarshal;
    public $withMayUnmarshal;
    public $withIgnoreKeys;
    public $withPatternProperties;
    public $withAdditionalProperties;

    public $patterns = [];
    public $regexNames = [];

    public function patternVarName($pattern)
    {
        if (isset($this->patterns[$pattern])) {
            return $this->patterns[$pattern];
        }
        $regexName = $this->goBuilder->codeBuilder->privateName('regex_' . $pattern);
        $i = 0;
        while (isset($this->regexNames[$regexName]) && $pattern !== $this->regexNames[$regexName]) {
            $i++;
            $regexName = $this->goBuilder->codeBuilder->privateName('regex_' . $pattern) . $i;
        }
        $this->patterns[$pattern] = $regexName;
        $this->regexNames[$regexName] = $pattern;
        return $regexName;
    }

    protected function toString()
    {
        $code = new Code();
        $code->imports()
            ->addByName('encoding/json');

        if ($this->withAdditionalProperties) {
            $code->imports()
                ->addByName('strings');
        }

        if ($this->withPatternProperties) {
            $code->imports()
                ->addByName('strings')
                ->addByName('regexp');

            $var = '';
            foreach ($this->patterns as $pattern => $regexName) {
                // var regex = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
                $var .= "\t$regexName = regexp.MustCompile({$this->escapeValue($pattern)})\n";
            }
            $code->addSnippet(<<<GO
// Regular expressions for pattern properties.
var (
$var)


GO
            );
        }

        $union = new StructDef('unionMap');
        if ($this->withMustUnmarshal) {
            $union->addProperty(new StructProperty('mustUnmarshal', new Type('[]interface{}')));
        }
        if ($this->withMayUnmarshal) {
            $union->addProperty(new StructProperty('mayUnmarshal', new Type('[]interface{}')));
        }
        if ($this->withIgnoreKeys) {
            $union->addProperty(new StructProperty('ignoreKeys', new Type('[]string')));
        }
        if ($this->withPatternProperties) {
            $union->addProperty(new StructProperty('patternProperties', new Type('map[*regexp.Regexp]interface{}')));
        }
        if ($this->withAdditionalProperties) {
            $union->addProperty(new StructProperty('additionalProperties', new Type('interface{}')));
        }

        $union->addProperty(new StructProperty('jsonData', new Type('[]byte')));


        $unionMap = <<<GO
{$union->render()}
func (u unionMap) unmarshal() error {
	{$this->ifThenElse($this->withMustUnmarshal,
            "for _, item := range u.mustUnmarshal {
		// Unmarshal to struct.
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			return err
		}
	}")}
	
	{$this->ifThenElse($this->withMayUnmarshal,
            "for i, item := range u.mayUnmarshal {
		// Unmarshal to struct.
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			u.mayUnmarshal[i] = nil
		}
	}")}
	
	{$this->ifThenElse($this->withPatternProperties && !$this->withAdditionalProperties,
            "if len(u.patternProperties) == 0 {
		return nil
	}")}
	
	{$this->ifThenElse(!$this->withPatternProperties && $this->withAdditionalProperties,
            "if u.additionalProperties == nil {
		return nil
	}")}
	
	{$this->ifThenElse($this->withPatternProperties && $this->withAdditionalProperties,
            "if len(u.patternProperties) == 0 && u.additionalProperties == nil {
		return nil
	}")}
	
	{$this->ifThenElse($this->withPatternProperties || $this->withAdditionalProperties,
            "// Unmarshal to a generic map.
	var m map[string]*json.RawMessage

	err := json.Unmarshal(u.jsonData, &m)
	if err != nil {
		return err
	}")}
	
	{$this->ifThenElse($this->withIgnoreKeys && ($this->withPatternProperties || $this->withAdditionalProperties),
            "// Remove ignored keys (defined in struct).
	for _, i := range u.ignoreKeys {
		delete(m, i)
	}
	
	// Return early on empty map.
	if len(m) == 0 {
		return nil
	}")}
	
	{$this->ifThenElse($this->withPatternProperties,
            "if len(u.patternProperties) != 0 {
		err = u.unmarshalPatternProperties(m)
		if err != nil {
			return err
		}
	}")}
	
	{$this->ifThenElse($this->withPatternProperties || $this->withAdditionalProperties,
            "// Returning early on empty map.
	if len(m) == 0 {
		return nil
	}")}
	
	{$this->ifThenElse($this->withAdditionalProperties,
            "if u.additionalProperties != nil {
		return u.unmarshalAdditionalProperties(m)
	}")}
	
	return nil
}

{$this->ifThenElse($this->withAdditionalProperties, <<<'GO'
func (u unionMap) unmarshalAdditionalProperties(m map[string]*json.RawMessage) error {
	var err error
	
	subMap := make([]byte, 1, 100)
	
	subMap[0] = '{'

	// Iterating map and filling additional properties.
	for key, val := range m {
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`
		
		if len(subMap) != 1 {
			subMap = append(subMap[:len(subMap)-1], ',')
		}
		
		subMap = append(subMap, []byte(keyEscaped)...)
		
		if val != nil {
			subMap = append(subMap, []byte(*val)...)
		} else {
			subMap = append(subMap, []byte("null")...)
		}
		
		subMap = append(subMap, '}')
	}

	if len(subMap) > 1 {
		err = json.Unmarshal(subMap, u.additionalProperties)
		if err != nil {
			return err
		}
	}
	
	return nil
}
GO
        )}
{$this->ifThenElse($this->withPatternProperties, <<<'GO'
func (u unionMap) unmarshalPatternProperties(m map[string]*json.RawMessage) error {
	patternMapsRaw := make(map[*regexp.Regexp][]byte, len(u.patternProperties))
	
	// Iterating map and filling pattern properties sub maps.
	for key, val := range m {
		matched := false
		ok := false
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex := range u.patternProperties {
			if regex.MatchString(key) {
				matched = true
				
				var subMap []byte
				
				if subMap, ok = patternMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)
				
				if val != nil {
					subMap = append(subMap, []byte(*val)...)
				} else {
					subMap = append(subMap, []byte("null")...)
				}
				
				subMap = append(subMap, '}')

				patternMapsRaw[regex] = subMap
			}
		}

		// Remove from properties map if matched to at least one regex.
		if matched {
			delete(m, key)
		}
	}

	for regex := range u.patternProperties {
		if subMap, ok := patternMapsRaw[regex]; !ok {
			continue
		} else {
			err := json.Unmarshal(subMap, u.patternProperties[regex])
			if err != nil {
				return err
			}
		}
	}
	
	return nil
}


GO
        )}
GO;

        $code->addSnippet($this->stripEmptyLines($unionMap));

        return $code;

    }
}