package main

import "fmt"

func main() {
	list := []int{}
	l := removeDuplicates(list)
	fmt.Println(list[:l])
}

//in-place
//two pointer
func removeDuplicates(nums []int) int {
	flag := 0
	for i := 1; i < len(nums); i++ {
		if nums[flag] != nums[i] {
			flag += 1
			nums[flag] = nums[i]
		}
	}
	return flag + 1
}
