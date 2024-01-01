package mathpack

import "fmt"

func AddInt[Type int | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr](num1, num2 Type) {
	sum := num1 + num2
	fmt.Println(sum)
}

func AddFloat[Type float32 | float64](floatnum1, floatnum2 Type) {
	sum := floatnum1 + floatnum2
	fmt.Println(sum)
}
