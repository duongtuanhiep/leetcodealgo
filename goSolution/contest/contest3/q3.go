package contest3

import (
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	var curNum int = 1
	// fmt.Println(arr)
	for i, val := range arr {
		if i == 0 && val != 1 {
			arr[i] = 1
		} else if !abs(val, curNum) {
			arr[i] = curNum + 1
		}
		curNum = arr[i]
	}
	// fmt.Println(arr)
	return arr[len(arr)-1]
}

func abs(a, b int) bool {
	// fmt.Println(a,b)
	if a-b <= 1 && a >= b {
		return true
	} else if a-b >= -1 && a <= b {
		return true
	}
	return false
}
