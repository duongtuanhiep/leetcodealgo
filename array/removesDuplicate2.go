package array

// Counting extra duplicate & modify inplace.
func removeDuplicates(nums []int) int {

	// base case
	if len(nums) < 2 {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if nums[i] == nums[i-1] {
			var u, extra = i, 0
			u++
			for nums[u] == nums[u-1] && u < len(nums) {
				u++
				extra++
			}
			nums = append(nums[0:i+extra], nums[i+extra:len(nums)]...)
		}
	}
	return len(nums)
}

// // counting accepted value only
// func removeDuplicates(nums []int) int {

// 	// base case
// 	if len(nums) < 2 {
// 		return len(nums)
// 	}

// 	res := 0
// 	for i := 0; i < len(nums); i++ {
// 		if i < 2 {
// 			continue
// 		}
// 		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
// 			continue
// 		}
// 		res++
// 	}
// 	return res
// }
