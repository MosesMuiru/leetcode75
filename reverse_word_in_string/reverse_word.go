package main

import (
	"fmt"
	"slices"
	"strings"
)

func ReverseWord(word string) string {
	// spilt the world

	word = strings.Trim(word, " ")
	new_word := strings.Split(word, " ")
	word_list := []string{}

	fmt.Println("before", len(new_word))
	// remvoe any wodls
	for _, value := range new_word {
		if value == "" {
			continue
		}
		fmt.Println("this is trimmeddd", value)
		word_list = append(word_list, strings.TrimSpace(value))
	}

	// reverse the orde

	slices.Reverse(word_list)
	fmt.Println(word_list)
	new := strings.Join(word_list, " ")

	return new
}
