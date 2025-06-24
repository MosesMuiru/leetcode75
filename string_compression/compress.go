package main

import (
	"fmt"
	"strconv"
)

func Compress(string_list []string) int {
	var repeat int

	for index, value := range string_list {

		if len(string_list) == 1 {
			break
		} else if index+1 < len(string_list) && value == string_list[index+1] {

			repeat++
			string_list[index] = value
			string_list[index+1] = strconv.Itoa(repeat)
		}

	}
	fmt.Println(string_list)

	return len(string_list)
}

func CompressString(chars []string) int {
	// count the repeating number,
	// check how many times that number appears
	// return the number length
	read := 0
	write := 0

	// read
	for read < len(chars) {
		currentChar := chars[read]
		count := 0

		// count the occurrence of the current character
		for read < len(chars) && chars[read] == currentChar {
			read++  // i have read that character
			count++ // i have counted the apprearance of that char

		}
		// write the char

		chars[write] = currentChar
		write++ // writing

		if count > 1 {
			for _, digit := range []byte(fmt.Sprintf("%d", count)) {
				chars[write] = string(digit)
				write++
			}
		}
	}
	fmt.Println(chars)
	return len(chars)
}

func main() {
	chars := []string{"a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}
	fmt.Println(Compress(chars))
}
