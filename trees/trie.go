package trees

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	n := t.root
	for _, r := range word {
		_, exists := n.children[r]
		if !exists {
			n.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		n = n.children[r]
	}
	n.isWord = true
}

func (t *Trie) Search(word string) bool {
	n := t.root
	for _, r := range word {
		_, exists := n.children[r]
		if !exists {
			return false
		}
		n = n.children[r]
	}
	return n.isWord
}

func (t *Trie) Suggestions(substring string, depth int) []string {
	n := t.root
	for _, r := range substring {
		_, exists := n.children[r]
		if !exists {
			return []string{}
		}
		n = n.children[r]
	}

	return t.collectWords(n, substring, len(substring)+depth)
}

func (t *Trie) collectWords(n *TrieNode, substring string, maxLen int) []string {
	var words []string
	if n.isWord {
		words = append(words, substring)
	}

	if len(substring) == maxLen {
		return words
	}

	for r, child := range n.children {
		newSubstring := substring + string(r)
		words = append(words, t.collectWords(child, newSubstring, maxLen)...)
	}
	return words
}
