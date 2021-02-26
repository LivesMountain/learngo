package main

import "fmt"

func lengthofnonrepeatingsubstr(s string) int {
	lastoccurred := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(s) {
		if lasti, ok := lastoccurred[ch]; ok && lasti >= start {
			start = lasti + 1
		}
		// if lastoccurred[ch] >=start{
		// 	start=lastoccurred[ch]+1
		// }
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastoccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(lengthofnonrepeatingsubstr("asdsazxcasasd"))
}
