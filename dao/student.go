package dao

import (
	"github.com/shuwenhe/shuwen-go-web/model"
	"github.com/shuwenhe/shuwen-go-web/utils"
)

func Login(num string) (*model.Student, error) {
	sql := "select id,num,name,pass,phone from student where num=?"
	row := utils.Db.QueryRow(sql, num)
	s := &model.Student{}
	err := row.Scan(&s.ID, &s.Num, &s.Name, &s.Pass, &s.Phone)
	if err != nil {
		return nil, err
	}
	return s, nil
}
