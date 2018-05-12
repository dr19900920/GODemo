package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpidePage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func SpidePage(i int, page chan<- int) {
	//	url := "http://tieba.baidu.com/f?kw=%e7%bb%9d%e5%9c%b0%e6%b1%82%e7%94%9f&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Println("url = ", url)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	// 把内容写入文件
	//	fileName := strconv.Itoa(i) + ".html"
	//	f, err1 := os.Create(fileName)
	//	if err1 != nil {
	//		fmt.Println("err1 =", err1)
	//		return
	//	}
	//	f.WriteString(result)
	//	f.Close()
	//	page <- i
	re1 := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))" target="_blank">`)
	if re1 == nil {
		fmt.Println("regexp.MustCompile error")
		return
	}
	joyUrls := re1.FindAllStringSubmatch(result, -1) // 最后一个参数为-1,找到所有
	//	fmt.Println("joyUrls", joyUrls)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	for m, data := range joyUrls {

		//		fmt.Println("data =", data[1])
		title, content, err3 := SpideOneJoy(data[1])

		if err3 != nil {
			fmt.Println("SpideOneJoy error =", err3)
			continue
		}
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
		//		fmt.Println("title = ", title)
		//		fmt.Println("content =", content)
	}
	StoreJoyToFile(i, fileTitle, fileContent)
	page <- i

}

func SpideOneJoy(url string) (title, content string, err error) {

	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile error")
		return
	}
	tmpTitle := re1.FindAllStringSubmatch(result, 1) // 最后一个参数为1,只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		//		title = strings.Replace(title,"\r","",-1)
		//		title = strings.Replace(title,"\n","",-1)
		//		title = strings.Replace(title," ","",-1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile error")
		return
	}
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		break
	}

	return
}

func StoreJoyToFile(i int, fileTitle, fileContent []string) {
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer f.Close()
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n")
		f.WriteString(fileContent[i] + "\n")
		f.WriteString("\n=============================\n")
	}
	f.Close()
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >=1 ) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	DoWork(start, end)
}
