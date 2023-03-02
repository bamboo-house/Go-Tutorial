package main

func SampleType() {
	// var sum int
	// sum = 5 + 6 + 3
	// avg := sum / 3
	// fmt.Println(avg)
	// if float64(avg) > 4.5 {
	// 	fmt.Println("good")
	// }

	// var a, b, c bool
	// if a && b || !c {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }

	// fmt.Println(false && true)


	// 構造体、配列、スライス
	// var ns []int
	// var m map[string]int
	// fmt.Println(ns, m)

	// var p struct {
	// 	name string
	// 	age int
	// }
	// p.name = "Taro"
	// fmt.Println(p)

	// j := struct {
	// 	name string
	// 	age int
	// } {
	// 	name: "shuto",
	// 	age: 22,
	// }
	// j.age++
	// fmt.Println(j)

	// var ns1 []int
	// ns1 = make([]int, 3, 10)
	// fmt.Println(ns1)

	// var ns2 = []int{10, 20, 30, 40, 50}
	// fmt.Println(ns2)

	// 配列とスライス
	// array := [5]int{1, 2, 3, 4, 5}
	// sli := array[1:3]
	// fmt.Println(reflect.TypeOf(array))
	// fmt.Println(reflect.TypeOf(sli))
	// fmt.Println(&array)

	// ns := []int{10, 20, 30, 40, 50}
	// // 要素にアクセス
	// fmt.Println(ns[3])
	// // 長さ
	// fmt.Println(len(ns))
	// // 容量
	// fmt.Println(cap(ns))
	// // 要素の追加
	// // 容量が足りない場合は背後の配列が再確保される
	// ns = append(ns, 60, 70)
	// fmt.Println(ns)
	// fmt.Println(len(ns), cap(ns)) // 長さと容量

	// a := []int{10, 20}
	// fmt.Println(a, cap(a))

	// b := append(a, 30)
	// a[0] = 100
	// fmt.Println(a, cap(a))
	// fmt.Println(b, cap(b))
	// fmt.Println(&a[0], &b[0])

	// c := append(b, 40)
	// b[1] = 200
	// c[2] = 300
	// fmt.Println(b, cap(b))
	// fmt.Println(c, cap(c))
	// fmt.Println(&a[0], &b[0], &c[0])

	// d := append(c, 50)
	// c[1] = 400
	// d[2] = 500
	// fmt.Println(c, cap(c))
	// fmt.Println(d, cap(d))
	// fmt.Println(&a[0], &b[0], &c[0], &d[0])

	// ns := []int{10, 20, 30, 40, 50}
	// ns = append(ns[:1], ns[2:]...)
	// fmt.Println(ns)

	// ns := []int{19, 86, 1, 12}
	// var sum int
	// for i := 0; i < len(ns); i++ {
	// 	sum += ns[i]
	// }
	// fmt.Println(sum)

	// m := map[string]int{"x": 10, "y": 20}

	// fmt.Println(m["x"])
	// m["z"] = 30
	// n, ok := m["z"]
	// fmt.Println(n, ok)

	// delete(m, "z")
	// fmt.Println(m)

	// n, ok = m["z"]
	// fmt.Println(n, ok)

	// ユーザー定義型
	// type MyInt int
	// var n int = 100
	// m := MyInt(n)
	// n = int(m)
	// fmt.Println(n, m)

	// type Score struct {
	// 	UserID string
	// 	GameID int
	// 	Pont int
	// }
	// s := Score{"shuto", 1, 200}
	// fmt.Println(s)
	// fmt.Printf("%T", n)
}