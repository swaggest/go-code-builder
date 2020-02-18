<?php


namespace Swaggest\GoCodeBuilder\Style;


class Comment
{
    public static function sentence($s) {
        $s = trim($s);
        if (substr($s, -1) !== '.') {
            $s = $s . '.';
        }
        $s = ucfirst($s);
        return $s;
    }

}