package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func permutations(num int, len int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		res := make([]int, len)
		if !yield(res) {
			return
		}
		numPermutations := uint64(math.Pow(float64(num), float64(len))) - 1
		for k := uint64(0); k < numPermutations; k++ {
			index := 0
			for {
				res[index] += 1
				if res[index] < num {
					break
				} else {
					res[index] = 0
					index += 1
				}
			}
			if !yield(res) {
				return
			}
		}
	}
}

func checkOperators(numbers []uint64, total uint64, numOperators int) bool {
	for operators := range permutations(numOperators, len(numbers)-1) {
		calculated := numbers[0]
		str := fmt.Sprintf("%d = %d", total, calculated)
		for k := 1; k < len(numbers); k++ {
			switch operators[k-1] {
			case 0:
				calculated *= numbers[k]
				str += fmt.Sprintf(" %s %d", "*", numbers[k])
			case 1:
				calculated += numbers[k]
				str += fmt.Sprintf(" %s %d", "+", numbers[k])
			case 2:
				calculated, _ = strconv.ParseUint(fmt.Sprintf("%d%d", calculated, numbers[k]), 10, 64)
				str += fmt.Sprintf(" %s %d", "||", numbers[k])
			}
		}
		if total == calculated {
			fmt.Println("=> " + str)
			return true
		} else {
			//fmt.Println("!" + str)
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sumA := uint64(0)
	sumB := uint64(0)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), ": ", 2)
		total, _ := strconv.ParseUint(parts[0], 10, 64)
		numbers := []uint64{}
		parts = strings.Split(parts[1], " ")
		for _, part := range parts {
			number, _ := strconv.ParseUint(part, 10, 64)
			numbers = append(numbers, number)
		}

		if checkOperators(numbers, total, 2) {
			sumA += total
			sumB += total
		} else if checkOperators(numbers, total, 3) {
			sumB += total
		}
	}

	fmt.Println(sumA)
	fmt.Println(sumB)
}
