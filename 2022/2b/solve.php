<?php

$winning = ['A' => 'B', 'B' => 'C', 'C' => 'A'];
$losing = ['A' => 'C', 'B' => 'A', 'C' => 'B'];

$play = [];
foreach (['A' => 1, 'B' => 2, 'C' => 3] as $shape_opponent => $shape_opponent_score) {
  $play[$shape_opponent . ' X'] = $shape_opponent . ' ' . $losing[$shape_opponent];
  $play[$shape_opponent . ' Y'] = $shape_opponent . ' ' . $shape_opponent;
  $play[$shape_opponent . ' Z'] = $shape_opponent . ' ' . $winning[$shape_opponent];
}

print_r($play);

$score = [];
foreach (['A' => 1, 'B' => 2, 'C' => 3] as $shape_opponent => $shape_opponent_score) {
  foreach (['A' => 1, 'B' => 2, 'C' => 3] as $shape_me => $shape_me_score) {
    $diff = $shape_me_score - $shape_opponent_score;
    if ($diff === 1 || ($shape_opponent == 'C' && $shape_me === 'A')) {
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
   $total_score += $score[$play[$line]];
}
echo $total_score, "\n";
