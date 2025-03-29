package dynamicp

/*

Question 1749: https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/

DP with memo
**/

func maxAbsoluteSum(nums []int) int {
	var res int
	minNum, maxNum := make([]int, len(nums)), make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			minNum[i], maxNum[i] = nums[i], nums[i]
		} else {
			if minNum[i-1]+nums[i] < nums[i] {
				minNum[i] = minNum[i-1] + nums[i]
			} else {
				minNum[i] = nums[i]
			}

			if maxNum[i-1]+nums[i] > nums[i] {
				maxNum[i] = maxNum[i-1] + nums[i]
			} else {
				maxNum[i] = nums[i]
			}
		}

		curAbs := max(absInt(maxNum[i]), absInt(minNum[i]))
		if curAbs > res {
			res = curAbs
		}
	}

	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
