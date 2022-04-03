package main

import "fmt"

func main() {
	s := []int{-1, 1, 0, -3, 3}
	result := productExceptSelf(s)
	fmt.Println(result)
}

func productExceptSelf(nums []int) []int {
	left := 1
	right := 1
	product := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		product[i] = left
		left *= nums[i]
	}

	for i := len(nums) - 1; i > -1; i-- {
		product[i] = right * product[i]
		right *= nums[i]
	}
	return product
}
