package main

import (
	"fmt"
	"strings"
)

type found_vowels_position struct {
	vowel    rune
	position int
}

func Reverse(word string) string {
	// split the word
	upper := []rune(strings.ToUpper("aeiou"))

	vowels := []rune("aeiou")
	fmt.Println("Vowels", vowels)
	//	newWord := strings.Split(word, ",")
	// convert the world to rune
	new_world := []rune(word)
	fmt.Println("old--- ", new_world)

	foundVowels := []found_vowels_position{}

	for i := range new_world {
		// check if the i is a vowel
		for j := range vowels {
			if new_world[i] == vowels[j] || new_world[i] == upper[j] {
				// i have to push the vowel that i found to found vowels
				fmt.Printf("comparing this %v and this --- %v\n", new_world[i], vowels[j])

				foundVowels = append(foundVowels, found_vowels_position{
					vowel:    new_world[i], // its rune
					position: i,
				})
			}

		}
	}

	//fmt.Println("old", foundVowels)
	n := len(foundVowels)
	for i := 0; i < n/2; i++ {
		foundVowels[i].vowel, foundVowels[n-1-i].vowel = foundVowels[n-1-i].vowel, foundVowels[i].vowel
	}

	//fmt.Println("new ", foundVowels)

	// new world should be rearranged
	for _, value := range foundVowels {
		//
		//	fmt.Println("rinnnnnnn")
		//	fmt.Println("position-- ", new_world[value.position])
		new_world[value.position] = value.vowel
	}

	//	fmt.Println("neww-- ", new_world)

	//	fmt.Println(new_world)
	//	fmt.Println("before", word)
	return string(new_world)
}
