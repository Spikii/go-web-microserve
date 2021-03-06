package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"beego/dao"
	"beego/models"
	"beego/utils"
)


//数据库crud操作
func GetClassByID(w http.ResponseWriter, r *http.Request) {
	stringID := r.FormValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "Data is error", err.Error()))
		return
	}
	class, err := dao.GetClassByID(id)
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "Get class fail", err.Error()))
		return
	}
	w.Write(utils.NewResult(utils.Success, "Get class success", class))
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := dao.GetClasses()
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "Don't get the classes", err.Error()))
		return
	}
	w.Write(utils.NewResult(utils.Success, "Get the classes", classes))
}

func DeleteClassByID(w http.ResponseWriter, r *http.Request) {
	stringID := r.FormValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "Delete the data is fail,please try it again!", err.Error()))
		return
	} else {
		dao.DeleteClassByID(id)
		w.Write(utils.NewResult(utils.Success, "Delete the data is success!"))
		return
	}
}

func AddClass(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Println("name = ", name)
	des := r.FormValue("des")
	fmt.Println("name = ", name)
	c := &models.Class{
		Name: name,
		Des:  des,
	}
	if c.Name == "" {
		w.Write(utils.NewResult(utils.Fail, "Class name is not null"))
		return
	}
	err := dao.AddClass(c)
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "Add class fail!"))
		return
	}
	w.Write(utils.NewResult(utils.Success, "Add class success!"))
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	stringID := r.FormValue("id")
	name := r.FormValue("name")
	des := r.FormValue("des")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "insert fail", "insert error"))
		return
	}
	c := &models.Class{
		ID:   id,
		Name: name,
		Des:  des,
	}
	err2 := dao.UpdateClass(c)
	if err2 != nil {
		w.Write(utils.NewResult(utils.Fail, "Update fail", "Update err"))
		return
	}
	w.Write(utils.NewResult(utils.Success, "Update succcess", "Update success"))
}

