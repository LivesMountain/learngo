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

func test(i int) string {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}
func main() {
	fmt.Println(
		convertToBin(5),
		reflect.TypeOf(convertToBin(13)),
	)
}
