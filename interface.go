package main

import "fmt"

func myfunc(arg interface{}) {
	fmt.Println("myFunc is called.. arg =", arg, "")
	fmt.Printf("type: %T", arg)
	// 判断arg是否是对应类型 emmmmm? type, value的pair
	value, ok := arg.(string) // 这是强转吗woc
	fmt.Println("value:", value, ", ok:", ok)
}

// interface是一个指针
type AnimalIF interface {
	Sleep()           // 动物睡觉
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 不需要实现接口
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {

	//var animal AnimalIF    // 接口的数据类型，父类的指针
	cat := &Cat{"Green"} // 但凡cat少实现了一个AnimalIF中的方法，这里就会报错 无法赋值
	dog := &Dog{"Blue"}
	showAnimal(cat)
	showAnimal(dog)

	myfunc("string你妹")
	myfunc(dog)

}
