package main

import "fmt"

func main() {
	nums := []int{3, 2, 6, 5, 0, 3}
	result := maxProfit(nums)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	max := 0
	min := int(^uint(0) >> 1)

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if (prices[i] - min) > max {
			max = prices[i] - min
		}
	}
	return max
}
