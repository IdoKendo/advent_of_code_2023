package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	output := day1(input, 1)
	expected := 142

	assert.Equal(t, output, expected, "The two words should be the same.")

}

func TestPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	output := day1(input, 2)
	expected := 281

	assert.Equal(t, output, expected, "The two words should be the same.")
}
