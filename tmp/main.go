package main

import (
	"fmt"
	"time"
)

func main() {
	num := time.Now()
	if num.Unix()%5 != 1 {
		fmt.Println(num)
	}
	fmt.Println(num.Unix())
}
