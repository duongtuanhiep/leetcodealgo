package array

func numIdenticalPairs(nums []int) int {
	var rec map[int]int
	var res int
	for _, val := range nums {
		// if val, ok := rec[nums[pos]]; ok && val < pos {
		// 	res++
		// }
		rec[val]++
	}

	for _, occur := range rec {
		if occur > 1 {
			res += occur * (occur - 1) / 2
		}
	}
	return res
}
