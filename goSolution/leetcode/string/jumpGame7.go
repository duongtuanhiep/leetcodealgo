package string

/*
Question 1871: https://leetcode.com/contest/weekly-contest-242/problems/jump-game-vii/

Idea: 
- Create a seperate array to mark position has checked so far. 
- Starting from beginning, on each iteration mark all position travelable from the initial spot.
- The key idea is to avoid marking spot that has marked before. For that we only check part that 
has not been searched by the previous iteration. 
- Note that we should only start the iteration at the position that has been marked (can travel 
	to from prev iteration)
*/

func canReach(s string, minJump int, maxJump int) bool {
	step := make([]bool, len(s))
    step[0] = true
    
	lastReach := 0
    for i := range s {
		if s[i] == '0' && step[i] == true {
			start := max(i+minJump, lastReach)
			for j := start; j <= i+maxJump && j < len(s); j++ {
                if s[j] == '0' {
    				step[j] = true
                }
			}
			lastReach = i + maxJump
		}
	}
	return step[len(step)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
