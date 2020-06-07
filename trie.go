package strmatch

import (
	"github.com/bradfitz/slice"
)

type trieNode struct {
	children []*trieNode
	isEnd bool
	frequency int
}

/*
* TrieNode constructor 
*/
func  newTrieNode() *trieNode {
	tn := new(trieNode)
	tn.children = make([]*trieNode, 26)
	tn.isEnd = false
	tn.frequency = 0
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

/*
* Inserts a word into the trie
*/
func (t *Trie) Insert(word string) {
	wordR := []rune(word)
	current := t.root
	l := len(word)
	for lvl := 0; lvl < l; lvl++ {
		index := ord(wordR[lvl])
		if current.children[index] == nil {
			current.children[index] = newTrieNode()
		}
		current = current.children[index]
	}
	current.frequency++
	current.isEnd = true
}

/*
* Checks if a word is in the trie
*/
func (t *Trie) Search(word string)  bool {
	wordR := []rune(word)
	current := t.root
	l := len(word)
	for lvl := 0; lvl < l; lvl++ {
		index := ord(wordR[lvl])
		if current.children[index] == nil{
			return false
		}
		current = current.children[index]
	} 
	return (current != nil) && current.isEnd
}

type prefixFrequency struct {
	p string
	f int
}

/*
* Helper function - enumerates all possible words from the starting trieNode
*/ 
func enumFromPrefix(tn *trieNode, prefix string, prefixes  *[]prefixFrequency) {
	if tn == nil {
		return
	}
	if tn.isEnd {
		*prefixes = append(*prefixes, prefixFrequency{p: prefix, f: tn.frequency,})		
	}
	for i := range tn.children {
		r := rune(i + int('a'))
		enumFromPrefix(tn.children[i], prefix + string(r), prefixes)
	}
}

/* 
* Returns a list of the trie entries that match the given prefix, sorted by frequency
*/ 
func (t *Trie) PrefixMatch(prefix string) []string {
	prefixR := []rune(prefix)
	current := t.root
	l := len(prefix) 
	for lvl := 0; lvl < l; lvl++ {
		index := ord(prefixR[lvl])
		if current.children[index] == nil{
			return []string{}
		}
		current = current.children[index]
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


