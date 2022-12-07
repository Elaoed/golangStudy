package main

import "fmt"

func main() {
	s := "String"
	fmt.Println(s[1:])
	n := len(s)
	for i := 0; i < n; i++ {
		ch := s[i]
		fmt.Println(i, ch)
	}
	// rune代表单个Unicode字符
	k := [32]byte{}
	//z := [2 * N]struct{ x, y int}
	fmt.Println(k)
	//fmt.Println(z)
}
