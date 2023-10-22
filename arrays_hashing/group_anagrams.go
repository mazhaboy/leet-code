package arrays_hashing

import (
	"sort"
)

type MyGroupAnagramMap struct {
	Dict map[string][]string
}

func NewMyGroupAnagramMap() *MyGroupAnagramMap {
	newMyGroupAnagramMap := new(MyGroupAnagramMap)
	dict := make(map[string][]string, 0)
	newMyGroupAnagramMap.Dict = dict
	return newMyGroupAnagramMap
}

func GroupAnagrams(strs []string) [][]string {
	dict := NewMyGroupAnagramMap()
	for _, v := range strs {
		_, ok := dict.Dict[SortSting(v)]
		if !ok {
			newSlice := make([]string, 0)
			newSlice = append(newSlice, v)
			dict.Dict[SortSting(v)] = newSlice
		} else {
			dict.Dict[SortSting(v)] = append(dict.Dict[SortSting(v)], v)
		}
	}

	res := make([][]string, 0)
	for k := range dict.Dict {
		res = append(res, dict.Dict[k])
	}

	return res
}

func SortSting(str string) string {
	bstr := []byte(str)
	sort.Slice(bstr, func(i, j int) bool {
		return bstr[i] > bstr[j]
	})
	res := string(bstr)
	return res
}
