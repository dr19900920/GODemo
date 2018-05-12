package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	buf := "abc azc a7c aac 888 a9c tac"

	//	reg1 := regexp.MustCompile("a.c")
	//	reg1 := regexp.MustCompile("a[0-9]c")
	reg1 := regexp.MustCompile("a\\dc")
	if reg1 == nil {
		fmt.Println("正则表达式解析失败")
		return
	}
	slice := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("slice = ", slice)

}

func test1() {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10) //10进制
	slice = strconv.AppendQuote(slice, "abcgohello")
	fmt.Println("slice =", string(slice))

	var str string
	str = strconv.FormatBool(true)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	str = strconv.Itoa(666)
	fmt.Println("str =", str)

	var flag bool
	var err error

	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag =", flag)
	} else {
		fmt.Println("err =", err)
	}

	a, _ := strconv.Atoi("567")
	fmt.Println("a =", a)

}
