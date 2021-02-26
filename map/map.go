package main

import "fmt"

func main() {
	// 第一个string是代表k的数据类型 第二个string是代表value的数据类型
	m := map[string]string{
		"name":   "ccomuse",
		"course": "golang",
		"site":   "imooc",
	}
	fmt.Println(m)

	m2 := make(map[string]int) //m2==empty map
	fmt.Println(m2)

	var m3 map[string]int //m3==nil
	fmt.Println(m3)

	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 已知key获取value
	fmt.Println(m["name"])
	// 判断key是否在map里
	coursename, ok := m["course"]
	coursename1, ok1 := m["cours"]
	if coursename, ok := m["course"]; ok {
		fmt.Println(coursename)
	} else {
		fmt.Println("没有")
	}
	fmt.Println(coursename, ok, coursename1, ok1)
	// 删除元素
	_, ok = m["name"]
	fmt.Println(ok)
	delete(m, "name")
	_, ok = m["name"]
	fmt.Println(ok)

	//
}
