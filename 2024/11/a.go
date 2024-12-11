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
	var calc func(stone string, numBlinks int, cache map[string]map[int]int) int
	calc = func(stone string, numBlinks int, cache map[string]map[int]int) int {
		key := stone
		if _, ok := cache[key]; !ok {
			cache[key] = make(map[int]int)
		}
		if res, ok := cache[key][numBlinks]; ok {
			return res
		} else {
			num := 1
			for blink := 0; blink < numBlinks; blink++ {
				if stone == "0" || stone == "" {
					stone = "1"
				} else {
					l := len(stone)
					if l%2 == 0 {
						mid := l / 2
						num += calc(strings.TrimLeft(stone[mid:], "0"), numBlinks-(blink+1), cache)
						stone = stone[:mid]
					} else {
						n, _ := strconv.ParseUint(stone, 10, 64)
						stone = strconv.FormatUint(n*2024, 10)
					}
				}
			}
			cache[key][numBlinks] = num
			return num
		}
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	calculated := map[string]map[int]int{}

	numBlinks := 25
	numStones := 0
	for _, stone := range parts {
		numStones += calc(stone, numBlinks, calculated)
	}
	fmt.Println(numStones)

	numBlinks = 75
	numStones = 0
	for _, stone := range parts {
		numStones += calc(stone, numBlinks, calculated)
	}
	fmt.Println(numStones)
}
