package main

import (
	"fmt"
	"math"
)

//force,find all sub array
//O(n^2)
func maxSubArray(nums []int) int {
	maxArray := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			maxArray = max(maxArray, sum)
		}
	}
	return maxArray
}

//dynamic programming
func maxSubArrayD(nums []int) int {
	maxArray := math.MinInt32
	sum := 0
	for _, v := range nums {
		sum += v
		if sum < 0 {
			maxArray = max(maxArray, sum)
			sum = 0
			continue
		} else {
			maxArray = max(maxArray, sum)
		}
	}
	return maxArray
}

//Divide-and-Conquer
func maxSubArrayDAC(nums []int) int {
	return maxSubArrayDivideAndConquer(nums, 0, len(nums)-1, (len(nums)-1)/2)

}

func maxSubArrayDivideAndConquer(nums []int, left int, right int, center int) int {
	if left == right {
		return nums[left]
	}
	maxLeft := maxSubArrayDivideAndConquer(nums, left, center, (left+center)/2)
	maxRight := maxSubArrayDivideAndConquer(nums, center+1, right, (center+right+1)/2)

	//判断横跨中间值的区间
	//包含左边界
	maxCenter := nums[center] + nums[center+1]
	maxLeftBorder := nums[center]
	maxRightBorder := nums[center+1]
	sumLeft := 0
	sumRight := 0
	for i := center; i >= left; i-- {
		sumLeft += nums[i]
		if sumLeft > maxLeftBorder {
			maxLeftBorder = sumLeft
		}
	}
	//包含右边界
	for i := center + 1; i <= right; i++ {
		sumRight += nums[i]
		if sumRight > maxRightBorder {
			maxRightBorder = sumRight
		}
	}
	maxCenter = maxLeftBorder + maxRightBorder
	return maxThree(maxLeft, maxRight, maxCenter)
}

func maxThree(num1 int, num2 int, num3 int) int {
	return max(max(num1, num2), num3)
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArrayDAC(nums))
}
