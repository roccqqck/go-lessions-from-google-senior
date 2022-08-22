package main

import (
	"fmt"
	"math"
)

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4 //不指定int或float
	)
	var c int
	c = int(math.Sqrt((a*a + b*b))) //不需轉換成float

	fmt.Println(filename, a, b, c)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

func enums1() {
	const (
		cpp = iota // 自增值
		java
		python
		_
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func enums2() {
	const (
		b  = 1 << (10 * iota) // a << 1 將a的二進位左移1位 右補0 相當於該數乘以2
		kb                    // a << 10  乘以2的10次方
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {

	consts()
	fmt.Println("------->")
	enums()
	fmt.Println("------->")
	enums1()

	fmt.Println("------->")
	enums2()
}
