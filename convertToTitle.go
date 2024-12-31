package main

import (
	"fmt"
	"strings"
)

func convertToTitle(n int) string {
	if n <= 26 {
		return "" + string(n+64)
	}
	var sb strings.Builder
	for n > 0 {
		n--
		fmt.Println("n", n)
		c := (n % 26) + 65
		fmt.Println("c", c)

		sb.WriteString("" + string(c))
		n = n / 26
		fmt.Println("n2", n)
	}

	return ""
}
