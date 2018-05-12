package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("resp.Status =", resp.Status)
	fmt.Println("resp.StatusCode =", resp.StatusCode)
	fmt.Println("resp.Header =", resp.Header)

	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err1 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err1", err1)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp =", tmp)
}
