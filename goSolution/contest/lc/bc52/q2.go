package bc52

func memLeak(memory1 int, memory2 int) []int {
	var i int
	for i = 0; memory1 >= i || memory2 >= i; i++ {
		if memory1 >= memory2 {
			memory1 -= i
		} else {
			memory2 -= i
		}
	}

	return []int{i, memory1, memory2}
}
