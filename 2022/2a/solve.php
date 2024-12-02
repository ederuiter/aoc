<?php

$score = [];
foreach (['A' => 1, 'B' => 2, 'C' => 3] as $shape_opponent => $shape_opponent_score) {
  foreach (['X' => 1, 'Y' => 2, 'Z' => 3] as $shape_me => $shape_me_score) {
    $diff = $shape_me_score - $shape_opponent_score;
    if ($diff === 1 || ($shape_opponent == 'C' && $shape_me === 'X')) {
      $s = 6;
    } else if ($diff === 0) {
      $s = 3;
    } else {
      $s = 0;
    }
    $score[$shape_opponent . ' ' . $shape_me] = $shape_me_score + $s;
  }
}

print_r($score);

$total_score = 0;
$lines = file("input.txt", FILE_IGNORE_NEW_LINES);
foreach ($lines as $line) {
   $total_score += $score[$line];
}
echo $total_score, "\n";
