// Find the First Non-Repeating Character in a String
// Problem Statement:

// Given a string s, return the first character that does not repeat (i.e., appears only once). If no such character exists, return "_".
// Input: "golang"
// Output: "g"

// Input: "swiss"
// Output: "w"

// Input: "aabbcc"
// Output: "_"

package main

import "fmt"

func main() {

	str := "golang" // swiss
	fmt.Println(str)

	charCount := make(map[rune]int)

	for _, ch := range str {
		charCount[ch]++
	}

	result := "_"

	for _, ch := range str {
		if charCount[ch] == 1 {
			result = string(ch)
			break
		}
	}

	fmt.Println("First non repeating character :", result)

}
