package main

import "fmt"

func main() {
	x := 1
	p := &x
	println(*p)
	*p = 2
	println(p)
	println(x)
	println(&x)
	println()
	var y, z int
	fmt.Println(&y == &y, &y == &z)

}
