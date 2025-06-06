package main

import "fmt"

func MergeStrings(word1 string, word2 string) string {
	fmt.Println("workd1", string(word1[1])+string(word2[1]))
	large_world, small_word := word1, word2

	var new_word string

	if len(word1) < len(word2) {
		large_world = word2
		small_word = word1
	}

	fmt.Println("large workld 000> ", large_world)
	fmt.Println("small workld 000> ", small_word)

	for i := 0; i < len(large_world); i++ {
		if i > len(small_word)-1 {
			new_word += string(large_world[i])
			continue
		}
		new_word += string(word1[i]) + string(word2[i])
		fmt.Println("new world in thte loop--->", new_word)
	}

	fmt.Println("new world --->", new_word)
	return new_word

}
