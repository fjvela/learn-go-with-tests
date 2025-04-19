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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(countdownDuration)
}

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buffer, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
