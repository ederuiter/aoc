<?php

function getDuplicate($line) {
  $items = str_split($line, 1);
  [$part1, $part2] = array_chunk($items, ceil(count($items) / 2));
  $res = array_values(array_intersect($part1, $part2))[0] ?? '';
  
  echo implode('', $part1) . ' ' . implode('', $part2) . ' => ' . $res, "\n";
  return $res;
}

$priority = array_flip([
  '', 
  'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
  'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
]);

$priorities = 0;
$lines = file("input.txt", FILE_IGNORE_NEW_LINES);
foreach ($lines as $line) {
   $priorities += $priority[getDuplicate($line)];
}
echo $priorities, "\n";
