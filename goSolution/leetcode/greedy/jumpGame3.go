package greedy

/*
Question 1306: https://leetcode.com/problems/jump-game-iii/

Greedily try to jump to every possible position, marked position as arrived
Stop when stack == 0 or reached 0

**/

func canReach(arr []int, start int) bool {
	cur := arr[start]
	if arr[start] == 0 {
		return true
	}
	arr[start] = -1

	var left, right bool
	if start+cur >= 0 && start+cur < len(arr) && start+cur > -1 {
		left = canReach(arr, start+cur)
	}
	if start-cur >= 0 && start-cur < len(arr) && start-cur > -1 {
		right = canReach(arr, start-cur)
	}
	return left || right
}
