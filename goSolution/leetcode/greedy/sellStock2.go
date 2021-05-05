package greedy

/*
Question 122: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

Greedy approach:

**/

/*
Window Sliding/2 pointers

**/

// func maxProfit(prices []int) int {
// 	var maxSum int
// 	if len(prices) == 0 {
// 		return 0
// 	}
// 	for i, j := 0, 0; i < len(prices) && j < len(prices); j++ {
// 		if prices[j]-prices[i] > 0 {
// 			maxSum += prices[j] - prices[i]
// 			i = j
// 			continue
// 		}
// 		if prices[j] < prices[i] {
// 			i = j
// 		}
// 	}
// 	return maxSum
// }

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxSum, buy := 0, prices[0]
	for _, sell := range prices[1:] {
		if sell > buy {
			maxSum += sell - buy
		}
		buy = sell
	}
	return maxSum
}
