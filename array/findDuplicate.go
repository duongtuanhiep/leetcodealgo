package array

// Naive apporach using map
func findDuplicates(nums []int) []int {
	dup := make(map[int]bool)
	res := []int{}
	for pos := range nums {
		if dup[nums[pos]] == true {
			res = append(res, nums[pos])
		}
		dup[nums[pos]] = true
	}
	return res
}
