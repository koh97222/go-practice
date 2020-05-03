package main

import (    
    "fmt"
    // 階層が1段階下だと、自分でimport書く必要がある？
    "os/user"
    "time"
    // "strconv"
)
// 定数はグローバルに。他のGoファイルから呼び出されるときは、大文字。
const Pi = 3.14

const (
    // 型を宣言していない。
    Username = "test_user"
    Password = "test_pass"
)

// 変数定義　関数外でも宣言可能。
var (
    i int           = 1
    f64 float64     = 1.2
    f32 float32     = 0.2
    // ※complex64　複素数型
    c64 complex64   = -5 + 12i
    s string = "test"
    t,f bool = true,false
)
func foo (){
    // 短縮形　この書き方は関数内でしかできない。
    xi := 1
    // 明示的に型を使う場合はこっち。
    var xf32 float32 = 1.2
    var (
    u8 uint8        = 255
    i8 int8         = 127
    )


    // %Tで変数の型を調べることができる。%vは値。
    fmt.Printf("%T\n", xf32)
    fmt.Printf("%T\n", xi)
    fmt.Printf("type=%T value=%v",u8,u8)
    fmt.Println(i8)

}
func main() {
    // foo()
    //x := 0+3
    fmt.Println(string("Hello world"[0])) // H
    fmt.Println(i,f64,s,t,f)
    fmt.Println(Pi,Username,Password)
    fmt.Println("Hello, World!", time.Now())
    fmt.Println(user.Current())
    // 型変換
    // var x int = 1
    // xx := float64(x)
    // fmt.Printf("%T %v %f\n", xx,xx,xx)

    // var y float64 = 1.2
    // yy := int(y)
    // fmt.Printf("%T %v %d\n", yy,yy,yy)

    // var s string = "14"
    // i, _ := strconv.Atoi(s)
    // fmt.Printf("%T %v\n", i,i)

    // 可変長の配列をスライスという？
    // n := []int {1,2,3,4,5}
    // fmt.Println(n)          //[1 2 3 4 5]
    // fmt.Println(n[2])       //3
    // fmt.Println(n[2:4])     //[3 4]
    // fmt.Println(n[:2])      //[1 2]
    // fmt.Println(n[2:])      //[3 4 5]
    // fmt.Println(n[:])       //[1 2 3 4 5]

    // n[2] = 100
    // fmt.Println(n)          //[1 2 100 4 5]

    // var board = [][]int{
    //     []int{0,1,2},
    //     []int{3,4,5},
    //     []int{6,7,8},
    // }

    // fmt.Println(board)      //[[0 1 2] [3 4 5] [6 7 8]]

    // n = append(n,100,200,300,400)
    // fmt.Println(n)          //[1 2 100 4 5 100 200 300 400]

    n := make([]int, 3,5)
    // len=3 cap=5 value=[0 0 0]
    fmt.Printf("len=%d cap=%d value=%v\n" , len(n),cap(n),n)
    n = append(n,0,0)
    fmt.Printf("len=%d cap=%d value=%v\n" , len(n),cap(n),n)

    var c [] int
    c = make([]int, 0,5)
    for i := 0; i<5; i++ {
        c = append(c,i)
        fmt.Println(c)
    }
    fmt.Println(c)

}