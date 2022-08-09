package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	fmt.Println(p)
	*p = 2
	fmt.Println(*p)
	x := new(int)
	y := new(int)

	fmt.Println("---------------")
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(x == y)
}
