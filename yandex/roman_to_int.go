package yandex

func RomanToInt(s string) int {
	res := 0
	dict := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	cur := 0
	next := 0
	for i := 0; i < len(s)-1; i++ {
		cur = dict[string(s[i])]
		next = dict[string(s[i+1])]
		if cur < next {
			res -= cur
		} else {
			res += cur
		}
	}
	res += dict[string(s[len(s)-1])]

	return res
}
