package thirtydaychallenge

func moveZeroesArr(nums []int) []int {
	newArr := make([]int, len(nums))
	oCounter := 0
	lenArr := len(nums)
	if lenArr <= 1 {
		return nums
	}
	for i, val := range nums {
		if val != 0 {
			newArr[i-oCounter] = val
		} else if val == 0 {
			newArr[lenArr-oCounter-1] = 0
			oCounter++
		}
	}
	for i := range nums {
		nums[i] = newArr[i]
	}

	return newArr
}
