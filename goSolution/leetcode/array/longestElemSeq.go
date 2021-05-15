package array

/*
Question 128: https://leetcode.com/problems/longest-consecutive-sequence/


Solution: https://leetcode.com/problems/longest-consecutive-sequence/solution/

1) Sorting:
	- If cur == last, res remains.
	- If cur == last +1, res++.
	- If cur != last +1, reset res=1

2) Put everything into a heap, init it O(N). Keeps track of last pop. Increase counter when number satify

3) Create a map for counting step where:
	- Key: integer value of number
	- Value: number of step that can perform starting from value of Key
We initialize all value in nums arr to our map with val 1.
Iterate through all the element in nums array. On each iteration finds the most step can be performed starting
from that node. We also update every nodes that original nodes pass through so that if we see that element again we
dont have to do the work. Each distinct number therefore is guaranteed to have value updated only once -> constant work -> O(N)
Runtime O(N), memory O(N)

Sample testcase:
[0,3,7,2,5,8,4,6,0,1]
[0,0,2,5,7,8,9,0,2,3,4,4,5,6,6,666,11,12,13,14,16,13,15,10]

**/

/*
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	sort.Ints(nums)
	var cur, res, last int = 1, 0, nums[0]
	for _, val := range nums {
		if val == last+1 {
			cur++
		} else if val != last {
			cur = 1
		}
		last = val
		if cur > res {
			res = cur
		}
	}
	return res
}
*/

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	step := make(map[int]int)
	for _, val := range nums {
		step[val] = 1
	}

	var cur, max int
	for _, val := range nums {
		cur = numberSreach(step, val)
		if cur > max {
			max = cur
		}
	}
	return max
}

func numberSreach(step map[int]int, curVal int) int {
	if step[curVal] == 0 {
		return 0
	} else if step[curVal] > 1 {
		return step[curVal]
	} else {
		step[curVal] = 1 + numberSreach(step, curVal+1)
	}
	return step[curVal]
}
