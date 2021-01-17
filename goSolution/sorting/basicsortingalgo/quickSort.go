package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// newArr := []int{7, 7, 8, 7, 5, 10, 8}
	// newArr := []int{7, 5, 7, 7, 8, 8, 5, 10, 7}
	newArr := []int{4, 4, 17, 7, 16, 16, 16, 15, 2, 3, 1, 17, 6, 12, 9}
	//[4,4,17,7,16,16,16,15,2,3,1,17,6,12,9]
	// sort.SliceStable(newArr, func(i, j int) bool { return newArr[i] < newArr[j] })
	quickSort(newArr)
	fmt.Println(newArr)
}

func quickSort(arr []int) {
	sortQuick(arr, 0, len(arr)-1)
}

// Running quickSort with in mem indices. Left and right is REACHABLE (not out of bound)
func sortQuick(arr []int, left, right int) {

	// array part with 1 element is sorted by default. NOT 2 ELEMENTS THX
	if right-left < 1 {
		return
	}

	// Generate Pivot: Lowerbound + random range from lower->upper
	pivot := left + rand.Intn(right-left)

	// Swapping pivot to front of arr
	arr[pivot], arr[left] = arr[left], arr[pivot]

	// partrition based on pivot
	// i holds the boundary between elements that is smaller and greater than pivot
	// j holds the partritioned part of array
	// i is the position to put in NEXT element that is smaller than pivot
	// SWAP smaller value with LEFTMOST elements that is greater than j
	i := left + 1
	var j int
	for j = left + 1; j <= right; j++ {
		if arr[j] < arr[left] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	// Moving the pivot to its actual position
	// SWAP the RIGHTMOST elements that is smaller than pivot with fronts
	arr[i-1], arr[left] = arr[left], arr[i-1]

	sortQuick(arr, left, i-2) // sorting left half
	sortQuick(arr, i, right)  // sorting right half
}
