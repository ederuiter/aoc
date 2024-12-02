package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	l1 := []uint64{}
	l2 := []uint64{}
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "   ", 2)
		id1, _ := strconv.ParseUint(parts[0], 10, 64)
		id2, _ := strconv.ParseUint(parts[1], 10, 64)
		l1 = append(l1, id1)
		l2 = append(l2, id2)
	}
	slices.Sort(l1)
	slices.Sort(l2)

	total := uint64(0)
	for index, id1 := range l1 {
		id2 := l2[index]
		if id1 > id2 {
			total += (id1 - id2)
		} else {
			total += (id2 - id1)
		}
	}
	fmt.Println(total)
}
