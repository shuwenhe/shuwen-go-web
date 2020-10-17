package dao

import (
	"fmt"
	"testing"
)

func testGetUsers(t *testing.T) {
	users, _ := GetUsers()
	fmt.Println("users = ", users)
	for _, user := range users {
		fmt.Println("user = ", user)
	}
}

func testGetUserByID(t *testing.T) {
	user, _ := GetUserByID(2)
	fmt.Println("user = ", user)
}

func testAddUser(t *testing.T) {
	AddUser("shuwen", "3", "123456", "http://www.xstiku.com/logo.png", 18)
}

func testDeleteUserByID(t *testing.T) {
	DeleteUserByUserID(3)
}

func testUpdateUserByID(t *testing.T) {
	UpdateUserByID("shuwen", "3", "654321", "http://www.google.com/logo.png", 18, 46)
}
