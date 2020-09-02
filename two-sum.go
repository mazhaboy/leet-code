package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 5, 5, 11}
	target := 10
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	array := []int{}
	i := 0
	j := 0
	for i = 0; i < len(nums)-1; i++ {
		for j = 0; j < len(nums)-1; j++ {
			if (nums[i]+nums[j+1]) == target && i != j+1 {
				array = append(array, i, j+1)
				return array

			}
		}

	}
	return array
}
