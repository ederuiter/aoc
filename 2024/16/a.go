package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	type Tile struct {
		pos       int
		costs     int
		direction int
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	world := map[int]*Tile{}
	row := 0
	numCols := 0
	var start *Tile
	var end *Tile
	for scanner.Scan() {
		line := []byte(scanner.Text())
		numCols = len(line)
		for col, tile := range line {
			pos := row*numCols + col
			if tile == '#' {
				continue
			}
			world[pos] = &Tile{
				pos:       pos,
				costs:     -1,
				direction: -1,
			}
			switch tile {
			case 'E':
				end = world[pos]
				end.costs = 1
			case 'S':
				start = world[pos]
			}
		}
	}

	directions := []int{-numCols, 1, numCols, -1}
	stack := []*Tile{start}
	for len(stack) > 0 {
		next := stack[len(stack)-1]
		if next.costs != -1 {
			stack = stack[:len(stack)-1]
			continue
		}

		complete := true
		neighbours := map[int]*Tile{}
		cheapest := -1
		cheapestDirection := -1
		for d, direction := range directions {
			if neighbour, ok := world[next.pos+direction]; ok {
				neighbours[d] = neighbour
				if neighbour.costs == -1 {
					complete = false
					stack = append(stack, neighbour)
				} else {
					myCosts := neighbour.costs + 1
					if neighbour.direction != -1 && (neighbour.direction-d)%2 != 0 {
						myCosts = neighbour.costs + 1000
					}
					if cheapest == -1 || myCosts < cheapest {
						cheapest = myCosts
						cheapestDirection = direction
					}
				}
			}
		}

		if complete {
			next.costs = cheapest
			next.direction = cheapestDirection
			stack = stack[:len(stack)-1]
		}
	}
}
