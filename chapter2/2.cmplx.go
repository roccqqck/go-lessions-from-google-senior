package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) // 複數絕對值

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // Pow(x, y) = x的y次方

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) // %.3f 只取小數點後三位

}

func main() {
	euler()

}
