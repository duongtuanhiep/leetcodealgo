package array

/*
IDEA:
Go through the array keep track of last elem last iteration.
Whenever there is a decrease in value, update the value itselfs
and add change++.
The goal of the update is to makes the processing part of the
array + the next element (i+1) to follows the rules of non-decreasing
while making the least changes to the value.
Plan: Update either i or i+1 pos to the "latest" value that is greater
than i-1
**/
func checkPossibility(nums []int) bool {
	var change int
	var last int
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}
		if nums[i] > nums[i+1] {
			change++
			if i == 0 {
				nums[i] = nums[i+1]
			} else if nums[i+1] > last {
				nums[i] = nums[i+1]
			} else if nums[i+1] < last {
				nums[i+1] = nums[i]
			}
		}
		last = nums[i]
		if change >= 2 {
			return false
		}
	}
	return change <= 1
}
