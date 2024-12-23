package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	connections := map[string]map[string]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		computers := strings.SplitN(line, "-", 2)
		if _, ok := connections[computers[0]]; !ok {
			connections[computers[0]] = make(map[string]bool, 0)
		}
		if _, ok := connections[computers[1]]; !ok {
			connections[computers[1]] = make(map[string]bool, 0)
		}
		connections[computers[0]][computers[1]] = true
		connections[computers[1]][computers[0]] = true
	}

	num := 0
	largest := ""
	seen := map[string]bool{}
	for computer, cons := range connections {
		if computer[0] != 't' {
			continue
		}
		sets := make([][]string, 0)
		for con, _ := range cons {
			for _, set := range sets {
				matches := true
				for _, c := range set {
					if _, ok := connections[con][c]; !ok {
						matches = false
						break
					}
				}
				if matches {
					newSet := make([]string, len(set))
					copy(newSet, set)
					newSet = append(newSet, con)
					sets = append(sets, newSet)
				}
			}
			sets = append(sets, []string{con})
		}
		//fmt.Printf("%s => %d\n", computer, len(sets))
		for _, set := range sets {
			a := make([]string, len(set))
			copy(a, set)
			a = append(a, computer)
			slices.Sort(a)
			key := strings.Join(a, ",")
			if len(key) > len(largest) {
				largest = key
			}
			if len(set) == 2 {
				if _, ok := seen[key]; !ok {
					fmt.Printf("[%d] %s => %s\n", len(set), computer, strings.Join(set, ", "))
					seen[key] = true
					num++
				}
			}
		}
	}
	fmt.Println("num sets =>", num)
	fmt.Println("password =>", largest)
}
