package main

import (
	"fmt"
)

var a int
type hotdog int // own type
var b hotdog
func main(){
	a = 25
	b=30

	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	a = int (b) //This is a conversion not a casting
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	//y = "Arda" // because go lang is static type not dynamic type you dont change int to string or other type
	//fmt.Println(y)

}

