package main

import "errors"

// Dictionary Dictionary
type Dictionary map[string]string

// ErrNotFound dictionaryのkeyが見つからず、正常にvalueを取れなかったとき。
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search Search
func (d Dictionary) Search(word string) (string, error) {

	// map 2番目の戻り値はkeyが正常に検出されたかを判別する。（bool）
	definication, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definication, nil
}
