package main

import "fmt"

// arr [5]int 是array  只能input 5個int的array array當input會複製
func printArray2(arr [5]int) {
	arr[0] = 100
	for k, v := range arr {
		fmt.Println(k, v)
	}

}

// arr []int 是 slice 可以input多個int的slice slice當input不會複製
func printArray(arr []int) {
	arr[0] = 101
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

	// range
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

	// 取array最大值
	fmt.Println("------->")
	maxi := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}

	fmt.Println(maxi, maxValue)

	fmt.Println("------->")
	printArray2(arr1)
	//printArray2(arr2)
	fmt.Println("------->")
	printArray2(arr3)

	fmt.Println("------->")
	fmt.Println(arr1, arr3)

	fmt.Println("------->")
	printArray(arr1[:])
	fmt.Println("------->")
	printArray(arr2[:])
	fmt.Println("------->")
	printArray(arr3[:])

	fmt.Println("------->")
	fmt.Println(arr1, arr2, arr3)

}
