package main

import (
	"fmt"
)

func main() {
	//fmt.Println(babki([]int{2, 2, 3, 9}))
	fmt.Println(LenWorld("salam ddaa"))
	fmt.Println(Plus([]int{1, 2, 3, 4}))
}

func babki(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0

	for _, price := range prices {
		if price < min {
			min = price
		}
		profit := price - min
		if profit > max {
			max = profit
		}
	}
	return max
}

func LenWorld(x string) int {
	y := len(x)
	len := 0
	i := y - 1
	for x[i] != ' ' {
		len++
		i--
	}
	return len
}

func Plus(digits []int) []int {
	m := len(digits)
	for i := m - 1; i >= 0; i-- {
		if digits[i] > 0 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return digits
}
