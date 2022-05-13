package main

import "fmt"

func main() {
	nums := []int{0, 6, 0, 25}
	result := duplicateHash(nums)
	fmt.Println(result)
}

//this is O(n)
func duplicateHash(nums []int) []int {
	cache := make(map[int]bool)
	var results []int
	if len(nums) == 0 {
		return results
	}
	for i := 0; i < len(nums); i++ {
		if cache[nums[i]] {
			results = append(results, nums[i])
		} else {
			cache[nums[i]] = true
		}
	}
	return results
}
