package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(buffer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(buffer, i)
	}
	fmt.Fprint(buffer, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
