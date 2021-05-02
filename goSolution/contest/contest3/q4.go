package contest3

func closestRoom(rooms [][]int, queries [][]int) []int {
	var res []int

	for _, q := range queries {
		roomNum := -1
		diff := -1
		for _, room := range rooms {
			// checksize
			if room[1] < q[1] {
				continue
			}

			absDiff := abs(room[0], q[0])
			if diff < 0 {
				diff = absDiff
				roomNum = room[0]
			} else if diff > absDiff {
				diff = absDiff
				roomNum = room[0]
			} else if diff == absDiff && room[0] < roomNum {
				roomNum = room[0]
			}
		}

		res = append(res, roomNum)
	}

	return res
}

func abs(a, b int) int {
	if a-b <= 0 {
		return (a - b) * -1
	}
	return a - b
}
