package two_pointers

func TwoSum2(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}
	l := 0
	r := len(numbers) - 1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}
