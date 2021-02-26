package main

import "fmt"

func printslice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}

func main() {
	var s []int //zero value for slice is nil   s=nil
	for i := 0; i < 100; i++ {
		printslice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 16)
	// s3 := make([]int, 16, 32)

	copy(s2, s1)
	fmt.Println(s2)
	printslice(s2)
	// 删除第三个元素
	s2 = append(s2[:3], s2[4:]...)
	printslice(s2)
	// 删除头尾元素
	tou := s2[1:]
	fmt.Println(tou)
	printslice(tou)
	tail := len(tou) - 1
	s2 = tou[:tail]
	printslice(s2)
	fmt.Println(s2)

}
