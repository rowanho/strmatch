package strlib

import (
	"errors"
)

var (
	ErrKeyNotFound = errors.New("Key not found in map")
	ErrWrongType = errors.New("Trying to populate the map with >1 type")
)

type trieNode struct {
	children []*trieNode
	isEnd bool
}

/*
* TrieNode constructor 
*/
func  newTrieNode() *trieNode {
	tn := new(trieNode)
	tn.children = make([]*trieNode, 26)
	tn.isEnd = false
	return tn
}


func ord(r rune) int {
	return int(r) - int('a')
}

type Trie struct {
	root *trieNode
}

/*
* Trie constructor 
*/
func NewTrie() *Trie {
	t := new(Trie)
	t.root = newTrieNode()
	return t
}


func (t *Trie) insert(key string) {
	keyr := []rune(key)
	current := t.root
	l := len(key)
	for lvl := 0; lvl < l; lvl++ {
		index := ord(keyr[lvl])
		if current.children[index] == nil {
			current.children[index] = newTrieNode()
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (t *Trie) search(key string)  bool {
	keyr := []rune(key)
	current := t.root
	l := len(key)
	for lvl := 0; lvl < l; lvl++ {
		index := ord(keyr[lvl])
		if current.children[index] == nil{
			return false
		}
		current = current.children[index]
	} 
	return (current != nil) && current.isEnd
}




