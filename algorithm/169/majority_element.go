package main

import "fmt"

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	maxNum := 0
	for i := 0;i < len(nums);i++ {
		if _,ok:=dict[nums[i]];!ok{
			dict[nums[i]] = 1
			if i == 0{
				maxNum = nums[i]
			}
		}else{
			dict[nums[i]]  += 1
			if dict[nums[i]] > dict[maxNum]{
				maxNum = nums[i]
			}
		}
	}
	return maxNum
}

func majorityElement2(nums []int) int {
	dict := make(map[int]int)
	l := len(nums)/2
	for _, v := range nums {
		dict[v] += 1
		if dict[v] >= l{
			return v
		}
	}
	return 0
}


func main() {
	nums := []int{3,2,3}
	fmt.Println(majorityElement2(nums))
}