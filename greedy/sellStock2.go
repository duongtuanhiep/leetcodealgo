package greedy

func maxProfit(prices []int) int {
	var result, localMin, localMax int
	result = 0
	lenArr := len(prices)
	if len(prices) <= 1 {
		return 0
	}
	localMin = prices[0]
	localMax = prices[0]
	for i := 1; i < lenArr; i++ {
		if localMax < prices[i] {
			localMax = prices[i]
		} else if localMax > prices[i] {
			result = result + localMax - localMin
			localMin = prices[i]
			localMax = prices[i]
		}
		if i == lenArr-1 {
			return result + localMax - localMin
		}
	}
	return result
}
