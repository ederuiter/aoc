package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	antennas := map[byte][][2]int{}
	room := [][]string{}
	row := 0
	numCols := 0
	for scanner.Scan() {
		parts := strings.SplitAfter(scanner.Text(), "")
		room = append(room, parts)
		numCols = len(parts)
		for col, part := range parts {
			if part != "." {
				if _, ok := antennas[part[0]]; !ok {
					antennas[part[0]] = [][2]int{}
				}
				antennas[part[0]] = append(antennas[part[0]], [2]int{row, col})
			}
		}
		row++
	}
	numRows := row

	antiPositions := map[int]bool{}
	for antenna, locations := range antennas {
		fmt.Println(string(antenna))
		fmt.Printf("%+v\n", locations)
		for i := 0; i < len(locations)-1; i++ {
			loc1 := locations[i]
			for j := i + 1; j < len(locations); j++ {
				loc2 := locations[j]
				delta := [2]int{loc2[0] - loc1[0], loc2[1] - loc1[1]}

				k := 0
				for {
					a1 := [2]int{loc1[0] - (delta[0] * k), loc1[1] - (delta[1] * k)}
					if a1[0] >= 0 && a1[0] < numRows && a1[1] >= 0 && a1[1] < numCols {
						index := a1[0]*numCols + a1[1]
						antiPositions[index] = true
						room[a1[0]][a1[1]] = "#"
					} else {
						break
					}
					k++
				}

				k = 0
				for {
					a2 := [2]int{loc2[0] + (delta[0] * k), loc2[1] + (delta[1] * k)}
					if a2[0] >= 0 && a2[0] < numRows && a2[1] >= 0 && a2[1] < numCols {
						index := a2[0]*numCols + a2[1]
						antiPositions[index] = true
						room[a2[0]][a2[1]] = "#"
					} else {
						break
					}
					k++
				}
			}
		}
	}
	fmt.Println(len(antiPositions))
	for _, r := range room {
		fmt.Println(strings.Join(r, ""))
	}

}
