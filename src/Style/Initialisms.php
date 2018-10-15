<?php

namespace Swaggest\GoCodeBuilder\Style;

class Initialisms
{
    public static function process($goName)
    {
        static $initialisms = null;
        if (null === $initialisms) {
            $initialisms = json_decode('{
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

        $words = preg_split('/(?=[A-Z])/', $goName);
        foreach ($words as &$word) {
            if ($word === strtolower($word)) { // skip lowercase words
                continue;
            }

            $uppercase = strtoupper($word);
            if (isset($initialisms[$uppercase])) {
                $word = $uppercase;
            }
        }
        $goName = implode('', $words);
        return $goName;
    }

}