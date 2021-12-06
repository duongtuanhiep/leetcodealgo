package hashmap

/*
Question 448: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

**/

func findDisappearedNumbers(nums []int) []int {
	exist := make(map[int]bool)
	var last int
	for _, val := range nums {
		exist[val] = true
		if val > last {
			last = val
		}
	}

	var res []int
	for i := 1; i <= len(nums); i++ {
		if !exist[i] {
			res = append(res, i)
		}
	}
	return res
}
