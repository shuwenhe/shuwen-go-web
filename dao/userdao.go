package dao

import (
	"github.com/shuwenhe/shuwen-go-web/model"
	"github.com/shuwenhe/shuwen-go-web/utils"
)

// GetUsers Get users
func GetUsers() ([]*model.User, error) {
	sql := "select * from users"
	users := []*model.User{}
	err := utils.Db.Select(&users, sql) // Here you need to modify the value in the user structure, so you need to take the address
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id int) (*model.User, error) {
	user := &model.User{}
	sql := "select * from users where id = ?"
	err := utils.Db.Get(user, sql, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(name, num, password, logo string, age int) error {
	sql := "INSERT INTO users(name,num,password,logo,age) VALUES(?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, name, num, password, logo, age)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByUserID(userID int) error {
	sql := "delete from users where id=?"
	_, err := utils.Db.Exec(sql, userID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserByID(name, num, password, logo string, age int, id int) error {
	sql := "UPDATE users SET name=?,num=?,password=?,logo=?,age=? WHERE id=?"
	_, err := utils.Db.Exec(sql, name, num, password, logo, age, id)
	if err != nil {
		return err
	}
	return nil
}
