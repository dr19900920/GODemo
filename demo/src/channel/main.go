package main

import (
	"fmt"
	"time"
)

//全局变量,创建一个channel
var ch = make(chan int)

// 定义一个打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func main() {

}

func producer(cha chan<- int) {
	for i := 0; i < 10; i++ {
		cha <- i * i
	}
	close(cha)
}

func consumer(cha <-chan int) {
	for num := range cha {
		fmt.Println("num =", num)
	}
}

//单向channel的应用
func test6() {
	c := make(chan int)
	go producer(c) // channel传参是引用传递
	consumer(c)
}

// 单向channel
func test5() {
	c := make(chan int)
	var writech chan<- int = c // 只能写
	writech <- 666
	var readch <-chan int = c // 只能读
	<-readch
	//单向无法转换成双向
}

//关闭通道
func test4() {
	c := make(chan int)
	go func() {
		for i := 0; i < 3000; i++ {
			c <- i
		}
		close(c)
		//		c <- 666 // 关闭后无法写数据,但是可以读数据

	}()

	for {
		if num, ok := <-c; ok == true {
			fmt.Println(num)
		} else {
			break
		}
	}

}

// 有缓存的channel 不会阻塞 除非大于管道容量会阻塞
func test3() {
	c := make(chan int, 3)
	fmt.Printf("len(ch) - %d, cap(ch) - %d\n", len(c), cap(c))
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i - ", i)
			c <- i
			fmt.Printf("len(ch) - %d, cap(ch) - %d\n", len(c), cap(c))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
}

// 无缓存的channel 会阻塞
func test2() {
	c := make(chan int, 0)
	fmt.Printf("len(ch) - %d, cap(ch) - %d\n", len(c), cap(c))
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i - ", i)
			c <- i
			fmt.Printf("len(ch) - %d, cap(ch) - %d\n", len(c), cap(c))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
}

// 通过channel实现数据交互
func test1() {

	cha := make(chan string)

	defer fmt.Println("主协程工作完毕")

	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i =", i)
			time.Sleep(time.Second)
		}

		cha <- "haha"
	}()

	str := <-cha
	fmt.Println(str)
}

// 通过channel实现同步
func test() {
	// 新建2个协程,代表两个人,2个人同时使用打印机 发生资源竞争
	go person1()
	go person2()

	for {

	}
}

// person1执行完后,才能到person2执行, 通过管道实现同步
func person1() {
	Printer("hello")
	ch <- 666 // 给管道写数据
}

func person2() {
	<-ch // 从管道取数据,接收,如果没有数据就会阻塞
	Printer("world")
}
