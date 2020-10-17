package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-go-web/model"
)

func testGetClasses(t *testing.T) {
	classes, _ := GetClasses()
	for _, class := range classes {
		fmt.Println("class =", class)
	}
}

func testDeleteClassByID(t *testing.T) {
	DeleteClassByID(1)
}

func testAddClass(t *testing.T) {
	c := &model.Class{
		Name: "Shuwen",
		Des:  "123456",
	}
	AddClass(c)
}

func testUpdateClass(t *testing.T) {
	class := &model.Class{
		ID:   7,
		Name: "Shuwen",
		Des:  "654321",
	}
	UpdateClass(class)
}

func TestGetClassByID(t *testing.T) {
	c, _ := GetClassByID(7)
	fmt.Println("c = ", c)
}
