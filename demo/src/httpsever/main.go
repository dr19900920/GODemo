package main

import (
	"fmt"
	//	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello go"))
	fmt.Println("r.Method =", r.Method)
	fmt.Println("r.URL =", r.URL)
	fmt.Println("r.Header =", r.Header)
	fmt.Println("r.Body =", r.Body)
	fileName := "1.html"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer f.Close()

	var result string
	buf := make([]byte, 1024*4)
	for {
		n, err1 := f.Read(buf)
		if n == 0 {
			fmt.Println("err1 =", err1)
			break
		}
		result += string(buf)
	}
	w.Write([]byte(result))
	//	fmt.Println(string(result))
}

func main() {
	http.HandleFunc("/1.html", handler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
