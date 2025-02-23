package solutions

type Trie struct {
	children [26]*Trie
	isEnd    bool
	word     string
}

func TrieConstructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for _, ch := range word {
		ch -= 'a'
		if t.children[ch] == nil {
			t.children[ch] = &Trie{}
		}
		t = t.children[ch]
	}
	t.isEnd = true
	t.word = word
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}
