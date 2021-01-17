package array

import "sort"

func threeSum(nums []int) [][]int {

	sort.SliceStable(nums, func(i, j int) bool { return nums[i] < nums[j] })
	numPos := make(map[int]int)
	newNums := []int{}
	for pos, val := range nums {
		numPos[val] = pos
	}

	var res [][]int
	for i := 0; i < len(newNums); i++ {
		for j := i + 1; j < len(newNums); j++ {
			if pos, ok := numPos[0-(newNums[i]+newNums[j])]; ok && pos > j {
				var subRes = []int{newNums[i], newNums[j], newNums[pos]}
				res = append(res, subRes)
				break
			}
		}
	}

	return res
}
