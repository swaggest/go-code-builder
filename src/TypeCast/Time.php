<?php

namespace Swaggest\GoCodeBuilder\TypeCast;


use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\Type\TypeCastException;

class Time implements Registry
{
    static $types = array('time.Time', 'string', 'int64');

    /**
     * @param string $toTypeString
     * @param string $fromTypeString
     * @return bool
     */
    public function canProcess($toTypeString, $fromTypeString)
    {
        return (
            in_array($toTypeString, self::$types)
            && in_array($fromTypeString, self::$types)
        );
    }

    /**
     * @param $toTypeString
     * @param $fromTypeString
     * @param $toVarName
     * @param $fromVarName
     * @param $assignOp
     * @return Code|string
     * @throws TypeCastException
     */
    public function process($toTypeString, $fromTypeString, $toVarName, $fromVarName, $assignOp)
    {
        $tmpVar = 'tmp' . substr(md5($toVarName), 0, 4);

        if (' = *' === $assignOp) {
            $assignOp = ' = ';
        }

        if ($fromTypeString === 'time.Time' && $toTypeString === 'string') {
            return <<<GO
{$tmpVar} := {$fromVarName}.Format(time.RFC3339Nano)
{$toVarName}{$assignOp}{$tmpVar}
GO;
        }

        if ($toTypeString === 'time.Time' && $fromTypeString === 'string') {
            return <<<GO
{$tmpVar}, _ := time.Parse(time.RFC3339Nano, {$fromVarName})
{$toVarName}{$assignOp}{$tmpVar}
GO;
        }


        if ($fromTypeString === 'time.Time' && $toTypeString === 'int64') {
            return <<<GO
{$tmpVar} := {$fromVarName}.Unix()
{$toVarName}{$assignOp}{$tmpVar}
GO;
        }

        if ($toTypeString === 'time.Time' && $fromTypeString === 'int64') {
            return <<<GO
{$tmpVar} := time.Unix({$fromVarName}, 0)
{$toVarName}{$assignOp}{$tmpVar}
GO;
        }

    }
}