package main

import "fmt"

func swap3(a, b int) {
	a, b = b, a
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {

	// go 指針無法運算
	var a int = 2
	var pa *int = &a

	*pa = 3
	fmt.Println(a)

	a, b := 3, 4

	swap3(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)

	a, b = 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
