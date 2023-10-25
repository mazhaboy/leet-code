package sliding_window

import (
	"reflect"
)

func CheckInclusion(s1 string, s2 string) bool {
	d1 := make(map[uint8]int)
	d2 := make(map[uint8]int)
	for i := range s1 {
		d1[s1[i]]++
	}

	k := len(s1)

	s := 0
	e := 0
	for r := range s2 {
		_, ok := d1[s2[r]]
		if ok {
			s = r
			e = s + k - 1
			for s < e && e < len(s2) {
				d2[s2[s]]++
				s++
			}
			if reflect.DeepEqual(d1, d2) {
				return true
			} else {
				continue
			}
		}
	}

	return false
}
