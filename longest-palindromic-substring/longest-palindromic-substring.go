/*
 * 最长回文子串
 */
package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	if s[0] == s[len(s) - 1] {
		return isPalindrome(s[1: len(s) - 1])
	} else {
		return false
	}
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	if isPalindrome(s) {
		return s
	}
	if len(longestPalindrome(s[1:])) > len(longestPalindrome(s[: len(s) - 1])) {
		return longestPalindrome(s[1:])
	} else {
		return longestPalindrome(s[: len(s) - 1])
	}
}

func main()  {
	fmt.Println(longestPalindrome("abcaba"))
}
