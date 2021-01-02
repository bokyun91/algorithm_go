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

	for i := t; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
}
