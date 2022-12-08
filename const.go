package main

import "fmt"

// const相对于var属性是不可修改
// iota只能配合const使用
const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

// 不要想着怎么去回复出彩 -> 而是吧心里的想法说出来就好了 一致性
// 如果一个包下面的不同文件的函数名重复了，引包的时候要怎么处理 不能重复会报错
func main() {

	const v int = 1
	var k int = 2
	fmt.Println(v, k)

}
