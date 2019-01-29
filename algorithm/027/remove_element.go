package main

import "fmt"

//method one
//time O(n) n
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := range nums {
		if nums[fast] != val{
			nums[slow] = nums[fast]
			slow += 1
		}
	}
	return slow
}

//when elements to remove are rare
func removeElement2(nums []int, val int)int{
	n := len(nums)
	i:=0
	for ;i<n;{
		if nums[i] == val{
			nums[i] = nums[n-1]
			n -= 1
			continue
		}else{
			i += 1
		}
	}
	return i
}


func main(){
	nums := []int{1,3,3,4,5}
	fmt.Println(nums[:removeElement(nums, 2)])

}
