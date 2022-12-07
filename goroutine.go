package main

import (
	"fmt"
	"os"
	"time"
)

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	f, err := os.Open("/tmp/test.go")
	if err != nil {
		return
	}
	fmt.Println("finish open a file ")
	defer f.Close()
	time.Sleep(2 * time.Second)
	resultChan <- sum

}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 为map/slice/channel分配内存并返回一个初始化的结果 有点像Array.fill
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	key := "Hello"
	switch key {
	case "Hell":
		fmt.Println()
	default:
		fmt.Println()
	}
	// For Select 是用select来做到IO操作 监听多个IO的读写操作 当有数据已经准备好的时候 (已经在内核空间的时候 就会有响应不然会阻塞'
	sum1 := <-resultChan
	sum2 := <-resultChan
	fmt.Println("Sum1: ", sum1)
	fmt.Println("Sum2: ", sum2)
	fmt.Println("Result: ", sum1+sum2)
	// 为什么go是直接执行吗 因为没有遇到IO所以结果会是这样吗，加点IO
	// 实现超时机制 专门有一个超时的goroutine -> sleep完了之后就往里面塞truea
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	}

	// select的default的功能是当其他的case都阻塞住的时候会触发，可以用来检测ch是否已满
}
