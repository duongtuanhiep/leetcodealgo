package greedy

/*
Question 763: https://leetcode.com/problems/partition-labels/

Idea: a map or an array to keep track of last occurence of character
Greadily tries to make a substring containing only existed elem

**/

func partitionLabels(s string) []int {
	lastLocation := make(map[byte]int)

	for i := len(s) - 1; i >= 0; i-- {
		if _, found := lastLocation[s[i]]; !found {
			lastLocation[s[i]] = i
		}
	}

	var res []int
	for i := 0; i < len(s); {
		// Getting the right bound of the partitions.
		bound := lastLocation[s[i]]
		count := 1
		for j := i + 1; j <= bound; j++ {
			if bound < lastLocation[s[j]] {
				bound = lastLocation[s[j]]
			}

			count++
		}

		res = append(res, count)
		i += count
	}

	return res
}
