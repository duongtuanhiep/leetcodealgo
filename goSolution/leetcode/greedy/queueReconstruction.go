package greedy

import "fmt"

/*
Question 406: https://leetcode.com/problems/queue-reconstruction-by-height/

Key idea: maintaining a variant of people who are shorter than a specific number of height
**/

/*
Brute Force: O(N^2) solution
For each iteration, iteration through the people to finds the person who are shortest
possible that satify height rules.
**/
func reconstructQueue(people [][]int) [][]int {
	heightMap := make(map[int]int)

	// preproc
	for i := range people {
		heightMap[people[i][0]] = 0
	}

	var result [][]int
	for len(people) != 0 {
		var lowest int = -1 // index of shortest person that satify the queue rule
		for i := range people {
			if heightMap[people[i][0]] == people[i][1] {
				if lowest < 0 {
					lowest = i
				}
				if people[i][0] <= people[lowest][0] {
					lowest = i
				}
			}
		}

		for height := range heightMap {
			if height <= people[lowest][0] {
				heightMap[height]++
			}
		}

		fmt.Printf("Index %v Val %v \n ", lowest, people[lowest])
		fmt.Println(heightMap)

		// append to res
		result = append(result, people[lowest])

		// Update people array
		people[0], people[lowest] = people[lowest], people[0]
		people = people[1:]
	}

	return result
}
