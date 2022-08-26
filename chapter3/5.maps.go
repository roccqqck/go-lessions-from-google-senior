package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "immoc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 empty map
	var m3 map[string]int      // m3 nil, go的nil可以參與運算
	var m4 = make(map[string]int)

	fmt.Println(m, m2, m3, m4)

	fmt.Println("Traversing map ------------>")
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
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
