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

	// Lower 1-26
	// Upper 27-52

	total := 0
	for scanner.Scan() {
		txt := scanner.Text()

		mid := len(txt) / 2
		a, b := txt[0:mid], txt[mid:]
		l := make(map[string]bool, len(alpha))
		r := make(map[string]bool, len(alpha))

		fmt.Println(txt, a, b)
		for i := 0; i < mid; i++ {
			charA := fmt.Sprintf("%c", a[i])
			charB := fmt.Sprintf("%c", b[i])

			l[charA] = true
			r[charB] = true

			if l[charA] && r[charA] {
				total += index(charA)
				fmt.Println(index(charA), charA)
				break
			}

			if l[charB] && r[charB] {
				total += index(charB)
				fmt.Println(index(charB), charB)
				break
			}
		}
	}

	fmt.Println(total)
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
