package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用于发送数据的管道
	Name string      // 用户名
	Addr string      // 网络地址
}

var onlineMap map[string]Client //保存在线用户
var message = make(chan string) //有消息来了 用管道进行通信

func HandleConn(con net.Conn) {
	defer con.Close()
	cliAddr := con.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli
	go WriteMsgToClient(cli, con)
	message <- MakeMsg(cli, "login")
	cli.C <- MakeMsg(cli, "i am here")

	isQuit := make(chan bool)  // 对方是否主动退出
	hasData := make(chan bool) // 对方是否有数据发送

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := con.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("con.Read ", err)
				return
			}

			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who" {
				con.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = "addr:" + tmp.Addr + "name:" + tmp.Name + "\n"
					con.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				// rename|name
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				con.Write([]byte("rename sucessful\n"))
			} else {
				message <- MakeMsg(cli, msg)
			}
			hasData <- true
		}

	}()

	for {
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)        // 当前用户从map移除
			message <- MakeMsg(cli, "logout") // 广播谁下线了
			return
		case <-hasData:

		case <-time.After(60 * time.Second):
			delete(onlineMap, cliAddr)        // 当前用户从map移除
			message <- MakeMsg(cli, "logout") // 广播谁下线了
			return
		}
	}
}

func WriteMsgToClient(cli Client, con net.Conn) {
	//每个Client 都循环等待管道来新的消息
	for msg := range cli.C {
		con.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	// 接收消息 并消息转发给每一个成员
	go Manager()

	for {
		con, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		go HandleConn(con)
	}

}
