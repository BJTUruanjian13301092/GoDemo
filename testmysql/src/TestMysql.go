package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306"//IP地址
	dbusername = "root"//用户名
	dbpassword = "haha123"//密码
	dbname     = "publicfund"//数据库名
)

func main() {
	db, err := sql.Open("mysql", "root:haha123@tcp(127.0.0.1:3306)/publicfund?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM fund_feature_category")
	checkErr(err)

	for rows.Next() {
		var id int64
		var category string
		var source string
		err = rows.Scan(&id, &category, &source)
		checkErr(err)
		fmt.Println(id, category, source,"\n")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}