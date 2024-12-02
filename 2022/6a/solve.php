<?php

$data = file_get_contents('input.txt');
$length = strlen($data);

$buf = [];
for ($i=0; $i<$length; $i++) {
    $char = substr($data, $i, 1);
    $found = array_search($char, $buf);
    if ($found !== false) {
        array_splice($buf, 0, $found+1, []);
    }
    $buf[] = $char;
    if (count($buf) == 4) {
        $i++;
        $marker = implode('', $buf);
        echo "{$marker} => {$i}\n";
        break;
    }
}
