package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solution(input string, part int) int {
	lines := strings.Split(input, "\n")
	sum := 0
	re := regexp.MustCompile("[0-9]")
	words := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		if part == 2 {
			for k, v := range words {
				line = strings.ReplaceAll(line, k, v)
			}
		}
		numbers := re.FindAllString(line, -1)
		i, _ := strconv.Atoi(numbers[0])
		j, _ := strconv.Atoi(numbers[len(numbers)-1])
		sum += i*10 + j
	}
	return sum
}

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(solution(input, 1))
	fmt.Println(solution(input, 2))
}
