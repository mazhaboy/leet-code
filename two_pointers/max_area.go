package two_pointers

type MyMaxAreaMap struct {
	Dict map[int]int
}

func NewMyMaxAreaMap() *MyMaxAreaMap {
	newMyMaxAreaMap := new(MyMaxAreaMap)
	dict := make(map[int]int, 0)
	newMyMaxAreaMap.Dict = dict
	return newMyMaxAreaMap
}

func MaxArea(height []int) int {
	maxA := 0
	l := 0
	r := len(height) - 1

	for l < r {
		area := Min(height[l], height[r]) * (r - l)
		maxA = MaxA(maxA, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxA
}

func MaxArea2(height []int) int {
	d := NewMyMaxAreaMap()
	for i := range height {
		d.Dict[i] = height[i]
	}

	maxA := 0

	for i := range d.Dict {
		h1 := d.Dict[i]
		x1 := i + 1
		area := 0
		count := 0
		for count < len(d.Dict) {
			h2 := d.Dict[count]
			x2 := count + 1
			if i == count {
				continue
			}
			area = Min(h2, h1) * (x2 - x1)
			maxA = MaxA(maxA, area)
		}
	}

	return maxA
}

func MaxArea3(height []int) int {
	maxA := 0
	for i := 0; i < len(height); i++ {
		h1 := height[i]
		x1 := i + 1
		area := 0
		for j := 0; j < len(height); j++ {
			h2 := height[j]
			x2 := j + 1
			area = Min(h2, h1) * (x2 - x1)
			maxA = MaxA(maxA, area)

		}
	}
	return maxA
}

func MaxA(a1, a2 int) int {
	if a1 >= a2 {
		return a1
	}
	return a2
}

func Min(x1, x2 int) int {
	if x1 >= x2 {
		return x2
	}
	return x1
}
