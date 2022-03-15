/*
 * 17. 电话号码的字母组合
 */

package main

import (
	"fmt"
	"strings"
)

var numToStrMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

var result []string

func letterCombinations(digits string) []string {
	var track string
	backtrack(digits, &track, digits)
	return result
}

func backtrack(restDigits string, track *string, digits string) {
	if len(restDigits) == 0 {
		result = append(result, *track)
		return
	}
	for i := 0; i < len(digits); i += 1 {
		var choice int
		for j := 0; j < len(numToStrMap[restDigits[i]]); j += 1 {
			if strings.Contains(*track, string(numToStrMap[restDigits[0]][j])) {
				continue
			}
			choice = j
			break
		}
		var tmp = *track + string(numToStrMap[restDigits[i]][choice])
		track = &tmp
		backtrack(restDigits[1:], track, digits)
		tmp = (*track)[:len(*track)-1]
		track = &tmp
	}
	return
}

func main() {
	fmt.Println(letterCombinations("23"))
}
