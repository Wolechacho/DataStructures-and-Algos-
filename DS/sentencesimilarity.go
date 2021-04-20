package DS

import (
	"fmt"
)

func Sentencesimilarity(sentence1, sentence2 []string, pairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	hashMap := map[string]string{}

	for _, s := range pairs {
		if _, found := hashMap[s[0]]; !found {
			hashMap[s[0]] = s[1]
		}

		if _, found := hashMap[s[1]]; !found {
			hashMap[s[1]] = s[0]
		}
	}

	for i := 0; i < len(sentence1); i++ {
		w1 := sentence1[i]
		w2 := sentence2[i]
		fmt.Println(w1)
		fmt.Println(w2)
		if w1 == w2 {
			continue
		}
		if val, found := hashMap[w1]; found {
			if val == w2 {
				continue
			}
		}

		if val, found := hashMap[w2]; found {
			if val == w1 {
				continue
			}
		} else {
			return false
		}
	}
	fmt.Println(hashMap)
	return true

}
