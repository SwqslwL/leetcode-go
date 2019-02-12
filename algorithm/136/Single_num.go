package main

//a^b^c^a^b = a^a^b^b^c = 0^0^c = c
func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}

func main(){
	nums := []int{5,2,2,1,5}
	singleNumber(nums)
}