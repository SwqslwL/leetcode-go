package main

import "fmt"

func addBinary(a string, b string) string {
	carry := "0"
	c := ""
	l1 := len(a) - 1
	l2 := len(b) - 1
	for l1 >= 0 || l2 >= 0 {
		num := 0
		if getStr(a, l1) == "1" {
			num += 1
		}
		if getStr(b, l2) == "1" {
			num += 1
		}
		if carry == "1" {
			num += 1
		}
		switch num {
		case 0:
			{
				c = "0" + c
				carry = "0"
			}
		case 1:
			{
				c = "1" + c
				carry = "0"
			}
		case 2:
			{
				c = "0" + c
				carry = "1"
			}
		case 3:
			{
				c = "1" + c
				carry = "1"
			}
		}
		l1 -= 1
		l2 -= 1
	}
	if carry == "1" {
		c = "1" + c
	}
	return c
}

func getStr(s string, index int) string {
	if index < 0 {
		return "0"
	}
	return string(s[index])
}

func maxLength(a string, b string) int {
	if len(a) >= len(b) {
		return len(a)
	} else {
		return len(b)
	}
}
func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}
