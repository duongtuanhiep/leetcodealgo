package greedy

/*
Question 1004: https://leetcode.com/problems/max-consecutive-ones-iii/

Idea:
Sliding windows: we try to place 0 in every position that we can find.
- If we can do that then its all good, just need to keep track of the current sub seq so far
- If we cant, the current subsequence is equal to current - (last_placed_1 - last_1_Elem)
**/

func longestOnes(nums []int, k int) int {
	var res, cur, last int
	pos := []int{}
	for i, val := range nums {
		if val == 1 {
			cur++
		} else {
			if k == 0 && len(pos) == 0 {
				cur = 0
			} else {
				if k > 0 {
					k, cur = k-1, cur+1
				} else {
					cur -= pos[0] - last
					last = pos[0] + 1
					pos = pos[1:]
				}
				pos = append(pos, i)
			}
		}

		if cur > res {
			res = cur
		}
	}
	return res
}
