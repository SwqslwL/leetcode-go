package main

import (
	"fmt"
	"math"
)

type MinStack struct{
	top int
	min int
	data []int
}

func Constructor() MinStack {
	return MinStack{top : -1, min : math.MaxInt32}
}

func (this *MinStack) Push(x int) {
	this.top += 1
	this.data = append(this.data, x)
	if this.top == 0{
		this.min = x
	}else{
		if this.min > x{
			this.min = x
		}
	}
}

func (this *MinStack) Pop() {
	if this.top < 0{
		return
	}
	if this.data[this.top] == this.min{
		this.min = math.MaxInt32
		for i:=0 ; i< this.top; i++ {
			if this.min > this.data[i]{
				this.min = this.data[i]
			}
		}
	}
	this.data = this.data[:this.top]
	this.top -= 1
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) Top() int {
	return this.data[this.top]
}

func main() {
	obj := Constructor()
	obj.Push(21)
	obj.Push(1)
	obj.Pop()
	fmt.Println(obj.min)
}