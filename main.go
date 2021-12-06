package main

import "fmt"

var (
	A    []int = []int{3, 2, 6, 4}
	arr1 []int = []int{3, 2, 6, 4}
	arr2 []int = []int{1, 3, 2, 4}
	arr3 []int = []int{3, 1, 2, 2}
	arr4 []int = []int{1, 2, 1, 6, 8, 5, 5}
	arr5 []int = []int{1, 2, 1, 6}
)

func main() {
	for i := 0; i < len(A); i++ {
		newArr := make([]int, len(A))
		copy(newArr, A)
		fmt.Println(newArr)
		if i > 0 {
			newArr = append(newArr[:i], newArr[i+1:]...)
		} else if i == 0 {
			newArr = newArr[i+1:]
		}
		fmt.Println(newArr, i)
	}

	// fmt.Println(checkGood(arr1, 0))
	// fmt.Println(checkGood(arr2, 0))
	// fmt.Println(checkGood(arr3, 0))
	// fmt.Println(checkGood(arr4, 0))
	// fmt.Println(checkGood(arr5, 0))
}

func checkGood(arr []int) bool {
	var up bool
	last := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if up && last > 0 {
				return false
			} else {
				up = true
				last = i
			}
		} else if arr[i] < arr[i-1] {
			if !up && last > 0 {
				return false
			} else {
				up = false
				last = i
			}
		} else {
			return false
		}
	}
	return true
}
