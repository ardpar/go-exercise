package main

func main() {

	i, j := 0, 1
	println(i + j)
	i, j = j, i
	println(i)
	println(j)
}
