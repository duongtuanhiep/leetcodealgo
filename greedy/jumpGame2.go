package greedy

func jump(nums []int) int {
	result, _ := canJumpStep(nums, 0)
	return result
}

func canJumpStep(nums []int, step int) (int, bool) {

	maxReach := 0
	currentVal := nums[0]
	nextIndex := 0
	// len = 1
	if len(nums) == 1 {
		return step + 1, true
	}
	if currentVal == 0 && len(nums) != 1 {
		return 0, false
	}
	for i := 1; i <= currentVal; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
			nextIndex = i
		}
	}
	return canJumpStep(nums[nextIndex:], step+1)
}
