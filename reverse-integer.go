/*
 * 整数反转
 */
package main

import "fmt"
import "math"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	reverseX := 0
	for x > 0 {
		reverseX = 10 * reverseX + x % 10
		x = int(x / 10)
	}
	reverseX = reverseX * flag
	if reverseX < int(math.Pow(-2, 31)) {
		return 0
	}
	if reverseX > int(math.Pow(2, 31) - 1) {
		return 0
	}
	return reverseX
}

func main()  {
	fmt.Println(reverse(123))
}
