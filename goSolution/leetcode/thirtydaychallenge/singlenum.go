package thirtydaychallenge

func singleNumber(nums []int) int {
	var numMap = make(map[int]int)
	for _, val := range nums {
		numMap[val]++
	}
	for _, val := range nums {
		if numMap[val] == 1 {
			return val
		}
	}
	return 0
}
