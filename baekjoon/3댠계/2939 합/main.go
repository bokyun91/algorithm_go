package main

import "fmt"

func main() {
	var num, sum int

	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		sum += i
	}

	fmt.Println(sum)
}
