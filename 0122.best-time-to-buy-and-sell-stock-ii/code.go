package besttimetobuyandsellstockii

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit += max(0, prices[i]-prices[i-1])
	}

	return profit
}
