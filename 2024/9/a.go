package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	type Sector struct {
		FileId    int
		NumBlocks int
		Prev      *Sector
		Next      *Sector
	}

	var checksumPreCalc = map[int]uint64{
		0: 0,
		1: 1,
		2: 3,
		3: 6,
		4: 10,
		5: 15,
		6: 21,
		7: 28,
		8: 36,
		9: 45,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	head := &Sector{-1, 0, nil, nil}
	tail := head
	isFree := true
	for scanner.Scan() {
		id := 0
		for _, char := range scanner.Text() {
			isFree = !isFree
			num, _ := strconv.ParseInt(string(char), 10, 64)
			sector := Sector{id, int(num), tail, nil}
			tail.Next = &sector
			tail = &sector
			if isFree {
				sector.FileId = -1
			} else {
				id++
			}
		}
	}
	//make sure we always have a free Sector at the end
	if !isFree {
		sector := Sector{-1, 0, tail, nil}
		tail.Next = &sector
		tail = &sector
	}

	current := head.Next
	currentPos := 0
	checksum := uint64(0)
	top := tail.Prev
	for current != tail {
		if current.FileId == -1 {
			//fmt.Printf("Found empty spot @ %d (%d)\n", currentPos, current.NumBlocks)
			// find next occupied sector from the end
			for top.FileId == -1 {
				top = top.Prev
			}
			if current.Next == tail {
				break
			}

			//fmt.Printf("Filling with %d (%d)\n", top.FileId, top.NumBlocks)
			if top.NumBlocks < current.NumBlocks {
				// top does not completely fill the free sector, insert new sector
				sector := &Sector{top.FileId, top.NumBlocks, current.Prev, current}
				current.Prev.Next = sector
				current.Prev = sector
				current.NumBlocks -= top.NumBlocks
				current = sector

				tail.NumBlocks += top.NumBlocks
				top.NumBlocks = 0
			} else {
				current.FileId = top.FileId
				top.NumBlocks -= current.NumBlocks
				tail.NumBlocks += current.NumBlocks
			}

			if top.NumBlocks == 0 {
				top.Prev.Next = tail
				top = top.Prev
			}
		}

		c := uint64(currentPos*current.NumBlocks) + checksumPreCalc[current.NumBlocks-1]
		checksum += uint64(current.FileId) * c
		currentPos += current.NumBlocks
		current = current.Next
	}

	fmt.Println(checksum)

	//res := ""
	//block := head.Next
	//for block != nil {
	//	if block.FileId == -1 {
	//		res += strings.Repeat(".", block.NumBlocks)
	//	} else {
	//		res += strings.Repeat(fmt.Sprintf("%d", block.FileId), block.NumBlocks)
	//	}
	//	block = block.Next
	//}
	//fmt.Println(res)

}
