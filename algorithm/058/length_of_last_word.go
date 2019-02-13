package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	str := strings.Split(s, " ")
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == "" {
			continue
		} else {
			return len(str[i])
		}
	}
	return 0
}

func lengthOfLastWord2(s string) int {
	length := 0
	f := false
	for i := len(s) - 1; i >= 0; i-- {
		if f && string(s[i]) == " " {
			break
		}
		if string(s[i]) != " " {
			length += 1
			f = true
		}
	}
	return length
}
func main() {
	s := "sad asdasd avvvvvvv    "
	fmt.Println(lengthOfLastWord2(s))
}
