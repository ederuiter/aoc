package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don\'t\(\)`)
	total := uint64(0)
	enabled := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				i1, _ := strconv.ParseUint(match[1], 10, 64)
				i2, _ := strconv.ParseUint(match[2], 10, 64)
				total += (i1 * i2)
			}
		}
	}

	fmt.Println(total)
}
