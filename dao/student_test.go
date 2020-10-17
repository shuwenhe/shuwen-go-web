package dao

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	num := "1"
	student, _ := Login(num)
	fmt.Println("student = ", student)
}
