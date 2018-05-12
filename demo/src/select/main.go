package main

import (
	"fmt"
	"time"
)

func fibonace(ch chan<- int, quit <-chan bool) {

	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}

}

func main() {

	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}

	<-quit
	fmt.Println("程序结束")

}

func test() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()
	fibonace(ch, quit)
}
