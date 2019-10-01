<?php

namespace Swaggest\GoCodeBuilder\Style;

class Initialisms
{
    public $values;

    public function __construct()
    {
        $this->values = json_decode('{
    "API":   true,
    "ASCII": true,
    "CPU":   true,
    "CSS":   true,
    "DNS":   true,
    "EOF":   true,
    "HTML":  true,
    "HTTP":  true,
    "HTTPS": true,
    "ID":    true,
    "IP":    true,
    "JSON":  true,
    "LHS":   true,
    "QPS":   true,
    "RAM":   true,
    "RHS":   true,
    "RPC":   true,
    "SLA":   true,
    "SMTP":  true,
    "SSH":   true,
    "TLS":   true,
    "TTL":   true,
    "UI":    true,
    "UID":   true,
    "URI":   true,
    "URL":   true,
    "UTF8":  true,
    "VM":    true,
    "XML":   true,
    
    
    "FK": true
}', true);
    }

    public function process($goName)
    {
        $words = preg_split('/(?=[A-Z])/', $goName);
        if (false !== $words) {
            foreach ($words as &$word) {
                if ($word === strtolower($word)) { // skip lowercase words
                    continue;
                }

                $uppercase = strtoupper($word);
                if (isset($this->values[$uppercase])) {
                    $word = $uppercase;
                }
            }
            $goName = implode('', $words);
            return $goName;
        }
        return '';
    }

}