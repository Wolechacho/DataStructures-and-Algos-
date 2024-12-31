package main

import "fmt"

type Trie struct {
	leaf     bool
	children []*Trie
}

func IndexPairs(text string, words []string) [][]int {
	root := &Trie{children: make([]*Trie, 26)}

	node := root
	for _, w := range words {
		node = root
		for i := 0; i < len(w); i++ {
			index := w[i] - 'a'
			if node.children[index] == nil {
				node.children[index] = &Trie{children: make([]*Trie, 26)}
			}
			node = node.children[index]
		}
		node.leaf = true
	}

	result := [][]int{}
	for i := 0; i < len(text); i++ {
		node = root
		fmt.Println("node", node)

		for j := i; j < len(text); j++ {
			index := text[j] - 'a'
			if node.children[index] == nil {
				break
			}
			node = node.children[index]
			if node.leaf {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
