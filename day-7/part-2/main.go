package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var pwd []string
	dirSizes := make(map[string]int64)
	cmdStartsWith := "$"

	for _, s := range input {

		// Run command
		// Controls which parent we're looking at
		if strings.HasPrefix(s, cmdStartsWith) {
			cmdStr := strings.Trim(s, "$ ")
			cmdSli := strings.Split(cmdStr, " ")
			switch cmdSli[0] {
			case "cd":
				pwd = changeDir(pwd, cmdSli[1])
			}

			continue
		}

		for i, _ := range pwd {
			if strings.HasPrefix(s, "dir") {
				continue
			}

			fileStr := strings.Split(s, " ")
			fileSizeStr, _ := fileStr[0], fileStr[1]
			fileSize, _ := strconv.Atoi(fileSizeStr)
			path := pwd[0 : i+1]

			dirSizes[strings.Join(path, "/")] += int64(fileSize)
		}
	}

	printTree(dirSizes)
}

func printTree(sizes map[string]int64) {
	totalSpace := int64(70000000)
	remaining := totalSpace - sizes["/"]
	requiredToDel := 30000000 - remaining

	var potentialDirs []int
	for _, val := range sizes {
		if val >= requiredToDel {
			potentialDirs = append(potentialDirs, int(val))
		}
	}

	sort.Ints(potentialDirs)
	fmt.Println(potentialDirs)
}

func changeDir(pwd []string, s string) []string {
	if s == ".." {
		return pwd[:len(pwd)-1]
	}

	return append(pwd, s)
}
