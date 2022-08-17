package utils

import (
	"strconv"
	"strings"
)

// IntMatrix s
func IntMatrix(s string) [][]int {
	input := make([][]int, 0)
	for _, row := range strings.Split(strings.Trim(strings.Trim(s, "]]"), "[["), "],[") {

		parRow := make([]int, 0)
		for _, char := range strings.Split(row, ",") {
			val, _ := strconv.Atoi(char)

			parRow = append(parRow, val)
		}
		input = append(input, parRow)
	}
	return input
}

// Min calculates the lowest between a: int and b: int
func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
