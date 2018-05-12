package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, con net.Conn) {

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err1 := f.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("err1 =", err1)
			}
			return
		}
		con.Write(buf[:n])

	}
}

func main() {
	fmt.Println("请输入要传输的文件:")
	var path string
	fmt.Scan(&path)

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	con, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer con.Close()

	_, err2 := con.Write([]byte(info.Name()))

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	buf := make([]byte, 1024)
	n, _ := con.Read(buf)
	if "ok" == string(buf[:n]) {
		SendFile(path, con)
	}

}
