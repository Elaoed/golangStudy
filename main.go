package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("Hello world")
	a := []int{1, 2}
	fmt.Println(a)
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
	for i, v := range b {
		fmt.Println(i, v)
	}

	m1 := map[string]any{"name": "张三", "age": 18, "gender": true}
	delete(m1, "name")
	fmt.Println(m1)
	m1 = make(map[string]any)
	go func() {
		for {
			m1["1"] = 1
		}
	}()
	go func() {
		for {
			_ = m1["1"]
		}
	}()
	//select {}

	var m sync.Map
	m.Store("name", "张三")
	m.Store("age", "18")
	var k int
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		k++
		return true
	})
	// Const常量的定义
	const v1 = 1
	// 从字符串解析到bool值
	res, err := strconv.ParseBool("1/t/T")
	fmt.Println(res, err)
	// 分清楚数组和数组切片
}
