<?php

namespace Swaggest\GoCodeBuilder\Templates;

use PhpLang\ScopeExit;
use Swaggest\GoCodeBuilder\Exception;
use Swaggest\GoCodeBuilder\Templates\Struct\StructDef;

class GoFile extends GoTemplate
{
    const IMPORTS_KEY = 'imports';

    private $package;
    private $importPath;
    /** @var Code */
    private $code;
    private $imports;

    private $skipImportComment = false;

    /** @var GoFile */
    private $transaction;


    public function startTransaction()
    {
        if ($this->transaction !== null) {
            throw new Exception('Transaction is already running');
        }

        $this->transaction = new GoFile($this->package, $this->importPath);
    }

    public function commitTransaction()
    {
        if ($this->transaction === null) {
            throw new Exception('Not running a transaction');
        }

        foreach ($this->transaction->imports->imports as $import) {
            $this->imports->add($import);
        }
        if ($this->transaction->code->snippets) {
            foreach ($this->transaction->code->snippets as $item) {
                $this->code->addSnippet($item);
            }
        }

        $this->dropTransaction();
    }

    public function dropTransaction($ignoreMissing = true)
    {
        if (!$ignoreMissing && $this->transaction === null) {
            throw new Exception('Not running a transaction');
        }

        $this->transaction = null;
    }

    /**
     * @param boolean $skipImportComment
     * @return GoFile
     */
    public function setSkipImportComment($skipImportComment)
    {
        $this->skipImportComment = $skipImportComment;
        return $this;
    }

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

    protected function toString()
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

        $codeResult = rtrim($codeResult);

        $result = <<<GO
{$this->renderComment()}package {$this->renderPackageName()}{$this->renderImportPath()}

{$this->imports->render()}

{$codeResult}

GO;

        return $result;
    }

    private function renderImportPath()
    {
        if ($this->importPath && !$this->skipImportComment) {
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
     * @param $currentGoFile
     * @return GoFile previous go file
     */
    public static function setCurrentGoFile(GoFile $currentGoFile = null)
    {
        $previous = self::$currentGoFile;
        if ($currentGoFile !== null && $currentGoFile->transaction !== null) {
            self::$currentGoFile = $currentGoFile->transaction;
        } else {
            self::$currentGoFile = $currentGoFile;
        }
        return $previous;
    }


    private $depCodes = array();

    public function setDependentCode($uniqueKey, $value)
    {
        $this->depCodes[$uniqueKey] = $value;
        return $this;
    }

}