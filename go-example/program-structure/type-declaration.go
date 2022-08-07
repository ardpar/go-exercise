package main

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

}