package main

import "fmt"

func main() {
	nums := []int{0, 6, 10, 25}
	result := twoSumBf(nums, 31)
	fmt.Println(result)
}

//brute force solution
func twoSumBf(nums []int, target int) [][]int {
	//make a slice of arrays.
	//We want a slice of arrays to account for multiple ways to hit the target
	var s [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				s = append(s, []int{i, j})
			} else {
				continue
			}
		}
	}
	return s
}
