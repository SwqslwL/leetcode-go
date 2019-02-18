package  main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i:=0; i<len(numbers)&&numbers[i]<=target;i++ {
		t := target - numbers[i]
		for j:=i+1; j<len(numbers)&&numbers[j]<=t; j++ {
			if numbers[j] == t{
				return []int{i+1,j+1}
			}
		}

	}
	return []int{}
}

//双指针
func twoSum2(numbers []int, target int) []int {
	i := 0
	j := len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i+1, j+1}
		}else if sum < target{
			i += 1
		}else{
			j -= 1
		}
	}
	return []int{}
}


func main(){
	numbers := []int{-1,0}
	fmt.Println(twoSum2(numbers, -1))
}