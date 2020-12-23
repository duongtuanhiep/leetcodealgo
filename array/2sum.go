package array

import "fmt"

// func main() {
// 	inputMap := []int{3, 2, 4}

// 	fmt.Print(twoSum(inputMap, 6))

// }

func twoSum(nums []int, target int) []int {

	// RETURNING INDICES

	// Preprocessing
	numberMap := make(map[int]int)
	for pos, value := range nums {
		numberMap[value] = pos
		fmt.Printf("Added num %v . Position %v \n", value, numberMap[value])
	}

	// answer
	for currentPos, current := range nums {
		fmt.Printf("Current num %v . try to look for %v - %v in arr \n", current, target, current)
		if numberMap[target-current] != 0 && currentPos != numberMap[target-current] { // Has to be distinct

			return []int{currentPos, numberMap[target-current]}
		}
	}

	return nil
}
