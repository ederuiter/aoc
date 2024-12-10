package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	type Tile struct {
		Height    int
		NumRoutes int
		Tops      map[int]bool
		Position  int
	}

	world := map[int]*Tile{}
	index := map[int][]*Tile{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	getIndex := func(row int, col int, numCols int) int {
		return (row * numCols) + col
	}

	row := 0
	numCols := 0
	for scanner.Scan() {
		parts := strings.SplitAfter(scanner.Text(), "")
		numCols = len(parts)
		for col, part := range parts {
			position := getIndex(row, col, numCols)
			height, _ := strconv.ParseInt(part, 10, 64)
			tile := &Tile{
				Height:    int(height),
				NumRoutes: 0,
				Tops:      map[int]bool{},
				Position:  position,
			}
			if int(height) == 9 {
				tile.NumRoutes = 1
				tile.Tops[position] = true
			}
			world[position] = tile
			index[int(height)] = append(index[int(height)], tile)
		}
		row++
	}

	total := 0
	rating := 0
	for i := 8; i >= 0; i-- {
		for _, tile := range index[i] {
			pos := tile.Position

			if neighbour, ok := world[pos-1]; pos%numCols != 0 && ok && neighbour.Height == i+1 {
				//fmt.Printf("neighbour[L]: %+v\n", neighbour)
				tile.NumRoutes += neighbour.NumRoutes
				for top, _ := range neighbour.Tops {
					tile.Tops[top] = true
				}
			}
			if neighbour, ok := world[pos+1]; pos%numCols != (numCols-1) && ok && neighbour.Height == i+1 {
				//fmt.Printf("neighbour[R]: %+v\n", neighbour)
				tile.NumRoutes += neighbour.NumRoutes
				for top, _ := range neighbour.Tops {
					tile.Tops[top] = true
				}
			}
			if neighbour, ok := world[pos+numCols]; ok && neighbour.Height == i+1 {
				//fmt.Printf("neighbour[D]: %+v\n", neighbour)
				tile.NumRoutes += neighbour.NumRoutes
				for top, _ := range neighbour.Tops {
					tile.Tops[top] = true
				}
			}
			if neighbour, ok := world[pos-numCols]; ok && neighbour.Height == i+1 {
				//fmt.Printf("neighbour[U]: %+v\n", neighbour)
				tile.NumRoutes += neighbour.NumRoutes
				for top, _ := range neighbour.Tops {
					tile.Tops[top] = true
				}
			}
			if i == 0 {
				//fmt.Printf("%+v => %d [%d]\n", tile, len(tile.Tops), total)
				total += len(tile.Tops)
				rating += tile.NumRoutes
			}
		}
	}

	fmt.Println(total)
	fmt.Println(rating)
}
