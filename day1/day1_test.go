package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	output := solution(input, 1)
	expected := 142
	if output != expected {
		t.Errorf("[%s] got %d; want %d", t.Name(), output, expected)
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	output := solution(input, 2)
	expected := 281
	if output != expected {
		t.Errorf("[%s] got %d; want %d", t.Name(), output, expected)
	}
}
