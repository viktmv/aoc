package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input_test")

	if err != nil {
		log.Fatalf("could not open file %e", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	ranges := [][]int{}
	ingredients := []int{}
	isIngredientLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// log.Printf("newline")
			isIngredientLine = true
			continue
		}
		if isIngredientLine {
			ingredients = append(ingredients, mustAtoI(line))
			continue
		}
		// ranges
		rangeLimits := strings.Split(line, "-")

		start := mustAtoI(rangeLimits[0])
		end := mustAtoI(rangeLimits[1])
		ranges = append(ranges, []int{start, end})
	}

	// pt1: naive approach - iterate and check
	counter := 0
	seen := map[int]bool{}
	for i := range ingredients {
		for j := range ranges {
			start, end := ranges[j][0], ranges[j][1]
			ingredient := ingredients[i]
			if !seen[ingredient] && ingredient >= start && ingredient <= end {
				counter++
				seen[ingredient] = true
			}
		}
	}

	// log.Printf("ingredients %v", ingredients)
	// log.Printf("ranges %v", ranges)
	// log.Printf("fresh %d", counter)

	// pt2 - same naive approach
	// seen2 := map[int]bool{}
	freshIDs := 0
	for i := range ranges {
		start, end := ranges[i][0], ranges[i][1]

		if i == 0 {
			log.Printf("fresh id %d", freshIDs)
			freshIDs += end - start + 1
			continue
		}

		for j := range i - 1 {
			start2, end2 := ranges[j][0], ranges[j][1]
			log.Printf("fresh id %d", freshIDs)

			if start > start2 && end < end2 {
				log.Printf("inside overlap %d, %d", start, start2)
				// inside overlap, skip
				continue
			} else if start < end2 && end > end2{
				// high overlap
				log.Printf("high overlap of %d, %d with %d, %d", start, end, start2, end2)
				freshIDs += end - end2
				break
			} else if start < start2 && end > end2 {
				// full outside overlap
				log.Printf("outside overlap of %d, %d with %d, %d", start, end, start2, end2)
				freshIDs -= end - start 
				break
			} else if end > start2 && start < start2{
				log.Printf("low overlap of %d, %d with %d, %d", start, end, start2, end2)
				// low overlap
				freshIDs += end - start2
				break
			} else {
				log.Printf("no overlap of %d, %d with %d, %d", start, end, start2, end2)
				// no overlap
				freshIDs += end - start + 1
				break
			}

		}
	}

	log.Printf("fresh IDs %d", freshIDs)
	// log.Printf("ranges with total %v", rangesWithTotal)
}

// p2 - collapse the ranges
// 3-5
// 10-14
// 16-20
// 12-18
// 12-14
// 10-22

// keep track of the largest offset
// smallest1 3, largest1 5
// largest2 14 , smallest2 3
// largest 20, smallest 3

// no overlap

// if  start2 > end
// if end2 > start

// inside overlap
// if start2 > start and end2 < end
//  -> discard
//
//  // high overlap
// if start2 > start and end2 > end
//  -> + end - end
//
//  // full outside overlap
// if start2 < start and end2  > end
//  ->  - original start
//
//  // low overlap
// if start2 < start and end2 < end
//  ->  + (start2 - start1)

func mustAtoI(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not atoi with %s", err)
	}
	return r
}
