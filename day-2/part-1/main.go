package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	defer f.Close()

	// A Rock
	// B Scissors
	// B Paper
	actionMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	total := 0
	for scanner.Scan() {
		txt := scanner.Text()
		actions := strings.Split(txt, " ")

		theirs := actions[0]
		ours := actions[1]

		a, _ := actionMap[theirs]
		switch ours {
		case "Y":
			total += a + 3
		case "Z":
			total += a + 6
		case "X":
			total += a + 0
		}
	}

	fmt.Println(total)
}
