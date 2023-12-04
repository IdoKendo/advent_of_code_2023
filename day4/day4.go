package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func solution(input string, part int) int {
	lines := strings.Split(input, "\n")
	sum := 0
	pileSize := 0
	multipliers := make(map[int]int)
	for i := range lines {
		multipliers[i] = 1
	}
	for i, line := range lines {
		if line == "" {
			continue
		}
		pileSize += multipliers[i]
		title := strings.Split(line, ":")
		numbers := strings.Split(title[1], "|")
		winningNumbers := numbers[0]
		hand := numbers[1]
		numbersList := strings.Split(winningNumbers, " ")
		handList := strings.Split(hand, " ")
		matches := -1
		for _, h := range handList {
			if h == "" {
				continue
			}
			if contains(numbersList, h) {
				matches += 1
			}
		}
		if part == 1 {
			sum += int(math.Pow(2, float64(matches)))
		} else if matches >= 0 {
			for j := 1; j < matches+2; j++ {
				multipliers[i+j] += multipliers[i]
			}
		}
	}
	if part == 1 {
		return sum
	} else {
		return pileSize
	}
}

func main() {
	b, err := os.ReadFile("input4.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(solution(input, 1))
	fmt.Println(solution(input, 2))
}
