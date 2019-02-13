package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}

	}
	return []int{}
}

func twoSumHash(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i, v := range nums {
		hmap[v] = i
	}
	for value, key := range hmap {
		result := target - value
		if _, ok := hmap[result]; ok && hmap[result] != key {
			if hmap[result] < key {
				return []int{hmap[result], key}
			} else {
				return []int{key, hmap[result]}
			}
		}
	}
	return []int{}
}

func twoSumHash2(nums []int, target int) []int {
	hmap := make(map[int]int)
	for key, value := range nums {
		result := target - value
		if _, ok := hmap[result]; ok {
			return []int{hmap[result], key}
		} else {
			hmap[value] = key
		}
	}
	return []int{}
}

func main() {
	nums := []int{3, 3, 5}
	//	fmt.Println(nums)
	result := []int{}
	//fmt.Println(result)
	target := 8
	result = twoSumHash2(nums, target)
	fmt.Println(result)
}
