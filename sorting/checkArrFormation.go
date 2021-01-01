package sorting

/*
Question 1640: https://leetcode.com/problems/check-array-formation-through-concatenation/

Observation:
Do a linear run through destination array, have all the pieces stored onto a map by the first
elements (since the element are guaranteed DISTINCT).

We then checks all the element in each pieces to see if its correnspond to the element in destination
if not return false
**/

func canFormArray(arr []int, pieces [][]int) bool {
	piece := make(map[int][]int)
	for i := range pieces {
		piece[pieces[i][0]] = pieces[i]
	}

	for i := 0; i < len(arr); i++ {
		if _, ok := piece[arr[i]]; !ok {
			return false
		} else {
			cur := arr[i]
			for u := range piece[cur] {
				if i == len(arr) || arr[i] != piece[cur][u] {
					return false
				}
				i++
			}
			i--
		}
	}
	return true
}
