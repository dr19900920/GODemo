package main

import (
	_ "fmt"
	_ "os"
)

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func Each(m map[string]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[string]interface{}, hf HandlerFunc) {
	Each(m, hf)
}

// func fetch(url string, ch chan string) {

// 		start := time.Now()
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			ch <- fmt.Sprint(err)
// 			return
// 		}
// 		nbyte, err := io.Copy(ioutil.Discard, resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			ch <- fmt.Sprint(err)
// 			return
// 		}
// 		secs := time.Since(start).Seconds()
// 		ch <- fmt.Sprintf("%.2fs  %7d   %s", secs, nbyte, url)
// 	}

// 	type Query struct {
// 		sql    chan string
// 		result chan string
// 	}

// 	func chanFucn() {
// 		q := Query{
// 			make(chan string),
// 			make(chan string),
// 		}

// 		sharded_var(q)

// 		fmt.Println(<-q.result)

// 		q.sql <- "123123"

// 		fmt.Println(<-q.result)
// 	}

// 	func execQuery(q Query) {
// 		go func() {
// 			fmt.Println("正在查询")
// 			//获取
// 			sql := <-q.sql
// 			//访问数据库, 输出查询结果
// 			q.result <- "get " + sql
// 		}()
// 	}

// 	func sharded_var(q Query) {
// 		go func() {
// 			value := "testString"
// 			for {
// 				select {
// 				case value = <-q.sql:
// 					fmt.Println("正在写入....")
// 				case q.result <- value:
// 					fmt.Println("正在载入....")
// 				default:
// 					break
// 				}
// 			}
// 		}()
// 	}

// 	type FileAction func(*Student) bool

// 	type FileFunc interface {
// 		FileAgeFunc(age int) FileAction
// 		FileHeightFunc(h int) FileAction
// 	}

// 	func File() {

// 		students := []*Student{
// 			&Student{"姚1", 15, 5},
// 			&Student{"姚2", 17, 3},
// 			&Student{"姚3", 18, 2},
// 		}

// 		result1 := Filter(students, FileAgeFunc(16))

// 		for _, s := range result1 {
// 			fmt.Println(s.Name)
// 		}

// 		result2 := Filter(students, FileHeightFunc(2))
// 		for _, s := range result2 {
// 			fmt.Println(s.Name)
// 		}
// 	}

// 	func FileAgeFunc(age int) FileAction {
// 		return func(s *Student) bool {
// 			return s.Age > age
// 		}
// 	}

// 	func FileHeightFunc(h int) FileAction {
// 		return func(s *Student) bool {
// 			return s.Height > h
// 		}
// 	}

// 	func Filter(l []*Student, f FileAction) []*Student {
// 		result := []*Student{}
// 		for _, s := range l {
// 			if f(s) {
// 				result = append(result, s)
// 			}
// 		}
// 		return result
// 	}
