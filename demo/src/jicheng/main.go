package main

import "fmt"

// 子集
type Humaner interface {
	sayHi()
}

// 超集
type Personer interface {
	Humaner // 匿名字段继承了sayHi()
	sing(lrc string)
}

type Student struct {
	name string
	age  int
}

func (tmp *Student) sayHi() {
	fmt.Println("sayHi")
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("sing")
}

func main() {
	test4()
}

func test1() {
	var i Personer
	s := &Student{"MIKE", 666}
	i = s
	i.sayHi()
	i.sing("hah")
}

func test2() {
	var iPro Personer // 超集
	var i Humaner     // 子集
	// 超集可以转换成子集
	i = iPro
	s := &Student{"MIKE", 666}
	i = s
	i.sayHi()
}

func test3() {
	var i interface{} = 1
	fmt.Println(i)
}

func test4() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student{"mike", 666}

	// 类型断言
	for index, data := range i {
		//		if value, ok := data.(int); ok == true {
		//			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		//		} else if value, ok := data.(string); ok == true {
		//			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		//		} else if value, ok := data.(Student); ok == true {
		//			fmt.Printf("x[%d] 类型为Student, 内容为name:%s,age:%d\n", index, value.name, value.age)
		//		}

		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d] 类型为Student, 内容为name:%s,age:%d\n", index, value.name, value.age)
		}
	}
}
