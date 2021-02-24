package main

import "fmt"

func printarray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	slices()
	var arr1 [5]int
	// arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	// var grid [4][5]int
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// range可以获得数组的下标i,v可以获得数组值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
	// 值类型传递
	// printarray(arr1)
	// printarray(arr3)
	fmt.Printf("数组是值传递类型的，在调用函数传入数组的时候，把变量拷贝了一份传给数组，但是实际拷贝前的数组不变", arr3)
	// 指针传递
	printarray(&arr1)
	printarray(&arr3)
	fmt.Printf("数组是值传递类型的，在调用函数传入数组的时候，把变量拷贝了一份传给数组，但是实际拷贝前的数组不变", arr3)
}
