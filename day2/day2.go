package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solution(input string, part int) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	p1 := 0
	p2 := 0
	redRe := regexp.MustCompile(`(\d+) red`)
	greenRe := regexp.MustCompile(`(\d+) green`)
	blueRe := regexp.MustCompile(`(\d+) blue`)
	gameRe := regexp.MustCompile(`(\d+)`)
	for _, line := range lines {
		minRed := 0
		minGreen := 0
		minBlue := 0
		splitLine := strings.Split(line, ":")
		gameID := splitLine[0]
		l := splitLine[1]
		possible := true
		for _, s := range strings.Split(l, ";") {
			red := redRe.FindStringSubmatch(s)
			if len(red) > 0 {
				n, _ := strconv.Atoi(red[1])
				minRed = max(minRed, n)
				if n > 12 {
					possible = false
				}
			}
			green := greenRe.FindStringSubmatch(s)
			if len(green) > 0 {
				n, _ := strconv.Atoi(green[1])
				minGreen = max(minGreen, n)
				if n > 13 {
					possible = false
				}
			}
			blue := blueRe.FindStringSubmatch(s)
			if len(blue) > 0 {
				n, _ := strconv.Atoi(blue[1])
				minBlue = max(minBlue, n)
				if n > 14 {
					possible = false
				}
			}
		}
		p2 += minRed * minGreen * minBlue
		if possible {
			gameID := gameRe.FindAllString(gameID, -1)
			n, _ := strconv.Atoi(gameID[0])
			p1 += n
		}
	}
	if part == 1 {
		return p1
	} else {
		return p2
	}
}

func main() {
	b, err := os.ReadFile("input2.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(solution(input, 1))
	fmt.Println(solution(input, 2))
}
