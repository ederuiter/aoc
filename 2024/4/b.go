package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func hasXMas(matrix [][]byte, x int, y int) bool {
	maxX := len(matrix) - 2
	maxY := len(matrix[0]) - 2
	if x < 1 || y < 1 || x > maxX || y > maxY {
		return false
	}

	if ((matrix[x-1][y-1] == 'M' && matrix[x+1][y+1] == 'S') || (matrix[x-1][y-1] == 'S' && matrix[x+1][y+1] == 'M')) && ((matrix[x+1][y-1] == 'M' && matrix[x-1][y+1] == 'S') || (matrix[x+1][y-1] == 'S' && matrix[x-1][y+1] == 'M')) {
		return true
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
	matrix := [][]byte{}
	for scanner.Scan() {
		matrix = append(matrix, []byte(scanner.Text()))
	}

	var total = int64(0)
	for x, row := range matrix {
		for y, char := range row {
			if char == 'A' {
				if hasXMas(matrix, x, y) {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}
