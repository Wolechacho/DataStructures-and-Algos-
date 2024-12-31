package main

import (
	"fmt"
)

func Sentencesimilarity(sentence1, sentence2 []string, pairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	hashset := map[string]struct{}{}

	for _, pair := range pairs {
		format := fmt.Sprintf("%s#%s", pair[0], pair[1])
		hashset[format] = struct{}{}
	}

	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] != sentence2[i] {
			format := fmt.Sprintf("%s#%s", sentence1[i], sentence2[i])
			if _, exist := hashset[format]; !exist {
				format2 := fmt.Sprintf("%s#%s", sentence2[i], sentence1[i])
				if _, exist2 := hashset[format2]; !exist2 {
					return false
				}
			}
		}
	}
	return true
}
