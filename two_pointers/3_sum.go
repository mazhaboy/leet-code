package two_pointers

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			threeSum := v + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				newS := make([]int, 0)
				newS = append(newS, v, nums[l], nums[r])
				res = append(res, newS)
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return res
}
