package main

//import "fmt"

//func main() {

//	fmt.Println("aaaaaaa")
//}

//import . "fmt"

//func main() {
//	Println("bbbbbb")
//}

//import io "fmt"

//func main() {
//	io.Println("3333333333")
//}

import (
	"calc"
	"fmt"
)

func main() {
	// 同一个目录直接可以调用
	haha()
	result := calc.Add(1, 2)
	fmt.Println("result = ", result)
}
