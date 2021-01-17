package array

/*
Question 1646: https://leetcode.com/problems/get-maximum-in-generated-array/

**/

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	max, i, resArr := 0, 1, []int{0, 1}
	for len(resArr) < n+1 {
		if len(resArr)%2 == 0 {
			resArr = append(resArr, resArr[i])
		} else {
			resArr = append(resArr, resArr[i]+resArr[i+1])
			i++
		}
		if resArr[len(resArr)-1] > max {
			max = resArr[len(resArr)-1]
		}
	}

	return max
}
