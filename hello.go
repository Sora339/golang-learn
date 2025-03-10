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

// package main

// import (
//     "fmt"
//     "math"
// )

// func main() {
//     var x,y int = 3,4
//     var f = math.Sqrt(float64(x*x + y*y))
//     var z  = f
//     fmt.Println(x,y,z)
// }

//---------------------------------------------
//Flow control statements: for, if, else, switch and defer
//---------------------------------------------

// package main

// import "fmt"

// func main() {
//     sum := 0
//     for i := 0; i < 10; i++ {
//         sum += i
//     }
//     fmt.Println(sum)
// }

// package main

// import "fmt"

// func main() {
//     sum := 1
//     for sum < 1000 {
//         sum += sum
//     }
//     fmt.Println(sum)
// }

// package main

// func main() {
//     for {

//     }
// }

// package main

// import (
//     "fmt"
//     "math"
// )

// func sqrt(x float64) string {
//     if x < 0 {
//         return sqrt(-x) + "i"
//     }
//     return fmt.Sprint(math.Sqrt(x))
// }

// func main() {
//     fmt.Println(sqrt(2), sqrt(-4))
// }

// package main

// import (
//     "fmt"
//     "math"
// )

// func pow(x,n,lim float64) float64 {
//     if v := math.Pow(x,n); v < lim {
//         return v
//     }
//     return lim
// }

// func main() {
//     fmt.Println(
//         pow(3,2,10),
//         pow(3,3,20),
//     )
// }


// package main

// import (
//     "fmt"
//     "math"
// )

// func pow(x, n, lim float64) float64 {
//     if v := math.Pow(x, n); v < lim {
//         return v
//     } else {
//         fmt.Printf("%g >= %g\n", v, lim)
//     }
//     return lim
// }

// func main() {
//     fmt.Println(
//         pow(3,2,10),
//         pow(3,3,20),
//     )
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func Sqrt(x float64) float64 {
//     z := 1.0
//     for i := 0; i < 10; i++ {
//         preZ := z
//         z -= (z*z - x)/(2 * z)
//         fmt.Println(fmt.Sprintf("%d回目",i + 1))
//         if math.Abs(z - preZ) < 1e-6 {
//             break
//         }
//     }
//     return z
// }

// func main() {
//     fmt.Println(Sqrt(2))
//     fmt.Println(math.Sqrt(2))
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Print("Go runs on")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X.")
//     case "linux":
//         fmt.Println(("Linux."))
//     default: fmt.Printf("%s.\n",os)    
// 	}
// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     fmt.Println("When's Saturday?")
//     today := time.Now().Weekday()
//     fmt.Println(today +1)
//     switch time.Saturday {
//     case today + 0:
//         fmt.Println("Today")
//     case today + 1:
//         fmt.Println("Tomorrow")
//     case today + 2:
//         fmt.Println("In two days")
//     default:
//         fmt.Println("Too far away.")
//     }

// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     t := time.Now()
//     switch {
//     case t.Hour() < 12:
//         fmt.Println(("Good moring!"))
//     case t.Hour() < 17:
//         fmt.Println("Good afternoon!")
//     default:
//         fmt.Println("Good evening.")
//     }
// }

// package main

// import "fmt"

// func main() {
//     defer fmt.Println("world")

//     fmt.Println(("Hello"))
// }

// package main

// import "fmt"

// func main() {
//     fmt.Println("counting")

//     for i:= 0; i < 10; i++ {
//         defer fmt.Println(i)
//     }

//     fmt.Println(("done"))
// }


//---------------------------------------------
//More types: structs, slices, and maps.
//---------------------------------------------

// package main

// import "fmt"

// func main() {
//     i,j := 42, 2701

//     p := &i
//     fmt.Println(*p)
//     *p = 21
//     fmt.Println(i)

//     p = &j
//     fmt.Println(*p)
//     *p = *p / 37
//     fmt.Println(j)
// }

// package main

// import "fmt"

// type Vertex struct {
//     X int
//     Y int
// }

// func main() {
//     fmt.Println(Vertex{1,2})
// }

// package main

// import "fmt"

// type Vertex struct {
//     X int
//     Y int
// }

// func main() {
//     v := Vertex{1, 2}
//     v.X = 4
//     fmt.Println(v)
// }

// package main

// import "fmt"

// type Vertex struct {
//     X int
//     Y int
// }

// func main() {
//     v:=Vertex{1,2}
//     p:=&v
//     p.X=1e9
//     fmt.Println(v)
// }

// package main

// import "fmt"

// type Vertex struct {
//     X,Y int
// }

// var (
//     v1 = Vertex{1,2}
//     v2 = Vertex{X:1}
//     v3 = Vertex{}
//     p = &Vertex{1,2}
// )

// func main() {
//     fmt.Println(v1,p,v2,v3)
// }

// package main

// import "fmt"

// func main() {
//     var a [2]string
//     a[0] = "Hello"
//     a[1] = "World"
//     fmt.Println(a[0],a[1])
//     fmt.Println(a)

//     primes := [6]int{2, 3, 5, 7, 11, 13}
//     fmt.Println(primes)
// }

package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)
}