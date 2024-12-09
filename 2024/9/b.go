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

	var fileMap map[int]*Sector
	head := &Sector{-1, 0, nil, nil}
	tail := head
	isFree := true
	id := 0
	for scanner.Scan() {
		fileMap = make(map[int]*Sector, len(scanner.Text()))
		for _, char := range scanner.Text() {
			isFree = !isFree
			num, _ := strconv.ParseInt(string(char), 10, 64)
			sector := Sector{id, int(num), tail, nil}
			tail.Next = &sector
			tail = &sector
			if isFree {
				sector.FileId = -1
			} else {
				fileMap[id] = &sector
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

	for i := id - 1; i > 0; i-- {
		sector := fileMap[i]
		var freeSector *Sector
		current := sector.Prev
		for current != nil {
			if current.FileId == -1 && current.NumBlocks >= sector.NumBlocks {
				freeSector = current
			}
			current = current.Prev
		}
		if freeSector != nil {
			if freeSector.NumBlocks > sector.NumBlocks {
				newSector := &Sector{sector.FileId, sector.NumBlocks, freeSector.Prev, freeSector}
				freeSector.Prev.Next = newSector
				freeSector.Prev = newSector
				freeSector.NumBlocks -= sector.NumBlocks
			} else {
				freeSector.FileId = sector.FileId
			}
			//fmt.Printf("Found empty spot for %d after %d\n", sector.FileId, freeSector.Prev.FileId)

			sector.FileId = -1
			if sector.Prev.FileId == -1 {
				sector.NumBlocks += sector.Prev.NumBlocks
				sector.Prev = sector.Prev.Prev
				sector.Prev.Next = sector
			}
			if sector.Next != nil && sector.Next.FileId == -1 {
				sector.NumBlocks += sector.Next.NumBlocks
				sector.Next = sector.Next.Next
				if sector.Next != nil {
					sector.Next.Prev = sector
				}
			}
		}
	}

	current := head.Next
	currentPos := 0
	checksum := uint64(0)
	for current != nil {
		if current.FileId != -1 {
			c := uint64(currentPos*current.NumBlocks) + checksumPreCalc[current.NumBlocks-1]
			checksum += uint64(current.FileId) * c
		}
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
