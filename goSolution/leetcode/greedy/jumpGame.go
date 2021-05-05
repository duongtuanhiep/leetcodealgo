package greedy

var (
	counter = 0
)

func canJump(nums []int) bool {

	maxReach := 0
	currentVal := nums[0]
	nextIndex := 0
	// len = 1
	if len(nums) == 1 {
		return true
	}
	if currentVal == 0 && len(nums) != 1 {
		return false
	}
	for i := 1; i <= currentVal; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
			nextIndex = i
		}
	}
	return canJump(nums[nextIndex:])
}
