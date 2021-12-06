package constructive

/*
Question 1899: https://leetcode.com/problems/merge-triplets-to-form-target-triplet/

Idea: 3 linear scan through target array 
**/

func mergeTriplets(triplets [][]int, target []int) bool {
	for i, val := range target {
		found := false
		for j, num := range triplets {
			if num[i] == val {
				if goodBound(triplets[j], target) {
					found = true
					break
				}
			}
		}
		if !found {
			return false
		}

	}
	return true
}

func goodBound(cur, target []int) bool {
	for i := range cur {
		if cur[i] > target[i] {
			return false
		}
	}
	return true
}
