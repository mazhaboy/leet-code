package yandex

import (
	"fmt"
)

func DeleteDuplicates(arr []int) []int {
	res := make([]int, 0)
	if len(arr) == 0 {
		return []int{}
	}
	prev := arr[0]
	res = append(res, prev)
	fmt.Println(prev)
	for i := 1; i < len(arr); i++ {
		if prev != arr[i] {
			fmt.Println(arr[i])
			res = append(res, arr[i])
		}
		prev = arr[i]
	}
	return res
}
