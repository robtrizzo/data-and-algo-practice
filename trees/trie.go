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

// TODO
func (t *Trie) Suggestions(substring string, depth int) []string {
	return []string{}
}
