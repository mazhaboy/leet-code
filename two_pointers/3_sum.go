package two_pointers

import (
	"fmt"
)

type MyThreeSumMap struct {
	Dict map[int]int
}

func NewMyThreeSumMap() *MyThreeSumMap {
	newMyThreeSumMap := new(MyThreeSumMap)
	dict := make(map[int]int, 0)
	newMyThreeSumMap.Dict = dict
	return newMyThreeSumMap
}

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	i := 0
	j := 1
	k := 2
	for k < len(nums) {
		newSlice := make([]int, 0)
		fmt.Printf("k: %d, j: %d, i: %d\n", k, j, i)
		if nums[i]+nums[j]+nums[k] == 0 {
			newSlice = append(newSlice, nums[i], nums[j], nums[k])
			res = append(res, newSlice)
		}
		k++
	}
	fmt.Println(k)

	for j < len(nums)-1 {
		newSlice := make([]int, 0)
		fmt.Printf("k: %d, j: %d, i: %d\n", k, j, i)
		if nums[i]+nums[j]+nums[len(nums)-1] == 0 {
			newSlice = append(newSlice, nums[i], nums[j], nums[len(nums)-1])
			res = append(res, newSlice)
		}
		j++
	}
	fmt.Println(j)

	for i < len(nums)-2 {
		newSlice := make([]int, 0)
		fmt.Printf("k: %d, j: %d, i: %d\n", k, j, i)
		if nums[i]+nums[len(nums)-2]+nums[len(nums)-1] == 0 {
			newSlice = append(newSlice, nums[i], nums[len(nums)-2], nums[len(nums)-1])
			res = append(res, newSlice)
		}
		i++
	}
	fmt.Println(i)

	return res
}
