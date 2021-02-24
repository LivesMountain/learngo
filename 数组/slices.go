package main

import "fmt"

func updateslice(s []int) {
	s[0] = 100
}

func slices() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr[2:6])
	s1 := arr[2:]
	updateslice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
}
