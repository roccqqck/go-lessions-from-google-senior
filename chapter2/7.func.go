package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

// 13 / 4 = 3...1
func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func eval2(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b) // 沒用到要 _
		return q
	default:
		panic("unsupported operation: " + op)
	}
}
func eval3(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()    // 函數指針
	opName := runtime.FuncForPC(p).Name() // 函數名

	fmt.Printf("calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)

}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}

// ...int 可input多個int
func sum(number ...int) int {
	s := 0
	for i := range number { // range number 是 number的array
		s += number[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))

	fmt.Println("------->")
	fmt.Println(div(13, 3))

	fmt.Println("------->")
	q, r := div2(13, 2)
	fmt.Println(q, r)

	fmt.Println("------->")
	fmt.Println(eval2(3, 4, "/"))

	fmt.Println("------->")
	fmt.Println(eval3(3, 4, "x"))

	fmt.Println("------->")
	if result, err := eval3(123, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("------->")
	fmt.Println(apply(pow, 3, 4))

	fmt.Println("------->")
	// 類似lambda 直接定義匿名的func
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("------->")
	fmt.Println(sum(1, 2, 3, 4))

}
