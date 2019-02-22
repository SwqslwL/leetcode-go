package main

import "fmt"

/*
27/26 =1 ... 1
1/26 = 0 ...1

----------------
701/26 = 26...25
26/26 = 1...0
ZY
----------------
702/26 = 27...0
27/26 = 1...1

----------------
24/27 = 0...24
W
----------------
*/


func convertToTitle(n int) string {
	res := ""
	for n != 0{
		n -= 1
		remain := n%26
		res = string(byte(remain+65))+res
		n = n/26
	}
	return res
}

func main(){
	fmt.Println(convertToTitle(28))
}