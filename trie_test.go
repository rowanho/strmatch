package strlib

import (
	"testing"
)


func TestSearch(t *testing.T) {
	keys := []string{"the", "a", "b", "cd", "he", "sh"}
	trie := NewTrie()
	for _, key := range keys {
		trie.insert(key)
	}

	foundStrings := []string{"the", "a", "cd"}
	wrongStrings := []string{"she", "ah", "e"}

	for _, found := range foundStrings {
		if trie.search(found) != true {
			t.Errorf("%s should be found in the trie but isn't", found)
		}		
	}

	for _, wrong := range wrongStrings {
		if trie.search(wrong) == true {
			t.Errorf("%s shouldn't  be found in the trie but is", wrong)
		}		
	}

}
