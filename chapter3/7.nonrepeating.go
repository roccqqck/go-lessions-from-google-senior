package main

import "fmt"

func length0fNonRepeatingSubStr2(s string) int {

	last0ccurred := make(map[byte]int) // 紀錄這個字上次出現的位置,empty map

	start := 0 // 此不連續句子的頭的位置
	maxLength := 0

	// []byte(s)  字串轉byte
	for i, ch := range []byte(s) {
		if lastI, ok := last0ccurred[ch]; ok && lastI >= start { // ok = true這個字出現過, 並且lastI >= start這個字位置>=start位置
			start = lastI + 1 // start 往左移
		} // ok = false這個字從未出現過, 或者 lastI < start這個字位置 < start位置  , start位置不變

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		last0ccurred[ch] = i // 每次都要記錄這個字 上次出現的位置
	}
	return maxLength
}

func length0fNonRepeatingSubStr3(s string) int {

	last0ccurred := make(map[rune]int)

	start := 0
	maxLength := 0

	for i, ch := range s {
		if lastI, ok := last0ccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i
	}
	return maxLength
}

func length0fNonRepeatingSubStr(s string) int {

	last0ccurred := make(map[rune]int)

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := last0ccurred[ch]; ok && lastI >= start {

			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println("length0fNonRepeatingSubStr2 -------->")
	fmt.Println(length0fNonRepeatingSubStr2("abcabcbb"))
	fmt.Println(length0fNonRepeatingSubStr2("bbbbb"))
	fmt.Println(length0fNonRepeatingSubStr2("pwwkew"))
	fmt.Println(length0fNonRepeatingSubStr2(""))
	fmt.Println(length0fNonRepeatingSubStr2("b"))
	fmt.Println(length0fNonRepeatingSubStr2("abcdef"))
	fmt.Println(length0fNonRepeatingSubStr2("Yes我喜歡你!"))
	fmt.Println(length0fNonRepeatingSubStr2("Yes我喜歡媽媽!"))
	fmt.Println(length0fNonRepeatingSubStr2("Yes我喜歡你媽媽!"))

	fmt.Println("length0fNonRepeatingSubStr3 -------->")
	fmt.Println(length0fNonRepeatingSubStr3("abcabcbb"))
	fmt.Println(length0fNonRepeatingSubStr3("bbbbb"))
	fmt.Println(length0fNonRepeatingSubStr3("pwwkew"))
	fmt.Println(length0fNonRepeatingSubStr3(""))
	fmt.Println(length0fNonRepeatingSubStr3("b"))
	fmt.Println(length0fNonRepeatingSubStr3("abcdef"))
	fmt.Println(length0fNonRepeatingSubStr3("Yes我喜歡你!"))
	fmt.Println(length0fNonRepeatingSubStr3("Yes我喜歡媽媽!"))
	fmt.Println(length0fNonRepeatingSubStr3("Yes我喜歡你媽媽!"))

	fmt.Println("length0fNonRepeatingSubStr -------->")
	fmt.Println(length0fNonRepeatingSubStr("abcabcbb"))
	fmt.Println(length0fNonRepeatingSubStr("bbbbb"))
	fmt.Println(length0fNonRepeatingSubStr("pwwkew"))
	fmt.Println(length0fNonRepeatingSubStr(""))
	fmt.Println(length0fNonRepeatingSubStr("b"))
	fmt.Println(length0fNonRepeatingSubStr("abcdef"))
	fmt.Println(length0fNonRepeatingSubStr("Yes我喜歡你!"))
	fmt.Println(length0fNonRepeatingSubStr("Yes我喜歡媽媽!"))
	fmt.Println(length0fNonRepeatingSubStr("Yes我喜歡你媽媽!"))

}
