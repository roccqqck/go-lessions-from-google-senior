package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我喜歡你媽!"
	s2 := "Yes, I do like your mother!"

	fmt.Println(s)
	fmt.Println(len(s)) // utf8編碼長度

	fmt.Println(s2)
	fmt.Println(len(s2)) // utf8編碼長度

	fmt.Println("[]byte(string)--------------------------------")
	fmt.Println([]byte(s))
	fmt.Printf("%X\n", []byte(s))

	fmt.Println([]byte(s2))
	fmt.Printf("%X\n", []byte(s2))

	fmt.Println("range []byte(string)--------------------------------")
	for i, b := range []byte(s) {
		fmt.Printf("(%d %X) ", i, b)
	}
	fmt.Println(" ")

	for i, b := range []byte(s2) {
		fmt.Printf("(%d %X) ", i, b)
	}
	fmt.Println(" ")

	fmt.Println("range string--------------------------------")
	for i, ch := range s { // ch 是 rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	for i, ch := range s2 { // ch 是 rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))
	fmt.Println("Rune count: ", utf8.RuneCountInString(s2))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) // 英文size 1, 中文size 3
		bytes = bytes[size:]               // bytes去頭
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 最常用
	fmt.Println("[]rune(string)--------------------------------")
	fmt.Println([]rune(s))
	fmt.Printf("%X\n", []rune(s))

	fmt.Println([]rune(s2))
	fmt.Printf("%X\n", []rune(s2))

	fmt.Println("range []rune(string)--------------------------------")
	for i, ch := range []rune(s) {
		fmt.Printf("(%d  %c) ", i, ch)
	}

	fmt.Println()

	for i, ch := range []rune(s2) {
		fmt.Printf("(%d  %c) ", i, ch)
	}

	fmt.Println()
}
