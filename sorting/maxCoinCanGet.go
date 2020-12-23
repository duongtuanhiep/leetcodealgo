package sorting

import (
	"math/rand"
)

func maxCoins(piles []int) int {
	var max int

	quickSort(piles)
	for pos := range piles {
		if pos > len(piles)-1-2*pos-1 {
			return max
		}
		max += piles[len(piles)-1-2*pos-1]
	}
	return max
}

func quickSort(arr []int) {
	sortQuick(arr, 0, len(arr)-1)
}

// Running quickSort with in mem indices. Left and right is REACHABLE (not out of bound)
func sortQuick(arr []int, left, right int) {
	// array part with 1 element is sorted
	if right-left < 1 {
		return
	}

	// Generate Pivot: Lowerbound + random range from lower->upper
	pivot := left + rand.Intn(right-left)

	// Swapping pivot to front of arr
	arr[pivot], arr[left] = arr[left], arr[pivot]

	// partrition based on pivot
	i := left + 1
	var j int
	for j = left + 1; j <= right; j++ {
		if arr[j] < arr[left] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	// Moving the pivot to its actual position
	arr[i-1], arr[left] = arr[left], arr[i-1]

	sortQuick(arr, left, i-2) // sorting left half
	sortQuick(arr, i, right)  // sorting right half
}
