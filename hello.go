//---------------------------------------------
//Goの並行処理の基本をサクッとまとめてみた(https://zenn.dev/y_yuita/articles/de09b33dad9bfb)
//---------------------------------------------


// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, World!");
// }


// package main

// import "fmt"

// func printNumbers(done chan struct{}) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// 	done <- struct{}{}
// }

// func main() {
// 	done := make(chan struct{})
// 	go printNumbers(done)
// 	<-done
// 	fmt.Println(("Done"))
// 	close(done)
// }

// package main

// import "fmt"

// func main() {
//     set := make(map[string]struct{}) // キーのみを持つセット

//     fmt.Println(set)
    
//     set["apple"] = struct{}{}
//     set["banana"] = struct{}{}
//     set["apple"] = struct{}{}

// 	fmt.Println(set)

//     // 存在確認
//     if _, exists := set["apple"]; exists {
//         fmt.Println("apple exists in the set")
//     }
// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     ch := make(chan int, 2)

//     go func() {
//         defer close(ch)
//         for i := 0; i< 5; i++ {
//             ch <- i
//             fmt.Printf("データ送信: %d\n", i)
//         }
//     }()

//     time.Sleep(time.Second * 1)
//     for v := range ch {
//         fmt.Printf("受信側にて出力: %d\n", v)
//         time.Sleep(time.Second * 1)
//     }
// }

//---------------------------------------------
//Packages, variables, and functions.
//---------------------------------------------

// package main

// import (
//     "fmt"
//     "math"
// )

// func main() {
//     fmt.Println(math.Pi)
// }

// package main

// import "fmt"

// func add(x int, y int) int {
//     return x + y
// }

// func main() {
//     fmt.Println((add(42, 13)))
// }

// package main

// import "fmt"

// func add (x,y int ) int {
//     return x + y
// }

// func main() {
//     fmt.Println((add(42,13)))
// }

// package main

// import "fmt"

// func swap(x, y string) (string, string) {
//     return y, x
// }

// func main() {
//     a,b := swap("hello", "world")
//     fmt.Println(a,b)
// }

// package main

// import "fmt"

// func split(sum int) (x, y int) {
//     x = sum * 4 / 9
//     y = sum - x
//     return
// }

// func main() {
//     fmt.Println(split(17))
// }

// package main

// import "fmt"

// var c, python, java bool

// func main() {
//     var i int
//     fmt.Println(i, c, python, java)
// }

// package main

// import "fmt"

// var i, j int = 1, 2

// func main() {
//     var c, python, java = true, false, "no!"
//     fmt.Println(i, j, c, python, java)
// }

// package main

// import "fmt"

// func main() {
//     var i, j int = 1, 2
//     k := 3
//     c, python, java := true, false, "no"

//     fmt.Println(i, j, k, c, python, java)
// }

// package main

// import (
//     "fmt"
//     "math/cmplx"
// )

// var (
//     ToBe bool = false
//     MaxInt uint64 = 1<<64 - 1
//     z complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
//     fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
//     fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
//     fmt.Printf("Type: %T Value: %v\n", z, z)
// }

package main

import (
    "fmt"
    "math"
)

func main() {
    var x,y int = 3,4
    var f = math.Sqrt(float64(x*x + y*y))
    var z  = f
    fmt.Println(x,y,z)
}

//---------------------------------------------
//Flow control statements: for, if, else, switch and defer
//---------------------------------------------