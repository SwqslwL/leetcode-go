package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	way1 := climbStairs(n - 1)
	way2 := climbStairs(n - 2)
	return way1 + way2
}

//dynamic programming
func climbStairsD(n int) int {
	var a, b = 1, 2
	if n < 3 {
		return n
	}
	var t int
	for i := 3; i <= n; i++ {
		t = a + b
		a = b
		b = t
	}
	return t
}

func main() {
	fmt.Println(climbStairsD(5))
}
