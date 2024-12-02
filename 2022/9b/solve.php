<?php

function distance($pos1, $pos2)
{
    return [
        $pos2[0] - $pos1[0],
        $pos2[1] - $pos1[1]
    ];
}

function move($pos, $direction)
{
    static $map = [
        'L' => [-1, 0],
        'R' => [1, 0],
        'D' => [0, 1],
        'U' => [0, -1],
    ];

    return [
        $pos[0] + $map[$direction][0],
        $pos[1] + $map[$direction][1]
    ];
}


function visualize($visited) {
    $minmax = array_reduce(
        array_keys($visited),
        function($minmax, $pos) {
            [$x, $y] = explode(',', $pos);
            return [min($x, $minmax[0]), min($y, $minmax[1]), max($x, $minmax[2]), max($y, $minmax[3])];
        },
        [0, 0, 0, 0]
    );

    [$startX, $startY] = explode(",", array_key_first($visited));
    for ($y=$minmax[1]; $y<=$minmax[3]; $y++) {
        for ($x=$minmax[0]; $x<=$minmax[2]; $x++) {
            if ($x == $startX && $y == $startY) {
                echo "s";
            } else {
                echo isset($visited["{$x},{$y}"]) ? '#' : '.';
            }
        }
        echo "\n";
    }

}

$visited = [];
$lines = file("input.txt", FILE_IGNORE_NEW_LINES);

//$lines = [
//"R 5",
//"U 8",
//"L 8",
//"D 3",
//"R 17",
//"D 10",
//"L 25",
//"U 20",
//];

$numFollowers = 9;
$currentHeadPos = [0, 0];
$knotPositions = array_fill(0, $numFollowers, $currentHeadPos);
foreach ($lines as $line) {
    [$direction, $num] = explode(' ', $line);
    for ($i=0; $i<$num; $i++) {
        $currentHeadPos = move($currentHeadPos, $direction);

        $currentLeaderPos = $currentHeadPos;
        $currentTailPos = $currentHeadPos;
        for ($k=0; $k<$numFollowers; $k++) {
            $currentTailPos = $knotPositions[$k];
            [$dx, $dy] = distance($currentTailPos, $currentLeaderPos);
            $notTouchingX = abs($dx) > 1;
            $notTouchingY = abs($dy) > 1;
            $notTouchingDiagonal = ($notTouchingY && abs($dx) > 0) || ($notTouchingX && abs($dy) > 0);
            if ($notTouchingY || $notTouchingDiagonal) {
                $currentTailPos = move($currentTailPos, $dy > 0 ? 'D' : 'U');
            }
            if ($notTouchingX || $notTouchingDiagonal) {
                $currentTailPos = move($currentTailPos, $dx > 0 ? 'R' : 'L');
            }
            $knotPositions[$k] = $currentTailPos;
            $currentLeaderPos = $currentTailPos;
        }

        $visited["{$currentTailPos[0]},{$currentTailPos[1]}"] = true;
    }
}

visualize($visited);

echo "Visited ", count($visited), " positions\n";
