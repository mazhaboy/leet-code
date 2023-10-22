package arrays_hashing

type MyLongestConsecutiveMap struct {
	Dict map[int]int
}

func NewMyLongestConsecutiveMap() *MyLongestConsecutiveMap {
	myLongestConsecutiveMap := new(MyLongestConsecutiveMap)
	dict := make(map[int]int, 0)
	myLongestConsecutiveMap.Dict = dict
	return myLongestConsecutiveMap
}

func LongestConsecutive(nums []int) int {
	d := NewMyLongestConsecutiveMap()
	longest := 0
	for i := range nums {
		d.Dict[nums[i]]++
	}

	for i := range nums {
		target := nums[i]
		_, ok1 := d.Dict[target-1]
		length := 0
		if !ok1 {
			ok2 := true
			for ok2 {
				_, ok2 = d.Dict[target+length]
				if ok2 {
					length++
				}
			}
		}
		longest = Max(longest, length)
	}

	return longest
}

func Max(longest, length int) int {
	if longest >= length {
		return longest
	}
	return length
}
