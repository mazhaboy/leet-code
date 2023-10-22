package arrays_hashing

import (
	"sort"
)

func IsAnagram(s string, t string) bool {
	str1 := []byte(s)
	str2 := []byte(t)
	sort.Slice(str1, func(i, j int) bool {
		return str1[i] > str1[j]
	})
	sort.Slice(str2, func(i, j int) bool {
		return str2[i] > str2[j]
	})
	s1 := string(str1)
	t2 := string(str2)
	if s1 == t2 {
		return true
	}

	return false
}
