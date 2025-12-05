package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("could not open file %e", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	banks := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		bank := strings.Split(line, "")
		banks = append(banks, bank)
	}

	log.Printf("banks %v\n", banks)

    sum := 0
	// day1
	// for i := range(banks) {
	// 	bank := banks[i]
	// 	maxPair := 0
	// 	for j := range(bank) {
	// 		for k := j+1; k < len(bank); k++ {
	// 			pair := mustAtoI(bank[j] + bank[k])
	// 			if pair > maxPair {
	// 				maxPair = pair
	// 			}
	// 		}
	// 	}
	// 	sum += maxPair
	// 	log.Printf("maxPair: %d", maxPair)
	// }

	// day 2
	for i := range(banks) {
		bank := banks[i]
		maxValue := ""
		maxNum := 0
		maxIdx := -1
		for k := 12; k > 0; k-- {
			for j := maxIdx + 1; j <= len(bank)-k; j++ {
				if maxNum < mustAtoI(bank[j]) {
					maxNum = mustAtoI(bank[j])
					maxIdx = j
				}
			}

			maxValue += strconv.Itoa(maxNum)
			maxNum = 0
		}

		log.Printf("bank %v", maxValue)
		sum += mustAtoI(maxValue)
	}

	log.Printf("sum: %d", sum)
}

func mustAtoI(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not atoi with %s", err)
	}
	return r
}

func toIntPair(s1, s2 int) int {
	return mustAtoI(strconv.Itoa(s1) + strconv.Itoa(s2))
}
