package binary_search

import (
	"fmt"
)

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	fmt.Println(nums)
	for l <= r {
		m := (l + r) / 2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else if nums[m] == target {
			return m
		}
	}

	return -1
}
