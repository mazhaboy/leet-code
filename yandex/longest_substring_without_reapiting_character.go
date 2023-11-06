package yandex

func LengthOfLongestSubstring(s string) int {
	//dict := make(map[uint8]int, 0)
	res := 0

	if len(s) == 0 {
		return res
	}

	for l := range s {
		dict := make(map[uint8]int, 0)
		r := l
		ok := false
		counter := 0
		for r < len(s) && !ok {
			_, ok = dict[s[r]]
			if !ok {
				dict[s[r]]++
			} else {
				break
			}
			counter++
			r++
		}
		res = Max(res, counter)
	}
	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
