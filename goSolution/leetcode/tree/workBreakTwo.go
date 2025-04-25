package tree

/*

Question 140: https://leetcode.com/problems/word-break-ii

**/

func wordBreak(s string, wordDict []string) []string {
	t := Constructor()
	for _, word := range wordDict {
		t.Insert(word)
	}

	return tryBuild(s, &t)
}

func tryBuild(s string, t *Trie) []string {
	var cur string
	var res []string
	for _, c := range s {
		if 
	}
}

type LetterNode struct {
	Char  rune
	IsEnd bool
	Next  []*LetterNode
}

type Trie struct {
	root *LetterNode
}

func Constructor() Trie {
	return Trie{root: &LetterNode{IsEnd: false, Next: make([]*LetterNode, 26)}}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		if cur.Next[c-'a'] == nil {
			new := &LetterNode{Char: c, IsEnd: false, Next: make([]*LetterNode, 26)}
			cur.Next[c-'a'] = new
		}
		cur = cur.Next[c-'a']
	}

	cur.IsEnd = true
}
