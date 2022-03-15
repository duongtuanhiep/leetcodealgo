package array

/*
Question 667: https://leetcode.com/problems/beautiful-arrangement-ii/

Idea:
- Maybe inversion has sth to do with it



**/

func constructArray(n int, k int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = i + 1
	}
}
