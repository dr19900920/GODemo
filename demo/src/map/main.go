package main

import "fmt"

func main() {
	d := map[int]string{110: "haha", 11: "go"}
	d[22] = "c++"
	fmt.Println(d)
	for key, value := range d {
		fmt.Println(key, value)
	}
	delete(d, 110)
	fmt.Println(d)
	//map是引用传递
}
