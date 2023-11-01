package binary_search

type MinArr struct {
	index int
	value int
}

func SearchRotatedArray(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] >= nums[l] {
			if nums[m] < target || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target > nums[r] || target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}

func SearchRotatedArray2(nums []int, target int) int {
	minIndex := FindMinIndex(nums)
	firstPart := Search(nums[minIndex:], target)
	if -1 != firstPart {
		return minIndex + firstPart
	}

	secondPart := Search(nums[:minIndex], target)
	if -1 != secondPart {
		return secondPart
	}
	return -1
}

func FindMinIndex(nums []int) int {
	l := 0
	r := len(nums) - 1
	res := MinArr{
		index: l,
		value: nums[l],
	}
	for l <= r {
		if nums[l] <= nums[r] {
			res = min(res, MinArr{
				index: l,
				value: nums[l],
			})
			break
		}

		m := (r + l) / 2
		res = min(res, MinArr{
			index: m,
			value: nums[m],
		})
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res.index
}

func min(x, y MinArr) MinArr {
	if x.value < y.value {
		return x
	}
	return y
}
