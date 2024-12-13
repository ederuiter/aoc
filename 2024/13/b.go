package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func win2(ax int64, ay int64, bx int64, by int64, prizeX int64, prizeY int64) (int64, int64, int64) {
	tokens := int64(0)
	winA := int64(0)
	winB := int64(0)

	maxA := max(prizeX/ax, prizeY/ay)
	for a := int64(0); a <= maxA; a++ {
		x := prizeX - (ax * a)
		y := prizeY - (ay * a)
		if x%bx == 0 {
			b := x / bx
			if y-(b*by) == 0 {
				costs := 3*a + b
				if tokens == 0 || costs < tokens {
					tokens = costs
					winA = a
					winB = b
				}
			}
		}
	}
	return tokens, winA, winB
}

func parse2(text string, offset int64, re *regexp.Regexp) (int64, int64) {
	match := re.FindStringSubmatch(text)
	x, _ := strconv.ParseInt(match[1], 10, 64)
	y, _ := strconv.ParseInt(match[2], 10, 64)
	return x + offset, y + offset
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	prizes := int64(0)
	costs := int64(0)
	re := regexp.MustCompile(`X.(\d+), Y.(\d+)`)
	for scanner.Scan() {
		ax, ay := parse2(scanner.Text(), 0, re)
		scanner.Scan()
		bx, by := parse2(scanner.Text(), 0, re)
		scanner.Scan()
		prizeX, prizeY := parse2(scanner.Text(), 10000000000000, re)
		scanner.Scan()
		fmt.Printf("a: %d %d\nb: %d %d\nprize: %d %d\n", ax, ay, bx, by, prizeX, prizeY)
		tokens, a, b := win2(ax, ay, bx, by, prizeX, prizeY)
		if tokens > 0 {
			costs += tokens
			prizes++
			fmt.Printf("can be won with %d tokens (A: %d B: %d)\n", tokens, a, b)
		}
	}
	fmt.Printf("You can win %d prizes with %d tokens\n", prizes, costs)
}
