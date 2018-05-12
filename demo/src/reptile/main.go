package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// http://tieba.baidu.com/f?kw=%e7%bb%9d%e5%9c%b0%e6%b1%82%e7%94%9f&ie=utf-8&pn=0

func HttpGet(url string) (result string, err error) {

	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			if err2 != nil && err2 == io.EOF {

			} else {
				err = err2
			}
			break
		}
		result += string(buf)
	}
	return
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 页数据\n", start, end)

	for i := start; i <= end; i++ {
		url := "http://tieba.baidu.com/f?kw=%e7%bb%9d%e5%9c%b0%e6%b1%82%e7%94%9f&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url = ", url)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("err =", err)
			continue
		}
		// 把内容写入文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("err1 =", err1)
			continue
		}
		f.WriteString(result)
		f.Close()
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >=1 ) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	DoWork(start, end)
}
