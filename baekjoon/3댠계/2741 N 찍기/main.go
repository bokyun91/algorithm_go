package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var t int
	fmt.Fscanf(r, "%d", &t)

	for i := 1; i <= t; i++ {
		fmt.Fprintln(w, i)
	}
}
