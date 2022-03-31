package main

import (
	"fmt"
	"reflect"
)

var a = "Hello"
var i = 123

type Name struct {
	A string
	B bool
	C int
}

func main() {
	fmt.Println("================================")
	b := 10
	fmt.Println("a =>", &a, a, "b =>", &b, b)

	a, b := 5, 20
	fmt.Println("a =>", &a, a, "b =>", &b, b)
	fmt.Println("================================")

	fmt.Println("i =>", &i, i)
	i = 123456
	fmt.Println("i =>", &i, i)

	i, j := 123, 456
	fmt.Println("i =>", &i, i, "j =>", &j, j)
	fmt.Println("================================")

	const w = "Hello, world"
	fmt.Println("w =", w)
	fmt.Println("================================")
	fmt.Println("a is", reflect.TypeOf(a))
	fmt.Println("b is", reflect.TypeOf(b))
	fmt.Printf("i is %T\n", i)
	fmt.Println("w is", reflect.TypeOf(w))
	fmt.Println("================================")

	fmt.Printf("%v\n", Name{})
	fmt.Printf("%+v\n", Name{})
	fmt.Printf("%#v\n", Name{})
	fmt.Println("================================")
}
