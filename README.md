# strmatch
A trie data structure for prefix matching 

## Usage 
```
// Create a new trie instance
t := NewTrie()

// Insert data (eg each time a user makes a search in a search engine)
// Strings should be lowercase alpbabet only
searchesMade := []string{"their", "there", "there", "these", "these", "these"}
for _, s := range searchesMade {
   t.Insert(s)
}
```

We can then match a prefix, the returned list is sorted in descending order of frequency
```
t.PrefixMatch("the") // Should return {"these", "there", "their"}
```
