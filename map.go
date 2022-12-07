package main

import (
	"fmt"
	"sync"
)

func main() {

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

	// sync map
	var m sync.Map
	m.Store("name", "张三")
	m.Store("age", "18")
	var k int
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		k++
		return true
	})

	mm1 := make(map[string]string)
	fmt.Println(mm1)
	mm1 = map[string]string{}
	fmt.Println(mm1)
}
