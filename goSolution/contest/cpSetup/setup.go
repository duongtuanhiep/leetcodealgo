package cpSetup

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

	}
}
