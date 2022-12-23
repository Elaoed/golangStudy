package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called..")
	fmt.Println("%v\n", this)
}
func main() {
	user := User{1, "Delphi", 10}
	DoFileAndMethod(user)
}

func DoFileAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is: ", inputType.Name())
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is: ", inputValue)

	fmt.Println(">>>>> 通过反射遍历打印入参对象字段")
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 为什么不起作用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}
