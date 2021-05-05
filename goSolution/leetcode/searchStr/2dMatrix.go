package searchStr

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, value := range matrix {
		if len(value) == 0 {
			return false
		}
		if target > value[0] && target < value[len(value)-1] && binarySearch(target, value) {
			return true
		}
	}
	return false
}

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2
		if haystack[median] == needle {
			return true
		}
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
