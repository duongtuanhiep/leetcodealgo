package bc52

/*
Test case:
[["#",".","#"]]

[["#",".","#","."],
["#",".","*","."],
["#",".",".","."]]

[[".","#","#"],
[".","#","#"],
["#","#","*"],
["#","*","."],
["#",".","*"],
["#",".","."]]


**/

func rotateTheBox(box [][]byte) [][]byte {
	var cur int
	for i := range box {
		for j := range box[i] {
			if box[i][j] == byte('*') {
				for k := j - 1; k >= 0 && box[i][k] != byte('*'); k-- {
					if cur > 0 {
						box[i][k] = byte('#')
						cur--
					} else {
						box[i][k] = byte('.')
					}
				}
			} else if box[i][j] == byte('#') {
				cur++
			}
		}
		for k := len(box[i]) - 1; k >= 0 && box[i][k] != byte('*'); k-- {
			if cur > 0 {
				box[i][k] = byte('#')
				cur--
			} else {
				box[i][k] = byte('.')
			}
		}
	}

	// DO LATER
	var newBox [][]byte
	for i := 0; i < len(box[0]); i++ {
		newArr := make([]byte, len(box))
		for j := 0; j < len(box); j++ {
			newArr[j] = box[len(box)-1-j][i]
		}
		newBox = append(newBox, newArr)
	}
	return newBox
}
