package main

import "fmt"
import "errors"

func main() {
	result, err := MyDiv(1, 0)
	fmt.Println(result, err)
}

func test1() {
	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println("err1 =", err1)
	err2 := errors.New("hehe")
	fmt.Println("err2 =", err2)
}

func test2() {
	// 使用revoer捕获异常 不会崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}
}

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}
