package models

import (
	"database/sql"
)

var (
	//数据库操作对象
	DataBase *sql.DB = nil
)