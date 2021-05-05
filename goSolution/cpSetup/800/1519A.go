package main

import (
	"bufio"
	"fmt"
	"os"
)

// tag math

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)
	for ; t > 0; t-- {
		var r, b, d int64
		scanf("%d %d %d\n", &r, &b, &d)

		// var res string
		if d > 0 {
			if r*(d+1) >= b && r <= b || b*(d+1) >= r && b <= r {
				printf("YES\n")
			} else {
				printf("NO\n")
			}
		} else if r != b {
			printf("NO\n")
		} else {
			printf("YES\n")
		}
	}
}
