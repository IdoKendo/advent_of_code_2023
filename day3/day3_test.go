package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	output := solution(input, 1)
	expected := 4361
	if output != expected {
		t.Errorf("[%s] got %d; want %d", t.Name(), output, expected)
	}
}

func TestPart2(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	output := solution(input, 2)
	expected := 467835
	if output != expected {
		t.Errorf("[%s] got %d; want %d", t.Name(), output, expected)
	}
}
