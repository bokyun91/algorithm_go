package main

import "fmt"

func main() {
	var dan int
	fmt.Scanf("%d", &dan)

	for i := 1; i <= 9; i++ {
		fmt.Println(dan, "*", i, "=", dan*i)
	}
}
