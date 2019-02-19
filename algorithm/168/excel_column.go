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
	for n != 0 {
		weight := 1

		head := 1
		tag := false
		for ;head<=26;head++{
			if head
		}
		for{
			if weight*26 >= n{
				break
			}
			weight *= 26
		}
		alp, tmp := getAlp(n, weight)
		fmt.Println(weight, alp, tmp)
		res += alp
		n -= tmp
	}
	return res
}

func getAlp(n int, weight int) (alp string, tmp int){
	for i:=1 ;i <=26; i++{
		if i*weight > n{
			return  string(byte(i-1)+64),(i-1)*weight
		}
	}
	return
}

func main(){
	fmt.Println(convertToTitle(702))
}