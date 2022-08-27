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
	for ; n > 0; n /= 2 { // while n>0 == true
		lsb := n % 2
		result = strconv.Itoa(lsb) + result // strconv.Itoa()  數字轉字串
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
	scanner := bufio.NewScanner(reader) // bufio.NewScanner()逐行讀取,    ioutil.ReadFile()全部讀取
	for scanner.Scan() {                // while   , Scan()在讀到一行時返回true，沒讀到返回false
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
