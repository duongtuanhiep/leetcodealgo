package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var (
	ArrStr = []string{}
)

func generateParenthesis(n int) []string {
	ArrStr = []string{}

	var singleStr = make([]string, n*2)
	if n > 0 {
		printParenthesis(0, n, 0, 0, singleStr)
	}
	return ArrStr
}

func printParenthesis(pos int, pair int, open int, close int, currStr []string) {
	if open < pair {
		currStr[pos] = "("
		printParenthesis(pos+1, pair, open+1, close, currStr)
	}
	if open > close {
		currStr[pos] = ")"
		printParenthesis(pos+1, pair, open, close+1, currStr)
	}
	if pos == pair*2 {
		ArrStr = append(ArrStr, toString(currStr))
	}
}

func toString(arrStr []string) string {
	result := ""
	for _, val := range arrStr {
		result = result + val
	}
	return result
}

// void printParenthesis(int n)
// {
//     if(n > 0)
//         _printParenthesis(0, n, 0, 0);
//     return;
// }

// void _printParenthesis(int pos, int n, int open, int close)
// {
//     static char str[MAX_SIZE];

//     if(close == n)
//     {
//         printf("%s \n", str);
//         return;
//     }
//     else
//     {
//         if(open > close)
//         {
//             str[pos] = '}';
//             _printParenthesis(pos+1, n, open, close+1);
//         }

//         if(open < n)
//         {
//         str[pos] = '{';
//         _printParenthesis(pos+1, n, open+1, close);
//         }
//     }
// }
