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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		l1, _ := strconv.ParseInt(parts[0], 10, 64)
		l2, _ := strconv.ParseInt(parts[1], 10, 64)
		if l1 == l2 {
			continue
		}
		min_diff := int64(1)
		max_diff := int64(3)
		if l1 > l2 {
			min_diff = -3
			max_diff = -1
		}
		is_safe := true
		for index, s2 := range parts[1:] {
			l2, _ = strconv.ParseInt(s2, 10, 64)
			l1, _ = strconv.ParseInt(parts[index], 10, 64)
			if (l2-l1) < min_diff || (l2-l1) > max_diff {
				is_safe = false
				break
			}
		}
		if is_safe {
			fmt.Println(scanner.Text())
			safe++
		}
	}

	fmt.Println(safe)
}
