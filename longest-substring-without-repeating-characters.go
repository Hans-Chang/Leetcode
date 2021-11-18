package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	longestSubstringLen := 0
	charMapPosition := make(map[byte]int)
	left := 0
	right := 0
	for i := 0; i < len(s); i += 1 {

		_, ok := charMapPosition[s[i]]
		if !ok {
			charMapPosition[s[i]] = i
			right += 1
			if right - left > longestSubstringLen {
				longestSubstringLen = right - left
			}
		} else {
			for j := left; j < charMapPosition[s[i]]; j += 1 {
				delete(charMapPosition, s[j])
			}
			left = charMapPosition[s[i]] + 1
			right += 1
			charMapPosition[s[i]] = i
		}
	}
	return longestSubstringLen
}

func main()  {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}


