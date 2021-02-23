package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		fmt.Println(lsb)
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func test1(i int) string {
	sum := 0
	result := ""
	for i := 1; i <= 100; i++ {
		sum += i
	}
	result = strconv.Itoa(sum)
	return result
}

func main() {
	fmt.Println(
		convertToBin(5),
		reflect.TypeOf(convertToBin(13)),
		test1(10),
	)
	// fmt.Println(test1(10))
}
