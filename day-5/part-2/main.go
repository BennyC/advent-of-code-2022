package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	split := 0
	for i, s := range input {
		if s == "" {
			split = i
			break
		}
	}

	toParse := input[0:split]
	graph := parse(toParse)

	fmt.Println(graph)
	for _, action := range input[split+1:] {
		fmt.Println(action)
		graph = act(action, graph)
		fmt.Println(graph)
		fmt.Println(strings.Repeat("-", 15))
	}

}

func act(action string, graph map[int][]string) map[int][]string {
	g := graph
	r := regexp.MustCompile("\\d+")
	matches := r.FindAllString(action, -1)
	numStr, fromStr, toStr := matches[0], matches[1], matches[2]

	num, _ := strconv.Atoi(numStr)
	from, _ := strconv.Atoi(fromStr)
	to, err := strconv.Atoi(toStr)
	if err != nil {
		fmt.Println(err)
	}

	fromSlice := make([]string, len(g[from]))
	copy(fromSlice, g[from])

	move := make([]string, len(fromSlice[:num]))
	copy(move, fromSlice[:num])

	keep := make([]string, len(fromSlice[num:]))
	copy(keep, fromSlice[num:])

	g[from] = keep
	g[to] = append(move, g[to]...)

	return g
}

func parse(input []string) map[int][]string {
	cols := strings.Split(strings.Trim(input[len(input)-1], " "), " ")
	numOfCols, _ := strconv.Atoi(cols[len(cols)-1])
	graph := make(map[int][]string)

	for i := 1; i <= numOfCols; i++ {
		graph[i] = make([]string, len(input))
	}

	for i := len(input) - 2; i >= 0; i-- {
		s := input[i]
		for i := 0; i < len(s); i += 4 {
			char := fmt.Sprintf("%c", s[i+1])
			pos := i / 4
			if char != " " {
				graph[pos+1] = append([]string{char}, graph[pos+1]...)
			}
		}
	}

	return graph
}
