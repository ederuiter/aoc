package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func arrangements(pos int, possible [][]int, towels []string, cache map[int]int) int {
	if cached, ok := cache[pos]; ok {
		return cached
	}

	total := 0
	length := len(possible)
	if pos == length {
		total++
		goto ret
	} else if pos > length || len(possible[pos]) == 0 {
		goto ret
	}

	for _, p := range possible[pos] {
		l := len(towels[p])
		total += arrangements(pos+l, possible, towels, cache)
	}
ret:
	cache[pos] = total
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numArrangements := 0
	numPossible := 0
	towels := []string{}
	parsedTowels := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsedTowels = true
		} else if !parsedTowels {
			towels = append(towels, strings.Split(line, ", ")...)
		} else {
			possible := make([][]int, len(line))
			for i := 0; i < len(line); i++ {
				possible[i] = []int{}
			}
			fmt.Printf("Checking pattern: %s", line)
			for towelIndex, towel := range towels {
				offset := -1
				for {
					index := strings.Index(line[offset+1:], towel)
					if index != -1 {
						offset = index + offset + 1
						possible[offset] = append(possible[offset], towelIndex)
					} else {
						break
					}
				}
			}

			cache := map[int]int{}
			myArrangements := arrangements(0, possible, towels, cache)
			fmt.Printf(" => %d arrangements\n", myArrangements)
			if myArrangements > 0 {
				numArrangements += myArrangements
				numPossible++
			}
		}
	}
	fmt.Printf("%d possible patterns with %d possible arrangments\n", numPossible, numArrangements)
}
