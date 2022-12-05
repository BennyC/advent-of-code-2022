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

	actionScores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	// A Rock
	// B Scissors
	// B Paper

	actionMap := map[string]struct {
		beats string
		loses string
		draws string
	}{
		"A": {
			beats: "C",
			loses: "B",
			draws: "A",
		},
		"B": {
			beats: "A",
			loses: "C",
			draws: "B",
		},
		"C": {
			beats: "B",
			loses: "A",
			draws: "C",
		},
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
			v, _ := actionScores[a.draws]
			total += v + 3
		case "Z":
			v, _ := actionScores[a.loses]
			total += v + 6
		case "X":
			v, _ := actionScores[a.beats]
			total += v + 0
		}
	}

	fmt.Println(total)
}
