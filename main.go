package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	var result [][]int
	rows := len(matrix)
	cols := len(matrix[0])

	// Loop over each diagonal starting from the first row
	for d := 0; d < cols; d++ {
		temp := []int{}
		for i, j := 0, d; i < rows && j >= 0; i, j = i+1, j-1 {
			temp = append(temp, matrix[i][j])
		}
		result = append(result, temp)
	}

	// Loop over each diagonal starting from the last column
	for d := 1; d < rows; d++ {
		temp := []int{}
		for i, j := d, cols-1; i < rows && j >= 0; i, j = i+1, j-1 {
			temp = append(temp, matrix[i][j])
		}
		result = append(result, temp)
	}

	fmt.Println(result)

}
