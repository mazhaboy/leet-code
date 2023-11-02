package binary_search

func FindMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	res := nums[l]
	for l <= r {
		if nums[l] <= nums[r] {
			res = Min(res, nums[l])
			break
		}

		m := (r + l) / 2
		res = Min(res, nums[m])
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}
