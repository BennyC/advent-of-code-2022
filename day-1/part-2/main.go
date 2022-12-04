package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	defer f.Close()

	var totals []int

	curr := 0
	for scanner.Scan() {
		txt := scanner.Text()
		next, _ := strconv.Atoi(txt)

		if txt == "" {
			totals = append(totals, curr)
			curr = 0

			continue
		}

		curr += next
	}

	sort.Ints(totals)

	total := 0
	for _, x := range totals[len(totals)-3:] {
		total += x
	}

	fmt.Println(total)
}
