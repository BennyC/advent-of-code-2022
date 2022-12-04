package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	defer f.Close()

	max := 0
	curr := 0

	for scanner.Scan() {
		txt := scanner.Text()
		next, _ := strconv.Atoi(txt)

		if curr >= max {
			max = curr
		}

		if txt == "" {
			curr = 0
		}

		curr += next
	}

	fmt.Println(max)
}
