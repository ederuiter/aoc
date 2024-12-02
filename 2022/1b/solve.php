<?php

$lines = file('input.txt', FILE_IGNORE_NEW_LINES);

$totals = [];
$current = 0;
foreach ($lines as $index => $line) {
    if ($line === '') {
        $totals[] = $current;
        $current = 0;
    }
    $current += intval($line);
}
rsort($totals, SORT_NUMERIC);
$top = array_slice($totals, 0, 3);
$total = array_sum($top);
echo $total, "\n";
