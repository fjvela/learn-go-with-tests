package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(buffer io.Writer) {
	fmt.Fprint(buffer, "3")
}

func main() {
	Countdown(os.Stdout)
}
