package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(parts []string) (bool, int) {
	l1, _ := strconv.ParseInt(parts[0], 10, 64)
	l2, _ := strconv.ParseInt(parts[1], 10, 64)
	if l1 == l2 {
		return false, 1
	}
	min_diff := int64(1)
	max_diff := int64(3)
	if l1 > l2 {
		min_diff = -3
		max_diff = -1
	}

	for index, s2 := range parts[1:] {
		l2, _ = strconv.ParseInt(s2, 10, 64)
		l1, _ = strconv.ParseInt(parts[index], 10, 64)
		if (l2-l1) < min_diff || (l2-l1) > max_diff {
			return false, index + 1
		}
	}
	return true, 0
}

func remove(slice []string, s int) []string {
	s2 := append([]string{}, slice...)
	return append(s2[:s], s2[s+1:]...)
}

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
		p := parts
		is_safe, _ := isSafe(parts)
		if !is_safe {
			for i := 0; i < len(parts); i++ {
				t := remove(parts, i)
				is_safe, _ = isSafe(t)
				if is_safe {
					break
				}
			}
		}

		if is_safe {
			fmt.Println("Safe: " + strings.Join(p, " "))
			safe++
		} else {
			fmt.Printf("Not safe at all: %s\n", scanner.Text())
		}
	}

	fmt.Println(safe)
}
