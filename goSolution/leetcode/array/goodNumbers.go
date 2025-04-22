package array

func goodTriplets(nums1 []int, nums2 []int) int64 {

	// Preprocess: index of each number in n1 and n2
	index1, index2 := make([]int, len(nums1)), make([]int, len(nums1))
	for i, val := range nums1 {
		index1[val] = i
	}
	for i, val := range nums2 {
		index2[val] = i
	}

	less, more := make([]int, len(nums1)), make([]int, len(nums1))
	// Build common lessArr: each index i tells how many number that has smaller index in both n1 and n2
	for i, val := range nums1 {
		less[val] =  

	}

	// build common moreArr

	// Preprocess: build arr2 where for each values k,
	// it shows how many numbers smaller than k with lower index in arr1 and arr2

	return 0
}
