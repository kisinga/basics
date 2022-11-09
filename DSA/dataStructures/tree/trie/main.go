package main

import (
	"fmt"
	"strings"
)

// AlphabetSize is the number of possible characters int the trie
const AlphabetSize = 26

// Node is a point in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie just stores a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie is just a helper function to initialize a trie correctly
func InitTrie() *Trie {
	t := Trie{
		root: &Node{},
	}
	return &t
}

// InsertWord inserts a word into the trie
func (t *Trie) InsertWord(w string) {
	// make sure that all the letters in our word are lowercase
	w = strings.ToLower(w)
	currentNode := t.root
	for _, letter := range w {
		charIndex := letter - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search returns true if the word exists in our trie
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, letter := range w {
		charIndex := letter - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false

}

func main() {
	trie := InitTrie()
	trie.InsertWord("someone")
	trie.InsertWord("Something")
	trie.InsertWord("some")
	trie.InsertWord("one")
	trie.InsertWord("thing")
	fmt.Println(trie.Search("thing"))
	fmt.Println(trie.Search("one"))
	fmt.Println(trie.Search("two"))
}
