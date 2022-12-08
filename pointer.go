package main

import "fmt"

func changeValue(p int) {
	p = 10
}
func changePointerValue(p *int) {
	// 下潜
	*p = 10
}

func swap(a *int, b *int) {
	p := *a
	*a = *b
	*b = p
}
func main() {
	a := 1
	changeValue(a)
	fmt.Println("a = ", a)
	// 上浮
	changePointerValue(&a)
	fmt.Println("a = ", a)
	b := 5
	// 形参和实参，利用引用来操作实参
	swap(&a, &b)
	fmt.Println("a = ", a, ", b = ", b)
}
