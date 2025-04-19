package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const countdownDuration = 1 * time.Second

func Countdown(buffer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		time.Sleep(countdownDuration)
	}
	fmt.Fprint(buffer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
