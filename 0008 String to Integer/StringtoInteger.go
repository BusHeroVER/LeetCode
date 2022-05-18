package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("-2147483648")) // -2147483648
	fmt.Println(myAtoi("-2147483649")) // -2147483648
	fmt.Println(myAtoi("2147483647"))  // 2147483647
	fmt.Println(myAtoi("2147483648"))  // 2147483647

	fmt.Println(myAtoi(" 4193 with words"))   // 4193
	fmt.Println(myAtoi("11 4193 with words")) // 11
}

func myAtoi(s string) int {
	ret := 0
	isReading := false
	isPostive := true

	for _, c := range s {
		if c == ' ' {
			if isReading {
				return ret
			}
		} else if c == '+' {
			if isReading {
				return ret
			} else {
				isReading = true
			}
		} else if c == '-' {
			if isReading {
				return ret
			} else {
				isReading = true
				isPostive = false
			}
		} else if '0' <= c && c <= '9' {
			isReading = true

			ret *= 10
			if isPostive {
				ret += int(c - '0')
			} else {
				ret -= int(c - '0')
			}

			if ret < math.MinInt32 {
				return math.MinInt32
			} else if math.MaxInt32 < ret {
				return math.MaxInt32
			}
		} else {
			return ret
		}
	}
	return ret
}
