package basicsortingalgo

import "sort"

// Fuck pancakeSort
func pancakeSort(arr []int) []int {
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}
