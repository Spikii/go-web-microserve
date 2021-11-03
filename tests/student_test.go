package tests

import (
	"fmt"
	"testing"
	"beego/dao"
)

func TestLogin(t *testing.T) {
	num := "1"
	student, _ := dao.Login(num)
	fmt.Println("student = ", student)
}
