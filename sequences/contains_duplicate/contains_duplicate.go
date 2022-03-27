package main

import "fmt"

func main() {
	nums := []int{0, 6, 0, 25}
	result := containsDuplicate(nums)
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
