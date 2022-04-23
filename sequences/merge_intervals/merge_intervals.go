package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := merge(nums)
	fmt.Println(result)
}

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	//returns sorted intervals
	sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

	for i, interval := range intervals {
        if i == 0 {
            result = append(result, interval)
            continue
        }

        lastInterval := result[len(result) - 1]

        if lastInterval[1] < interval[0] {
			fmt.Println(result)
            result = append(result, interval)
			fmt.Println(result)
        } else if interval[1] > lastInterval[1] {
            lastInterval[1] = interval[1]
        }
    }

    return result
}