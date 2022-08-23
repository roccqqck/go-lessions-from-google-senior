package main

import "fmt"

func printArray(arr []int) {
	arr[0] = 100
	for k, v := range arr {
		fmt.Println(k, v)
	}

}

func main() {
	//fmt.Println("Hello,World!")
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid1 [4][5]int
	var grid2 [4][5]bool

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid1)
	fmt.Println(grid2)

	fmt.Println("------->")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	fmt.Println("------->")
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	fmt.Println("------->")
	for _, k := range arr3 {
		fmt.Println(k)
	}

	fmt.Println("------->")
	for i, k := range arr3 {
		fmt.Println(i, k)
	}

	fmt.Println("------->")
	printArray(arr1[:])
	//printArray(arr2)
	fmt.Println("------->")
	printArray(arr3[:])

	fmt.Println("------->")
	fmt.Println(arr1, arr3)

}
