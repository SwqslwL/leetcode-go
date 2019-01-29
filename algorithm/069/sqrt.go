package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	for i := 0;i<=x;i++{
		if(i+1)*(i+1) > x{
			return i
		}
	}
	return 0
}

func mySqrtNewton(x int) int {
	t := float64(x)
	flag := 0.1
	oldT := 0.0
	if x == 0 {
		return 0
	}
	for ;math.Abs(t-oldT)>flag;{
		fmt.Println("xxx")
		oldT = t
		t = (t + float64(x)/t)/2
	}

	return int(t)
}

func mySqrtDivide(x int) int {
	if x == 0{
		return 0
	}
	left := 0
	right := x
	for ;left<right;{
		mid := (left + right)/2
		fmt.Println("XXX")
		if mid * mid == x{
			return mid
		}
		if mid * mid < x{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return left-1
}
func mySqrtDivide2(x int) int {
	left := 0
	right := x
	mid := x/2
	for ;;{
		fmt.Println("xxx")
		if mid * mid == x{
			return mid
		}
		if (mid - 1) * (mid - 1) <= x && mid * mid >= x{
			return mid - 1
		}
		if mid * mid < x{
			left = mid + 1
		}else{
			right = mid - 1
		}

		mid = (left + right)/2
	}
}

func main(){
	fmt.Println(mySqrtNewton(1213124123123))
}
