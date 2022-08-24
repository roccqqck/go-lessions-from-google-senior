package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我喜歡你媽!"
	fmt.Println(s)
	fmt.Println(len(s)) // utf8編碼長度

	fmt.Println([]byte(s))
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println(" ")

	for i, ch := range s { // ch 是 rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) // 英文size 1, 中文size 3
		bytes = bytes[size:]               // bytes去頭
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 最常用
	for i, ch := range []rune(s) {
		fmt.Printf("(%d  %c) ", i, ch)
	}

	fmt.Println()
}
