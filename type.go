package main

import "fmt"

type Bird struct {
}
type IFly interface {
	Fly()
}

func (b *Bird) Fly() {
	fmt.Println("Im flying")

}
func main() {
	fly := new(Bird)
	fly.Fly()
}
