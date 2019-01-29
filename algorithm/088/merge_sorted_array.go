package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i := m+n-1;i >= 0;i--{
		if n - 1 < 0{
			nums1[i] = nums1[m-1]
			m--
			continue
		}
		if m - 1 < 0 {
			nums1[i] = nums2[n-1]
			n--
			continue
		}
		if nums1[m-1] >= nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
			continue
		}
		if  nums1[m-1] < nums2[n-1]{
			nums1[i] = nums2[n-1]
			n--
			continue
		}
	}
}


func main(){
	nums1 := []int{0}
	nums2 := []int{1}
	m, n := 0, 1
	merge(nums1, m, nums2, n)
}
