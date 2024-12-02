<?php

class Monkey
{
    public $name;
    public $items;
    public $operation;
    public $test;
    public $actions;
    public $numInspections = 0;

    public function __construct($name, $items, $operation, $test, $actions)
    {
        $this->name = $name;
        $this->items = $items;
        $this->operation = $operation;
        $this->test = $test;
        $this->actions = $actions;
    }

    public function play(array $others)
    {
        echo $this->name, ":\n";
        while ($item = array_shift($this->items)) {
            $this->numInspections += 1;
            echo "  Monkey inspects an item with a worry level of $item\n";
            $old = $item;
            $new = 0;
            eval($this->operation);
            echo "    Worry level update ({$this->operation}) from {$old} to {$new}\n";
            $new = floor($new / 3);
            echo "    Monkey gets bored with item. Your worry level is divided by 3 to {$new}\n";
            $result = ($new % $this->test) == 0;
            echo "    Current worry level is " . ($result ? '' : 'not ') . "divisible by {$this->test}\n";
            $recipient = $this->actions[$result ? 1 : 0];
            echo "    Item with worry level {$new} is thrown to {$recipient}\n";
            $others[$recipient]->items[] = $new;
        }
    }
}

function getValue($key, $line)
{
    $line = ltrim($line, ' ');
    if (substr($line, 0, strlen($key) + 2) != $key . ': ') {
        throw new \Exception('Unexpected input: ' . $line . ' -- expected -- ' . $key);
    }
    return substr($line, strlen($key) + 2);
}

function readMonkeys()
{
    $lines = file('input.txt', FILE_IGNORE_NEW_LINES);

    $res = [];
    while (count($lines) > 0) {
        $monkey = strtolower(rtrim(array_shift($lines), ':'));
        $startingItems = explode(', ', getValue('Starting items', array_shift($lines)));
        $operation = getValue('Operation', array_shift($lines));
        $operation = str_replace(["old", "new"], ['$old', '$new'], $operation) . ";";
        $test = str_replace('divisible by ', '', getValue('Test', array_shift($lines)));
        $resultTrue = str_replace('throw to ', '', getValue('If true', array_shift($lines)));
        $resultFalse = str_replace('throw to ', '', getValue('If false', array_shift($lines)));
        array_shift($lines);

        $res[$monkey] = new Monkey($monkey, $startingItems, $operation, $test, [0 => $resultFalse, 1 => $resultTrue]);
    }
    return $res;
}

$monkeys = readMonkeys();
for ($i=0; $i<20; $i++) {
    foreach ($monkeys as $monkey) {
        $monkey->play($monkeys);
    }
}

$inspections = [];
foreach ($monkeys as $monkey) {
    $inspections[] = $monkey->numInspections;
}
sort($inspections, SORT_NUMERIC);
$top1 = array_pop($inspections);
$top2 = array_pop($inspections);

echo ($top1 * $top2), "\n";
