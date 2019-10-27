<?php

namespace Swaggest\GoCodeBuilder\Templates\Func;

use Swaggest\GoCodeBuilder\Templates\Code;
use Swaggest\GoCodeBuilder\Templates\GoTemplate;

class FuncDef extends GoTemplate
{
    /** @var Arguments */
    private $arguments;
    /** @var Result|null */
    private $result;
    /** @var Code|null */
    private $body;
    /** @var Argument|null */
    private $self;

    /** @var string */
    private $name;

    /**
     * FuncDef constructor.
     * @param string $name
     * @param string $comment
     */
    public function __construct($name, $comment = '')
    {
        $this->name = $name;
        $this->comment = $comment;
    }

    const RENDER_FUNC = 'func';
    const RENDER_IFACE = 'iface';
    private $renderMode = self::RENDER_FUNC;

    /**
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * @return Argument|null
     */
    public function getSelf()
    {
        return $this->self;
    }

    /**
     * @param string $name
     * @return FuncDef
     */
    public function setName($name)
    {
        $this->name = $name;
        return $this;
    }

    /**
     * @return string
     */
    public function getRenderMode()
    {
        return $this->renderMode;
    }

    /**
     * @param string $renderMode
     * @return FuncDef
     */
    public function setRenderMode($renderMode)
    {
        $this->renderMode = $renderMode;
        return $this;
    }


    protected function toString()
    {
        $result = (string)$this->result;
        if ($result) {
            $result = ' ' . $result;
        }

        if ($this->renderMode === self::RENDER_FUNC) {
            $body = $this->body ? $this->body->render() : '';
            $code = <<<GO
{$this->renderComment()}func {$this->renderSelf()}{$this->name}({$this->arguments}){$result} {
{$this->padLines("\t", $this->tabIndents($this->stripEmptyLines(trim($body))), false)}
}


GO;
        } else {
            $code = <<<GO
{$this->renderComment()}{$this->name}({$this->arguments}){$result}
GO;
        }
        return $code;
    }

    private function renderSelf()
    {
        if ($this->self) {
            return "($this->self) ";
        }
        return '';
    }

    /**
     * @param Arguments $arguments
     * @return FuncDef
     */
    public function setArguments(Arguments $arguments)
    {
        $this->arguments = $arguments;
        return $this;
    }

    /**
     * @return Arguments
     */
    public function getArguments()
    {
        return $this->arguments;
    }

    /**
     * @return Result|null
     */
    public function getResult()
    {
        return $this->result;
    }


    /**
     * @param Result $result
     * @return FuncDef
     */
    public function setResult(Result $result = null)
    {
        $this->result = $result;
        return $this;
    }

    /**
     * @param Code $body
     * @return FuncDef
     */
    public function setBody(Code $body)
    {
        $this->body = $body;
        return $this;
    }

    /**
     * @param Argument $self
     * @return FuncDef
     */
    public function setSelf(Argument $self)
    {
        $this->self = $self;
        return $this;
    }
}