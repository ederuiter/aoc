package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var directions = map[byte][]int{
	'N': {-1, 0},
	'E': {0, 1},
	'S': {1, 0},
	'W': {0, -1},
}

var turn = map[byte]byte{
	'E': 'S',
	'S': 'W',
	'W': 'N',
	'N': 'E',
}

var step = map[byte]byte{
	'N': '^',
	'S': '|',
	'E': '>',
	'W': '<',
}

func checkLoop(room map[int]map[int]byte, startPos [2]int, startDirection byte) bool {
	pos := startPos
	direction := startDirection
	num := 0
	size := len(room) * len(room[0])
	used := map[int]map[int]bool{}
	for {
		nextPos := [2]int{
			pos[0] + directions[direction][0],
			pos[1] + directions[direction][1],
		}

		nextTile, ok := room[nextPos[0]][nextPos[1]]
		if !ok {
			break
		} else if nextTile == '#' || nextTile == 'O' {
			direction = turn[direction]
		} else {
			pos = nextPos
			if _, ok := used[pos[0]]; !ok {
				used[pos[0]] = map[int]bool{}
			}
			used[pos[0]][pos[1]] = true

			num++
			if num > size {
				//printRoom(room, map[int]map[int]bool{})
				//printRoom(room, used)
				return true
			}
		}
	}
	return false
}

func printRoom(room map[int]map[int]byte, used map[int]map[int]bool) {
	for r := 0; r < len(room); r++ {
		line := fmt.Sprintf("%03d ", r+1)
		for c := 0; c < len(room[r]); c++ {
			u, ok := used[r][c]
			if ok && u {
				line += "x"
			} else {
				line += string(room[r][c])
			}
		}
		fmt.Println(line)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	direction := byte('N')
	room := map[int]map[int]byte{}

	start := [2]int{0, 0}
	row := 0
	for scanner.Scan() {
		parts := strings.SplitAfter(scanner.Text(), "")
		room[row] = map[int]byte{}
		for col, part := range parts {
			switch part {
			case "^":
				direction = 'N'
				start[0] = row
				start[1] = col
				room[row][col] = '^'
			case ".":
				room[row][col] = ' '
			case "#":
				room[row][col] = '#'
			}
		}
		row++
	}

	pos := start
	num := 1
	numLoops := 0
	done := map[int]bool{}
	for {
		nextPos := [2]int{
			pos[0] + directions[direction][0],
			pos[1] + directions[direction][1],
		}

		nextTile, ok := room[nextPos[0]][nextPos[1]]
		if !ok {
			break
		} else if nextTile == '#' {
			direction = turn[direction]
			room[pos[0]][pos[1]] = '+'
		} else {
			if nextTile == ' ' {
				index := (nextPos[0] * len(room[0])) + nextPos[1]
				room[nextPos[0]][nextPos[1]] = 'O'
				if _, ok := done[index]; !ok && (nextPos[0] != start[0] || nextPos[1] != start[1]) && checkLoop(room, pos, direction) {
					numLoops++
					done[index] = true
				}
			}

			pos = nextPos
			if nextTile == ' ' {
				room[pos[0]][pos[1]] = step[direction]
				num++
			} else {
				room[pos[0]][pos[1]] = '*'
			}
		}
	}

	//printRoom(room, map[int]map[int]bool{})

	fmt.Println(num)
	fmt.Println(numLoops)
}
