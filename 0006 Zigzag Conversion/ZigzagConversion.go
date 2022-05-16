package main

import (
	"bytes"
)

func main() {
	convert("PAYPALISHIRING", 1)
}

func convert(s string, numRows int) string {
	len := len(s)
	if len <= numRows || numRows < 2 {
		return s
	}

	var ret bytes.Buffer
	for r := 0; r < numRows; r++ {
		index := r
		offset0 := 2 * (numRows - r - 1)
		offset1 := 2 * r

		ret.WriteString(string(s[index]))
		for index < len {
			if offset0 > 0 {
				index += offset0
				if index < len {
					ret.WriteString(string(s[index]))
				}
			}
			if offset1 > 0 {
				index += offset1
				if index < len {
					ret.WriteString(string(s[index]))
				}
			}
		}
	}
	return ret.String()
}
