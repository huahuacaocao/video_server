/*
@Time : 07/04/2020 8:00 PM 
@Author : GC
*/
package dbops

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

// 初始化 DB
func init() {
	var err error
	if db, err = sql.Open("mysql", "root:a00000@tcp(127.0.0.1:33060)/video?charset=utf8"); err != nil {
		panic(errors.New("数据库连接信息无效"))
	} else {
		if db.Ping() != nil {
			panic(errors.New("数据库连接失败"))
		}
	}
}
