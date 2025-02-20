package solutions

type WordDictionary struct {
	trieRoot *Trie
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{
		trieRoot: &Trie{},
	}
}

func (w *WordDictionary) AddWord(word string) {
	w.trieRoot.Insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	var dfs func(int, *Trie) bool
	dfs = func(index int, node *Trie) bool {
		if index == len(word) {
			return node.isEnd
		}

		if word[index] == '.' {
			for _, child := range node.children {
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		} else {
			ch := word[index] - 'a'
			child := node.children[ch]
			if child != nil && dfs(index+1, child) {
				return true
			}
		}
		return false
	}
	return dfs(0, w.trieRoot)
}
