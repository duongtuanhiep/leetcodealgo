package prefixsum

/*

Question 152: https://leetcode.com/problems/maximum-product-subarray/

Kadane algorithm with left right scan
Can be solved with dp.
**/

func maxProduct(nums []int) int {
	res := -9999999
	prod := 1
	for i := range nums {
		prod *= nums[i]
		res = max(res, prod)
		if prod == 0 {
			prod = 1
		}
	}

	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod *= nums[i]
		res = max(res, prod)
		if prod == 0 {
			prod = 1
		}
	}

	return res
}
