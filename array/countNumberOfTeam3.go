package array

// Brute Force solution
// func numTeams(rating []int) int {
// 	var totalTeam int

// 	for i, _ := range rating {
// 		for j, _ := range rating {
// 			for k, _ := range rating {
// 				if i < j && j < k && (rating[i] > rating[j] && rating[j] > rating[k] || rating[i] < rating[j] && rating[j] < rating[k]) {
// 					totalTeam++
// 				}
// 			}
// 		}
// 	}
// 	return totalTeam
// }

// Smarter brute force solution
func numTeams(rating []int) int {
	var totalTeam int
	for i := range rating {
		for j := i; j <= len(rating)-1; j++ {
			for k := j; k <= len(rating)-1; k++ {
				if i < j && j < k && (rating[i] > rating[j] && rating[j] > rating[k] || rating[i] < rating[j] && rating[j] < rating[k]) {
					totalTeam++
				}
			}
		}
	}
	return totalTeam
}
