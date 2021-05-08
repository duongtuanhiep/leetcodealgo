package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var testCase int
	scanf("%d\n", &testCase)
	for ; testCase > 0; testCase-- {

		var len int
		var taskSolved string
		var res string = "YES\n"
		scanf("%v \n %v \n", &len, &taskSolved)

		solved := make(map[rune]bool)
		prev := '1'
		for _, char := range taskSolved {
			if solved[char] && char != prev {
				res = "NO\n"
				break
			} else if char != prev {
				solved[prev] = true
				prev = char
			}
		}

		printf(res)
	}
}
