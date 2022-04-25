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

	arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println("arr =>", &arr, arr)
	fmt.Println("arr[4] =", arr[4])
	fmt.Printf("arr => %p, %#v\n", &arr, arr)
	fmt.Printf("arr is %T\n", arr) //same as fmt.Println("ar is", reflect.TypeOf(ar))
	s := []int{6, 7, 8, 9, 10}
	fmt.Printf("s => %p, %#v\n", &s, s)
	fmt.Printf("s is %T\n", s)
	s2 := make([]int, 5, 10)
	fmt.Printf("s2 => %p, %v\n", &s2, s2)
	fmt.Printf("s2 is %T\n", s2)
	s3 := arr[1:3:4]
	fmt.Printf("s3 => %p, %#v\n", &s3, s3)
	fmt.Printf("s3 is %T\n", s3)
	fmt.Println("s3 len =", len(s3))
	fmt.Println("s3 cap =", cap(s3))
	fmt.Println("s3[0] =", s3[0])
	fmt.Println("s3[1] =", s3[1])
	fmt.Println("================================")
}
