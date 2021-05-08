package array

/**
Question 1014: https://leetcode.com/problems/best-sightseeing-pair/submissions/

Idea: Greedy solution
Back < cur changes: no problem, the current max we have improved.
Front < cur changes: we make sure we have the maximum up until that point (combination of front and back).
- We changes the front to current when sum(cur) > sum(front) and there are remaining elements at the back
- On each iteration, we try to "extend" the back to see if it is greater than our current max
*/

// Dumb O(n^2)
func maxScoreSightseeingPair(A []int) int {
	var max int
	var front int = 0
	for i := range A {
		if A[front]+front+A[i]-i > max && front != i {
			max = A[front] + front + A[i] - i
		}
		if A[i]+i >= A[front]+front && i < len(A)-1 {
			front = i
		}
	}
	return max
}
