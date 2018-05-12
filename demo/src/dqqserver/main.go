package main

import (
	"fmt"
	"net"

	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	//	"strings"
	//	"database/sql"
	//	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var UserName string = "root"
var Password string = "abc#123"
var Host string = "39.106.54.49:3306"
var Name string = "drsql"

type DBCofig struct {
	Host         string `json:"host"`           //连接地址
	Username     string `json:"username"`       //用户名
	Password     string `json:"password"`       //用户密码
	Name         string `json:"name"`           //数据库名
	Charset      string `json:"charset"`        //
	MaxIdleConns int    `json:"max_idle_conns"` //连接池最大空闲连接数
	MaxOpenConns int    `json:"max_open_conns"` //连接池最大连接数
}

//{
//  "host": "mysql.test.neoteched.com",
//  "username": "root",
//  "password": "yaochizaocan",
//  "charset": "utf8mb4",
//  "name": "crm",
//  "max_idle_conns": 10,
//  "max_open_conns": 50
//}

type User struct {
	gorm.Model
	id         int
	name       string
	decripsion string
	password   string
	user_name  string
}

type Respo struct {
	Code    int    `json:"code"`
	IsOK    bool   `json:"isok"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form.Get("user")
	pasword := r.Form.Get("password")
	fmt.Println("r.Method =", r.Method)
	fmt.Println("r.URL =", r.URL)
	fmt.Println("r.Header =", r.Header)
	fmt.Println("r.Body =", r.Body)
	fmt.Println("user =", user)
	fmt.Println("password =", pasword)

	db, err1 := Connect()
	defer db.Close()

	if err1 != nil {
		fmt.Println("err1 =", err1)
		BackToClient(w, 404, false, err1.Error())
		return
	}
	fmt.Println("数据库访问成功")
	if !db.HasTable("users") {
		fmt.Println("数据库不包含user表")
		BackToClient(w, 404, false, "数据库不包含user表")
		return
	}

	user_db := db.Table("users")
	var user_info User
	user_db.Where("user_name = ?", user).First(&user_info)
	fmt.Println(user_info.name)

}

//链接数据库
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Host, Name))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func BackToClient(w http.ResponseWriter, code int, isok bool, msg string) {
	respo := Respo{Code: code, IsOK: isok, Message: msg}
	data, err2 := json.Marshal(respo)
	if err2 != nil {
		fmt.Println("err2 =", err2)
		return
	}
	fmt.Println("data =", string(data))
	w.Write(data)
}

func main() {
	http.HandleFunc("/login", handler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func test1() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()
	for {
		con, er := listener.Accept() //类似channel 阻塞等待用户链接
		if er != nil {
			fmt.Println(er)
			return
		}

		go HandleCon(con) // 新建协程,避免Conn被替换
	}
}

func HandleCon(con net.Conn) {
	defer con.Close()
	addr := con.RemoteAddr().String()
	fmt.Println("addr connect sucessful", addr)
	for {
		buf := make([]byte, 1024)
		n, err1 := con.Read(buf)

		if err1 != nil {
			fmt.Println(err1)
			return
		}
		fmt.Println("buf =", string(buf[:n]))
		con.Write([]byte("loginsuccess"))
	}
}
