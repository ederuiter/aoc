package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func win(ax int, ay int, bx int, by int, prizeX int, prizeY int) (int, int, int) {
	tokens := 0
	winA := 0
	winB := 0
	for a := 0; a <= 100; a++ {
		x := prizeX - (ax * a)
		y := prizeY - (ay * a)
		if x%bx == 0 {
			b := x / bx
			if b <= 100 && y-(b*by) == 0 {
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

func parse(text string, re *regexp.Regexp) (int, int) {
	match := re.FindStringSubmatch(text)
	x, _ := strconv.ParseInt(match[1], 10, 64)
	y, _ := strconv.ParseInt(match[2], 10, 64)
	return int(x), int(y)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	prizes := 0
	costs := 0
	re := regexp.MustCompile(`X.(\d+), Y.(\d+)`)
	for scanner.Scan() {
		ax, ay := parse(scanner.Text(), re)
		scanner.Scan()
		bx, by := parse(scanner.Text(), re)
		scanner.Scan()
		prizeX, prizeY := parse(scanner.Text(), re)
		scanner.Scan()
		//fmt.Printf("a: %d %d\nb: %d %d\nprize: %d %d\n", ax, ay, bx, by, prizeX, prizeY)
		tokens, a, b := win(ax, ay, bx, by, prizeX, prizeY)
		if tokens > 0 {
			costs += tokens
			prizes++
			fmt.Printf("can be won with %d tokens (A: %d B: %d)\n", tokens, a, b)
		}
	}
	fmt.Printf("You can win %d prizes with %d tokens\n", prizes, costs)
}
