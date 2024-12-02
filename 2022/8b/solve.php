<?php

function getScore($grid, $x, $y, $output = false)
{
    $myHeight = $grid[$y][$x];
    $maxX = count($grid[0]) - 1;
    $maxY = count($grid) - 1;
    $score = 1;

    if ($x === 0 || $y === 0 || $x === $maxX || $y === $maxY) {
        return 0;
    }

    //left
    $viewingDistance = 0;
    for ($i=$x-1; $i>=0; $i--) {
        $viewingDistance++;
        if ($grid[$y][$i] >= $myHeight) {
            break;
        }
    }
    $score = $score * $viewingDistance;

    //right
    $viewingDistance = 0;
    for ($i=$x+1; $i<=$maxX; $i++) {
        $viewingDistance++;
        if ($grid[$y][$i] >= $myHeight) {
            break;
        }
    }
    $score = $score * $viewingDistance;

    //top
    $viewingDistance = 0;
    for ($i=$y-1; $i>=0; $i--) {
        $viewingDistance++;
        if ($grid[$i][$x] >= $myHeight) {
            break;
        }
    }
    $score = $score * $viewingDistance;

    //bottom
    $viewingDistance = 0;
    for ($i=$y+1; $i<=$maxY; $i++) {
        $viewingDistance++;
        if ($grid[$i][$x] >= $myHeight) {
            break;
        }
    }
    $score = $score * $viewingDistance;

    return $score;
}

$lines = file("input.txt", FILE_IGNORE_NEW_LINES);
//
//$lines = [
//    "30373",
//    "25512",
//    "65332",
//    "33549",
//    "35390",
//];

$grid = [];
foreach ($lines as $line) {
    $grid[] = str_split($line, 1);
}

$output_grid = "";
$maxScore = 0;
foreach ($grid as $y => $line) {
    foreach ($line as $x => $cell) {
        $maxScore = max($maxScore, getScore($grid, $x, $y));
        $output_grid .= $cell;
        $output_grid .= " ";
    }
    $output_grid .= "\n";
}

echo str_repeat("--", count($grid[0])), "\n";
echo $output_grid;
echo " => {$maxScore}\n";
