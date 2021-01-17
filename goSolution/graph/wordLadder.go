package graph

import (
	"fmt"
	"math"
)

/*
"lost"
"miss"
["most","mist","miss","lost","fist","fish"]

"hit"
"cog"
["hot","dot","dog","lot","log","cog"]

"game"
"thee"
["frye","heat","tree","thee","game","free","hell","fame","faye"]
**/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var step = 0

	// Preproc, creates list of node
	var words []WordNode
	var wordCharMap []map[byte]int
	var existed = false
	wordList = append(wordList, beginWord) // implied
	wordList[len(wordList)-1], wordList[0] = wordList[0], wordList[len(wordList)-1]
	for i := range wordList {
		if endWord == wordList[i] {
			existed = true
		}
		var word WordNode
		wordMap := make(map[byte]int)
		for j := range wordList[i] {
			wordMap[wordList[i][j]]++
		}
		word.Word = wordList[i]
		wordCharMap = append(wordCharMap, wordMap)
		words = append(words, word)
	}

	if !existed {
		return step
	}

	// linking
	for i := range words {
		for j := range words {
			if i != j {
				match, req := 0, len(words[j].Word)-1
				curMap := make(map[byte]int)
				for key, value := range wordCharMap[i] {
					curMap[key] = value
				}
				for k := range words[j].Word {
					if _, ok := wordCharMap[i][words[j].Word[k]]; ok && curMap[words[j].Word[k]] > 0 {
						match += int(math.Min(float64(curMap[words[j].Word[k]]), float64(wordCharMap[j][words[j].Word[k]])))
						curMap[words[j].Word[k]]--
					}
				}
				if match >= req {
					words[i].Reachable = append(words[i].Reachable, &words[j])
				}
			}
		}
	}

	// BFS now
	var queue []WordNode
	visited := make(map[string]bool)
	queue = append(queue, words[0])
	for len(queue) != 0 {
		length := len(queue)
		step++
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			visited[cur.Word] = true
			if cur.Word == endWord {
				return step
			}
			for i := range cur.Reachable {
				if !visited[cur.Reachable[i].Word] {
					queue = append(queue, *cur.Reachable[i])
					visited[cur.Reachable[i].Word] = true
				}
			}
			fmt.Printf("Current Word : %v\t step: %v\n", cur.Word, step)
		}
	}
	if visited[endWord] {
		return step
	}
	return 0
}

type WordNode struct {
	Word      string
	Reachable []*WordNode
}
