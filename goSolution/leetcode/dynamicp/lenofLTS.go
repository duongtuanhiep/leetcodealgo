package dynamicp

/*

Question 300: https://leetcode.com/problems/longest-increasing-subsequence/


DP - O(N^2):
- Go from bottom, try to consider number of valid max subsequences (first one is always 0)
- On each iteration, try to get the longest valid subsequence for index i: index i is part of subsequence that continues with some i+n index
- Update index[i] with its max value.

FWTrees - O(NlogN):

This comment explains multiple ways to accomplish this:
https://leetcode.com/problems/longest-increasing-subsequence/solutions/1326308/c-python-dp-binary-search-bit-segment-tree-solutions-picture-explain-o-nlogn/?envType=problem-list-v2&envId=27ol1gds


**/

// func lengthOfLIS(nums []int) int {
// 	dp := make([]int, len(nums))
// 	var res int
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		dp[i] = 1
// 		for k := i + 1; k < len(nums); k++ {
// 			if nums[k] > nums[i] && dp[k]+1 > dp[i] {
// 				dp[i] = dp[k] + 1
// 			}
// 		}
// 		if dp[i] > res {
// 			res = dp[i]
// 		}
// 	}

// 	return res
// }

func lengthOfLIS(nums []int) int {
	fwTree := fenwickTree(1e4*2 + 1) // cause value is from -1e4 <= x <= 1e4
	for i := range nums {
		cur := fwTree.query(nums[i] + 1e4 - 1) // finding counts of elem less than nums[i]
		fwTree.update(nums[i]+1e4, cur+1)
	}

	return fwTree.query(1e4 * 2)
}

type FenwickTree struct {
	tree []int
}

func fenwickTree(size int) *FenwickTree {
	return &FenwickTree{tree: make([]int, size+1)}
}

func (ft *FenwickTree) update(index, val int) {
	index++
	for index < len(ft.tree) {
		ft.tree[index] = max(ft.tree[index], val)
		index += index & -index
	}
}

func (ft *FenwickTree) query(index int) int {
	index++
	res := 0
	for index > 0 {
		res = max(ft.tree[index], res)
		index -= index & -index
	}
	return res
}
