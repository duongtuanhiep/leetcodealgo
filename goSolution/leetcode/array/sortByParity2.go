package array

// store them numba to 2 distincguish Arr
//
func sortArrayByParityII(A []int) []int {
	var even, odd []int
	for i := range A {
		if A[i]%2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}
	for i := range A {
		if i%2 == 0 {
			A[i] = even[i/2]
		} else {
			A[i] = odd[i/2]
		}
	}
	return A
}
