package array

func maxArea(height []int) int {
	var maxArea = 0
	var leftPos = 0
	var rightPos = len(height) - 1
	var currentHeight, currentMax int
	for leftPos < rightPos {
		currentHeight = height[leftPos]
		if height[leftPos] > height[rightPos] {
			currentHeight = height[rightPos]
		}
		currentMax = (rightPos - leftPos) * currentHeight
		if currentMax > maxArea {
			maxArea = currentMax
		}
		if height[leftPos] < height[rightPos] {
			leftPos++
		} else {
			rightPos--
		}
	}
	return maxArea
}
