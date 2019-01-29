package main

import "fmt"

func maxProfit(profit []int) int {
	buy, sell := profit[0], profit[0]
	for _,v := range profit {
		if buy > v{
			buy = v
			sell = v
		}
		if sell < v {
			sell = v
		}
	}
	pro := sell - buy
	if pro < 0 {
		return 0
	}
	return sell - buy
}

func main(){
	profit := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(profit))
}