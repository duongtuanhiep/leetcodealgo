package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, K int) string {
	// write your code in Go 1.4

	dayInt := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	dayMapInt := make(map[string]int)
	dayMapInt["Mon"] = 0
	dayMapInt["Tue"] = 1
	dayMapInt["Wed"] = 2
	dayMapInt["Thu"] = 3
	dayMapInt["Fri"] = 4
	dayMapInt["Sat"] = 5
	dayMapInt["Sun"] = 6

	return dayInt[(dayMapInt[S]+K)%7]
}
