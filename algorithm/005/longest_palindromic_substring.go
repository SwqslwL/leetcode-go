package main

import "fmt"

func longestPalindrome(s string) string {
	//find all palindrome
	//二次优化后
	str := ""
	ROOT:for l:=len(s);l>=1;l--{
		for i,j:=0,l;j<=len(s);{
			if isPalindrome(s[i:j]){
				str = s[i:j]
				break ROOT
			}
			i += 1
			j += 1
		}
	}
	return str
}

//TODO:zhen luo suo
// dynamic programming
func longestPalindrome2(s string)string{
	//find strings where length is odd or even
	odd, even := make([]int,0), make([]int,0)
	//initiation 1
	for i := range s{
		odd = append(odd, i)
	}
	//initiation 2
	for _,v := range odd{
		if v+1<len(s) && s[v]==s[v+1]{
			even = append(even, v)
		}
	}
	fmt.Println(odd, even)
	s1:=getLongestPalindromeOdd(s,1, odd)
	s2:=getLongestPalindromeEven(s,2, even)

	if len(s1)>=len(s2) {
		return s1
	}else{
		return s2
	}

}

func getLongestPalindromeOdd(s string, length int,list []int)string{
	str := ""
	if len(list)!=0{
		str = string(s[list[0]])
	}
	var hlength int
	var tmp []int
	for ;len(list)!=0;{
		hlength = length/2
		tmp = []int{}
		for _,v :=range list{
			if v-hlength-1>=0 && v+hlength+1<len(s) && s[v-hlength-1]==s[v+hlength+1]{
				tmp = append(tmp, v)
			}
		}
		if len(tmp) != 0 {
			length += 2
			str = s[tmp[0]-hlength-1:tmp[0]+hlength+1+1]
		}
		list = tmp
	}
	fmt.Println("***",length,str)
	return  str
}

func getLongestPalindromeEven(s string, length int,list []int)string{
	str := ""
	if len(list)!=0{
		str = string(s[list[0]:list[0]+1+1])
	}
	var hlength int
	var tmp []int
	for ;len(list)!=0;{
		hlength = length/2
		tmp = []int{}
		for _,v :=range list{
			if v-hlength>=0 && v+hlength+1<len(s) && s[v-hlength]==s[v+hlength+1]{
				tmp = append(tmp, v)
			}
		}
		if len(tmp) != 0 {
			length += 2
			str = s[tmp[0]-hlength:tmp[0]+hlength+1+1]
		}
		list = tmp
	}
	fmt.Println("***",length, str)
	return str
}


func isPalindrome(s string) bool{
	for elem:=0;elem<=len(s)/2-1;elem++{
		if s[elem]==s[len(s)-elem-1] {
			continue
		}
		return false
	}
	return true
}

func main(){
	str := "cbbd"
	res := longestPalindrome2(str)
	fmt.Println(res)
}