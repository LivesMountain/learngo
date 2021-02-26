package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes你很牛逼"
	fmt.Printf("%s\n%X\n", []byte(s), []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune 是一个int32
		fmt.Printf("(%d,%x)", i, ch)
	}
	// E4 BD A0是utf8的编码，4f60是一个Unicode 4字节
	fmt.Println()
	// 打印文本长度
	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Println(ch, size)
		fmt.Println(bytes)
		bytes = bytes[size:]
		fmt.Println(bytes)
		fmt.Printf("%c ", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", i, ch)

	}
}
