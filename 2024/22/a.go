package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := int64(0)
	changes := map[[4]int64]int{}
	for scanner.Scan() {
		secret, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		//fmt.Printf("%d: ", secret)
		buyerPrices := []int64{secret}
		seen := map[[4]int64]bool{}
		for i := 1; i <= 2000; i++ {
			secret = (secret ^ (secret * 64)) % 16777216
			secret = (secret ^ (secret / 32)) % 16777216
			secret = (secret ^ (secret * 2048)) % 16777216
			price := secret % 10
			buyerPrices = append(buyerPrices, price)
			if i > 3 {
				myChanges := [4]int64{buyerPrices[i-3] - buyerPrices[i-4], buyerPrices[i-2] - buyerPrices[i-3], buyerPrices[i-1] - buyerPrices[i-2], buyerPrices[i] - buyerPrices[i-1]}
				if _, ok := seen[myChanges]; ok {
					continue
				}
				if _, ok := changes[myChanges]; !ok {
					changes[myChanges] = 0
				}
				changes[myChanges] += int(price)
				seen[myChanges] = true
			}
		}
		//fmt.Printf("%d\n", secret)
		total += secret
	}
	fmt.Println(" =>", total)
	maxBananas := 0
	var maxChange [4]int64
	for change, bananas := range changes {
		if bananas > maxBananas {
			maxBananas = bananas
			maxChange = change
		}
	}
	fmt.Println(" =>", maxBananas)
	fmt.Printf("%+v\n", maxChange)
}
