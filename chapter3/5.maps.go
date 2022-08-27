package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "immoc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 make(map[string]int)是empty map , 不能宣告值
	var m3 map[string]int      // m3 nil, go的nil可以參與運算
	var m4 = make(map[string]int)

	var m5 = map[string]int{
		"name":    1,
		"course":  2,
		"site":    3,
		"quality": 4,
	}

	var m6 = map[int]string{
		1: "ccmouse",
		2: "golang",
		3: "immoc",
		4: "notbad",
	}

	fmt.Println(m, m2, m3, m4, m5, m6)

	fmt.Println("Traversing map m------------>")
	for k, v := range m {
		fmt.Println(k, v) // map順序是隨機的
	}

	fmt.Println("------------>")
	for k := range m {
		fmt.Println(k)
	}

	fmt.Println("------------>")
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("Traversing map m5------------>")
	for k, v := range m5 {
		fmt.Println(k, v) // map順序是隨機的
	}

	fmt.Println("------------>")
	for k := range m5 {
		fmt.Println(k)
	}

	fmt.Println("------------>")
	for _, v := range m5 {
		fmt.Println(v)
	}

	fmt.Println("Traversing map m6------------>")
	for k, v := range m6 {
		fmt.Println(k, v) // map順序是隨機的
	}

	fmt.Println("------------>")
	for k := range m6 {
		fmt.Println(k)
	}

	fmt.Println("------------>")
	for _, v := range m6 {
		fmt.Println(v)
	}

	fmt.Println("Getting values------------>")
	courseName := m["course"]
	fmt.Println(courseName)

	causeName := m["cause"]
	fmt.Println(causeName) // 不會出錯

	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	causeName2, ok := m["cause"] // 只要有一個變數未被宣告就可以用 :=
	fmt.Println(causeName2, ok)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting Values------------>")
	fmt.Println(m)
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	fmt.Println(m)
}
