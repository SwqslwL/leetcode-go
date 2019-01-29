package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	y := abs(x)
	var res int64 = 0
	var pop int64 = 0
	for ;y!=0;{
		res *= 10
		pop = int64(y)%10
		res += pop
		y /= 10
	}
	fmt.Println(x)
	if x>0&&res<=math.MaxInt32 {
		return int(res)
	} else if x<0&&-res>=math.MinInt32{
		return -int(res)
	}else{
		return 0
	}
	return int(res)
}

func abs(num int)(n int){
	return (num^num>>31)-num>>31
}

func main(){
	a := reverse(
		-2147483648)
	fmt.Println(a)
}
