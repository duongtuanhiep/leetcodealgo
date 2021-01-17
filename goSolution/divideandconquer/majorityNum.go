package divideandconquer

// func majorityElement(nums []int) int {
// 	freq := make(map[int]int)
// 	datVal := 0
// 	for _, val := range nums {
// 		freq[val]++
// 		if freq[val] >= len(nums)/2+1 && len(nums)%2 == 1 || freq[val] > len(nums)/2 && len(nums)%2 == 0 {
// 			datVal = val
// 		}
// 	}
// 	return datVal
// }

// Base case
// len() = 1 -> return i
// len() = 0 -> return 0

func majorityElement(nums []int) int {
	return majElement(nums, len(nums)-1, 0)
}

func majElement(nums []int, hi, lo int) int {
	if hi == lo {
		return nums[hi]
	}

	var newMid int
	newMid = (hi + lo) / 2

	// Divide
	rightMaj := majElement(nums, hi, newMid+1)
	leftMaj := majElement(nums, newMid, lo)

	// Combine res
	if leftMaj == rightMaj {
		return leftMaj
	}

	// If res of sub prob different, do count.
	leftOcc := countOccurence(nums, hi, lo, leftMaj)
	rightOcc := countOccurence(nums, hi, lo, rightMaj)
	if leftOcc > rightOcc {
		return leftMaj
	}
	return rightMaj
}

func countOccurence(nums []int, hi, lo, num int) int {
	var occu int

	for pos := lo; pos < hi; pos++ {
		if nums[pos] == num {
			occu++
		}
	}
	return occu
}
