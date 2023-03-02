package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

func main() {
	// SampleFunc()
	SampleInterFace()
	// SampleType()
	// 標準入力で整数を受け取り、その値を出力する
		// var price int
		// fmt.Print("値段>")
		// fmt.Scanln(&price)
		// fmt.Printf("%d円ですね\n", price)


	// 変数に関して
		// var n int = 100
		// println(n)

		// var m int
		// println(m)

		// i, k := 300, 400
		// j := i + k
		// println(j)

		// var (
		// 	u = 2000
		// 	t = 1999
		// )
		// println(u, t)

		// var gg string = "世界"
		// ii := "unko"
		// println("hello", gg, ii)
		// fmt.Printf("hello, %s, %s", gg, ii)

		// const run = 'よ'
		// println(run)

	// 定数に関して
		// 型のある定数
		// const n int = 100
		// println(n)

		// // 型のない定数
		// const m = 200
		// println(m)

		// // 定数式の使用
		// const s = "hello, " + "世"
		// const j = "ii" + s
		// println(s)
		// println(j)
		
		// // デフォルトの型
		// const v = 10000000000000000000 / 10000000000000000000
		// var c = v
		// println(c)

		// // iotaを利用したフラグ
		// const (
		// 	FlagA = 1 << iota
		// 	FlagB
		// 	FlagC
		// 	FlagD
		// )
		// print(FlagA, FlagB, FlagC, FlagD)


	// ポインタに関して
		// var o int = 100
		// var o_pointer *int = &o
		// println("ポインタ型変数：", o_pointer)
		// println("ポインタ型変数の中身：", *o_pointer)

	// 演算に関して
		// var a bool = true
		// var b bool = false
		// println(a == b)

		// n := 100 + 200
		// m := n - 200
		// msg := "hoge" + "fuga"
		// boooo := n > m
		// println(n, m, msg, boooo)

	// 制御構文に関して
		// if文
		// x := 2
		// if x == 1 {
		// 	println("xは1です")
		// } else if x == 2 {
		// 	println("xは2です")
		// } else {
		// 	println("xは1でも2でもありません")
		// }

		// if a := -3; a > 0 {
		// 	fmt.Println(a)
		// } else {
		// 	fmt.Println(2*a)
		// }

		// switch文
		// aa := 22
		// switch aa {
		// 	// 1又は2の場合
		// 	case 1, 2:
		// 		fmt.Println("a is 1 or 2")
		// 	default:
		// 		fmt.Println("default")
		// }

		// switch {
		// case aa == 22:
		// 	fmt.Println("aa is 22")
		// }

		// for文
		// for i := 0; i <= 10; i++ {
		// 	fmt.Println(i)
		// }

		// for i, v := range []int{1, 2, 3} {
		// 	fmt.Println(i)
		// 	fmt.Println(v)
		// }

		// var i int
		// for {
		// 	if i == 22 {
		// 		break
		// 	} else if i == 10 {
		// 		fmt.Println("i is 10", i)
		// 		i++
		// 		continue
		// 	}
		// 	i++
		// }
		// fmt.Println(i)

		// for i := 1; i <= 100; i++ {
		// 	if i % 2 == 0 {
		// 		fmt.Printf("%d-偶数\n", i)
		// 	} else if i % 2 != 0 {
		// 		fmt.Printf("%d-奇数\n", i)
		// 	}
		// }

		// t := time.Now().UnixNano()
		// fmt.Println(t)
		// rand.Seed(t)
		// n := rand.Intn(6)
		// switch n + 1 {
		// 	case 1:
		// 		fmt.Println("凶")
		// 	case 2, 3:
		// 		fmt.Println("吉")
		// 	case 4, 5:
		// 		fmt.Println("中吉")
		// 	case 6:
		// 		fmt.Println("大吉")
		// }


		// 配列について
		// var array1[2] string
		// array1[0] = "hoge"
		// array1[1] = "fuga"
		// fmt.Println(array1)

		// var array2[2] int = [2]int{1, 2}
		// fmt.Println(array2)

		// array3 := [3]string{"hoge", "fuga", "piyo"}
		// fmt.Println(array3)



}