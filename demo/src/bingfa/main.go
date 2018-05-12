package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) // 延迟1秒
	}
}

func main() {

}

func test1() {
	go newTask() // 新建一个协程, 新建一个任务
	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second) // 延迟1秒
	}
}

//主协程退出子协程不复存在
func test2() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i=", i)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i=", i)
		if i == 2 {
			break
		}
	}
}
