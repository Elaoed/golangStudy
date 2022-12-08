package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 66, 77
}

// 显示制定返回值的名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 666
	r2 = 777
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 666
	r2 = 777
	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)
	ret1, ret2 := foo2("abc", 555)
	fmt.Println("ret1: ", ret1, ",ret2: ", ret2)
}
