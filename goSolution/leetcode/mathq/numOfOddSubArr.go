package mathq

/*

Question 1524: https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/

Math:
- total possible sub arr n(n-1)/2
- odd sum arr is when theres an odd number of odd number in an arr.
**/

func numOfSubarrays(arr []int) int {
	var res int
	prefix, even, odd := 0, 1, 0
	for _, val := range arr {
		prefix += val
		if prefix%2 == 1 {
			res += even
			odd++
		} else {
			res += odd
			even++
		}
		prefix, res = prefix%(1e9+7), res%(1e9+7)
	}

	return res
}
