package models

import (
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	RoleId   int    `json:"roleId"`
	SchoolId int    `json:"schoolId"`
}



/**
* 检查账号是否存在
*/
func CheckAccount(email string) (int, error){

	var num int
	rows, err := DataBase.Query("SELECT COUNT(*) FROM user WHERE email=?", email)
	if err != nil {
		log.Print(err)
		return -1, &DatabaseError{"输入的邮箱不存在"}
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&num)
	}
	return num, nil
}

//获取用户信息
func GetUserInfo(email string) (*User, error) {
	var user User
	row := DataBase.QueryRow("SELECT id, name, email, mobile, role_id, school_id FROM user WHERE email=?", email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Mobile, &user.RoleId, &user.SchoolId)
	if err != nil {
		log.Print("获取用户信息:", err)
		return nil, &DatabaseError{"查询用户数据错误"}
	}
	return &user, nil
}
