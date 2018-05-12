package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type IT struct {
	Company  string // `json:"company"` //tag
	Subjects []string
	IsOK     bool // `json:",string"` value类型转换
	Price    float64
}

func main() {
	//	path := "/Users/dengrui/Desktop/tmp.txt"
	//	os.Stdout.Close() // 关闭后无法输出
	//	fmt.Println("haha")
	//	WriteFile(path)
	//	ReadFile(path)
	//	ReadFileLine(path)

	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFileName dstFileName")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]

	if srcFileName == dstFileName {
		fmt.Println("路径不能相同")
		return
	}
	sf, err1 := os.Open(srcFileName)

	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}

	df, err2 := os.Create(dstFileName)

	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}

	defer sf.Close()
	defer df.Close()

	buf := make([]byte, 4*1024)
	for {
		n, err := sf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err=", err)
		}
		df.Write(buf[:n])
	}

}

func ReadFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		buf, err1 := r.ReadBytes('\n')
		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			fmt.Println("err1=", err1)
		}
		fmt.Printf("buf=#%s#\n", string(buf))
	}

}

func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*2)

	n, err1 := f.Read(buf)

	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1=", err1)
		return
	}
	fmt.Println("buf =\n", string(buf[:n]))

}

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		n, err1 := f.WriteString(buf)
		if err1 != nil {
			fmt.Println("err =", err1)
		}
		fmt.Println("n = ", n)
	}

}

func test1() {
	s := IT{"go", []string{"c++", "swift", "oc"}, true, 66666.66}
	data, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Println(string(data))

	m := make(map[string]interface{}, 4)
	m["Company"] = "go"
	m["Subjects"] = []string{"c++", "swift", "oc"}
	m["IsOK"] = true
	m["Price"] = 44.44

	// 格式化indent
	result, _ := json.MarshalIndent(m, "", "	")
	fmt.Println(string(result))

	jsonBuf := `
	{
	"Company": "go",
	"IsOK": true,
	"Price": 44.44,
	"Subjects": [
		"c++",
		"swift",
		"oc"
	] []interface{}//万能指针 空切片
    }`

	var tmp IT
	err1 := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(tmp)
	fmt.Println(tmp.Subjects)

	tmp2 := make(map[string]interface{}, 4)
	err2 := json.Unmarshal([]byte(jsonBuf), &tmp2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(tmp2)
}
