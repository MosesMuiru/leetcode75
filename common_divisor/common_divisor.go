package main

import "fmt"

func CommonDivisor(str1, str2 string) string {
	var new_string string
	dividable := false
	largest_str, smallest := str1, str2
	if len(str2) > len(str1) {
		largest_str = str2
		smallest = str1
	}

	for i := 0; i < len(largest_str); i++ {
		// the smallest string
		if i < len(smallest) && string(largest_str[i]) == string(smallest[i]) {
			dividable = true
			continue
		}

		if dividable {
			new_string += string(largest_str[i])
			continue
		}
		new_string = ""
	}

	if new_string != "" {
		new_string = Divided(largest_str, new_string)

	}
	return new_string
}

func Divisor(str1, str2 string) string {
	fmt.Println("This is working")
	var new_string string
	dividable := false
	largest_str, smallest := str1, str2
	if len(str2) > len(str1) {
		largest_str = str2
		smallest = str1
	}

	finish := len(smallest)

	for i := 0; i < len(smallest); i++ {
		finish += len(smallest)

		for j := 0; j < len(largest_str); j++ {
			if largest_str[:j] == smallest[:i] {
				fmt.Println("This is the largest value----->", largest_str[:j])
				fmt.Println("This is smallest value", smallest[:i])

				dividable = true
				continue
			}
			dividable = false
		}
	}

	fmt.Println("This is dividabel ---->", dividable)

	return new_string

}

func Divided(largest, checker string) string {
	new_string := checker
	dividable := true
	finish := len(checker)
	fmt.Println("The lenght of the  finish", finish)

	for i := 0; i <= len(largest); i += finish {
		fmt.Println("this si hte iiiii_", i)
		if i+finish <= len(largest) && checker != largest[i:i+finish] {
			fmt.Println("this si hte checker---->", checker)
			fmt.Println("the largest valueli --", largest[i:i+finish])
			dividable = false
			break
		}
	}

	if dividable == false {
		new_string = ""
	}

	fmt.Println("new sing", dividable)

	return new_string
}

func Largest(str1, str2 string) (largest, smallest string) {
	if len(str2) > len(str1) {
		return str2, str1
	}
	return str1, str2

}
