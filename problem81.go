package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	f, err := ioutil.ReadFile("p081_matrix.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(f), "\n")
	matrix := make([][]int, 0)
	for _, row_str := range rows {
		if row_str != "" {
			cols := strings.Split(row_str, ",")
			row := make([]int, 0)
			for _, col := range cols {
				c, _ := strconv.Atoi(col)
				row = append(row, c)
			}
			matrix = append(matrix, row)
		}
	}

	for i := 78; i >= 0; i-- {
		matrix[79][i] += matrix[79][i+1]
		matrix[i][79] += matrix[i+1][79]
	}
	for i := 78; i >= 0; i-- {
		for j := 78; j >= 0; j-- {
			matrix[i][j] += min(matrix[i+1][j], matrix[i][j+1])
		}
	}
	fmt.Println(matrix[0][0])
}
