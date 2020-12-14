package main

import "errors"

// Dictionary Dictionary
type Dictionary map[string]string

// DictionaryErr DictionaryErr
type DictionaryErr string

// ErrNotFound dictionaryのkeyが見つからず、正常にvalueを取れなかったとき。
var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExist        = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search Search
func (d Dictionary) Search(word string) (string, error) {

	// map 2番目の戻り値はkeyが正常に検出されたかを判別する。（bool）
	definication, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definication, nil
}

// Add 辞書にwordを加えます
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// mapは参照型なので、ポインタを渡さなくても変更できる。
		// 基礎となるデータ構造の参照を保持します。
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

// Update 辞書のwordを更新します。
func (d Dictionary) Update(word, newDifinition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDifinition
	default:
		return err
	}

	return err
}

// Delete 辞書のwordを削除します。
func (d Dictionary) Delete(word string) {
	// mapで機能する組み込み関数
	// 第一引数：map 第二引数：key
	delete(d, word)
}
