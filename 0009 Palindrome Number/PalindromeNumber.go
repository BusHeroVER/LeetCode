package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1234567899))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(131))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	org := x
	var reverse int64 = 0
	for x != 0 {
		reverse = reverse*10 + int64(x%10)
		x /= 10
	}
	return org == int(reverse)
}
