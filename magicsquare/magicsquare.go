package magicsquare

import "leetcode-go/utils"

// LargestMagicSquare grid
func LargestMagicSquare(grid [][]int) int {
	return largestMagicSquare(grid)
}

func largestMagicSquare(grid [][]int) int {

	for i := utils.Min(len(grid), len(grid[0])) - 1; i >= 0; i-- {

		for x0 := 0; x0 < len(grid)-i; x0++ {

			for y0 := 0; y0 < len(grid[0])-i; y0++ {

				sum := 0
				for d := 0; d <= i; d++ {
					sum += grid[x0+d][y0+d]
				}
				partial := 0
				for d := i; d >= 0; d-- {
					partial += grid[x0+d][y0+d]
				}
				if partial != sum {
					continue
				}

				exit := false

				for y := 0; y <= i && !exit; y++ {
					partial := 0
					for x := 0; x <= i; x++ {
						partial += grid[x0+x][y0+y]
					}
					if partial != sum {
						exit = true
						break
					}
				}

				for x := 0; x <= i && !exit; x++ {
					partial := 0
					for y := 0; y <= i; y++ {
						partial += grid[x0+x][y0+y]
					}
					if partial != sum {
						exit = true
						break
					}
				}

				if !exit {
					return i + 1
				}
			}
		}
	}

	return 0
}
