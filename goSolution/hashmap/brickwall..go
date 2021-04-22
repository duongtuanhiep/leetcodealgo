package hashmap

/*
Question 554: https://leetcode.com/problems/brick-wall/

Idea:
- Create a hashmap to store all the possible "line" on each row
- On each line, increase the number of "line" at a spot by 1.

Scan through the last hash map and output least cut

*/

func leastBricks(wall [][]int) int {
	cracks := make(map[int]int)
	for i := range wall {
		var spot int
		for j := 0; j < len(wall[i])-1; j++ {
			spot += wall[i][j]
			cracks[spot]++
		}
	}

	var cut int
	for _, val := range cracks {
		if val > cut {
			cut = val
		}
	}

	return len(wall) - cut
}
