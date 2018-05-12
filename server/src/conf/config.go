package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type DBCofig struct {
	Host         string `json:"host"`           //连接地址
	Username     string `json:"username"`       //用户名
	Password     string `json:"password"`       //用户密码
	Name         string `json:"name"`           //数据库名
	Charset      string `json:"charset"`        //
	MaxIdleConns int    `json:"max_idle_conns"` //连接池最大空闲连接数
	MaxOpenConns int    `json:"max_open_conns"` //连接池最大连接数
}


//读取配置文件
func ReadConfig(path string) (*DBCofig, error){
	config := new(DBCofig)
	err := config.Parse(path)
	if err != nil {
		return nil, err
	}

	return config, nil
}

//解析配置文件
func (this *DBCofig) Parse(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &this)
	if err != nil {
		return err
	}
	return nil
}

//链接数据库
func (this *DBCofig)Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", this.Username, this.Password, this.Host, this.Name))
	if err != nil {
		return nil, err
	}
	return db, nil
}