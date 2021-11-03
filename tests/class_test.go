package tests

import (
	"fmt"
	"testing"

	"beego/models"
	"beego/dao"
)

func testGetClasses(t *testing.T) {
	classes, _ := dao.GetClasses()
	for _, class := range classes {
		fmt.Println("class =", class)
	}
}

func testDeleteClassByID(t *testing.T) {
	dao.DeleteClassByID(1)
}

func testAddClass(t *testing.T) {
	c := &models.Class{
		Name: "Shuwen",
		Des:  "123456",
	}
	dao.AddClass(c)
}

func testUpdateClass(t *testing.T) {
	class := &models.Class{
		ID:   7,
		Name: "Shuwen",
		Des:  "654321",
	}
	dao.UpdateClass(class)
}

func TestGetClassByID(t *testing.T) {
	c, _ := dao.GetClassByID(7)
	fmt.Println("c = ", c)
}
