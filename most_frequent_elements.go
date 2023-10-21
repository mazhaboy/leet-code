package main

import (
	"sort"
)

type MyTopKFrequentMap struct {
	Dict map[int]int
}

func NewMyTopKFrequentMap() *MyTopKFrequentMap {
	newMyTopKFrequentMap := new(MyTopKFrequentMap)
	dict := make(map[int]int, 0)
	newMyTopKFrequentMap.Dict = dict
	return newMyTopKFrequentMap
}

type NumFrequent struct {
	Num   int
	Count int
}

func TopKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	d := NewMyTopKFrequentMap()
	for _, v := range nums {
		_, ok := d.Dict[v]
		if ok {
			d.Dict[v]++
		} else {
			d.Dict[v] = 1
		}
	}

	nms := make([]NumFrequent, 0)

	for key, value := range d.Dict {
		new := NumFrequent{
			Num:   key,
			Count: value,
		}
		nms = append(nms, new)
	}

	sort.Slice(nms, func(i, j int) bool {
		return nms[i].Count > nms[j].Count
	})

	for i := 0; i < k; i++ {
		res = append(res, nms[i].Num)
	}

	return res
}
