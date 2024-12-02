<?php

function overlaps($range1, $range2) {
  return ($range1[0] >= $range2[0] && $range1[0] <= $range2[1]);
}

function needsReschedule($assignment) {
  [$a1, $a2] = explode(',', $assignment, 2);
  $r1 = explode('-', $a1);
  $r2 = explode('-', $a2);
  return (overlaps($r1, $r2) || overlaps($r2, $r1));
}

$lines = file("input.txt", FILE_IGNORE_NEW_LINES);
$num = 0;
foreach ($lines as $assignment) {
  if (needsReschedule($assignment)) {
    echo " => $assignment \n";
    $num += 1;
  }
}
echo $num, "\n";
