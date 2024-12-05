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

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	total := uint64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			fmt.Println(match[0])
			i1, _ := strconv.ParseUint(match[1], 10, 64)
			i2, _ := strconv.ParseUint(match[2], 10, 64)
			total += (i1 * i2)
		}
	}

	fmt.Println(total)
}
