package greedy

/*
Question 714: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

Implementation follow greedy approach.
Observation:
- Creates a graph of valley(low/buy) and hill(high/sell) it will become easy to know when to sell
or buy.

Idea: compares the sum of possible gain from tx and initial buy vs new buy to decide whether to
perform a transaction. Improved from sellStock 2
**/

func maxProfit(prices []int, fee int) int {
	var maxProfit, buy, curProfit int
	if len(prices) < 2 {
		return 0
	}
	buy = prices[0]
	for _, sell := range prices[1:] {
		if sell <= buy+curProfit {
			maxProfit += curProfit
			buy, curProfit = sell, 0
		} else if sell-buy-fee > curProfit {
			curProfit = sell - buy - fee
		}
		if sell == prices[len(prices)-1] {
			maxProfit += curProfit
		}

	}
	return maxProfit
}
