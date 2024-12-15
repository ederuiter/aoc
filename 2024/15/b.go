package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func printWorld2(world []byte, numCols int) {
	line := ""
	for i, c := range world {
		if i%numCols == 0 && i > 0 {
			fmt.Println(line)
			line = ""
		}
		line += string(c)
	}
	fmt.Println(line)
	fmt.Println("")
}

func getAffected(world []byte, affected map[int]byte, pos int, move int, replacement byte) bool {
	tile := world[pos]
	affected[pos] = replacement
	if tile == '.' {
		return true
	} else if tile == '#' {
		return false
	} else if tile == '@' || tile == 'O' || ((move == -1 || move == 1) && (tile == '[' || tile == ']')) {
		return getAffected(world, affected, pos+move, move, tile)
	} else if tile == '[' {
		if getAffected(world, affected, pos+move, move, '[') &&
			getAffected(world, affected, pos+move+1, move, ']') {
			if _, ok := affected[pos+1]; !ok {
				affected[pos+1] = '.'
			}
			return true
		}
	} else if tile == ']' {
		if getAffected(world, affected, pos+move, move, ']') &&
			getAffected(world, affected, pos+move-1, move, '[') {
			if _, ok := affected[pos-1]; !ok {
				affected[pos-1] = '.'
			}
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	world := []byte{}
	row := 0
	numCols := 0
	parsedMap := false
	robotPos := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if len(line) == 0 {
			parsedMap = true
			printWorld2(world, numCols)
		} else if !parsedMap {
			numCols = 2 * len(line)
			for col, tile := range line {
				switch tile {
				case '@':
					world = append(world, '@', '.')
					robotPos = row*numCols + col*2
				case '.':
					world = append(world, '.', '.')
				case '#':
					world = append(world, '#', '#')
				case 'O':
					world = append(world, '[', ']')
				}
			}
			row++
		} else {
			prev := make([]byte, len(world))
			for _, move := range line {
				offset := 0
				switch move {
				case '<':
					offset = -1
				case '>':
					offset = 1
				case '^':
					offset = -numCols
				case 'v':
					offset = numCols
				}
				affected := map[int]byte{}
				possible := getAffected(world, affected, robotPos, offset, '.')
				if possible {
					copy(prev, world)
					for pos, value := range affected {
						world[pos] = value
					}
					robotPos += offset
				}
				//
				//fmt.Print("\033[H\033[2J")
				//fmt.Printf("Move %s (possible: %t %+v)\n", string(move), possible, *affected)
				//printWorld2(world, numCols)
				//fmt.Scanln()
			}
		}
	}

	printWorld2(world, numCols)

	sum := 0
	num := 0
	for i, box := range world {
		if box == '[' {
			num++
			x := i % numCols
			y := (i - x) / numCols
			//fmt.Printf("Box at %d, %d\n", y, x)
			sum += (y * 100) + x
		}
	}
	fmt.Println(num)
	fmt.Println(sum)
}
