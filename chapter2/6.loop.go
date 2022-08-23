package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func sum0To99() int {
	sum := 0
	for i := 0; i < 100; i++ {
		sum = sum + i
	}
	return sum
}

// int -> 2進位
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { // while
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for { // infinite loop
		fmt.Println("abc")
		time.Sleep(time.Second * 3)
	}

}

func main() {

	fmt.Println(sum0To99())

	fmt.Println(
		convertToBin(5),  //11
		convertToBin(13), //1011 -->1101
		convertToBin(111111),
		convertToBin(0),
	)

	printFile("go.mod")

	s := `
		asda
		asda
		asfas       fff
		aaa
		`

	printFileContents(strings.NewReader(s))

	//forever()

}
