<?php

namespace Swaggest\GoCodeBuilder\Templates;

use PhpLang\ScopeExit;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;

class GoFile extends GoTemplate
{
    const IMPORTS_KEY = 'imports';

    private $package;
    private $importPath;
    /** @var Code */
    private $code;
    private $imports;

    /**
     * @return Imports
     */
    public function getImports()
    {
        return $this->imports;
    }

    /**
     * GoFile constructor.
     * @param $package
     * @param $importPath
     */
    public function __construct($package, $importPath = null)
    {
        $this->package = $package;
        $this->importPath = $importPath;
        $this->imports = new Imports();
        $this->code = new Code();
    }

    /**
     * @return mixed
     */
    public function getComment()
    {
        return $this->comment;
    }

    /**
     * @return Code
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * @param Code|StructDef $code
     * @return GoFile
     */
    public function setCode($code)
    {
        $this->code = $code;
        return $this;
    }

    /**
     * @return mixed
     */
    public function getPackage()
    {
        return $this->package;
    }

    /**
     * @param mixed $package
     * @return GoFile
     */
    public function setPackage($package)
    {
        $this->package = $package;
        return $this;
    }

    /**
     * @return null
     */
    public function getImportPath()
    {
        return $this->importPath;
    }

    /**
     * @param null $importPath
     * @return GoFile
     */
    public function setImportPath($importPath)
    {
        $this->importPath = $importPath;
        return $this;
    }

    private function renderPackageName()
    {
        $path = explode('/', $this->package);
        return array_pop($path);
    }

    public function toString()
    {
        $prevGoFile = self::getCurrentGoFile();
        self::setCurrentGoFile($this);
        /** @noinspection PhpUnusedLocalVariableInspection */
        $_ = new ScopeExit(function () use ($prevGoFile) {
            self::setCurrentGoFile($prevGoFile);
        });

        $codeResult = (string)$this->code;


        $depCodesProcessed = array();
        while (!empty($this->depCodes)) {
            $depCodes = $this->depCodes;
            $this->depCodes = array();
            foreach ($depCodes as $key => $depCode) {
                if (!isset($depCodesProcessed[$key])) {
                    $codeResult .= $depCode;
                    $depCodesProcessed[$key] = 1;
                }
            }
        }

        $result = <<<GO
{$this->renderComment()}package {$this->renderPackageName()}{$this->renderImportPath()}

{$this->imports->toString()}

{$codeResult}
GO;

        return rtrim($result);
    }

    private function renderImportPath()
    {
        if ($this->importPath) {
            return ' // import "' . $this->importPath . '"';
        }
        return '';
    }

    /** @var GoFile */
    private static $currentGoFile;

    /**
     * @return GoFile
     */
    public static function getCurrentGoFile()
    {
        return self::$currentGoFile;
    }

    /**
     * @param GoFile $currentGoFile
     */
    public static function setCurrentGoFile($currentGoFile)
    {
        self::$currentGoFile = $currentGoFile;
    }


    private $depCodes = array();

    public function setDependentCode($uniqueKey, $value)
    {
        $this->depCodes[$uniqueKey] = $value;
        return $this;
    }

}