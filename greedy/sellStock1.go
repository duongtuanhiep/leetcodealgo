package greedy

/*
Question 121: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Idea: We are only allowed to do one transaction only, meaning we can find a pair
of day that has the largest gap between buying and selling.

Formal Statement: given arr []int, finds Max(arr[j]-arr[i]) that i<j.
**/

// Approach 1: Brute force
// The naive approach is to get an element i, on each element i finds element j
// so that i<j and max(arr[i]-arr[j]), after that returns max
// Runtime: O(N^2) / Space: O(1)

func maxProfit(prices []int) int {
	var maxVal int
	for i := range prices {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > maxVal {
				maxVal = prices[j] - prices[i]
			}
		}
	}
	return maxVal
}

// Approach 2
// If we can find the maximum value arr[j] an from where index i can "reach"
// then we only needs to compare the sum of (arrMaxReach[i] - arr[i]) to find
// the maximum profit.
// We can then create a holder array, iterate the given array backward, store
// the maximum value arr[j] from that current location i can "reach" including
// itselfs.
// Runtime O(N), Space O(N)
func maxProfit(prices []int) int {
	// Preprocessing, finds max val an index can "reach" including itself
	var maxVal int
	maxReach := make([]int, len(prices))
	var maxReachLocal int
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] < maxReachLocal {
			maxReach[i] = maxReachLocal
		} else {
			maxReach[i] = prices[i]
			maxReachLocal = prices[i]
		}
	}
	for i := range prices {
		if maxReach[i]-prices[i] > maxVal {
			maxVal = maxReach[i] - prices[i]
		}
	}
	return maxVal
}

/* Approach 3: Sliding windows
One thing we can try to do is to :
Get LocalMin, LocalMax index i,j then LocalTxn = LocalMax - LocalMin
We drops the localMin when we can find a smaller value than previous
localMin, update LocalTxn when there is a new localMax thats greater
than the previous localMax.

**/
// Runtime: O(N), Space O(1)
func maxProfit(prices []int) int {
	// Handle case len 0
	if len(prices) == 0 {
		return 0
	}

	var maxVal, localVal int
	var low, high int = 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[low] {
			low, high, localVal = i, i, 0
		} else if prices[i] > prices[high] {
			high, localVal = i, prices[high]-prices[low]
		}
		if maxVal < localVal {
			maxVal = localVal
		}
	}
	return maxVal
}
