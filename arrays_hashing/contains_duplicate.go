package arrays_hashing

type MyMap struct {
	Dict map[int]int
}

func NewMyMap() *MyMap {
	d := make(map[int]int, 0)
	newMap := new(MyMap)
	newMap.Dict = d
	return newMap
}

func ContainsDuplicate(nums []int) bool {
	myMap := NewMyMap()
	for i := range nums {
		_, ok := myMap.Dict[nums[i]]
		if !ok {
			myMap.Dict[nums[i]]++
		} else {
			return true
		}
	}

	return false
}
