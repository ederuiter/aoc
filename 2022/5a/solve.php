<?php

$lines = file("input.txt", FILE_IGNORE_NEW_LINES);

$num = 0;
$stack_lines = [];
$stacks = [];
$stack_names = [];
$reading_crates = true;

foreach ($lines as $line) {
   if ($reading_crates) {
       if (strpos($line, '[') === false) {
           $stack_names = array_values(array_filter(explode(' ', $line)));
           foreach ($stack_names as $stack_index => $stack_name) {
               $stacks[$stack_name] = array_values(array_filter(array_map(fn ($item) => trim($item[$stack_index], '[] '), $stack_lines)));
           }
           $reading_crates = false;
       } else {
           $stack_lines[] = str_split($line, 4); //TODO: assumes max 9 stacks + single character boxes
       }
   } else if (trim($line) != '') {
       [$action, $num, $from, $src, $to, $dst] = explode(' ', $line);
       if ($action != 'move' || $from != 'from' || $to != 'to') {
           echo "ignoring invalid line: {$line}\n";
       }

       for ($i=0; $i<$num; $i++) {
           array_splice($stacks[$dst], 0, 0, array_splice($stacks[$src], 0, 1, []));
       }

       echo " ==> ", implode('', array_map(fn($stack) => $stack[0] ?? '', $stacks)), "\n";

   }
}

echo implode('', array_map(fn($stack) => $stack[0] ?? '', $stacks)), "\n";
