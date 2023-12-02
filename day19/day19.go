package main

import (
	"fmt"
	"os"
)

func solution(input string, part int) int {
	return 0
}

func main() {
	b, err := os.ReadFile("input19.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(solution(input, 1))
	fmt.Println(solution(input, 2))
}
