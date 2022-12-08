package main

import (
	"fmt"
)

// array和slice的区别

// [4]int 和[10]int 是两个不同的数据类型
// 静态数组的问题, 而且还是一个值拷贝的过程, 修改不生效
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index: ", index, ", value: ", value)
	}
}

// 切片是引用传递 传递的是指针，就根本没有理由用静态数组
// 那动态指针怎么扩容呢, 用append并接受
func printDynamicArray(myArray []int) {
	// make是重新开辟空间
	myArray = make([]int, 3)
	for _, value := range myArray {
		fmt.Println("value: ", value)
	}
	// 定义slice没有开辟空间的时候 slice = nil
}

func main() {

	// 0, 0 as default
	array := [4][2]int{1: {10, 11}, 3: {5, 6}}
	//array := [4][2]int{{10, 11}, {1, 2}, {3, 4}, {5, 6}}
	fmt.Println(array)
	s := []int{1, 2}
	fmt.Println(s)
	s = append(s, 3, 4, 5)
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	// 取地址
	fmt.Printf("%p\n", s)
	fmt.Printf("%p %p %p", s, &s[0], &s[1])
	fmt.Println()
	//array = append(array, {1, 2})
	// 从切片里面删除数据 copy和append
	seq := []string{"a", "b", "c", "d"}
	fmt.Println(seq)
	fmt.Println(cap(seq))
	seq = append(seq[:3], "ma")
	fmt.Println(seq)

	b := [][]int{{11}, {1, 2}}
	b[0] = append(b[0], 22)
	fmt.Println(b)

	// ================== 两种遍历的方式

	// 定长数组, 通过索引值赋值
	var myArray1 [10]int
	myArray1[0] = 1
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1)
	}

	for index, value := range myArray1 {
		fmt.Println(index, value)
	}

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := []int{1, 2}
	fmt.Printf("myArray2 = %T\n", myArray2)
	fmt.Printf("myArray3 = %T\n", myArray3)
	myArray3 = append(myArray3, 2)
	myArray3[2] = 1
	fmt.Println("111")
	//myArray3[3] = 2 // 会报错
	fmt.Println("myArray3 = ", myArray3)
	printDynamicArray(myArray3)
	fmt.Println("myArray3 = ", myArray3) // 和上面的选项相等

	myArray4 := make([]int, 3, 5)
	fmt.Println("myArray4: ", myArray4, ", len: ", len(myArray4), ", cap: ", cap(myArray4))
	//myArray4[3] = 1 // 虽然有空间但是不给用，会报错，只能用append 这种情况下不用重新开辟空间
	fmt.Println(&myArray4)
	myArray4 = append(myArray4, 4)
	fmt.Println("myArray4: ", myArray4, ", len: ", len(myArray4), ", cap: ", cap(myArray4))
	fmt.Println(&myArray4) // ???
	// append超过cap之后会double cap

}
