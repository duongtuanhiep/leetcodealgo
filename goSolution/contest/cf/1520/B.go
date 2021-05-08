package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		var limit int
		scanf("%v\n", &limit)

		var res int
		limitStr := len(strconv.Itoa(limit))

		if limit > 1 {
			res += (limitStr - 1) * 9
		}

		// build int
		var agaisnt string
		for i := 0; i < limitStr; i++ {
			agaisnt += "1"
		}

		numA, _ := strconv.Atoi(agaisnt)
		for limit >= 0 {
			limit -= numA
			res++
		}

		printf("%v\n", res-1)
	}
}
