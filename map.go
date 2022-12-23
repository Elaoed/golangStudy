package main

import (
	"fmt"
	"sync"
)

// map是引用传递, 传递的是map的指针和slice一样
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key: ", key, ", value: ", value)
	}
}

func main() {

	var myMap map[string]string
	myIntMap := map[string]int{"1": 1, "2": 2}
	if myMap == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 开辟空间
	myMap = make(map[string]string)
	// slice和map就是最常用的
	fmt.Println(myMap)

	// 默认是""
	fmt.Println("Key不存在的时候: ", myMap["hello"])
	// int默认是0 都是一样的
	fmt.Println("Key不存在的时候: ", myIntMap["3"])

	myMap2 := map[string]interface{}{"name": "张三", "age": 18, "gender": true}
	//delete(m1, "name")
	fmt.Println(myMap2)

	go func() {
		for {
			myMap["1"] = "1"
		}
	}()
	go func() {
		for {
			_ = myMap["1"]
		}
	}()
	//select {}

	// sync map
	var m sync.Map
	m.Store("name", "张三")
	m.Store("age", "18")
	var k int
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		k++
		return true
	})

}
