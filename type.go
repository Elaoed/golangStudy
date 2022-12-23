package main

import "fmt"

// 作为一种别名，实际上没啥用
type myint int

// 类名首字母大写表示共有，对外能够访问，否则只能在类的内部使用
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string {
	fmt.Println("Name = ", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	// 当前的this是对象，调用该方法对象的一个拷贝
	// 加指针和不加指针的区别，形参和实参
	this.Name = newName
}

func printObj(book Book) {
	book.title = "666"
	fmt.Println(book)
}

// 定义是指针，但是使用的时候不需要用*
func changeBook(book *Book) {
	book.title = "777"
}

// 就是一个对象
type Bird struct {
}

type Book struct {
	title  string
	author string
}

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// 可以多继承吗？
type SuperMan struct {
	Human // SuperMan集成Human类的方法
	level int
}

// 重定义父类的方法Eat
func (this *SuperMan) Eat() {
	fmt.Println("Superman.Eat()...")
}
func (this *SuperMan) Fly() {
	fmt.Println("Superman.Fly()...")
}

func main() {

	var book1 Book
	book1.title = "小王子"
	book1.author = "圣什么什么"
	fmt.Println(book1)
	printObj(book1)
	fmt.Println(book1)
	changeBook(&book1)
	fmt.Println(book1)

	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("嘟嘟嘟")
	hero.Show()

	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	s := SuperMan{Human{"li4", "male"}, 5}
	s.Walk()
	s.Eat()
	s.Fly()

	hero2 := Hero{Ad: 100}
	fmt.Println(hero2, hero2.Name == "", hero2.Level == 0)

}
