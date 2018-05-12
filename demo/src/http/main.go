package main

import (
	"fmt"
	"net"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello,world")
}

func main() {
	test2()
}

func test2() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// 捕获请求头
func test1() {
	listener, err := net.Listen("tcp", ":8000")
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
	if n == 0 {
		fmt.Println(err2)
		return
	}
	fmt.Printf("#%v#", string(buf[:n]))
}
