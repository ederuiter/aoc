package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	StartX int
	StartY int
	vX     int
	vY     int
}

func parse(text string, re *regexp.Regexp) *Robot {
	match := re.FindStringSubmatch(text)
	x, _ := strconv.ParseInt(match[1], 10, 64)
	y, _ := strconv.ParseInt(match[2], 10, 64)
	vx, _ := strconv.ParseInt(match[3], 10, 64)
	vy, _ := strconv.ParseInt(match[4], 10, 64)
	return &Robot{int(x), int(y), int(vx), int(vy)}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	width := 101
	height := 103
	middleX := 50
	middleY := 51
	steps := 100
	robots := []*Robot{}
	q := []int{0, 0, 0, 0}

	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	for scanner.Scan() {
		r := parse(scanner.Text(), re)
		robots = append(robots, r)
	}

	for _, r := range robots {
		x := (((r.StartX + (steps * r.vX)) % width) + width) % width
		y := (((r.StartY + (steps * r.vY)) % height) + height) % height

		if x < middleX && y < middleY {
			q[0]++
		} else if x > middleX && y < middleY {
			q[1]++
		} else if x < middleX && y > middleY {
			q[2]++
		} else if x > middleX && y > middleY {
			q[3]++
		}
	}

	fmt.Printf("%+v\n", q)
	fmt.Println(q[0] * q[1] * q[2] * q[3])

	for steps = 1; steps < 100000; steps++ {
		world := map[int]bool{}
		for _, r := range robots {
			x := (((r.StartX + (steps * r.vX)) % width) + width) % width
			y := (((r.StartY + (steps * r.vY)) % height) + height) % height
			world[y*width+x] = true
		}
		lines := ""
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if world[y*width+x] {
					lines += "X"
				} else {
					lines += " "
				}
			}
			lines += "\n"
		}
		if strings.Contains(lines, "XXXXX") {
			fmt.Print("\033[H\033[2J")
			fmt.Print(lines)
			fmt.Printf(" => %d\n", steps)
		}
	}
}
