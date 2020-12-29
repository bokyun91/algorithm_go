package main

import "fmt"

func main() {
	var h, m int
	fmt.Scanf("%d %d", &h, &m)

	if m < 45 {
		if h == 0 {
			h = 23
		} else {
			h--
		}
		m += 60
	}
	m -= 45
	fmt.Printf("%d %d\n", h, m)
}
