package main

import (
	"container/list"
	"fmt"
)

func main() {

	li := list.New()
	li.PushBack("last")
	e := li.PushFront("first")
	fmt.Println("e: ", e)
	e2 := li.InsertAfter("2", e)
	fmt.Println("e2: ", e2)
	li.InsertBefore("3", e2)
	head := li.Front()
	fmt.Println(">>>>> li: ", li)

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next()
	}

	//fmt.Println(li.Front().Value)

}
