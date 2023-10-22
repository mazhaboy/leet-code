package main

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1
	postfix := 1
	for i := range nums {
		res[i] = prefix
		prefix = prefix * nums[i]
	}

	for i := len(res) - 1; i >= 0; i-- {
		res[i] = res[i] * postfix
		postfix = postfix * nums[i]
	}

	return res
}
