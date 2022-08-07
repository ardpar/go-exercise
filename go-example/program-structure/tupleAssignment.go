package main

import "os"

func main() {

	var i, j, k int = 2, 3, 5

	a, b, c := 6, 7, 8

	println(i)
	println(j)
	println(k)

	println(a)
	println(b)
	println(c)

	gdcVar := gcd(3, 2)
	println(gdcVar)

	f, err := os.Open("foo.txt")
	println(f)
	println(err)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
