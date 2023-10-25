package sliding_window

type MyCharacterReplacementMap struct {
	Dict map[uint8]int
}

func NewMyCharacterReplacementMap() *MyCharacterReplacementMap {
	newMyCharacterReplacementMap := new(MyCharacterReplacementMap)
	dict := make(map[uint8]int, 0)
	newMyCharacterReplacementMap.Dict = dict
	return newMyCharacterReplacementMap
}

func (d *MyCharacterReplacementMap) Max() int {
	max := 0
	for key := range d.Dict {
		max = Max(max, d.Dict[key])
	}
	return max
}

func CharacterReplacement(s string, k int) int {
	d := NewMyCharacterReplacementMap()
	res := 0
	l := 0
	maxF := 0
	for r := range s {
		d.Dict[s[r]]++
		maxF = Max(maxF, d.Dict[s[r]])
		for (r-l+1)-maxF > k {
			d.Dict[s[l]]--
			l++
		}
		res = Max(res, r-l+1)
	}

	//a := map[int]int{1: 2, 2: 2}
	//b := map[int]int{1: 2, 2: 2}
	//
	//reflect.DeepEqual(a, b)

	return res
}
