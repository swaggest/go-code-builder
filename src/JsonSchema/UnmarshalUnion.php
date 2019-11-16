<?php

namespace Swaggest\GoCodeBuilder\JsonSchema;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class UnmarshalUnion extends GoTemplate
{
    /** @var GoBuilder */
    public $goBuilder;

    public $withPatternProperties;

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

        if ($this->withPatternProperties) {
            $code->imports()
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

        return $code;
    }
}