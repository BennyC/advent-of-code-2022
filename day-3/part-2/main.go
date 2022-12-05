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

	total := 0
	for i := 0; i < len(input); i += 3 {
		x, y, z := input[i], input[i+1], input[i+2]
		fmt.Println(x, y, z)
		char := repeatedChar(x, y, z)

		total += index(char)
		fmt.Println(index(char), char)
	}

	fmt.Println(total)
}

func repeatedChar(strs ...string) string {
	l := len(strs)
	m := make(map[string][]bool, len(alpha))

	for _, a := range alpha {
		m[a] = make([]bool, l)
	}

	for i, str := range strs {
		for x := 0; x < len(str); x++ {
			char := fmt.Sprintf("%c", str[x])
			m[char][i] = true

			if len(filter(m[char])) == l {
				return char
			}
		}
	}

	return ""
}

func filter(bools []bool) []bool {
	var r []bool
	for _, b := range bools {
		if b {
			r = append(r, b)
		}
	}

	return r
}

var alpha = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func index(a string) int {
	for i, b := range alpha {
		if a == b {
			return i + 1
		}
	}

	return -1
}
