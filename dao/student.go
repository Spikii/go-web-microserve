package dao

import (
	"beego/models"
	"beego/utils"
)

func Login(num string) (*models.Student, error) {
	sql := "select id,num,name,pass,phone from student where num=?"
	row := utils.Db.QueryRow(sql, num)
	s := &models.Student{}
	err := row.Scan(&s.ID, &s.Num, &s.Name, &s.Pass, &s.Phone)
	if err != nil {
		return nil, err
	}
	return s, nil
}
