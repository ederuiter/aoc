package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var word = "XMAS"

var directions = map[byte][]int{
	'E': {1, 0},
	'W': {-1, 0},
	'N': {0, 1},
	'S': {0, -1},
	'Y': {1, 1},
	'K': {-1, 1},
	'Z': {1, -1},
	'X': {-1, -1},
}

func hasWord(matrix [][]byte, x int, y int, direction []int) bool {
	maxX := len(matrix)
	maxY := len(matrix[0])
	for i := 1; i < len(word); i++ {
		x = x + direction[0]
		y = y + direction[1]
		if x < 0 || y < 0 || x >= maxX || y >= maxY || matrix[x][y] != word[i] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	matrix := [][]byte{}
	for scanner.Scan() {
		matrix = append(matrix, []byte(scanner.Text()))
	}

	var total = int64(0)
	for x, row := range matrix {
		for y, char := range row {
			if char == word[0] {
				for _, direction := range directions {
					if hasWord(matrix, x, y, direction) {
						total++
					}
				}
			}
		}
	}

	fmt.Println(total)
}
