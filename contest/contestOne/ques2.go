package contestOne

func getSumAbsoluteDifferences(nums []int) []int {
	var res []int
	sum := make([][]int, len(nums))
	for i, _ := range sum {
		var holder []int
		holder = make([]int, i)
		sum[i] = holder
	}

	// Preproc
	for i := range nums {
		for j := 0; j < i; j++ {
			res := nums[i] - nums[j]
			if res > 0 {
				sum[i][j] = res

			} else {
				sum[i][j] = res * -1

			}
		}
	}

	// fmt.Println(sum)

	for i := range nums {
		cur := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			} else if j > i {
				cur += sum[j][i]
			} else {
				cur += sum[i][j]
			}

		}
		res = append(res, cur)
	}
	return res
}
