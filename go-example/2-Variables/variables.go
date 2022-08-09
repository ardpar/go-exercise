package main

import "fmt"

var a = 500

func main() {
	var roomNumber, floorNumber int = 154, 3
	fmt.Println(roomNumber, floorNumber)
	var password = "notSecured"
	fmt.Println(password)

	buildingNumber, streetNumber := 20, 7
	fmt.Println(buildingNumber, streetNumber)

	var (
		name, location, age = "Arda", "Istanbul", 34
	)
	addName, addLocation := "Ev", "Istanbul"
	fmt.Println(name, location, age)
	fmt.Println(addName, addLocation)
	var name1 = "Arda"
	name2 := "Arda"
	fmt.Println(name1, name2)
	var a = 600
	fmt.Println(a)
	a = 300
	fmt.Println(a)
	testGlobal()
}
func testGlobal() {
	fmt.Println(a)
}
func testLocal() {
	fmt.Println(name)
}
