package bitmanipulation

/*

Question 1863: https://leetcode.com/problems/sum-of-all-subset-xor-totals/?envType=daily-question&envId=2025-04-05

101
100
001

**/

func subsetXORSum(nums []int) int {
	res := 0
	subsets := genSubsetXOR(0, nums)
	for _, subSum := range subsets {
		res += subSum
	}
	return res
}

func genSubsetXOR(index int, nums []int) []int {
	var res []int
	for i := index; i < len(nums); i++ {
		subsets := genSubsetXOR(i+1, nums)
		for _, subsetXORSum := range subsets {
			res = append(res, subsetXORSum^nums[i])
		}
		res = append(res, nums[i])
	}
	return res
}
