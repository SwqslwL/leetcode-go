package main

import (
	"fmt"
)

//just the validity

func isValid(s string) bool {
	dict := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := make([]string, 0)
	top := -1
	for _, v := range s {

		if _, ok := dict[string(v)]; top != -1 && ok && dict[string(v)] == stack[top] {
			stack = stack[:top]
			top -= 1
		} else {
			stack = append(stack, string(v))
			top += 1
		}
	}
	if top < 0 {
		return true
	} else {
		return false
	}
}

func main() {
	parentheses := "([])"
	fmt.Println(isValid(parentheses))
}
