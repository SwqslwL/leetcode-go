package main

import (
	"fmt"

	//"reflect"
)

func lengthOfLongestSubstringForce(s string) int {
	length := 0
	for i:=0;i<len(s);i++{
		sub := make(map[string]int,1)
		lengthTmp := 0
		for j:=i;j<len(s);j++{
			sub[string(s[j])] = j

			if lengthTmp == len(sub){//发生重复
				break
			}else{
				lengthTmp += 1
			}
		}
		if lengthTmp > length{
			length = lengthTmp
		}
	}
	return length
}

func lengthOfLongestSubstringSlide(s string)int{
	i, j := 0,0
	length := 0
	sub := make(map[string]int,1)
	for ;i<len(s) && j<len(s);{
		if _,ok:=sub[string(s[j])];ok{
			//表示存在,删除i
			delete(sub,string(s[i]))
			i++
			continue
		}
		sub[string(s[j])] = j
		if length<len(sub){
			length = len(sub)
		}
		j++
	}
	return length
}

func lengthOfLongestSubstringSlide2(s string)int{
	i, j := 0,0
	length := 0
	sub := make(map[byte]int,1)

	for ;j<len(s);{
		if _,ok:=sub[s[j]];ok{
			//新元素存在表示存在,转移
			if i<sub[s[j]]+1{
				i = sub[s[j]]+1
			}
		}
		sub[s[j]] = j
		if length<(j-i+1){
			length = j-i+1
		}
		j++

	}
	return length
}


func main(){
	str := string("asdads")
	length := lengthOfLongestSubstringSlide2(str)
	fmt.Println(length)
}
