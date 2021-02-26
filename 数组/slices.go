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
	// 数组切片的slices是对底层数组的一个view所以，s2是可以取到的，是因为slice内部是有三个元素组成
	// ptr只想slice开头的元素，len数字的长度，cap是代表了ptr的到后面所有的长度
	// 如下s1的ptr是100坐标0 ，len为4 ，cap为7
	// s2 len为2，cap为4
	// clice可以向后扩展，不能向前扩展
	s1 = arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s1=%v,len(s2)=%d,cap(s2)=%d", s2, len(s2), cap(s2))
	fmt.Println(s1, s2)
}
