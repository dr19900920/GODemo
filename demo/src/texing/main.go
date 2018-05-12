package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person
	id   int
	addr string
	name string
}

func (tmp Student) PrintInfo() {
	fmt.Println(tmp)
}

func (p *Student) setInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {

	var s1 Student = Student{Person{"mike", 'm', 23}, 1, "bj", "aa"}
	fmt.Println(s1)

	s2 := Student{Person: Person{name: "haha"}, id: 1}
	s2.sex = 'f'
	s2.age = 18
	s2.name = "bvb"
	s2.Person.name = "vvv"
	s2.PrintInfo()

	var s3 Student
	(&s3).setInfo("mike", 'm', 23)
	s3.PrintInfo()

}
