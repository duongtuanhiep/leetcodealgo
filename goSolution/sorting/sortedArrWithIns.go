package sorting

/*
Question 1649: https://leetcode.com/problems/create-sorted-array-through-instructions/

Idea:
- Modify binary search to keep tracks of number of left and right and perform insertion to that position
 of array.

TEST CASE:
[4,14,10,2,5,3,8,19,7,20,12,1,9,15,13,11,18,6,16,17]
[1,3,3,3,2,4,2,1,2]
[1,5,6,2]

**/
func createSortedArray(instructions []int) int {
	var cost int
	var resArr []int

	if len(instructions) < 2 {
		return 0
	}
	resArr = append(resArr, instructions[0])
	for _, val := range instructions[1:] {
		left, right := leftBound(resArr, len(resArr), val), rightBound(resArr, len(resArr), val)
		resArr = append(resArr, 0)
		copy(resArr[left+1:], resArr[left:])
		resArr[left] = val
		cost += getMin(left, len(resArr)-right-2)
	}
	return cost % (1000000000 + 7)
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func leftBound(arr []int, n, num int) int {
	hi, lo := n, 0
	for lo < hi {
		var pivot = (hi + lo) / 2
		if arr[pivot] < num {
			lo = pivot + 1
		} else {
			hi = pivot
		}
	}
	return lo
}

func rightBound(arr []int, n, num int) int {
	hi, lo := n, 0
	for lo < hi {
		var pivot = (hi + lo) / 2
		if arr[pivot] > num {
			hi = pivot
		} else {
			lo = pivot + 1
		}
	}
	if hi == 0 {
		return 0
	}
	return hi - 1
}
