package main

import "fmt"

func main() {
	nums := []int{0, 6, 10, 25}
	result := twoSum(nums, 31)
	fmt.Println(result)
}

//two pointer solution
func twoSum(nums []int, target int) []int {
	var s []int
	left := 0
	right := len(nums) - 1
	var tmpSum int

	for left < right {
		tmpSum = nums[left] + nums[right]
		if tmpSum == target {
			s = []int{left, right}
			return s
		} else if tmpSum > target {
			right--
		} else {
			left++
		}
	}
	return s
}
