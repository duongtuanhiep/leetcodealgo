package cardPoints

/*
Question 1423: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/


Idea: Sliding windows
- Since we can only take from 2 ends, that means the "middle" part is consecutive.
The question becomes finds the min sum of the sub array length size - k.
- We assume the sub arr is at leftmost. On each iteration we adds the current elem in
that sum and substract the leftmost element of the sub arrays out. Keep track of
max = totalSum - subArrSum


Test case
[1,79,80,1,1,1,200,1]
5
**/

func maxScore(cardPoints []int, k int) int {

	var max, middle, cardSum int
	for i, val := range cardPoints {
		cardSum += val
		if i < len(cardPoints)-k {
			middle += val
		}
	}

	max = cardSum - middle
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		middle += cardPoints[i] - cardPoints[i-len(cardPoints)+k]
		if cardSum-middle > max {
			max = cardSum - middle
		}
	}

	return max
}
