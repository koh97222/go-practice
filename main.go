package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

// Hello ...
func Hello() string {
	return "Hello,World"
}

func main() {
	Countdown(os.Stdout)
}

// Countdown Countdown
func Countdown(out io.Writer) {

	for i := countDownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)

}
