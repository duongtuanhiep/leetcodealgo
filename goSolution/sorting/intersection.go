package sorting

func intersection(nums1 []int, nums2 []int) []int {
	occ := make(map[int]int)
	var res []int
	for _, val := range nums1 {
		occ[val] = 1
	}
	for _, val := range nums2 {
		if occ[val] == 1 {
			res = append(res, val)
			occ[val] = 2
		}
	}
	return res
}
