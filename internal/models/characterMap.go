package models

import (
	"sort"
	"strings"
)

// KeyValuePair ..
type KeyValuePair struct {
	Key   string `json:"letter"`
	Value int    `json:"number_of_times_seen"`
}

// CharacterMapSlice ..
type CharacterMapSlice []KeyValuePair

// NewCharacterMapSlice ..
func NewCharacterMapSlice(word string) CharacterMapSlice {
	characterMap := make(map[string]int)
	for _, c := range word {
		var char = strings.ToLower(string(c))
		if count, ok := characterMap[char]; ok {
			count++
			characterMap[char] = count
		} else {
			characterMap[char] = 1
		}
	}

	cm := make(CharacterMapSlice, len(characterMap))
	i := 0
	for k, v := range characterMap {
		cm[i] = KeyValuePair{Key: k, Value: v}
		i++
	}
	return cm
}

// SortByValue ..
func (cm CharacterMapSlice) SortByValue() {
	sort.Stable(cm)
}

// Len ..
func (cm CharacterMapSlice) Len() int {
	return len(cm)
}

// Less ..
func (cm CharacterMapSlice) Less(i, j int) bool {
	return cm[i].Value < cm[j].Value
}

// Swap ..
func (cm CharacterMapSlice) Swap(i, j int) {
	cm[i], cm[j] = cm[j], cm[i]
}
