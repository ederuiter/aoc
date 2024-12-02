<?php

class CPU
{
    protected int $programCounter = 0;
    protected array $registers = [];

    public function __construct(array $registers)
    {
        $this->registers = $registers;
    }

    protected function tick($callback)
    {
        $this->programCounter += 1;
        $callback($this->programCounter, $this->registers);
    }

    protected function decodeInstruction($line)
    {
        $args = explode(' ', $line);
        $instruction = array_shift($args);
        return [$instruction, $args];
    }

    public function executeInstruction($line, $callback)
    {
        [$instruction, $args] = $this->decodeInstruction($line);
        switch ($instruction) {
            case 'noop':
                $this->tick($callback);
                break;
            case 'addx':
                $this->tick($callback);
                $this->tick($callback);
                $this->registers['x'] += $args[0];
                break;
            default:
                throw new \Exception('Invalid instruction');
        }
    }
}

$ss = 0;
$cpu = new CPU(['x' => 1]);
$lines = file("input.txt", FILE_IGNORE_NEW_LINES);
//$lines =explode("\n",
//"addx 15
//addx -11
//addx 6
//addx -3
//addx 5
//addx -1
//addx -8
//addx 13
//addx 4
//noop
//addx -1
//addx 5
//addx -1
//addx 5
//addx -1
//addx 5
//addx -1
//addx 5
//addx -1
//addx -35
//addx 1
//addx 24
//addx -19
//addx 1
//addx 16
//addx -11
//noop
//noop
//addx 21
//addx -15
//noop
//noop
//addx -3
//addx 9
//addx 1
//addx -3
//addx 8
//addx 1
//addx 5
//noop
//noop
//noop
//noop
//noop
//addx -36
//noop
//addx 1
//addx 7
//noop
//noop
//noop
//addx 2
//addx 6
//noop
//noop
//noop
//noop
//noop
//addx 1
//noop
//noop
//addx 7
//addx 1
//noop
//addx -13
//addx 13
//addx 7
//noop
//addx 1
//addx -33
//noop
//noop
//noop
//addx 2
//noop
//noop
//noop
//addx 8
//noop
//addx -1
//addx 2
//addx 1
//noop
//addx 17
//addx -9
//addx 1
//addx 1
//addx -3
//addx 11
//noop
//noop
//addx 1
//noop
//addx 1
//noop
//noop
//addx -13
//addx -19
//addx 1
//addx 3
//addx 26
//addx -30
//addx 12
//addx -1
//addx 3
//addx 1
//noop
//noop
//noop
//addx -9
//addx 18
//addx 1
//addx 2
//noop
//noop
//addx 9
//noop
//noop
//noop
//addx -1
//addx 2
//addx -37
//addx 1
//addx 3
//noop
//addx 15
//addx -21
//addx 22
//addx -6
//addx 1
//noop
//addx 2
//addx 1
//noop
//addx -10
//noop
//noop
//addx 20
//addx 1
//addx 2
//addx 2
//addx -6
//addx -11
//noop
//noop
//noop");

$screen = array_fill(0, 6, array_fill(0, 40, "."));
foreach ($lines as $line) {
    $cpu->executeInstruction($line, function($tick, $registers) use (&$screen) {
        $cycle = $tick - 1;
        $row = floor($cycle / 40);
        $column = $cycle % 40;
        $draw = $column >= $registers['x'] -1 && $column <= $registers['x'] + 1;
        $screen[$row][$column] = $draw ? '#' : '.';
    });
}
foreach ($screen as $row) {
    echo implode("", $row), "\n";
}
