package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineType := ""
	locks := [][5]int{}
	locksByPosAndHeight := [5][7][]int{
		0: {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
		1: {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
		2: {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
		3: {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
		4: {0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}},
	}

	keys := [][5]int{}
	current := [5]int{0, 0, 0, 0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if lineType == "lock" {
				locks = append(locks, current)
				for index, depth := range current {
					maxKeyHeight := 7 - depth
					for h := 1; h <= maxKeyHeight; h++ {
						locksByPosAndHeight[index][h] = append(locksByPosAndHeight[index][h], len(locks)-1)
					}
				}
			} else if lineType == "key" {
				keys = append(keys, current)
			}
			lineType = ""
			current = [5]int{0, 0, 0, 0, 0}
		} else {
			if lineType == "" {
				if line[0] == '.' {
					lineType = "key"
				} else {
					lineType = "lock"
				}
			}
			for index, char := range line {
				if char == '#' {
					current[index]++
				}
			}
		}
	}
	if lineType == "lock" {
		locks = append(locks, current)
		for index, depth := range current {
			maxKeyHeight := 7 - depth
			for h := 1; h <= maxKeyHeight; h++ {
				locksByPosAndHeight[index][h] = append(locksByPosAndHeight[index][h], len(locks)-1)
			}
		}
	} else if lineType == "key" {
		keys = append(keys, current)
	}

	//fmt.Printf("%v\n", locks)
	//fmt.Printf("%v\n", keys)
	//fmt.Printf("%v\n", locksByPosAndHeight)

	fittingCombination := 0
	for _, key := range keys {
		possible := map[int]int{}
		for pos, height := range key {
			for _, lock := range locksByPosAndHeight[pos][height] {
				if _, ok := possible[lock]; !ok {
					possible[lock] = 1
				} else {
					possible[lock] += 1
				}
			}
		}
		for _, num := range possible {
			if num == 5 {
				fittingCombination += 1
			}
		}
	}
	fmt.Println(fittingCombination)
}
