package main

import (
	"bufio"
	"fmt"
	"os"
)

// TAG: dp/math

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

		var k, m, n, i, j int64
		scanf("%d %d %d \n", &n, &m, &k)

		// base
		if n == 1 && m == 1 && k != 0 {
			printf("NO\n")
		}

		var matrix [][]int64
		for i = n; i > 0; i-- {
			matrix = append(matrix, make([]int64, m))
		}

		// fmt.Println(matrix)
		// matrix[1][1] = 1
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				if i != 0 && j == 0 {
					matrix[i][j] += matrix[i-1][j] + 1
				} else if j != 0 {
					matrix[i][j] += matrix[i][j-1] + i + 1
				}
			}
		}

		// fmt.Println(matrix)
		if matrix[n-1][m-1] == k {
			printf("YES\n")
		} else {
			printf("NO\n")
		}

	}
}
