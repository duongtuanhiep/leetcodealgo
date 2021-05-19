package string

import (
	"fmt"
	"strings"
)

/*
Question 609: https://leetcode.com/problems/find-duplicate-file-in-system/

Idea:
- We create a map to store all the file. Key is the content, value is an array holds all file location that has that files.
- Iterate through the map again and returns all the location of file has duplicate contents.

Answering followup:
- We can keep the original file strategy but modify our saving method.
- Create a bloom filter, Hash the raw content as the key. Instead storing the directory serparately we go through bloom filter first.
- Paralelize it or do it concurrently. Split file into smaller chunk and hash those instead
- same complexity
- Keep a map to store the directory of last dup. Do a check before putting them in ? Combination of hashmap and bloomfilter

Check the solution discussion for more detailed follow up:
https://leetcode.com/problems/find-duplicate-file-in-system/solution/
**/

func findDuplicate(paths []string) [][]string {
	loc := make(map[string][]string)
	for _, val := range paths {
		combo := strings.Split(val, " ")
		for _, f := range combo[1:] {
			fileData := strings.Split(f, "(")
			if _, ok := loc[fileData[1]]; !ok {
				loc[fileData[1]] = []string{fmt.Sprintf("%v/%v", combo[0], fileData[0])}
			} else {
				loc[fileData[1]] = append(loc[fileData[1]], fmt.Sprintf("%v/%v", combo[0], fileData[0]))
			}
		}
	}

	var res [][]string
	for _, val := range loc {
		if len(val) > 1 {
			res = append(res, val)
		}
	}

	return res
}
