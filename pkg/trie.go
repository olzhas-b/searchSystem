package pkg

import "errors"

var NotFoundError = errors.New("Not found")

type Trie struct {
	children         map[rune]*Trie
	isEndOfWord      bool
	numberOfSearched int64
}

func NewTrie() *Trie {
	return &Trie{children: make(map[rune]*Trie)}
}

type Pair struct {
	NumberOfSearched int64
	Word             string
}

func (t *Trie) insert(word string) (err error) {
	curr := t
	for _, letter := range word {
		if curr.children[letter] == nil {
			curr.children[letter] = NewTrie()
		}
		curr = curr.children[letter]
	}
	curr.numberOfSearched++
	curr.isEndOfWord = true
	return err
}

func (t *Trie) search(word string) (result []Pair, err error) {
	curr := t
	for _, letter := range word {
		if curr.children == nil {
			return result, NotFoundError
		}
		curr = curr.children[letter]
	}

	recursivelyAddWords(&result, curr, word)
	return
}

func recursivelyAddWords(result *[]Pair, root *Trie, word string) {
	if result == nil || root == nil {
		return
	}

	for letter, curr := range root.children {
		if curr.isEndOfWord {
			*result = append(*result, Pair{Word: word + string(letter), NumberOfSearched: curr.numberOfSearched})
		}
		recursivelyAddWords(result, curr, word+string(letter))
	}
}
