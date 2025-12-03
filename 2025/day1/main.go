package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("could not open file %e", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	dial := 50
	counter := 0

	for scanner.Scan() {
		char := scanner.Text()
		dir := char[0]
		amount := mustAtoI(char[1:])

		if string(dir) == "L" {
			for range amount {
				if dial == 0 {
					counter += 1
					dial = 100
				}
				dial -= 1
			}
		} else {
			for range amount {
				if dial == 0 {
					counter += 1
				}
				if dial == 99 {
					dial = -1
				}
				dial += 1
			}
		}
	}
	log.Printf("result:  %d", counter)
}

func mustAtoI(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not atoi with %s", err)
	}
	return r
}
