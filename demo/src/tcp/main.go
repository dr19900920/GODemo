package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()

	con, er := listener.Accept() //类似channel 阻塞等待用户链接
	if er != nil {
		fmt.Println(er)
		return
	}

	buf := make([]byte, 1024)
	n, err1 := con.Read(buf)
	defer con.Close()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("n =", n)
	fmt.Println("buf =", string(buf[:n]))

}
