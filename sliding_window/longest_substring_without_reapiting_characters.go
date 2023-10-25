package sliding_window

import (
	"fmt"
)

type MySubstringMap struct {
	Dict map[string]int
}

func NewMySubstringMap() *MySubstringMap {
	newMySubstringMap := new(MySubstringMap)
	dict := make(map[string]int, 0)
	newMySubstringMap.Dict = dict
	return newMySubstringMap
}

func LengthOfLongestSubstring(s string) int {
	max := 0
	d := NewMySubstringMap()
	start := 0
	for end := range s {
		duplicateIndex, ok := d.Dict[string(s[end])]
		if ok {
			max = Max(max, end-start)
			for i := start; i <= duplicateIndex; i++ {
				delete(d.Dict, string(s[i]))
			}
			fmt.Println(duplicateIndex)
			start = duplicateIndex + 1
		}
		d.Dict[string(s[end])] = end
	}

	max = Max(max, len(s)-start)
	return max
}
