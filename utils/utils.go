package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := 0; index < 10; index++ {
		fmt.Println(r.Intn(10))
	}
	return r.Intn(100)
}

func main() {
	fmt.Println(GetRandNum())
}
