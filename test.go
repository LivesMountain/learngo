package main

import "fmt"

var (
	aaa = 3
	ss  = "222"
	bb  = true
)

func variable() {
	var a int
	var b string
	fmt.Println(a, b) //Println一般可以直接跟变量的名字，printf可以格式化字符串
	fmt.Printf("%d %q\n", a, b)
}
func varinit() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func vartype() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func varshorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s) //第一次定义变量的时候可以省略var 变为:=后面再变量的赋值或者转换的过程中可以不用：=直接=,:=只能在函数内部使用
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(a, b)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	const (
		b= 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func sayok() {
	fmt.Println("ok")
}
func main() {
	variable()
	sayok()
	varinit()
	fmt.Println(aaa, ss, bb)
	enums()
}




