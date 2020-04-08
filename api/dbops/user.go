/*
@Time : 06/04/2020 10:43 PM 
@Author : GC
*/
package dbops

import (
	"database/sql"
	"video_server_api/defs/models"
)

// 新增用户
func AddUser(name, password string) error {
	stmt, err := db.Prepare("INSERT INTO users (username,password) VALUES(?,?)")
	if err != nil {
		return err
	} else {
		// prepare 时出错后调用 Close会造成 panic
		defer stmt.Close()
	}
	if _, err = stmt.Exec(name, password); err != nil {
		return err
	}
	return nil
}

// 查询单个用户
func GetUser(name string) (*models.User, error) {
	stmt, err := db.Prepare("SELECT id, password FROM users WHERE username = ? ORDER BY id ASC LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var id int
	var password string
	if err := stmt.QueryRow(name).Scan(&id, &password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &models.User{Id: id, Username: name, Password: password}, nil
}
