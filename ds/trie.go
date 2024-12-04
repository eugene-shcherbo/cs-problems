package ds

type Trie struct {
	children map[rune](*Trie)
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{make(map[rune]*Trie), false}
}

func (t *Trie) Insert(word string) {
	currNode := t

	for _, r := range word {
		if _, exists := currNode.children[r]; !exists {
			currNode.children[r] = NewTrie()
		}
		currNode = currNode.children[r]
	}

	currNode.isWord = true
}

func (t *Trie) HasWord(word string) bool {
	node := findNode(t, word)
	if node == nil {
		return false
	}
	return node.isWord
}

func (t *Trie) HasPrefix(prefix string) bool {
	return findNode(t, prefix) != nil
}

func (t *Trie) Delete(word string) {

}

func findNode(root *Trie, prefix string) *Trie {
	curr := root

	for _, r := range prefix {
		if _, exists := curr.children[r]; !exists {
			return nil
		}
		curr = curr.children[r]
	}

	return curr
}
