package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	// 4가지 경우의 수
	if x > 0 && y > 0 {
		fmt.Println(1)
	} else if x < 0 && y > 0 {
		fmt.Println(2)
	} else if x < 0 && y < 0 {
		fmt.Println(3)
	} else if x > 0 && y < 0 {
		fmt.Println(4)
	}

}
