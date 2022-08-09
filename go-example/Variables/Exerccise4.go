package main

import (
	"fmt"
)
type ownType int
var x ownType
var y int
func main() {
	fmt.Println(x)

	x = 42
	fmt.Println(x)
	y = int (x)
	fmt.Println(y)
}

