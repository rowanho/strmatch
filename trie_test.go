package strmatch

import (
	"testing"
)


func TestSearch(t *testing.T) {
	words := []string{"the", "a", "b", "cd", "he", "sh"}
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	foundStrings := []string{"the", "a", "cd"}
	wrongStrings := []string{"she", "ah", "e"}

	for _, found := range foundStrings {
		if trie.Search(found) != true {
			t.Errorf("%s should be found in the trie but isn't", found)
		}		
	}

	for _, wrong := range wrongStrings {
		if trie.Search(wrong) == true {
			t.Errorf("%s shouldn't  be found in the trie but is", wrong)
		}		
	}

}

/*
* Check slice of strings is the same (disregard order)
*/
func sameStringsSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true  
}

func TestPrefixMatch(t *testing.T) {
	words := []string{"their", "the", "there", "these", "bolt", "bowl", "boolean"}
	frequencies := []int{5, 4, 3, 7, 6, 8, 2}
	trie := NewTrie()
	for i, word := range words {
		for j := 0; j < frequencies[i]; j++ {
			trie.Insert(word)			
		} 
	}

	prefixes := []string{"the", "bo", "", "h", "a"}
	prefixMatches := [][]string{
		{"these", "their", "the","there"},
		{"bowl", "bolt", "boolean"},
		{"bowl", "these", "bolt", "their", "the", "there", "boolean"},
		{},
		{},
	}

	for i, prefix := range prefixes {
		r := trie.PrefixMatch(prefix) 
		if !sameStringsSlice(r, prefixMatches[i]) {
			t.Errorf("Matches found %v don't match correct matches %v", r, prefixMatches[i])			
		}
	}
}