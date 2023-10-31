package binary_search

func MinEatingSpeed(piles []int, h int) int {
	kMin := 1
	kMax := 0
	for i := range piles {
		kMax = Max(kMax, piles[i])
	}

	res := kMax

	for kMin <= kMax {
		m := (kMax + kMin) / 2
		time := 0
		for i := range piles {
			if piles[i]%m == 0 {
				time += piles[i] / m
			} else {
				time += (piles[i] / m) + 1
			}
		}

		if time <= h {
			res = Min(res, m)
			kMax = m - 1
		} else if time > h {
			kMin = m + 1
		}
	}

	return res
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
