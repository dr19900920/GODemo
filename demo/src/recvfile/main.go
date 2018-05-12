package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func WriteFile(fileName string, con net.Conn) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 4*1024)

	for {
		n, err1 := con.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println(err1)
			}
			return
		}
		f.Write(buf[:n])
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	con, err1 := listener.Accept()

	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer con.Close()

	buf := make([]byte, 1024)
	n, err2 := con.Read(buf)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fileName := string(buf[:n])
	con.Write([]byte("ok"))

	WriteFile(fileName, con)
}
