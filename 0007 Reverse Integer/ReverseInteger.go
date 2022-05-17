package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

var max = math.MaxInt32 / 10
var min = math.MinInt32 / 10

func reverse(x int) int {
	ret := 0

	for x != 0 {
		digit := x % 10

		if ret > max || (ret == max && digit > 7) {
			return 0
		}
		if ret < min || (ret == min && digit < -8) {
			return 0
		}

		ret = ret*10 + digit
		x /= 10
	}
	return ret
}
