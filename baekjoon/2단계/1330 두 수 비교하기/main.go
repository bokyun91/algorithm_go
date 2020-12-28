package main

import "fmt"

func main() {
	var a, b int

	// 두 정수 입력 받기
	fmt.Scanf("%d %d", &a, &b)

	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}
