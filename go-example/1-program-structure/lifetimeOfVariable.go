package main

var global *int

func main() {
	var x int
	x = 1
	global = &x
	print(global)

}

func g() {
	y := new(int)
	*y = 1
}
