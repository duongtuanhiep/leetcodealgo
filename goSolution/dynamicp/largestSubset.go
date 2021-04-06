package dynamicp

/*
Question 474: https://leetcode.com/problems/ones-and-zeroes/

Idea:
- Create a "cost" array,
**/

// func findMaxForm(strs []string, m int, n int) int {

// 	if len(strs) == 0 {
// 		return 0
// 	}

// 	zero, one := countLen(strs[0])
// 	if zero > m || one > n {
// 		return findMaxForm(strs[1:], m, n)
// 	}
// 	// fmt.Println(strs)
// 	pick := 1 + findMaxForm(strs[1:], m-zero, n-one)
// 	notPick := findMaxForm(strs[1:], m, n)
// 	if pick > notPick {
// 		return pick
// 	}
// 	return notPick
// }

func findMaxForm(strs []string, m int, n int) int {

	var cost [][]int
	for _, str := range strs {
		zero, one := countLen(str)
		cost = append(cost, []int{zero, one})
	}

	return findMaxComboInt(cost, m, n, 0, len(cost)-1)
}

func findMaxComboInt(nums [][]int, m, n, cur, pos int) int {

	if pos < 0 {
		return cur
	}

	if nums[pos][0] > m || nums[pos][1] > n {
		return cur + findMaxComboInt(nums, m, n, cur, pos-1)
	}

	pick := findMaxComboInt(nums, m-nums[pos][0], n-nums[pos][1], cur+1, pos-1)
	notPick := findMaxComboInt(nums, m, n, cur, pos-1)
	if pick > notPick {
		return cur + pick
	}
	return notPick
}

func countLen(str string) (int, int) {
	zero, one := 0, 0
	for _, char := range str {
		if char == '0' {
			zero++
		} else {
			one++
		}
	}
	// fmt.Println(str, zero, one)
	return zero, one
}
