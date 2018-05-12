package main

import (
	"fmt"
	"time"
)

func main() {
	test5()
}

//循环定时器
func test5() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		num := <-ticker.C
		fmt.Println("num =", num)
	}

}

//重置时间
func test4() {
	timer := time.NewTimer(3 * time.Second)
	ok := timer.Reset(1 * time.Second)
	<-timer.C // channel没有数据前后阻塞
	fmt.Println("时间到", ok)
}

func test3() {
	timer := time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("zi")
	}()
	timer.Stop() // 定时器停止, 子协程的代码是延迟两秒执行的所以打印无效
	for {

	}
}

// 延迟
func test2() {
	<-time.After(2 * time.Second)
	fmt.Println("时间到")
}

//channel造成的死锁 timer只会响应一次
func test1() {
	timer := time.NewTimer(2 * time.Second)
	for {
		<-timer.C
		fmt.Println("时间到")
		//		timer = time.NewTimer(2 * time.Second)
	}
}

func test() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间:", time.Now())
	t := <-timer.C // channel没有数据前后阻塞
	fmt.Println("t = ", t)
}
