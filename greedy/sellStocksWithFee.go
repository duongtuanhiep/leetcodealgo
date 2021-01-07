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
	var maxProfit, curBuy, curProfit int
	if len(prices) < 2 {
		return 0
	}
	curBuy = prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] <= curBuy+curProfit {
			maxProfit += curProfit
			curBuy, curProfit = prices[i], 0
		} else if prices[i]-curBuy-fee > curProfit {
			curProfit = prices[i] - curBuy - fee
		}
		if i == len(prices)-1 && curProfit > 0 {
			maxProfit += curProfit
		}

	}
	return maxProfit
}
