package array

// O(n) time, O(n) space solution. 2 pointers
func sortedSquares(A []int) []int {
	// Finds first non-negative integer. -> First Pointer
	var i, j int
	for pos := range A {
		if A[pos] >= 0 {
			j = pos
			i = pos - 1
			break
		}
	}

	// No negative case
	if j == 0 {
		for pos := range A {
			// newArr? = append(,A[pos])
			A[pos] = A[pos] * A[pos]
		}
		return A
	}

	var newArr []int
	for j <= len(A)-1 || i >= 0 {
		if j > len(A)-1 && i >= 0 {
			newArr = append(newArr, A[i]*A[i])
			i--
			continue
		} else if i < 0 && j <= len(A)-1 {
			newArr = append(newArr, A[j]*A[j])
			j++
			continue
		}
		if A[j]*A[j] < A[i]*A[i] {
			newArr = append(newArr, A[j]*A[j])
			j++
		} else {
			newArr = append(newArr, A[i]*A[i])
			i--
		}
	}
	return newArr
}
