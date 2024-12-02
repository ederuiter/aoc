<?php

$map = [
    'a' => 1,
    'b' => 2,
    'c' => 3,
    'd' => 4,
    'e' => 5,
    'f' => 6,
    'g' => 7,
    'h' => 8,
    'i' => 9,
    'j' => 10,
    'k' => 11,
    'l' => 12,
    'm' => 13,
    'n' => 14,
    'o' => 15,
    'p' => 16,
    'q' => 17,
    'r' => 18,
    's' => 19,
    't' => 20,
    'u' => 21,
    'v' => 22,
    'w' => 23,
    'x' => 24,
    'y' => 25,
    'z' => 26,
    'S' => 1,
    'E' => 26,
];

function walk($grid, $track, $endPos)
{
    //echo implode(" => ", $track), "\n";
    static $steps = [
        [-1, 0],
        [1, 0],
        [0, 1],
        [0, -1],
    ];
    $stop = false;
    $min = PHP_INT_MAX;
    $minTrack = [];
    $current = $track[count($track) - 1];
    $currentPos = explode(",", $current);
    foreach ($steps as $step) {
        $newPos = [$currentPos[0] + $step[0], $currentPos[1] + $step[1]];
        $new = "{$newPos[0]},{$newPos[1]}";
        if (isset($grid[$newPos[0]][$newPos[1]]) && ($grid[$newPos[0]][$newPos[1]] <= ($grid[$currentPos[0]][$currentPos[1]] + 1)) && !in_array($new, $track)) {
            $newTrack = $track;
            $newTrack[] = $new;
            if ($new === $endPos) {
                $num = count($newTrack);
                $stop = true;
            } else {
                [$num, $newTrack] = walk($grid, $newTrack, $endPos);
            }

            if ($num < $min) {
                $min = $num;
                $minTrack = $newTrack;
                echo "Found a track of length: $min\n";
                echo implode(" => ", $minTrack), "\n";
            }

            if ($stop) {
                break;
            }
        }
    }

    return [$min, $minTrack];
}

function nextTo($grid, $pos, $char)
{
    static $steps = [
        [-1, 0],
        [1, 0],
        [0, 1],
        [0, -1],
    ];
    foreach ($steps as $step) {
        $newPos = [$pos[0] + $step[0], $pos[1] + $step[1]];
        if (isset($grid[$newPos[0]][$newPos[1]]) && $grid[$newPos[0]][$newPos[1]] === $char) {
            return true;
        }
    }
    return false;
}

$start = [];
$end = [];
$grid = [];
$waypoints = array_fill_keys(array_keys($map), []);
$lines = file('input.txt', FILE_IGNORE_NEW_LINES);
foreach ($lines as $y => $line) {
    $splitted = str_split($line, 1);
    foreach ($splitted as $x => $char) {
        $grid[$y][$x] = $char;
        if ($char === 'S') {
            $start = [$y, $x];
            $char = 'a';
        } else if ($char === 'E') {
            $end = [$y, $x];
            $char = 'z';
        }
        $waypoints[$char][] = [$y, $x];
    }
}

$charByIndex = array_keys($waypoints);
$indexByChar = array_flip($charByIndex);

$starts = ['a' => [$start]];
foreach ($waypoints as $char => $points) {
    if ($char === 'a') {
        continue;
    }
    $prevChar = $charByIndex[$indexByChar[$char]-1] ?? null;
    echo "$char => ", count($points), "\n";

    foreach ($points as $point) {
        if (nextTo($grid, $point, $prevChar)) {
            $starts[$char][] = $point;
        }
    }
}

walkBack($grid, $waypoints, )



echo $startPos, " => ", $endPos, "\n";
$result = walk($grid, [$startPos], $endPos);
print_r($result);



