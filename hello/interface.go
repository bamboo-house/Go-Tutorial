package main

import (
	"fmt"
)

// インターフェース
type Animal interface {
	// 吠える
	Bark() string
	// 走る
	Run() string
	// 食べる
	Eat() string
}

// 構造体
type Dog struct {
	name string
	kind string
}

type Cat struct {
	name string
	kind string
}

// インターフェースの実装
func (d Dog) Bark() string {
	fmt.Println("わんわん")
	fmt.Println(d.name)
	return "わんわん"
}

func (d Dog) Run() string {
	fmt.Println("飼い主のもとへ走る")
	return "飼い主のもとへ走る"
}

func (d Dog) Eat() string {
	fmt.Println("ドッグフードを食べる")
	return "ドッグフードを食べる"
}

func (d Dog) Bite() string {
	fmt.Println("飼い主を噛む")
	return "飼い主を噛む"
}

func (c Cat) Bark() string {
	fmt.Println("にゃーにゃー")
	return "にゃーにゃー"
}

func (c Cat) Run() string {
	fmt.Println("ベランダの窓から出て走る")
	return "ベランダの窓から出て走る"
}

func (c Cat) Eat() string {
	fmt.Println("キャットフードを食べる")
	return "キャットフードを食べる"
}

// インターフェースが引数となるメソッド
func AllAnimal(a Animal) {
	a.Bark()
	a.Run()
	a.Eat()
}

func SampleInterFace() {
	dog := Dog{name: "ポチ", kind: "柴犬"}
	cat := Cat{name: "タマ", kind: "スコティッシュホールド"}
	var animal Animal
	fmt.Println("<犬の動作>")
	animal = dog // Dog型がインターフェースであるAnimalを満たしているので、代入できる
	AllAnimal(animal)

	fmt.Println("<猫の動作>")
	animal = cat
	AllAnimal(animal)
}