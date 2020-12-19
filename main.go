package main

import (
	"fmt"
	"io"
	"os"
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

	// for i := countDownStart; i > 0; i-- {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Fprintln(out, i)
	// }
	// time.Sleep(1 * time.Second)
	// fmt.Fprint(out, finalWord)

	// defer
	// 関数の遅延実行
	// 関数の終了時に実行される。
	// 引数の評価はdefer呼び出し時。
	// スタック形式で実行（最後に呼び出しされたものが、最初に実行される。）

	msg := "!!!"
	defer fmt.Println(msg)
	msg = "World"
	defer fmt.Println(msg)
	fmt.Println("Hello")

}
