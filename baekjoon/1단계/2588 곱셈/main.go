package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	// 자릿수 구하기(b)
	c := strconv.Itoa(b)
	d := b

	for i := 1; i <= len(c); i++ {
		target := d % 10
		fmt.Println(a * target)
		d /= 10
	}

	fmt.Println(a * b)
}
