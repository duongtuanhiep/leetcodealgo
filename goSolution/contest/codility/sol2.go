package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

/*
Formal requirement: Given an array of []int in the range 4-200
return total number of ways if possible to remove a single value
so that its in the format of Hi-lo-hi-lo... or lo-hi-lo...

Since the bound is so small, perhaps we can try to remove each numbers
and see if the remaining is good
**/

func Solution(A []int) int {
	if checkGood(A) {
		return 0
	}

	var res int
	for i := 0; i < len(A); i++ {
		newArr := make([]int, len(A))
		copy(newArr, A)
		if i > 0 {
			newArr = append(newArr[:i], newArr[i+1:]...)
		} else if i == 0 {
			newArr = newArr[i+1:]
		}
		if checkGood(newArr) {
			res++
		}
	}

	if res == 0 {
		return -1
	}
	return res
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
