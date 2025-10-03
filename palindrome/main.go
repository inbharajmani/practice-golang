package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	if isPalindrome("maa") {
		fmt.Println("Yes, it is palindrome")
	} else {
		fmt.Println("No, it is not palindrome")
	}
}