package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str :=  strconv.Itoa(x)
	for i := range str{
		if str[i] != str[len(str)-1-i]{
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool{
	if x<0{
		return false
	}
	y := x
	res ,pop:= 0, 0
	for ;y!=0;{
		pop = y%10
		y/=10
		res = res*10 +pop
	}
	if res == x{
		return true
	}
	return false
}

func main(){
	fmt.Println(isPalindrome(1301))
}