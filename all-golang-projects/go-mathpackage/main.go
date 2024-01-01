package main

import "github.com/go-mathpackage/MathPack"

func main() {
	var number1 int = 36
	var number2 int = 4987

	var float float64 = 3
	var float2 float64 = 3

	mathpack.AddInt[int](number1, number2)
	mathpack.AddFloat[float64](float, float2) 
}
