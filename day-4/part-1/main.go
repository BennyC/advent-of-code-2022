package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	defer f.Close()

	var input []string
	for scanner.Scan() {
		txt := scanner.Text()
		input = append(input, txt)
	}

	solution(input)
}

func solution(input []string) {
	total := 0
	for _, pair := range input {
		splits := strings.Split(pair, ",")
		a, b := splits[0], splits[1]

		rangeA, rangeB := strRange(a), strRange(b)
		if rangeA.Contains(rangeB) || rangeB.Contains(rangeA) {
			total++
		}
	}

	fmt.Println(total)
}

type minMax struct {
	Min int
	Max int
}

func (m minMax) Contains(a minMax) bool {
	return m.Min <= a.Min && m.Max >= a.Max
}

func strRange(str string) minMax {
	splits := strings.Split(str, "-")
	min, _ := strconv.Atoi(splits[0])
	max, _ := strconv.Atoi(splits[1])

	return minMax{
		Min: min,
		Max: max,
	}
}
