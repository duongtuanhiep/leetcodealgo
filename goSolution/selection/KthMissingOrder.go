package selection

/* Observation:
	- Possibe number is from 1 to N
	- Number of missing element can either greater than biggest
	element, smaller than smallest elem or wedge somewhere in between.
	- Can identify this by finding possible missing element from 1 to
	max (last items) of the array K1 and possible missing element before
	first element (index 0) K2.
	Possible Case:
	- If Kth > K1 : to right of array. Returns arr[len(-1)] + Kth - K1
	- If Kth < K2 : to left of array. Returns Kth
	- If K1 > Kth > K2: recursively solves with boundary lo-hi and no of
	missing elem. This recursion can be stopped when hi-lo == 1 which means
	these 2 elems
**/
func findKthPositive(arr []int, k int) int {
	leftOfLast := arr[len(arr)-1] - len(arr)
	leftOfFirst := arr[0] - 1

	if k > leftOfLast {
		return k + len(arr)
	}
	if k <= leftOfFirst {
		return k
	}
	return binarySearchkTh(arr, len(arr)-1, 0, k-leftOfFirst)
}

func binarySearchkTh(arr []int, hi, lo, k int) int {
	if hi-lo < 2 {
		return arr[lo] + k
	}
	leftMiss := arr[(hi+lo)/2] - arr[lo] - ((hi+lo)/2 - lo)
	if k <= leftMiss {
		return binarySearchkTh(arr, (hi+lo)/2, lo, k)
	}
	return binarySearchkTh(arr, hi, (hi+lo)/2, k-leftMiss)
}
