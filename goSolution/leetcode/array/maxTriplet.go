package array

/*

Question 2874: https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/?envType=daily-question&envId=2025-04-04

If we can find the highest element up to index i and the highest element after index i, problems become:
result = max( (highestBefore[i-1] - nums[i]) * highestAfter[i+1])

**/

func maximumTripletValue(nums []int) int64 {
	// sort.Ints(nums)

	// make new Arr for storing the highest values of i since i
	highestAfterIndex := make([]int, len(nums))
	highestBeforeIndex := make([]int, len(nums))
	copy(highestAfterIndex, nums)
	copy(highestBeforeIndex, nums)

	lastIndex := len(highestAfterIndex) - 1
	highest := highestAfterIndex[lastIndex]
	for i := lastIndex; i >= 0; i-- {
		if highest > highestAfterIndex[i] {
			highestAfterIndex[i] = highest
		}
		if highest < highestAfterIndex[i] {
			highest = highestAfterIndex[i]
		}
	}

	highest = highestBeforeIndex[0]
	for i, current := range highestBeforeIndex {
		if highest > current {
			highestBeforeIndex[i] = highest
		}
		if highest < current {
			highest = current
		}
	}

	var res int64
	for j := 1; j < len(nums)-1; j++ {
		if highestBeforeIndex[j-1] < nums[j] {
			continue
		}
		cur := int64((highestBeforeIndex[j-1] - nums[j]) * highestAfterIndex[j+1])
		if cur > res {
			res = cur
		}
	}

	return res
}

// Or to simplify
// func maximumTripletValue(nums []int) int64 {
// 	x := 0
// 	y := 0
// 	z := 0
// 	for _, v := range nums {
// 		x = max(x, y*v)
// 		y = max(y, z-v)
// 		z = max(z, v)
// 	}
// 	return int64(x)
// }
