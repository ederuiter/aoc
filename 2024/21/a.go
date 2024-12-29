package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func decode(sequence string, pad map[byte][2]int) string {
	location := pad['A']
	rev := map[[2]int]byte{}
	for i, loc := range pad {
		rev[loc] = i
	}

	res := ""
	for i, _ := range sequence {
		char := sequence[i]
		switch char {
		case '^':
			location[0] -= 1
		case 'v':
			location[0] += 1
		case '>':
			location[1] += 1
		case '<':
			location[1] -= 1
		case 'A':
			res += string(rev[location])
		}
	}
	return res
}

func moves(sequence string, pad map[byte][2]int, order string) string {
	//fmt.Printf("%+v\n", pad)
	prevLocation := pad['A']
	//avoid := pad['X']
	res := ""
	//prevChar := "A"
	for i, _ := range sequence {
		char := sequence[i]
		newLocation, ok := pad[char]
		if !ok {
			panic("uhhmm" + string(char))
		}
		diffY := newLocation[0] - prevLocation[0]
		diffX := newLocation[1] - prevLocation[1]
		//fmt.Printf("%+v %s => %s dY: %d dX: %d\n", newLocation, prevChar, string(char), diffY, diffX)
		for _, direction := range order {
			switch direction {
			case '<':
				if diffX < 0 {
					res += strings.Repeat("<", -diffX)
				}
			case '>':
				if diffX > 0 {
					res += strings.Repeat(">", diffX)
				}
			case '^':
				if diffY < 0 {
					res += strings.Repeat("^", -diffY)
				}
			case 'v':
				if diffY > 0 {
					res += strings.Repeat("v", diffY)
				}
			}
		}
		res += "A"

		prevLocation = newLocation
		//prevChar = string(char)
	}
	return res
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	/*
		+---+---+---+
		| 7 | 8 | 9 |
		+---+---+---+
		| 4 | 5 | 6 |
		+---+---+---+
		| 1 | 2 | 3 |
		+---+---+---+
		    | 0 | A |
		    +---+---+
	*/

	pad1 := map[byte][2]int{
		'X': {3, 0},
		'0': {3, 1},
		'A': {3, 2},
		'1': {2, 0},
		'2': {2, 1},
		'3': {2, 2},
		'4': {1, 0},
		'5': {1, 1},
		'6': {1, 2},
		'7': {0, 0},
		'8': {0, 1},
		'9': {0, 2},
	}
	order1 := "^>v<"
	/*
		    +---+---+
		    | ^ | A |
		+---+---+---+
		| < | v | > |
		+---+---+---+
	*/

	//EXAMPLE <v<A>>^AvA^A <vA  <   AA>>^ A  AvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
	//           <   A > A   v      <<    A
	//               ^   A
	//MINE    v<<A>>^AvA^A v<<A >>^ AAv<  A <A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A

	pad2 := map[byte][2]int{
		'X': {0, 0},
		'^': {0, 1},
		'A': {0, 2},
		'<': {1, 0},
		'v': {1, 1},
		'>': {1, 2},
	}
	order2 := "v<>^"

	totalComplexity := 0
	for scanner.Scan() {
		input := scanner.Text()
		numInput, _ := strconv.ParseInt(input[0:3], 10, 64)
		seq1 := moves(input, pad1, order1)
		seq2 := moves(seq1, pad2, order2)
		seq3 := moves(seq2, pad2, order2)
		complexity := len(seq3) * int(numInput)
		totalComplexity += complexity
		fmt.Printf("input: %s\n------------\nseq1: %s\nseq2: %s\nseq3: %s\ncomplexity: %d * %d = %d\n", input, seq1, seq2, seq3, len(seq3), int(numInput), complexity)
	}
	fmt.Println(totalComplexity)

	seq3 := "<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"
	seq2 := decode(seq3, pad2)
	seq1 := decode(seq2, pad2)
	input := decode(seq1, pad1)
	fmt.Printf("seq3: %s\nseq2: %s\nseq1: %s\ninput: %s\n", seq3, seq2, seq1, input)

	seq3 = "v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A"
	seq2 = decode(seq3, pad2)
	seq1 = decode(seq2, pad2)
	input = decode(seq1, pad1)
	fmt.Printf("seq3: %s\nseq2: %s\nseq1: %s\ninput: %s\n", seq3, seq2, seq1, input)
}
