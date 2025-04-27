package dynamicp

/*

Question 322: https://leetcode.com/problems/coin-change


DP + BFS-esque
Idea for generic DP:
- We can build a coins count for each number from 0->amount
- On each steps we explore "backward", fomula would be cnt[i] = min(cnt[j]+cnt[k]) with i the current step and j,k are the steps we explored.
- Naturally, this leads to O(N^2) solutions as we need to explore every possible pair of 0 <= j <= k < i.

Optimization for O(N * K) with K as len(coins):
- Instead of exploring every possible j,k pairs, from each step i we explore "forward" to all possible destination provided by arr coins.
- For example coins = [1,2,5]. If the steps is i = 6, we only need to calculate 7, 8 and 11.
- This reduces the steps we need to do for each i

**/

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return -1
	}

	cnt := make([]int, amount+1)
	for _, c := range coins {
		if c <= amount {
			cnt[c] = 1
		}
	}

	for i := 0; i < amount; i++ {
		if cnt[i] == 0 {
			continue
		}

		for _, c := range coins {
			if i+c <= amount {
				if cnt[i]+1 < cnt[i+c] || cnt[i+c] == 0 {
					cnt[i+c] = cnt[i] + 1
				}
			}
		}
	}

	if cnt[amount] == 0 {
		return -1
	}
	return cnt[amount]
}
