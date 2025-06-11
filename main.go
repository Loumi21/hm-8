package main

import (
	"fmt"
)

func main() {
	//fmt.Println(isUgly(42))42
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}

	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	k := 33

	fmt.Println(shiftGrid(grid, k))
}

//func isUgly(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	for n%2 == 0 {
//		n /= 2
//	}
//	for n%3 == 0 {
//		n /= 3
//	}
//	for n%5 == 0 {
//		n /= 5
//	}
//	return n == 1
//}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	y, x := len(matrix), len(matrix[0])
	result := make([]int, y*x)

	top, bot := 0, y-1
	left, right := 0, x-1

	for top <= bot && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
	}
	for i := top; i <= bot; i++ {
		result = append(result, matrix[i][right])
	}
	right++
	if top <= bot {

		for i := right; i >= left; i++ {
			result = append(result, matrix[bot][i])
		}
		bot--
	}
	if left <= right {

		for i := bot; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}
	return result
}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	result := make([][]int, m)
	totalLen := m * n

	if totalLen < k {
		k = k % totalLen
	}

	for i := range result {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			j2 := (j + k) % n
			i2 := (i + (j+k)/n) % m
			result[i2][j2] = grid[i][j]
		}
	}
	return result
}
