package main

import (
	"fmt"
	"math"
)

func variableZeroValue() {
	var a int
	var s string

	fmt.Println(a, s)
	fmt.Printf("%d    %q\n", a, s)
}

func varibleInitalValue() {
	var b, c int = 11, 221
	var s1 string = "abc" +
		"hahhaha"
	fmt.Println(b, c, s1)

}

func varibleTypeDedution1() {
	var b, c = 11, 222
	var s1 = "abc" +
		"hahhaha"
	fmt.Println(b, c, s1)
	var i1, i2, b1, s2 = 11, 22, "def", false
	fmt.Println(i1, i2, b1, s2)

}

func varibleTypeDedution2() {
	b, c := 11, 223
	s1 := "abc" +
		"hahhaha"
	fmt.Println(b, c, s1)

	i1, i2, b1, s2 := 11, 22, false, "def"
	i1 = 12
	fmt.Println(i1, i2, b1, s2)
}

// package varibles
var (
	a = 001
	s = "hello"
)

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // Sqrt(x) 只吃float64,  return float64

	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("Hello,World!!")

	variableZeroValue()

	varibleInitalValue()

	varibleTypeDedution1()

	varibleTypeDedution2()

	fmt.Println(a, s)

	triangle()

}
