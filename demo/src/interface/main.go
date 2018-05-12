package main

import "fmt"

type Humaner interface {
	sayHi()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	addr  string
	group int
}

func (tmp *Student) sayHi() {
	fmt.Println(tmp.name, tmp.id)
}

func (tmp *Teacher) sayHi() {
	fmt.Println(tmp.addr, tmp.group)
}

func WhoSayHi(i Humaner) {
	i.sayHi()
}

func main() {

	s := &Student{"haha", 1}
	WhoSayHi(s)
	t := &Teacher{"bj", 11}
	WhoSayHi(t)
	x := make([]Humaner, 2)
	x[0] = s
	x[1] = t

	for i := range x {
		WhoSayHi(x[i])
		fmt.Println(i, x[i])
	}
}
