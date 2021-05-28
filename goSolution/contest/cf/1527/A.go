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

		var res, n, cur int
		scanf("%d\n", &n)

		cur = n
		for i := 1; i < n; i++ {
			cur = cur & (n - i)
			if cur == 0 {
				// printf("%v %v \n", n, i)
				res = n - i
				break
			}
		}

		printf("%d\n", res)
	}
}
