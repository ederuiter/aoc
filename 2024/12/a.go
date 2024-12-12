package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Region struct {
	Perimeter int
	Corners   int
	Tiles     []*Tile
}

type Tile struct {
	Type   string
	Region *Region
}

func getNeighbours(world map[int]*Tile, pos int, numCols int) []*Tile {
	res := make([]*Tile, 0)
	tile := world[pos]
	if neighbour, ok := world[pos-1]; pos%numCols != 0 && ok && neighbour.Type == tile.Type {
		res = append(res, neighbour)
	}
	if neighbour, ok := world[pos+1]; pos%numCols != (numCols-1) && ok && neighbour.Type == tile.Type {
		res = append(res, neighbour)
	}
	if neighbour, ok := world[pos+numCols]; ok && neighbour.Type == tile.Type {
		res = append(res, neighbour)
	}
	if neighbour, ok := world[pos-numCols]; ok && neighbour.Type == tile.Type {
		res = append(res, neighbour)
	}
	return res
}

func corners(world map[int]*Tile, pos int, numCols int) int {
	res := 0
	tile := world[pos]

	/*
			 looks for inner && outer corners:
			 .X.
			 XOX
			 .X.
			 inner: check each X.X in each corner
		     outer: check for ... on each corner
	*/
	offsets := [][3]int{{-numCols, -numCols + 1, 1}, {1, numCols + 1, numCols}, {numCols, numCols - 1, -1}, {-1, -numCols - 1, -numCols}}

	for _, tileOffsets := range offsets {
		pattern := ""
		for _, offset := range tileOffsets {
			newPos := pos + offset
			collDiff := (newPos % numCols) - (pos % numCols)
			t, ok := world[newPos]
			if !ok || collDiff < -1 || collDiff > 1 {
				pattern += "0"
			} else if t.Region == tile.Region {
				pattern += "1"
			} else {
				pattern += "0"
			}
		}
		if pattern == "010" || pattern == "101" || pattern == "000" {
			res++
		}
	}
	return res

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	world := map[int]*Tile{}
	regions := []*Region{}

	row := 0
	numCols := 0
	for scanner.Scan() {
		parts := strings.SplitAfter(scanner.Text(), "")
		numCols = len(parts)
		for col, tileType := range parts {
			position := (row * numCols) + col
			tile := &Tile{
				Type: tileType,
			}
			world[position] = tile
		}
		row++
	}

	for pos := 0; pos < len(world); pos++ {
		tile := world[pos]
		neighbours := getNeighbours(world, pos, numCols)
		perimeter := 4 - len(neighbours)
		var region *Region
		for _, neighbour := range neighbours {
			if neighbour.Region == nil {
				continue
			} else if region == nil {
				region = neighbour.Region
			} else if region != neighbour.Region {
				// merge regions
				r := neighbour.Region
				neighbour.Region = region
				region.Perimeter += r.Perimeter
				for _, t := range r.Tiles {
					region.Tiles = append(region.Tiles, t)
					t.Region = region
				}
				r.Tiles = []*Tile{}
				r.Perimeter = 0
			}
		}
		if region == nil {
			region = &Region{Tiles: []*Tile{}, Perimeter: 0, Corners: 0}
			regions = append(regions, region)
			for _, neighbour := range neighbours {
				neighbour.Region = region
			}
		}
		tile.Region = region
		region.Perimeter += perimeter
		region.Tiles = append(region.Tiles, tile)
	}

	for pos := 0; pos < len(world); pos++ {
		numCorners := corners(world, pos, numCols)
		world[pos].Region.Corners += numCorners
	}

	costs := 0
	for _, region := range regions {
		area := len(region.Tiles)
		if area == 0 {
			// skip empty regions (due to merge)
			continue
		}
		costs += area * region.Perimeter
		fmt.Printf("A region of %s plants with price %d * %d = %d\n", region.Tiles[0].Type, area, region.Perimeter, area*region.Perimeter)
	}
	fmt.Println(costs)

	fmt.Println("With bulk discount:")
	costs = 0
	for _, region := range regions {
		area := len(region.Tiles)
		if area == 0 {
			// skip empty regions (due to merge)
			continue
		}
		costs += area * region.Corners
		fmt.Printf("A region of %s plants with price %d * %d = %d\n", region.Tiles[0].Type, area, region.Corners, area*region.Corners)
	}
	fmt.Println(costs)
}
