package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score :%d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "c"

	case score < 90:
		g = "b"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score :%d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(65),
		grade(70),
		grade(100),
		grade(101),
	)
}
