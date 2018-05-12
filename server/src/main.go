package main

import (
	"conf"
	"flag"
	"log"
	"models"
	"net/http"
	"router"
)

const (
	Name    string = "Yao Go"
	Version string = "1.0"
)

var (
	portStr    = flag.String("port", ":8080", "默认端口")
	configPath = flag.String("config", "/Users/dengrui/Applications/Go/yaozi/src/config.json", "使用配置文件")
)

func main() {

	log.Println("*********************************************")
	log.Printf("           系统:[%s]版本:[%s]", Name, Version)
	log.Println("*********************************************")

	flag.Parse()

	//读取配置文件
	config, err := conf.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("读取配置文件错误: %s", err)
	}

	//开启服务
	Start(config)
}

func Start(config *conf.DBCofig) {

	var err error
	models.DataBase, err = config.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer models.DataBase.Close()

	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(*portStr, r))
}

//supervisor golang 部署工具
