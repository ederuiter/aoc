<?php

function isVisible($grid, $x, $y)
{
    $myHeight = $grid[$y][$x];
    $maxX = count($grid[0]) - 1;
    $maxY = count($grid) - 1;
    if ($x === 0 || $y === 0 || $x === $maxX || $y === $maxY) {
        return [true, "=" . $myHeight];
    }

    //left
    $allShorter = true;
    for ($i=0; $i<$x; $i++) {
        if ($grid[$y][$i] >= $myHeight) {
            $allShorter = false;
            break;
        }
    }
    if ($allShorter) {
        return [true, "<" . $myHeight];
    }

    //right
    $allShorter = true;
    for ($i=$x+1; $i<=$maxX; $i++) {
        if ($grid[$y][$i] >= $myHeight) {
            $allShorter = false;
            break;
        }
    }
    if ($allShorter) {
        return [true, ">" . $myHeight];
    }

    //top
    $allShorter = true;
    for ($i=0; $i<$y; $i++) {
        if ($grid[$i][$x] >= $myHeight) {
            $allShorter = false;
            break;
        }
    }
    if ($allShorter) {
        return [true, "^" . $myHeight];
    }

    //bottom
    $allShorter = true;
    for ($i=$y+1; $i<=$maxY; $i++) {
        if ($grid[$i][$x] >= $myHeight) {
            $allShorter = false;
            break;
        }
    }
    if ($allShorter) {
        return [true, "_" . $myHeight];
    }

    return [false, "++"];
}

$lines = file("input.txt", FILE_IGNORE_NEW_LINES);

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
$output_visible = "";
$numVisible = 0;
foreach ($grid as $y => $line) {
    foreach ($line as $x => $cell) {
        [$v, $o] = isVisible($grid, $x, $y);
        $output_visible .= $o;
        $output_grid .= $cell;
        $output_grid .= " ";
        if ($v) {
            $numVisible += 1;
        }
    }
    $output_visible .= "\n";
    $output_grid .= "\n";
}
echo $output_visible;
echo str_repeat("--", count($grid[0])), "\n";
echo $output_grid;

echo " => {$numVisible}\n";
