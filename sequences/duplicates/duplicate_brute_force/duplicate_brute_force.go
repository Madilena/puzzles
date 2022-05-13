package main

import "fmt"

func main() {
	nums := []int{0, 6, 0, 25, 25}
	result := duplicateBruteForce(nums)
	fmt.Println(result)
}

//this is O(n^2)
func duplicateBruteForce(nums []int) []int {
	var results []int

	for i := 0 ; i < len(nums); i++ {
		for j := i+1 ; j < len(nums); j++ {
			if nums[i] == nums[j]{
				results = append(results, nums[i])
			}
		}
	}
	return results
}
