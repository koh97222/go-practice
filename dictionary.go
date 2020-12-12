package main

// Dictionary Dictionary
type Dictionary map[string]string

// Search Search
func (d Dictionary) Search(word string) string {
	return d[word]
}
