package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	n := len(nums)

	// Sort the given array
	sort.Ints(nums)

	var result [][]int

	for firstPointer := 0; firstPointer < n-2; firstPointer++ {
		// Skip all duplicates from left
		// firstPointer>0 ensures this check is made only from 2nd element onwards
		if firstPointer > 0 && nums[firstPointer] == nums[firstPointer-1] {
			continue
		}

		secondPointer := firstPointer + 1
		thirdPointer := n - 1
		for secondPointer < thirdPointer {
			sum := nums[secondPointer] + nums[thirdPointer] + nums[firstPointer]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[firstPointer], nums[secondPointer], nums[thirdPointer]})

				thirdPointer--

				// Skip all duplicates from right
				for secondPointer < thirdPointer && nums[thirdPointer] == nums[thirdPointer+1] {
					thirdPointer--
				}
			} else if sum > 0 {
				// Decrement thirdPointer to reduce sum value
				thirdPointer--
			} else {
				// Increment secondPointer to increase sum value
				secondPointer++
			}
		}
	}
	return result
}
