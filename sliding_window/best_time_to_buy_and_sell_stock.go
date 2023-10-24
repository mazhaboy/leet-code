package sliding_window

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxChange := 0
	min := prices[0]

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			maxChange = Max(maxChange, prices[i+1]-min)
		} else {
			min = Min(min, prices[i+1])
		}
	}

	return maxChange
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MaxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0

	for i := 0; i < len(prices); i++ {
		for j := 1; j < len(prices); j++ {
			if j > i {
				max = Max(max, prices[j]-prices[i])
			}
		}
	}

	return max
}
