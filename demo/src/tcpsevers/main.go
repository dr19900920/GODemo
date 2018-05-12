package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleCon(con net.Conn) {
	defer con.Close()
	addr := con.RemoteAddr().String()
	fmt.Println("addr connect sucessful", addr)
	for {
		buf := make([]byte, 1024)
		n, err1 := con.Read(buf)

		if err1 != nil {
			fmt.Println(err1)
			return
		}
		fmt.Println("buf =", string(buf[:n]))
		if string(buf[:n-1]) == "exit" {
			fmt.Println(addr, " exit")
			return
		}
		con.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()
	for {
		con, er := listener.Accept() //类似channel 阻塞等待用户链接
		if er != nil {
			fmt.Println(er)
			return
		}

		go HandleCon(con) // 新建协程,避免Conn被替换
	}

}
