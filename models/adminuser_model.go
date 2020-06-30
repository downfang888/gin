package models

import (
	"fmt"
	"blogweb_gin/database"
)

type Adminuser struct {
	Id         int
	Username   string
	Password   string
	Status     int	// 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

//插入
func InsertAdminuser(adminuser Adminuser) (int64, error) {
	return database.ModifyDB("insert into adminuser(username,password,status,createtime) values (?,?,?,?)",
		adminuser.Username, adminuser.Password, adminuser.Status, adminuser.Createtime)
}

//按条件查询
func QueryAdminuserWightCon(con string) int {
	sql := fmt.Sprintf("select id from adminuser %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryAdminuserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryAdminuserWightCon(sql)
}


//根据用户名和密码，查询id
func QueryAdminuserWithParam(username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryAdminuserWightCon(sql)
}


