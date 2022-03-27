package main

import "fmt"

func main() {
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(s)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	maxSum := -1 * int(^uint(0) >> 1)
	currSum := 0
	for i := 0; i < len(nums); i++ {
		currSum = currSum + nums[i] 
		if currSum > maxSum{
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}
