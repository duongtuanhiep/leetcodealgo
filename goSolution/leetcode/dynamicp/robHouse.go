package dynamicp

/*
Question 198: https://leetcode.com/problems/house-robber/

Dynamic P:
- create an array shows max you can rob up to that point

maxRob[i] = maxRob[i-1] | maxRob[i-2] + num

[1,2,3,1]
[2,7,9,3,1]
[2,1,1,2]
[1,2]
[2,1]
**/

func rob(nums []int) int {
	maxRob := make([]int, len(nums))
	for i := range maxRob {
		if i == 0 {
			maxRob[i] = nums[i]
		} else if i == 1 {
			if nums[i] > nums[i-1] {
				maxRob[i] = nums[i]
			} else {
				maxRob[i] = nums[i-1]
			}
		} else {
			if nums[i]+maxRob[i-2] > maxRob[i-1] {
				maxRob[i] = maxRob[i-2] + nums[i]
			} else {
				maxRob[i] = maxRob[i-1]
			}
		}
	}
	return maxRob[len(maxRob)-1]
}
