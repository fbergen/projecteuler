package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	f, err := ioutil.ReadFile("p096_sudoku.txt")
	if err != nil {
		panic(err)
	}
	ps := regexp.MustCompile("Grid [0-9]*\n").Split(string(f), -1)
	solution := 0

	for puzzle := 1; puzzle < len(ps); puzzle++ {
		a := regexp.MustCompile("[0-9]").FindAllString(ps[puzzle], -1)
		var b [9][9]int
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				b[i][j], _ = strconv.Atoi(a[i*9+j])
			}
		}
		if !solve(&b, 0, 0) {
			fmt.Println("Couln't solve puzzle #", puzzle)
			return
		}
		solution += b[0][0]*100 + b[0][1]*10 + b[0][2]
	}
	fmt.Println(solution)
}

func isValid(matrix [9][9]int, row, col, val int) bool {
	// Check the row
	for i := 0; i < 9; i++ {
		if matrix[row][i] == val && i != col {
			return false
		}
	}
	// Check the col
	for i := 0; i < 9; i++ {
		if matrix[i][col] == val && i != row {
			return false
		}
	}
	// Check the 3x3 square
	startRow := 3 * (row / 3)
	startCol := 3 * (col / 3)
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if i != row || j != col {
				if matrix[i][j] == val {
					return false
				}
			}
		}
	}
	return true
}

func solve(matrix *[9][9]int, row, col int) bool {
	if col >= 9 {
		col = 0
		row++
	}
	if row >= 9 {
		return true
	}
	if matrix[row][col] == 0 {
		for val := 1; val < 10; val++ {
			if isValid(*matrix, row, col, val) {
				matrix[row][col] = val
				if solve(matrix, row, col+1) {
					return true
				}
			}
		}
		matrix[row][col] = 0
		return false
	}
	return solve(matrix, row, col+1)
}
