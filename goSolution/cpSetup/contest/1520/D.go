package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReaderSize(os.Stdin, 4096)
var writer *bufio.Writer = bufio.NewWriterSize(os.Stdout, 4096)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var testCase int
	scanf("%d\n", &testCase)
	for ; testCase > 0; testCase-- {

		var arrLen int
		scanf("%d\n", &arrLen)
		var arr []int
		for i := 0; i < arrLen; i++ {
			var val int
			scanf("%d ", &val)
			arr = append(arr, val)
		}
		scanf("\n")

		var res int64
		count := make(map[int]int)
		for i, val := range arr {
			count[i-val]++
		}

		for _, val := range count {
			if val > 1 {
				res += int64(int64(val) * int64(val-1) / 2)
			}
		}

		printf("%d\n", res)
	}
}
