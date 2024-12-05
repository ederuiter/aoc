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

func parseRow(line string, sep string) []int {
	var parts []string
	if sep == "" {
		parts = strings.SplitAfter(line, "")
	} else {
		parts = strings.Split(line, sep)
	}

	res := make([]int, len(parts))
	for i, part := range parts {
		p, _ := strconv.ParseInt(part, 10, 64)
		res[i] = int(p)
	}
	return res
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsing_rules := true
	rules := map[int]map[int]int{}
	total := uint64(0)
	total_unsorted := uint64(0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			parsing_rules = false
			continue
		} else if parsing_rules {
			row := parseRow(scanner.Text(), "|")
			if _, ok := rules[row[0]]; !ok {
				rules[row[0]] = map[int]int{}
			}
			rules[row[0]][row[1]] = 1

			if _, ok := rules[row[1]]; !ok {
				rules[row[1]] = map[int]int{}
			}
			rules[row[1]][row[0]] = -1
		} else {
			row := parseRow(scanner.Text(), ",")
			positions := map[int]int{}
			for i, item := range row {
				positions[item] = i
			}

			fmt.Println("Checking row: " + scanner.Text())
			matches_rules := true
			for pos, item := range row {
				for item2, rule := range rules[item] {
					pos2, ok := positions[item2]
					if !ok {
						continue
					}

					if rule > 0 && pos2 < pos {
						fmt.Printf("Not matching found %d before %d\n", item2, item)
						matches_rules = false
						break
					}

					if rule < 0 && pos2 > pos {
						fmt.Printf("Not matching found %d after %d\n", item2, item)
						matches_rules = false
						break
					}
				}
				if !matches_rules {
					break
				}
			}

			if matches_rules {
				fmt.Println(" => ok")
				total += uint64(row[((len(row) - 1) / 2)])
			} else {
				slices.SortFunc(row, func(a, b int) int {
					if _, ok := rules[a][b]; !ok {
						return 0
					}
					return rules[a][b]
				})
				fmt.Printf(" => ok: %+v\n", row)

				total_unsorted += uint64(row[((len(row) - 1) / 2)])
			}
		}
	}

	fmt.Println(total)
	fmt.Println(total_unsorted)
}
