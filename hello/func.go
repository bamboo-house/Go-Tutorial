package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}


func SampleFunc() {
	// fmt.Println(add(1, 2))
	// fmt.Println(swap(1, 2))
	// x, _ := swap(1, 2)
	// fmt.Println(x)
	// fmt.Println(swap2(1, 2))

	// msg := "Hello, World"
	// func() {
	// 	fmt.Println(msg)
	// }()

	// f := make([]func() string, 2)
	// f[0] = func() string { return "Hello" }
	// f[1] = func() string { return "World" }
	// fmt.Println(f)
	// for _, v := range f {
	// 	fmt.Println(v())
	// }

	// fs := make([]func(), 3)
	// fmt.Println(fs)

	// for i := range fs {
	// 	fs[i] = func() { fmt.Println(i) }
	// }
	// fmt.Println(fs)
	// for _, f := range fs {
	// 	fs[2]()
	// 	f()
	// 	// fmt.Println(i) 
	// }

	// p := struct{age int; name string} {age: 10, name: "Taro"}
	// p2 := p
	// p2.age = 20
	// fmt.Println(p.age, p.name)
	// fmt.Println(p2.age, p2.name)

	// var x int
	// f(&x)
	// fmt.Println(x)

	// ns := []int{10, 20, 30}
	// ns2 := ns
	// ns[1] = 100
	// fmt.Println(ns[0], ns[1], ns[2])
	// fmt.Println(ns2[0], ns2[1], ns2[2])

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(i)
	// 	if i % 2 == 0 {
	// 		fmt.Println("-偶数")
	// 	} else {
	// 		fmt.Println("-奇数")
	// 	}
	// }

	// n, m := swap(10, 20)

	// fmt.Println(n, m)

	// var hex Hex = 108
	// fmt.Println(hex.string())

	// var n MyInt
	// fmt.Println(n)
	// (&n).Inc()
	// fmt.Println(n)

	var hex Hex = 102
	f := hex.String
	fmt.Println(f())
}

// type Hex int
// func (h Hex) string() string {
// 	return fmt.Sprintf("%x", int(h))
// }

type MyInt int
func (n *MyInt) Inc() { *n++ }

type Hex int
func (h Hex) String() string {
	return fmt.Sprintf("%d", int(h))
}

// func f(xp *int) {
// 	fmt.Println(*xp)
// 	*xp = 100
// }