package sorting

import (
	"math/rand"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var res []bool
	for pos := range l {
		subArr := sortSubArr(l[pos], r[pos], nums)
		if len(subArr) == 2 {
			res = append(res, true)
			continue
		}
		arith := subArr[1] - subArr[0]
		ok := true
		for pos := 0; pos < len(subArr)-1; pos++ {
			if subArr[pos+1]-subArr[pos] != arith {
				ok = false
				break
			}
		}
		res = append(res, ok)
	}
	return res
}

func sortSubArr(left, right int, nums []int) []int {
	var newArr []int
	for p := left; p < right+1; p++ {
		newArr = append(newArr, nums[p])
	}
	sortAssist(newArr, 0, len(newArr)-1)
	return newArr
}

func sortAssist(arr []int, lo, hi int) {

	// if only contains 1 elem dont sort
	if hi-lo < 1 {
		return
	}

	// Generate pivot
	pivot := lo + rand.Intn(hi-lo+1)

	arr[pivot], arr[lo] = arr[lo], arr[pivot]

	i := lo + 1 // Boundary between elem > arr[pivot] and elem < arr[pivot]
	for j := lo + 1; j <= hi; j++ {
		if arr[j] < arr[lo] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Swapping back
	arr[lo], arr[i-1] = arr[i-1], arr[lo]
	sortAssist(arr, lo, i-2)
	sortAssist(arr, i, hi)
}
