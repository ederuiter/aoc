<?php

class Item
{
    protected ?Item $parent = null;
    protected string $name = '';
    protected string $type = '';
    protected int $size = 0;
    protected array $children = [];

    public function __construct(?Item $parent, string $type, string $name, int $size)
    {
        $this->parent = $parent;
        $this->type = $type;
        $this->name = $name;
        $this->size = $size;
    }

    public function getType(): string
    {
        return $this->type;
    }

    public function getSize(): int
    {
        return array_reduce(
            $this->children,
            fn($carry, Item $item) => $carry + $item->getSize(),
            $this->size
        );
    }

    public function addChild($type, $name, $size)
    {
        if ($this->type !== 'dir') {
            throw new \Exception('Not a directory');
        }
        $this->children[$name] = new Item($this, $type, $name, $size);
    }

    public function getRoot()
    {
        $current = $this;
        while ($current->parent !== null) {
            $current = $current->parent;
        }
        return $current;
    }

    public function getRealPath()
    {
        $res = '';
        if ($this->parent) {
            $res = $this->parent->getRealPath();
        }
        if ($res === '' || $res[-1] !== '/') {
            $res .= '/';
        }
        $res .= $this->name;
        return $res;
    }

    public function get($name)
    {
        $current = $this;
        if (substr($name, 0, 1) === '/') {
            $current =  $this->getRoot();
        }

        $parts = explode('/', $name);
        foreach ($parts as $part) {
            if ($current->type === 'file') {
                throw new \Exception('Not a directory: ' . $name);
            }

            if ($part === '.' || $part === '') {
                continue;
            } else if ($part === '..') {
                $current = $current->parent ?? $current;
            } else if (isset($current->children[$name])) {
                $current = $current->children[$name];
            } else {
                throw new \Exception('File not found: ' . $name);
            }
        }

        return $current;
    }

    public function dump()
    {
        foreach ($this->children as $child) {
            $child->dump();
        }
        echo "[{$this->type}] {$this->size} {$this->getSize()} {$this->getRealPath()}\n";
    }

    public function iterate($reducer, $initial)
    {
        $value = $reducer($initial, $this);
        foreach ($this->children as $child) {
            $value = $child->iterate($reducer, $value);
        }
        return $value;
    }
}

$fs = new Item(null, 'dir', '', 0);
$lines = file("input.txt", FILE_IGNORE_NEW_LINES);
$currentCmd = '';
$currentDir = $fs;
foreach ($lines as $line) {
    if ($line[0] === '$') {
        $args = explode(' ', $line, 3);
        $cmd = $args[1];
        if ($cmd === 'cd') {
            $currentDir = $currentDir->get($args[2]);
        } else if ($cmd !== 'ls') {
            throw new \Exception("Unknown command {$cmd}");
        }
        $currentCmd = $cmd;
    } else if ($currentCmd === 'ls') {
        [$size, $name] = explode(' ', $line, 2);
        $type = 'file';
        if ($size === 'dir') {
            $type = 'dir';
            $size = 0;
        }
        $currentDir->addChild($type, $name, $size);
    } else {
        throw new \Exception("Did not expect output of {$currentCmd}");
    }
}

$size = $fs->iterate(function ($carry, Item $item) {
    $size = $item->getSize();
    if ($item->getType() === 'dir' && $size <= 100000) {
        echo $item->getRealPath(), " => ", $size, "\n";
        return $carry + $size;
    }
    return $carry;
}, 0);

echo $size, "\n";
