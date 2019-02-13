package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums)
}

func searchInsertB(nums []int, target int) int {
	//二分法
	low, high := 0, len(nums)-1
	var mid int
	if target < nums[0] {
		return 0
	}
	if target > nums[high] {
		return high + 1
	}
	for low < high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func searchInsertBOP(nums []int, target int) int {
	//optimization
	low, high := 0, len(nums)
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	nums := []int{1}

	fmt.Println(searchInsertB(nums, 5))

}
