package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func printWorld(world []byte, numCols int) {
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
		line := scanner.Text()
		if line == "" {
			parsedMap = true
		} else if !parsedMap {
			numCols = len(line)
			index := strings.Index(line, "@")
			if index != -1 {
				robotPos = row*numCols + index
			}
			world = append(world, line...)
			row++
		} else {
			for _, move := range []byte(line) {
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

				boxOffset := 1
				possible := true
				for {
					switch world[robotPos+boxOffset*offset] {
					case '#':
						possible = false
						goto done
					case 'O':
						boxOffset++
					case '.':
						goto done
					}
				}

			done:
				if possible {
					if boxOffset > 1 {
						world[robotPos+boxOffset*offset] = world[robotPos+offset]
					}
					world[robotPos+offset] = '@'
					world[robotPos] = '.'
					robotPos = robotPos + offset
				}
				//fmt.Printf("Move %s (possible: %t)\n", string(move), possible)
				//printWorld(world, numCols)
			}
		}
	}

	printWorld(world, numCols)

	sum := 0
	for i, box := range world {
		if box == 'O' {
			x := i % numCols
			y := (i - x) / numCols
			//fmt.Printf("Box at %d, %d\n", y, x)
			sum += (y * 100) + x
		}
	}
	fmt.Println(sum)
}
