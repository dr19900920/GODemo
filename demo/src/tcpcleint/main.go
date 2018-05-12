package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	test1()
}

func test1() {
	con, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer con.Close()

	go func() {
		str := make([]byte, 1024)
		for {
			n, err1 := os.Stdin.Read(str)
			if err1 != nil {
				fmt.Println("con read err1", err1)
				return
			}

			con.Write(str[:n])
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, er := con.Read(buf)
		if er != nil {
			fmt.Println("con read er", er)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
