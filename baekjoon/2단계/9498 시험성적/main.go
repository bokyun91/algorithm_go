package main

import "fmt"

func main() {
	var a int

	fmt.Scanf("%d", &a)

	switch {
	case a >= 90:
		fmt.Println("A")
		break
	case a >= 80:
		fmt.Println("B")
		break
	case a >= 70:
		fmt.Println("C")
		break
	case a >= 60:
		fmt.Println("D")
		break
	default:
		fmt.Println("F")
		break
	}
}
