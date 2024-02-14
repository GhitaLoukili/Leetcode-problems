package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	numsWithIndex := make([][2]int, len(nums))
	for i, v := range nums {
		numsWithIndex[i] = [2]int{v, i}
	}
	sort.Slice(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i][0] < numsWithIndex[j][0]
	})

	l, r := 0, len(numsWithIndex)-1
	for l < r {
		sm := numsWithIndex[l][0] + numsWithIndex[r][0]
		if sm > target {
			r--
		} else if sm < target {
			l++
		} else {
			break
		}
	}

	return []int{numsWithIndex[l][1], numsWithIndex[r][1]}

}
