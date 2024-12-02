<?php

$lines = file('input.txt', FILE_IGNORE_NEW_LINES);

$max = 0;
$line_max = 0;
$current = 0;
foreach ($lines as $index => $line) {
    if ($line === '') {
        $max = max($max, $current);
        if ($max === $current) {
           $line_max = $index;
        }
        $current = 0;
    }
    $current += intval($line);
}
echo $max, "\n";
