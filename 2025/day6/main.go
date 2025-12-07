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

	// pt1(f)
	pt2(f)

}

func pt1(f *os.File) {
	problems := [][]string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := []string{}
		for _, v := range strings.Split(scanner.Text(), " ") {
			if v == "" {
				continue
			}
			line = append(line, strings.Trim(v, " "))
		}
		problems = append(problems, line)
	}

	sum := 0
	for i := range problems[0] {
		running_total := 0
		for _, line := range problems {
			value := line[i]
			operator := problems[len(problems)-1][i]

			if value == operator {
				continue
			}

			if running_total == 0 {
				running_total = mustAtoI(value)
				continue
			}

			running_total = op(running_total, mustAtoI(value), operator)
		}
		sum += running_total
	}
	log.Printf("sum: %d", sum)
}

func pt2(f *os.File) {
	lines := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ops := lines[len(lines)-1]
	sum := 0
	digits := []string{}

	for i := len(lines[0]) - 1; i >= 0; i-- {
		digit := ""
		for j := range len(lines) - 1 {
			char := lines[j][i]
			digit += string(char)
		}

		digits = append(digits, strings.Trim(digit, " "))
		digit = ""

		if i > len(ops)-1 {
			// last row might have less elems?
			continue
		}

		if string(ops[i]) != " " {
			res := 0
			for _, digit := range digits {
				if res == 0 {
					res = mustAtoI(digit)
					continue
				}
				res = op(res, mustAtoI(digit), string(ops[i]))
			}

			sum += res
			digits = []string{}

			// skip empty row right after the op
			i--
		}
	}

	log.Printf("sum %d", sum)
}

func op(v1, v2 int, operator string) int {
	switch operator {
	case "*":
		return v1 * v2
	case "+":
		return v1 + v2
	default:
		return 0

	}
}

func mustAtoI(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("could not atoi with %s", err)
	}
	return r
}
