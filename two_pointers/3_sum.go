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
	d := NewMyThreeSumMap()
	res := make([][]int, 0)
	for i := range nums {
		d.Dict[nums[i]]++
	}

	fmt.Println(d.Dict)

	return res
}
