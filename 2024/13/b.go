package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func win2(ax int64, ay int64, bx int64, by int64, prizeX int64, prizeY int64) (int64, int64, int64) {
	/*
			(ax*a)+(bx*b) == prizeX
			a=(priceX - bx*b)/ax
			a=(priceX/ax) - b*(bx/ax)

			(ay*a)+(by*b) == prizeY
			a=(priceY -by*b)/ay
			a=(priceY/ay - b*(by/ay)

			(priceX/ax) - b*(bx/ax) = (priceY/ay) - b*(by/ay)
		    b = ((priceY/ay) - (priceX/ax)) / ((by/ay)-(bx/ax))

	*/

	b := int64(math.Round(((float64(prizeY) / float64(ay)) - (float64(prizeX) / float64(ax))) / ((float64(by) / float64(ay)) - (float64(bx) / float64(ax)))))
	a := int64(math.Round((float64(prizeX) / float64(ax)) - float64(b)*(float64(bx)/float64(ax))))

	if (ax*a)+(bx*b) == prizeX && (ay*a)+(by*b) == prizeY {
		t := (3 * a) + b
		return t, a, b
	}

	return 0, -1, -1
}

func parse2(text string, offset int64, re *regexp.Regexp) (int64, int64) {
	match := re.FindStringSubmatch(text)
	x, _ := strconv.ParseInt(match[1], 10, 64)
	y, _ := strconv.ParseInt(match[2], 10, 64)
	return x + offset, y + offset
}

func main() {
	file, err := os.Open("test.txt")
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
