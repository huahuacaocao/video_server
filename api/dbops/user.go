/*
@Time : 06/04/2020 10:43 PM 
@Author : GC
*/
package dbops

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
