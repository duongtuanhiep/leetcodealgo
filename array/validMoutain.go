package array

func validMountainArray(arr []int) bool {

	// base case
	if len(arr) < 2 {
		return false
	}

	up := true

	// main loop
	for i := range arr {
		if i == len(arr)-1 && up {
			return false
		} else if i == len(arr)-1 && !up {
			return true
		}
		if arr[i] >= arr[i+1] && up {
			up = false
			if i == 0 {
				return false
			}
		}
		if arr[i] <= arr[i+1] && !up {
			return false
		}
	}
	return true
}
