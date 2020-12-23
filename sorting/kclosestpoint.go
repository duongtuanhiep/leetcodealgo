package sorting

import "math"

func kClosest(points [][]int, K int) [][]int {
	if len(points) == 0 {
		return nil
	}

	sortedArr := MergeSort(points)

	result := make([][]int, K)
	for count := 0; count < K; count++ {
		result[count] = sortedArr[count]
	}
	return result
}

func MergeSort(slice [][]int) [][]int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right [][]int) [][]int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([][]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if math.Hypot(float64(left[i][0]), float64(left[i][1])) < math.Hypot(float64(right[j][0]), float64(right[j][1])) {
			slice[k] = make([]int, 2)
			slice[k][0] = left[i][0]
			slice[k][1] = left[i][1]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
