package main

// 导包的三种方式
import (
	"fmt"
	"golangStudy/init5"
	_ "golangStudy/init5/lib1"
	mylib2 "golangStudy/init5/lib2"
	. "golangStudy/init5/lib3"
)

func main() {
	fmt.Println("Hello world")
	//lib1.Lib1Test()
	mylib2.Lib2Test()
	Lib3Test()
	// 只能使用这个package下面直接go文件中的函数 package中的package因为package路劲不一样了 所以不行
	init5.Init55()

	// 从字符串解析到bool值
	//res, err := strconv.ParseBool("1/t/T")
	//fmt.Println(res, err)
	// 分清楚数组和数组切片

}
