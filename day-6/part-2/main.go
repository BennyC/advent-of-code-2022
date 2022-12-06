package main

import (
	"bufio"
	"fmt"
	"os"
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
	for _, s := range input {
		for i := 0; i < len(s)-1; i++ {
			if allUnique(s[i : i+14]) {
				fmt.Println(i + 14)
				break
			}
		}
	}
}

func allUnique(s string) bool {
	chars := make(map[string]int, len(s))
	for _, c := range s {
		char := fmt.Sprintf("%c", c)
		chars[char]++

		if chars[char] > 1 {
			return false
		}
	}

	return true
}
