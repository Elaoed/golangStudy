package lib1

import "fmt"

// 潜规则 函数名大写说明是对外开放的public, 小写说明是内部函数只能在包内调用
func init() {
	fmt.Println("lib1.init()...")
}

func Lib1Test() {
	fmt.Println("lib1.test()...")
}
