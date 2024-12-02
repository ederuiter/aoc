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
	l2 := map[uint64]int{}
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "   ", 2)
		id1, _ := strconv.ParseUint(parts[0], 10, 64)
		id2, _ := strconv.ParseUint(parts[1], 10, 64)
		l1 = append(l1, id1)
		if _, ok := l2[id2]; ok {
			l2[id2]++
		} else {
			l2[id2] = 1
		}
	}
	slices.Sort(l1)

	total := uint64(0)
	for _, id1 := range l1 {
		num, ok := l2[id1]
		if ok {
			total += id1 * uint64(num)
		}
	}
	fmt.Println(total)
}
