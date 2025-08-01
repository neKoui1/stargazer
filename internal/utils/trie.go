package utils

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(s string) {
	node := t.root
	for _, char := range s {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(s string) bool {
	node := t.root
	for _, char := range s {
		if next, exists := node.children[char]; exists {
			node = next
		} else {
			node = nil
			break
		}
	}
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) (flag bool) {
	node := t.root
	for _, char := range prefix {
		if next, exists := node.children[char]; exists {
			node = next
		} else {
			return false
		}
	}
	return true
}
