package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("could not open file %e", err)
	}

	trimmed := strings.TrimSpace(string(data))
	pairs := [][]string{}

	ranges := strings.Split(trimmed, ",")
	for i := range ranges {
		r := strings.Split(ranges[i], "-")
		start := r[0]
		end := r[1]
		pairs = append(pairs, []string{start, end})
	}

	sum := 0
	matches := []int{}
	for i := range pairs {
		pair := pairs[i]

		for value := mustAtoI(pair[0]); value <= mustAtoI(pair[1]); value++ {
			s := strconv.Itoa(value)
			// pt 1
			//	if len(s)%2 != 0 {
			//		continue
			//	}

			//	first := s[:len(s)/2]
			//	second := s[len(s)/2:]
			//	if first == second {
			//		log.Printf("invalid id %s", s)
			//		sum += value
			//	}

			// pt 2
			for k := 1; k <= len(s) / 2; k += 1 {
				pattern := s[:k]

				if len(pattern) == len(s) {
					break
				}

				match := true
				for j := k; j <= len(s)-k; j += k {
					// if any part of the value doesn't match
					if s[j:j+k] != pattern {
						match = false
					}
				}

				// if we cannot evenly divide - it's not a match
				if len(s) % len(pattern) != 0 {
					match = false
				}

				if match {
					matches = append(matches, value)
					log.Printf("🚨 matched pattern %s, string: %s", pattern, s)
				}
			}
		}
	}

	dups := map[int]bool{}
	for i := range(matches) {
		if dups[matches[i]] {
			continue
		}

		dups[matches[i]] = true
		sum += matches[i]
	}
	log.Printf("matches %v\n", matches)
	log.Printf("result %d\n", sum)
}

func mustAtoI(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not atoi with %s", err)
	}
	return r
}
