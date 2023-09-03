package wordtrie

const (
	//ALPHABET_SIZE total chars in english
	ALPHABET_SIZE = 26
)

type trieNode struct {
	children  [ALPHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

const (
	ARUNE = 'a'
)

func (t *trie) insert(word string) {
	wordLen := len(word)
	current := t.root
	for i := 0; i < wordLen; i++ {
		//a == 0, so reduce rune(unicode) value by a (97)
		index := word[i] - ARUNE

		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word string) bool {
	wordLen := len(word)
	current := t.root
	for i := 0; i < wordLen; i++ {
		index := word[i] - ARUNE

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}

	if current.isWordEnd == true {
		return true
	}

	return false
}
