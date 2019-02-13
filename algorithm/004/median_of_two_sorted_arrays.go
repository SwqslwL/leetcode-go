package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1) + len(nums2)
	odd := true
	if (m)%2 == 0 {
		t := m / 2

	} else {
		t := math.Floor(float64(m))
		odd = false
	}
	if odd {
		fmt.Println(t)
	}
}

func main() {
	nums1 := []int{1, 2, 4, 2, 3, 473, 3, 42, 5, 74}
	nums2 := []int{12}
	fmt.Println(nums1, nums2)

}
