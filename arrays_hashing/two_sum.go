package arrays_hashing

type MyTwoSumMap struct {
	Dict map[int]int
}

func NewMyTwoSumMap() *MyTwoSumMap {
	newMyTwoSumMap := new(MyTwoSumMap)
	newDict := make(map[int]int, 0)
	newMyTwoSumMap.Dict = newDict
	return newMyTwoSumMap
}

func TwoSum(nums []int, target int) []int {
	d := NewMyTwoSumMap()

	for i, v := range nums {
		_, ok := d.Dict[target-v]
		if !ok {
			d.Dict[v] = i
		} else {
			return []int{d.Dict[target-v], i}
		}

	}

	return []int{}
}
