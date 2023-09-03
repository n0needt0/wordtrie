package main

import (
	"fmt"
	"github.com/n0needt0/wordtrie"
)

func main() {
	trie := initTrie()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.insert(words[i])
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.find(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}
