package DS

import (
	"strconv"
	"strings"
)

func ConfusingNumber(num int) bool {
	s := strconv.Itoa(num)
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		d := string(s[i])
		if d != "8" && d != "1" && d != "9" && d != "6" && d != "0" {
			return false
		}

		if d == "6" {
			b.WriteString("9")
		} else if d == "9" {
			b.WriteString("6")
		} else {
			b.WriteString(d)
		}
	}
	return s != b.String()
}
