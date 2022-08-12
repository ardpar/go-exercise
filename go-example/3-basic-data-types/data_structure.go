package main

import "fmt"

func main() {
	var isEmpty bool = true
	var isClosed = false
	isOk := true
	fmt.Println(isEmpty, isClosed, isOk)
	result := isEmpty == isClosed
	fmt.Println(result)
	//var bitExample int8 = 127
	//var shortExample int16
	//var intExample int32
	//var longExample int64
	//var float32Example float32
	//var float64Example float64

	var stringExample string = "HelloWorld"
	for _, c := range stringExample {
		fmt.Print(string(c))
	}
}
