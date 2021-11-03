package tests

import (
	"fmt"
	"testing"
	"beego/dao"
)

func testGetUsers(t *testing.T) {
	users, _ := dao.GetUsers()
	fmt.Println("users = ", users)
	for _, user := range users {
		fmt.Println("user = ", user)
	}
}

func testGetUserByID(t *testing.T) {
	user, _ := dao.GetUserByID(2)
	fmt.Println("user = ", user)
}

func testAddUser(t *testing.T) {
	dao.AddUser("shuwen", "3", "123456", "http://www.xstiku.com/logo.png", 18)
}

func testDeleteUserByID(t *testing.T) {
	dao.DeleteUserByUserID(3)
}

func testUpdateUserByID(t *testing.T) {
	dao.UpdateUserByID("shuwen", "3", "654321", "http://www.google.com/logo.png", 18, 46)
}
