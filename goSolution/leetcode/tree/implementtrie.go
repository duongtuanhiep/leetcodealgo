package tree

/*
Question 208: https://leetcode.com/problems/implement-trie-prefix-tree
*/

// Comment for Trie / Prefix Tree / Dictionary
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

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		if cur.Next[c-'a'] == nil {
			return false
		}
		cur = cur.Next[c-'a']
	}
	return cur.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		if cur.Next[c-'a'] == nil {
			return false
		}
		cur = cur.Next[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
