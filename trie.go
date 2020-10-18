package strmatch

import (
	"github.com/bradfitz/slice"
)

type trieNode struct {
	children map[rune]*trieNode
	isEnd bool
	frequency int
}

// TrieNode constructor 
func  newTrieNode() *trieNode {
	tn := new(trieNode)
	tn.children = make(map[rune]*trieNode)
	tn.isEnd = false
	tn.frequency = 0
	return tn
}


type Trie struct {
	root *trieNode
}


// Trie constructor 
func NewTrie() *Trie {
	t := new(Trie)
	t.root = newTrieNode()
	return t
}

// Inserts a word into the trie
func (t *Trie) Insert(word string) {
	wordR := []rune(word)
	current := t.root
	l := len(word)
	for lvl := 0; lvl < l; lvl++ {
		r:= wordR[lvl]
		if _, exists := current.children[r]; !exists {
			current.children[r] = newTrieNode()
		}
		current = current.children[r]
	}
	current.frequency++
	current.isEnd = true
}


// Checks if a word is in the trie
func (t *Trie) Search(word string)  bool {
	wordR := []rune(word)
	current := t.root
	l := len(word)
	for lvl := 0; lvl < l; lvl++ {
		r:= wordR[lvl]
		if _, exists := current.children[r]; !exists {
			current.children[r] = newTrieNode()
		}
		current = current.children[r]
	} 
	return (current != nil) && current.isEnd
}

type prefixFrequency struct {
	p string
	f int
}

// Helper function - enumerates all possible words from the starting trieNode
func enumFromPrefix(tn *trieNode, prefix string, prefixes  *[]prefixFrequency) {
	if tn.isEnd {
		*prefixes = append(*prefixes, prefixFrequency{p: prefix, f: tn.frequency,})		
	}
	for r := range tn.children {
		enumFromPrefix(tn.children[r], prefix + string(r), prefixes)
	}
}

// Returns a list of the trie entries that match the given prefix, sorted by frequency
func (t *Trie) PrefixMatch(prefix string) []string {
	prefixR := []rune(prefix)
	current := t.root
	l := len(prefix) 
	for lvl := 0; lvl < l; lvl++ {
		r:= prefixR[lvl]
		if _, exists := current.children[r]; !exists {
			current.children[r] = newTrieNode()
		}
		current = current.children[r]
	} 
	if current == nil {
		return []string{}
	}
	prefixes := make([]prefixFrequency, 0)
	enumFromPrefix(current, prefix, &prefixes)
	slice.Sort(prefixes, func(i, j int) bool {
		return prefixes[i].f < prefixes[j].f
	})
	l = len(prefixes)
	res := make([]string, l)
	for i := 0; i < l; i++ {
		res[l - 1 -i] = prefixes[i].p
	}
	return res
}


