package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsolutezeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 52) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f == 0)
	//fmt.Println(c == f)
	fmt.Println(c == Celsius(f))
	fmt.Println(CToF(Boiling))
	fmt.Println(Fahrenheit(Boiling))
	fmt.Println(FToC(Fahrenheit(Boiling)))
}
