package dao

import (
	"github.com/shuwenhe/shuwen-go-web/model"
	"github.com/shuwenhe/shuwen-go-web/utils"
)

func GetClasses() ([]*model.Class, error) {
	sql := "select * from class"
	classes := []*model.Class{}
	err := utils.Db.Select(&classes, sql)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func DeleteClassByID(id int) error {
	sql := "delete from class where id=?"
	_, err := utils.Db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func AddClass(c *model.Class) error {
	sql := "insert into class(name,des) values(?,?)"
	_, err := utils.Db.Exec(sql, c.Name, c.Des)
	if err != nil {
		return err
	}
	return nil
}

func UpdateClass(c *model.Class) error {
	sql := "update class set name=?,des=? where id=?"
	_, err := utils.Db.Exec(sql, c.Name, c.Des, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetClassByID(id int) (*model.Class, error) {
	sql := "select id,name,des from class where id=?"
	row := utils.Db.QueryRow(sql, id)
	c := &model.Class{}
	err := row.Scan(&c.ID, &c.Name, &c.Des)
	if err != nil {
		return nil, err
	}
	return c, nil
}
