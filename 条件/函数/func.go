package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// q,r= div(a,b)
		q, _ := div1(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}
func div1(a, b int) (q, r int) {
	return a / b, a % b
}
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// op 是一个函数  func是描述op这个函数是一个怎样的函数，如下就是op是接受两个int类型的函数，返回一个int
func apply(op func(int int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("opname:%s"+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply1(op func(int, int) int, a, b int) int {
	// p := reflect.ValueOf(op).Pointer()
	// opName := runtime.FuncForPC(p).Name()
	// fmt.Printf("Calling function %s with args(%d, %d)", opName, a, b)
	return op(a, b)
}
func pow1(a, b int) int {

	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	// if result, err := eval(3, 4, "x"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
	// fmt.Println(eval(3, 4, "x"))
	// fmt.Println(div(13, 3))
	// a, b := div1(13, 3)
	// fmt.Println(a, b)
	fmt.Println(apply1(pow1, 3, 4))
	// fmt.Println(apply(pow, 3, 4))
}
